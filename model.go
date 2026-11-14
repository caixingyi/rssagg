package main

import (
	"time"

	"github.com/caixingyi/rssagg/internal/database"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ApiKey    string `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID.String(),
		Name:      feed.Name,
		URL:       feed.Url,
		CreatedAt: feed.CreatedAt.Format(time.RFC3339),
		UpdatedAt: feed.UpdatedAt.Format(time.RFC3339),
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	feedList := make([]Feed, len(feeds))
	for i, feed := range feeds {
		feedList[i] = databaseFeedToFeed(feed)
	}
	return feedList
}

type FeedFollow struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	FeedID    string `json:"feed_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func databaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID.String(),
		UserID:    feedFollow.UserID.String(),
		FeedID:    feedFollow.FeedID.String(),
		CreatedAt: feedFollow.CreatedAt.Format(time.RFC3339),
		UpdatedAt: feedFollow.UpdatedAt.Format(time.RFC3339),
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	feedFollowList := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		feedFollowList[i] = databaseFeedFollowToFeedFollow(feedFollow)
	}
	return feedFollowList
}

type Post struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	URL         string `json:"url"`
	FeedID      string `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID.String(),
		CreatedAt:   post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   post.UpdatedAt.Format(time.RFC3339),
		Title:       post.Title,
		Description: post.Description.String,
		PublishedAt: post.PublishedAt.Format(time.RFC3339),
		URL:         post.Url,
		FeedID:      post.FeedID.String(),
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	postList := make([]Post, len(posts))
	for i, post := range posts {
		postList[i] = databasePostToPost(post)
	}
	return postList
}
