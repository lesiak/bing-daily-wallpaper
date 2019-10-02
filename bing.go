package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	// BingURL is the base url for bing
	BingURL = "http://www.bing.com"
	// ImageAPIEndpoint is the API endpoint to get today's wallpaper
	ImageAPIEndpoint = "/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-IN"
)

// ImageResponse is the response obtained from ImageApiEndpoint
type ImageResponse struct {
	Images []Image `json:"images"`
}

// Image is the properties of the image in the ImageResponse
type Image struct {
	URL       string `json:"url"`
	CopyRight string `json:"copyright"`
	StartDate int    `json:"startdate,string"`
}

// GetWallpaperURL returns the url for today's wallpaper from bing
func GetWallpaperURL() string {
	res, err := http.Get(BingURL + ImageAPIEndpoint)
	if err != nil {
		log.Fatalf("Failed to get response.\nError is: %v\n", err)
	}
	defer res.Body.Close()
	// Decode json
	decoder := json.NewDecoder(res.Body)
	var imgResponse ImageResponse
	err = decoder.Decode(&imgResponse)
	if err != nil {
		log.Fatalf("Failed to decode json.\nError is: %v\n", err)
	}
	return imgResponse.Images[0].URL
}