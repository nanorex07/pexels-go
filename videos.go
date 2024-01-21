package pexels

import (
	"context"
	"fmt"
	"net/http"
)

// VideoFile represents a file of a video in the Pexels API.
type VideoFile struct {
	ID       int     `json:"id"`        // Unique identifier for the file
	Quality  string  `json:"quality"`   // Quality of the file
	FileType string  `json:"file_type"` // Type of the file
	Width    int     `json:"width"`     // Width of the video in pixels
	Height   int     `json:"height"`    // Height of the video in pixels
	Fps      float64 `json:"fps"`       // Frames per second of the video
	Link     string  `json:"link"`      // URL to the video file
}

// VideoPicture represents a picture of a video in the Pexels API.
type VideoPicture struct {
	ID      int    `json:"id"`      // Unique identifier for the picture
	Picture string `json:"picture"` // URL to the picture
	Nr      int    `json:"nr"`      // Number of the picture
}

// Video represents a video in the Pexels API.
type Video struct {
	ID            int            `json:"id"`             // Unique identifier for the video
	Width         int            `json:"width"`          // Width of the video in pixels
	Height        int            `json:"height"`         // Height of the video in pixels
	URL           string         `json:"url"`            // URL to the video
	Image         string         `json:"image"`          // URL to the video's image
	FullRes       any            `json:"full_res"`       // Full resolution of the video
	Tags          []any          `json:"tags"`           // Tags of the video
	Duration      int            `json:"duration"`       // Duration of the video in seconds
	User          User           `json:"user"`           // User who uploaded the video
	VideoFiles    []VideoFile    `json:"video_files"`    // Files of the video
	VideoPictures []VideoPicture `json:"video_pictures"` // Pictures of the video
}

// GetVideosResponse represents the response from the GetVideos function.
type GetVideosResponse struct {
	Page         int     `json:"page"`          // Current page number
	PerPage      int     `json:"per_page"`      // Number of results per page
	TotalResults int     `json:"total_results"` // Total number of results for the query
	URL          string  `json:"url"`           // URL to the video
	Videos       []Video `json:"videos"`        // List of videos matching the query
}

// GetVideosParams represents the parameters for the GetVideos function.
type GetVideosParams struct {
	Query       string `url:"query"`       // Search query for videos
	Orientation string `url:"orientation"` // Desired orientation of videos (e.g., landscape, portrait)
	Size        string `url:"size"`        // Desired size of videos (e.g., small, medium, large)
	Locale      string `url:"locale"`      // Locale for the search query
	Page        int    `url:"page"`        // Page number for paginated results
	PerPage     int    `url:"per_page"`    // Number of results per page
}

// GetPopularVideosParams represents the parameters for the GetPopularVideos function.
type GetPopularVideosParams struct {
	MinWidth    int `url:"min_width"`    // Minimum width of the videos
	MinHeight   int `url:"min_height"`   // Minimum height of the videos
	MinDuration int `url:"min_duration"` // Minimum duration of the videos
	MaxDuration int `url:"max_duration"` // Maximum duration of the videos
	Page        int `url:"page"`         // Page number for paginated results
	PerPage     int `url:"per_page"`     // Number of results per page
}

// GetVideo retrieves a video from the Pexels API.
// It takes a context and an ID as input and returns a Video and an error.
// The ID is the unique identifier for the video.
// The Video contains the ID, width, height, URL, image URL, full resolution, tags, duration, user, video files, and video pictures of the video.
func (c *Client) GetVideo(ctx context.Context, id string) (*Video, error) {
	url := fmt.Sprintf("%s/videos/videos/%s", c.BaseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp Video = Video{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPopularVideos retrieves a list of popular videos from the Pexels API.
// It takes a context and GetPopularVideosParams as input and returns a GetVideosResponse and an error.
// The GetPopularVideosParams specify the minimum width, minimum height, minimum duration, maximum duration, page, and per page parameters.
// The GetVideosResponse contains the current page number, the number of results per page, the total number of results, a URL to the video, and a list of videos matching the query.
func (c *Client) GetPopularVideos(ctx context.Context, params *GetPopularVideosParams) (*GetVideosResponse, error) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 2
	}
	url := fmt.Sprintf("%svideos/popular?%s", c.BaseURL, c.structToURLValues(*params).Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp GetVideosResponse = GetVideosResponse{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVideos retrieves a list of videos from the Pexels API.
// It takes a context and GetVideosParams as input and returns a GetVideosResponse and an error.
// The GetVideosParams specify the search query, orientation, size, locale, page, and per page parameters.
// The GetVideosResponse contains the current page number, the number of results per page, the total number of results, a URL to the video, and a list of videos matching the query.
func (c *Client) GetVideos(ctx context.Context, params *GetVideosParams) (*GetVideosResponse, error) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 5
	}
	if params.Query == "" {
		return nil, fmt.Errorf("Query field cannot be empty.")
	}
	url := fmt.Sprintf("%s/videos/search?%s", c.BaseURL, c.structToURLValues(*params).Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp GetVideosResponse = GetVideosResponse{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
