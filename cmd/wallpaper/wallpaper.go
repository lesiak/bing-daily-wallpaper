package main

import (
	"../../pkg/bing"
	"../../pkg/wallpaper"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	img := bing.GetLastImage("pl-PL")
	log.Println("Downloading wallpaper")
	wallpaperPath := wallpaper.GetWallpaperPath(img.FileName())
	DownloadWallpaper(img.FullUrl(), wallpaperPath)
	log.Println("Wallpaper downloaded")
	log.Printf("Setting wallpaper to %s\n", wallpaperPath)
	wallpaper.SetWallpaper(wallpaperPath)
}

// DownloadWallpaper downloads the wallpaper from the provided url
// It stores the wallpaper in the path provided
func DownloadWallpaper(url string, path string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to download image.\nError is: %v\n", err)
	}
	wallpaper, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to create file.\nError is: %v\n", err)
	}
	defer res.Body.Close()
	defer wallpaper.Close()
	io.Copy(wallpaper, res.Body)
}
