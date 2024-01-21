package pexels

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

// BaseURL is the base URL for the Pexels API.
var BaseURL = "https://api.pexels.com/"

// Version is the version of the Pexels API being used.
var Version = "v1"

// Client represents a client for the Pexels API.
type Client struct {
	BaseURL    string       // The base URL for the Pexels API
	ApiKey     string       // The API key for accessing the Pexels API
	HTTPClient *http.Client // The HTTP client for making requests
	Version    string       // The version of the Pexels API being used
}

// User represents a user in the Pexels API.
type User struct {
	ID   int    `json:"id"`   // Unique identifier for the user
	Name string `json:"name"` // Name of the user
	URL  string `json:"url"`  // URL to the user's profile
}

// NewClient creates a new Pexels API client.
// It takes an API key as input and returns a new Client instance.
func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURL,
		ApiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute * 2,
		},
		Version: Version,
	}
}

// sendRequest sends an HTTP request to the Pexels API.
// It takes a context, an HTTP request, and a variable to store the response data as input and returns an error.
func (c *Client) sendRequest(ctx context.Context, req *http.Request, vals interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Unknown API error: %d %s", res.StatusCode, string(bytes))
	}
	if err := json.NewDecoder(res.Body).Decode(&vals); err != nil {
		return err
	}
	return nil
}

// structToURLValues converts a struct to URL values for use in HTTP requests.
// It takes a struct as input and returns URL values representing the struct fields.
func (c *Client) structToURLValues(s interface{}) url.Values {
	val := url.Values{}
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		urlTag := t.Field(i).Tag.Get("url")
		fieldValue := fmt.Sprint(field.Interface())
		fieldKind := field.Kind()
		if urlTag != "" && ((fieldKind == reflect.Int && fieldValue != "0") || (fieldKind == reflect.String && fieldValue != "")) {
			val.Set(urlTag, fieldValue)
		}
	}
	return val
}
