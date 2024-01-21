package pexels

import (
	"context"
	"os"
	"testing"
)

func TestGetPopularVideos(t *testing.T) {
	// Initialize a new Pexels API client
	client := NewClient(os.Getenv("PEXELS_API_KEY"))

	// Set up the parameters for the GetPopularVideos function
	params := &GetPopularVideosParams{
		Page:    1,
		PerPage: 10,
	}

	// Call the GetPopularVideos function
	resp, err := client.GetPopularVideos(context.Background(), params)
	if err != nil {
		t.Errorf("GetPopularVideos failed: %v", err)
	}

	// Check the response
	if resp == nil {
		t.Errorf("GetPopularVideos failed: response is nil")
	}
	if len(resp.Videos) == 0 {
		t.Errorf("GetPopularVideos failed: no videos returned")
	}
}

func TestGetVideos(t *testing.T) {
	// Initialize a new Pexels API client
	client := NewClient(os.Getenv("PEXELS_API_KEY"))

	// Set up the parameters for the GetVideos function
	params := &GetVideosParams{
		Query:       "nature",
		Orientation: "landscape",
		Size:        "medium",
		Locale:      "en-US",
		Page:        1,
		PerPage:     10,
	}

	// Call the GetVideos function
	resp, err := client.GetVideos(context.Background(), params)
	if err != nil {
		t.Errorf("GetVideos failed: %v", err)
	}

	// Check the response
	if resp == nil {
		t.Errorf("GetVideos failed: response is nil")
	}
	if len(resp.Videos) == 0 {
		t.Errorf("GetVideos failed: no videos returned")
	}
}

func TestGetVideo(t *testing.T) {
	// Initialize a new Pexels API client
	client := NewClient(os.Getenv("PEXELS_API_KEY"))

	// Set up the parameters for the GetVideo function
	id := "2499611" // Replace with a valid video ID

	// Call the GetVideo function
	resp, err := client.GetVideo(context.Background(), id)
	if err != nil {
		t.Errorf("GetVideo failed: %v", err)
	}

	// Check the response
	if resp == nil {
		t.Errorf("GetVideo failed: response is nil")
	}
}
