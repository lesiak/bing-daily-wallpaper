#!/bin/bash

OS="$(uname)"
Region="pl-PL"
BingURL="https://www.bing.com"
DaysAgo=${1:-0}
#ImageAPIEndpoint is the API endpoint to get today's wallpaper
ImageAPIEndpoint="/HPImageArchive.aspx?format=js&idx=0&n=8&mkt="
BingImageAPIUrl="${BingURL}${ImageAPIEndpoint}${Region}"

if (( $DaysAgo > 15 )); then
    echo "Only 16 days supported"
    exit
fi

if (( $DaysAgo > 7 )); then
    DaysAgo=$((DaysAgo - 8))
    ImageAPIEndpoint="/HPImageArchive.aspx?format=js&idx=8&n=8&mkt="
    BingImageAPIUrl="${BingURL}${ImageAPIEndpoint}${Region}"
fi

check_internet() {
  ping -c 1 8.8.8.8 > /dev/null 2>&1
}

while ! check_internet; do
  echo "Waiting for internet..."
  sleep 5
done

WALLPAPER_DATA=`curl -sSL "$BingImageAPIUrl" | jq ".images[${DaysAgo}]"`


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
