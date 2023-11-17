#!/bin/sh

OS="$(uname)"
Region="pl-PL"
BingURL="https://www.bing.com"
#ImageAPIEndpoint is the API endpoint to get today's wallpaper
ImageAPIEndpoint="/HPImageArchive.aspx?format=js&idx=0&n=1&mkt="
MultipleImagesAPIEndpoint="/HPImageArchive.aspx?format=js&idx=0&n=100&mkt="
BingImageAPIUrl="${BingURL}${ImageAPIEndpoint}${Region}"

check_internet() {
  ping -c 1 8.8.8.8 > /dev/null 2>&1
}

while ! check_internet; do
  echo "Waiting for internet..."
  sleep 5
done

WALLPAPER_DATA=`curl -sSL "$BingImageAPIUrl" | jq '.images[0]'`


WALLPAPER_URL_TITLE=`echo $WALLPAPER_DATA | jq -r '.copyright' | sed 's/ (©.*//'`
WALLPAPER_URL_BASE=`echo $WALLPAPER_DATA | jq -r '.urlbase'`
WALLPAPER_START_DATE=`echo $WALLPAPER_DATA | jq -r '.startdate'`

Resolution="UHD"

FullUrl="${BingURL}${WALLPAPER_URL_BASE}_${Resolution}.jpg"
FileName="${WALLPAPER_START_DATE} ${WALLPAPER_URL_TITLE} ${Resolution}.jpg"

echo $FullUrl
echo $FileName

case "$OS" in
  *Darwin*)
    echo "Downloading wallpaper"
    curl -sSLo "$HOME/Pictures/$FileName" "$FullUrl"
    echo "Setting wallpaper"
    osascript -e "tell application \"System Events\" to tell every desktop to set picture to \"$HOME/Pictures/$FileName\" as POSIX file"
  #  killall Dock
    ;;
esac
