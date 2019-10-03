package main

import (
	"../../pkg/bing"
	"../../pkg/wallpaper"
)

const (
	ResolutionHorizontal = "1920x1080"
	ResolutionVertical = "1080x1920"
)


func main() {
	images := bing.GetRecentImages("pl-PL")
	for _, img := range images {
		downloadPathHorizontal := wallpaper.GetWallpaperPath(img.FileName(ResolutionHorizontal))
		downloadPathVertical := wallpaper.GetWallpaperPath(img.FileName(ResolutionVertical))
		wallpaper.DownloadWallpaper(img.FullUrl(ResolutionHorizontal), downloadPathHorizontal)
		wallpaper.DownloadWallpaper(img.FullUrl(ResolutionVertical), downloadPathVertical)
	}
}
