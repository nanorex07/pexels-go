package pexels

import (
	"context"
	"os"
	"testing"
)

func TestGetPhotos(t *testing.T) {
	// Initialize a new Pexels API client
	client := NewClient(os.Getenv("PEXELS_API_KEY"))

	// Set up the parameters for the GetPhotos function
	params := &GetPhotosParams{
		Query:       "nature",
		Orientation: "landscape",
		Size:        "medium",
		Color:       "green",
		Locale:      "en-US",
		Page:        1,
		PerPage:     10,
	}

	// Call the GetPhotos function
	resp, err := client.GetPhotos(context.Background(), params)
	if err != nil {
		t.Errorf("GetPhotos failed: %v", err)
	}

	// Check the response
	if resp == nil {
		t.Errorf("GetPhotos failed: response is nil")
	}
	if len(resp.Photos) == 0 {
		t.Errorf("GetPhotos failed: no photos returned")
	}
}

func TestGetCurated(t *testing.T) {
	// Initialize a new Pexels API client
	client := NewClient(os.Getenv("PEXELS_API_KEY"))

	// Set up the parameters for the GetCurated function
	params := &GetCuratedPhotoParams{
		Page:    1,
		PerPage: 10,
	}

	// Call the GetCurated function
	resp, err := client.GetCurated(context.Background(), params)
	if err != nil {
		t.Errorf("GetCurated failed: %v", err)
	}

	// Check the response
	if resp == nil {
		t.Errorf("GetCurated failed: response is nil")
	}
	if len(resp.Photos) == 0 {
		t.Errorf("GetCurated failed: no photos returned")
	}
}

func TestGetPhoto(t *testing.T) {
	// Initialize a new Pexels API client
	client := NewClient(os.Getenv("PEXELS_API_KEY"))

	// Set up the parameters for the GetPhoto function
	id := "2014422"

	// Call the GetPhoto function
	resp, err := client.GetPhoto(context.Background(), id)
	if err != nil {
		t.Errorf("GetPhoto failed: %v", err)
	}

	// Check the response
	if resp == nil {
		t.Errorf("GetPhoto failed: response is nil")
	}
}
