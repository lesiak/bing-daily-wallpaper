package main

import (
	"../../pkg/bing"
	"../../pkg/wallpaper"
	"log"
)

func main() {
	img := bing.GetLastImage("pl-PL")
	log.Println("Downloading wallpaper")
	wallpaperPath := wallpaper.GetWallpaperPath(img.FileName())
	wallpaper.DownloadWallpaper(img.FullUrl(), wallpaperPath)
	log.Println("Wallpaper downloaded")
	log.Printf("Setting wallpaper to %s\n", wallpaperPath)
	wallpaper.SetWallpaper(wallpaperPath)
}
