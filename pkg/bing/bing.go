package bing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	// BingURL is the base url for bing
	BingURL = "http://www.bing.com"
	// ImageAPIEndpoint is the API endpoint to get today's wallpaper
	ImageAPIEndpoint = "/HPImageArchive.aspx?format=js&idx=0&n=1&mkt="
	MultipleImagesAPIEndpoint = "/HPImageArchive.aspx?format=js&idx=0&n=100&mkt="
)

// ImageResponse is the response obtained from ImageApiEndpoint
type ImageResponse struct {
	Images []Image `json:"images"`
}

// Image is the properties of the image in the ImageResponse
type Image struct {
	Url       string `json:"url"`
	UrlBase   string `json:"urlbase"`
	CopyRight string `json:"copyright"`
	StartDate int    `json:"startdate,string"`
}

func (img Image) FullUrl(resolution string) string {
	return fmt.Sprintf("%s%s_%s.jpg", BingURL, img.UrlBase, resolution)
}

func (img Image) FileName(resolution string) string {
	title := strings.Split(img.CopyRight, " (©")[0]
	return fmt.Sprintf("%d %s %s.jpg", img.StartDate, title, resolution)
}

func GetLastImage(region string) Image {
	res, err := http.Get(BingURL + ImageAPIEndpoint + region)
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
	return imgResponse.Images[0]
}

func GetRecentImages(region string) []Image {
	res, err := http.Get(BingURL + MultipleImagesAPIEndpoint + region)
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
	return imgResponse.Images
}