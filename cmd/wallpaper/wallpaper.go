package main

import (
	"../../pkg/bing"
	"../../pkg/wallpaper"
	"log"
)

const (
	ResolutionHorizontal = "1920x1080"
	ResolutionVertical = "1080x1920"
)


func main() {
	img := bing.GetLastImage("pl-PL")
	log.Println("Downloading wallpaper")

	downloadPathHorizontal := wallpaper.GetWallpaperPath(img.FileName(ResolutionHorizontal))
	downloadPathVertical := wallpaper.GetWallpaperPath(img.FileName(ResolutionVertical))
	wallpaper.DownloadWallpaper(img.FullUrl(ResolutionHorizontal), downloadPathHorizontal)
	wallpaper.DownloadWallpaper(img.FullUrl(ResolutionVertical), downloadPathVertical)
	log.Println("Wallpaper downloaded")
	log.Printf("Setting wallpaper to %s\n", downloadPathHorizontal)
	wallpaper.SetWallpaper(downloadPathHorizontal)
}
