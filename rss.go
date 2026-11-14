package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Items       []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(ctx context.Context, url string) (RssFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return RssFeed{}, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return RssFeed{}, err
	}
	// Check the HTTP response status code before proceeding
	if resp.StatusCode != http.StatusOK {
		return RssFeed{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	defer resp.Body.Close()
	decoder := xml.NewDecoder(resp.Body)
	feed := RssFeed{}
	err = decoder.Decode(&feed)
	return feed, err
}
