package pexels

import (
	"context"
	"fmt"
	"net/http"
)

// PhotoSrc represents the different sizes of a photo.
// Each field is a URL to the corresponding size of the photo.
type PhotoSrc struct {
	Original  string `json:"original"`  // URL to the original size photo
	Large2X   string `json:"large2x"`   // URL to the large 2x size photo
	Large     string `json:"large"`     // URL to the large size photo
	Medium    string `json:"medium"`    // URL to the medium size photo
	Small     string `json:"small"`     // URL to the small size photo
	Portrait  string `json:"portrait"`  // URL to the portrait size photoß
	Landscape string `json:"landscape"` // URL to the landscape size photo
	Tiny      string `json:"tiny"`      // URL to the tiny size photo
}

// Photo represents a photo from the Pexels API.
type Photo struct {
	ID              int      `json:"id"`               // Unique identifier for the photo
	Width           int      `json:"width"`            // Width of the photo in pixels
	Height          int      `json:"height"`           // Height of the photo in pixels
	URL             string   `json:"url"`              // URL to the photo
	Photographer    string   `json:"photographer"`     // Name of the photographer
	PhotographerURL string   `json:"photographer_url"` // URL to the photographer's profile
	PhotographerID  int      `json:"photographer_id"`  // Unique identifier for the photographer
	AvgColor        string   `json:"avg_color"`        // Average color of the photo in hexadecimal format
	Src             PhotoSrc `json:"src"`              // Object containing URLs to different sizes of the photo
	Liked           bool     `json:"liked"`            // Indicates if the photo is liked
	Alt             string   `json:"alt"`              // Alternative description of the photo
}

// GetPhotosParams represents the parameters for the GetPhotos function.
type GetPhotosParams struct {
	Query       string `url:"query"`       // Search query for photos
	Orientation string `url:"orientation"` // Desired orientation of photos (e.g., landscape, portrait)
	Size        string `url:"size"`        // Desired size of photos (e.g., small, medium, large)
	Color       string `url:"color"`       // Desired color of photos (e.g., red, blue, green)
	Locale      string `url:"locale"`      // Locale for the search query
	Page        int    `url:"page"`        // Page number for paginated results
	PerPage     int    `url:"per_page"`    // Number of results per page
}

// GetCuratedPhotoParams represents the parameters for the GetCurated function.
type GetCuratedPhotoParams struct {
	Page    int `url:"page"`     // Page number for paginated results
	PerPage int `url:"per_page"` // Number of results per page
}

// GetPhotoResponse represents the response from the GetPhotos function.
type GetPhotoResponse struct {
	TotalResults int     `json:"total_results"` // Total number of results for the query
	Page         int     `json:"page"`          // Current page number
	PerPage      int     `json:"per_page"`      // Number of results per page
	Photos       []Photo `json:"photos"`        // List of photos matching the query
	NextPage     string  `json:"next_page"`     // URL to the next page of results
	PrevPage     string  `json:"prev_page"`     // URL to the previous page of results
}

// GetPhotos retrieves a list of photos from the Pexels API.
// It takes a context and GetPhotosParams as input and returns a GetPhotoResponse and an error.
// The GetPhotosParams specify the search query, orientation, size, color, locale, page, and per page parameters.
// The GetPhotoResponse contains the total number of results, the current page number, the number of results per page, a list of photos matching the query, and URLs to the next and previous pages of results.
func (c *Client) GetPhotos(ctx context.Context, params *GetPhotosParams) (*GetPhotoResponse, error) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 5
	}
	if params.Query == "" {
		return nil, fmt.Errorf("Query field cannot be empty.")
	}
	url := fmt.Sprintf("%s%s/search?%s", c.BaseURL, c.Version, c.structToURLValues(*params).Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp GetPhotoResponse = GetPhotoResponse{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCurated retrieves a list of curated photos from the Pexels API.
// It takes a context and GetCuratedPhotoParams as input and returns a GetPhotoResponse and an error.
// The GetCuratedPhotoParams specify the page and per page parameters.
// The GetPhotoResponse contains the total number of results, the current page number, the number of results per page, a list of photos matching the query, and URLs to the next and previous pages of results.
func (c *Client) GetCurated(ctx context.Context, params *GetCuratedPhotoParams) (*GetPhotoResponse, error) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 5
	}
	url := fmt.Sprintf("%s%s/curated?%s", c.BaseURL, c.Version, c.structToURLValues(*params).Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp GetPhotoResponse = GetPhotoResponse{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPhoto retrieves a photo from the Pexels API.
// It takes a context and an ID as input and returns a Photo and an error.
// The ID is the unique identifier for the photo.
// The Photo contains the ID, width, height, URL, photographer, photographer URL, photographer ID, average color, source, liked status, and alternative description of the photo.
func (c *Client) GetPhoto(ctx context.Context, id string) (*Photo, error) {
	url := fmt.Sprintf("%s%s/photos/%s", c.BaseURL, c.Version, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	var resp Photo = Photo{}
	err = c.sendRequest(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
