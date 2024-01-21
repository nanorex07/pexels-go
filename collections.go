package pexels

import (
	"context"
	"fmt"
	"net/http"
)

// Collection represents a collection in the Pexels API.
type Collection struct {
	ID          string `json:"id"`           // Unique identifier for the collection
	Title       string `json:"title"`        // Title of the collection
	Description string `json:"description"`  // Description of the collection
	Private     bool   `json:"private"`      // Indicates if the collection is private
	MediaCount  int    `json:"media_count"`  // Number of media in the collection
	PhotosCount int    `json:"photos_count"` // Number of photos in the collection
	VideosCount int    `json:"videos_count"` // Number of videos in the collection
}

// GetCollectionsResponse represents the response from the GetCollections function.
type GetCollectionsResponse struct {
	Collections  []Collection `json:"collections"`   // List of collections
	Page         int          `json:"page"`          // Current page number
	PerPage      int          `json:"per_page"`      // Number of results per page
	TotalResults int          `json:"total_results"` // Total number of results for the query
	NextPage     string       `json:"next_page"`     // URL to the next page of results
	PrevPage     string       `json:"prev_page"`     // URL to the previous page of results
}

// GetFeaturedCollectionParams represents the parameters for the GetFeaturedCollection function.
type GetFeaturedCollectionParams struct {
	Page    int `url:"page"`     // Page number for paginated results
	PerPage int `url:"per_page"` // Number of results per page
}

// GetCollectionMediaParams represents the parameters for the GetCollectionMedia function.
type GetCollectionMediaParams struct {
	Type    string `url:"type"`     // Type of media to retrieve (e.g., photos, videos)
	Sort    string `url:"sort"`     // Sorting order of the media (e.g., popular, latest)
	Page    int    `url:"page"`     // Page number for paginated results
	PerPage int    `url:"per_page"` // Number of results per page
}

// CollectionMedia represents the media in a collection in the Pexels API.
type CollectionMedia struct {
	Type            string         `json:"type"`             // Type of the media
	ID              int            `json:"id"`               // Unique identifier for the media
	Width           int            `json:"width"`            // Width of the media in pixels
	Height          int            `json:"height"`           // Height of the media in pixels
	URL             string         `json:"url"`              // URL to the media
	Photographer    string         `json:"photographer"`     // Name of the photographer
	PhotographerURL string         `json:"photographer_url"` // URL to the photographer's profile
	PhotographerID  int            `json:"photographer_id"`  // Unique identifier for the photographer
	AvgColor        string         `json:"avg_color"`        // Average color of the media in hexadecimal format
	Src             PhotoSrc       `json:"src"`              // Object containing URLs to different sizes of the media
	Liked           bool           `json:"liked"`            // Indicates if the media is liked
	Duration        int            `json:"duration"`         // Duration of the video in seconds
	FullRes         any            `json:"full_res"`         // Full resolution of the video
	Tags            []any          `json:"tags"`             // Tags of the media
	Image           string         `json:"image"`            // URL to the video's image
	User            User           `json:"user"`             // User who uploaded the media
	VideoFiles      VideoFile      `json:"video_files"`      // Files of the video
	VideoPictures   []VideoPicture `json:"video_pictures"`   // Pictures of the video
}

// GetCollectionMedia represents the response from the GetCollectionMedia function.
type GetCollectionMedia struct {
	ID           string            `json:"id"`            // Unique identifier for the collection
	Media        []CollectionMedia `json:"media"`         // List of media in the collection
	Page         int               `json:"page"`          // Current page number
	PerPage      int               `json:"per_page"`      // Number of results per page
	TotalResults int               `json:"total_results"` // Total number of results for the query
	NextPage     string            `json:"next_page"`     // URL to the next page of results
	PrevPage     string            `json:"prev_page"`     // URL to the previous page of results
}

func (c *Client) getCollections(ctx context.Context, params *GetFeaturedCollectionParams, own bool) (*GetCollectionsResponse, error) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 5
	}
	url := fmt.Sprintf("%s%s/collections/featured?%s", c.BaseURL, c.Version, c.structToURLValues(*params).Encode())
	if own {
		url = fmt.Sprintf("%s%s/collections?%s", c.BaseURL, c.Version, c.structToURLValues(*params).Encode())
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp GetCollectionsResponse = GetCollectionsResponse{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetCollection retrieves a collection from the Pexels API.
// It takes a context, GetCollectionMediaParams, and an ID as input and returns a CollectionMedia and an error.
// The GetCollectionMediaParams specify the type, sort, page, and per page parameters.
// The ID is the unique identifier for the collection.
// The CollectionMedia contains the type, ID, width, height, URL, photographer, photographer URL, photographer ID, average color, source, liked status, duration, full resolution, tags, image URL, user, video files, and video pictures of the media in the collection.
func (c *Client) GetCollection(ctx context.Context, params *GetCollectionMediaParams, id string) (*CollectionMedia, error) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 5
	}
	url := fmt.Sprintf("%s%s/collections/%s?%s", c.BaseURL, c.Version, id, c.structToURLValues(*params).Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp CollectionMedia = CollectionMedia{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFeaturedCollections retrieves a list of featured collections from the Pexels API.
// It takes a context and GetFeaturedCollectionParams as input and returns a GetCollectionsResponse and an error.
// The GetFeaturedCollectionParams specify the page and per page parameters.
// The GetCollectionsResponse contains the current page number, the number of results per page, the total number of results, a URL to the collection, and a list of collections matching the query.
func (c *Client) GetFeaturedCollections(ctx context.Context, params *GetFeaturedCollectionParams) (*GetCollectionsResponse, error) {
	return c.getCollections(ctx, params, false)
}

// GetUserCollections retrieves a list of user's collections from the Pexels API.
// It takes a context and GetFeaturedCollectionParams as input and returns a GetCollectionsResponse and an error.
// The GetFeaturedCollectionParams specify the page and per page parameters.
// The GetCollectionsResponse contains the current page number, the number of results per page, the total number of results, a URL to the collection, and a list of collections matching the query.
func (c *Client) GetUserCollections(ctx context.Context, params *GetFeaturedCollectionParams) (*GetCollectionsResponse, error) {
	return c.getCollections(ctx, params, true)
}
