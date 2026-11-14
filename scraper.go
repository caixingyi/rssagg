package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/caixingyi/rssagg/internal/database"
	"github.com/google/uuid"
)

// startScraping scrapes feeds concurrently at the specified interval.
// It stops when ctx is cancelled (e.g. on graceful shutdown).
func startScraping(ctx context.Context, db *database.Queries, concurrent int, timeBetweenRequests time.Duration) {
	log.Printf("Scraping on %v goroutines every %v duration", concurrent, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Scraper stopping:", ctx.Err())
			return
		case <-ticker.C:
			scrapeFeeds(ctx, db, concurrent)
		}
	}
}

// scrapeFeeds fetches the next batch of feeds and scrapes them concurrently.
func scrapeFeeds(ctx context.Context, db *database.Queries, concurrent int) {
	feeds, err := db.GetNextFeedsToFetch(ctx, int32(concurrent))
	if err != nil {
		log.Println("Error fetching feeds:", err)
		return
	}
	log.Printf("Got %v feeds to fetch", len(feeds))

	wg := sync.WaitGroup{}
	for _, feed := range feeds {
		wg.Add(1)
		go func(feed database.Feed) {
			defer wg.Done()
			scrapeFeed(ctx, db, feed)
		}(feed)
	}
	wg.Wait()
}

// scrapeFeed fetches a single feed, saves its posts and marks it as fetched.
func scrapeFeed(ctx context.Context, db *database.Queries, feed database.Feed) {
	// Bound the whole scrape of a single feed so one slow feed can't block the batch forever.
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	rssFeed, err := urlToFeed(ctx, feed.Url)
	if err != nil {
		log.Println("Error scraping feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %q: %v", item.PubDate, err)
			continue
		}
		_, err = db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: publishedAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Println("Error saving post:", err)
		}
	}

	if _, err := db.MarkFeedAsFetched(ctx, feed.ID); err != nil {
		log.Printf("Error marking feed %v as fetched: %v", feed.Url, err)
		return
	}
	log.Printf("Finished scraping feed: %v", feed.Url)
	log.Printf("Total posts saved for feed: %v", len(rssFeed.Channel.Items))
}
