package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

const sampleRSS = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>Test Feed</title>
    <link>https://example.com</link>
    <description>Test description</description>
    <language>en-us</language>
    <item>
      <title>First Post</title>
      <link>https://example.com/1</link>
      <description>hello</description>
      <pubDate>Mon, 02 Jan 2006 15:04:05 -0700</pubDate>
    </item>
    <item>
      <title>Second Post</title>
      <link>https://example.com/2</link>
      <description>world</description>
      <pubDate>Tue, 03 Jan 2006 15:04:05 -0700</pubDate>
    </item>
  </channel>
</rss>`

func TestURLToFeed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/rss+xml")
		w.Write([]byte(sampleRSS))
	}))
	defer srv.Close()

	feed, err := urlToFeed(context.Background(), srv.URL)
	if err != nil {
		t.Fatalf("urlToFeed returned error: %v", err)
	}
	if feed.Channel.Title != "Test Feed" {
		t.Errorf("title = %q, want %q", feed.Channel.Title, "Test Feed")
	}
	if len(feed.Channel.Items) != 2 {
		t.Fatalf("got %d items, want 2", len(feed.Channel.Items))
	}
	if feed.Channel.Items[0].Title != "First Post" {
		t.Errorf("first item title = %q, want %q", feed.Channel.Items[0].Title, "First Post")
	}
}

func TestURLToFeedNonOKStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer srv.Close()

	_, err := urlToFeed(context.Background(), srv.URL)
	if err == nil {
		t.Fatal("expected error for non-200 status")
	}
}
