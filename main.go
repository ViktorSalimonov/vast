package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

type Creative struct {
	path     string
	duration string
	format   string
}

// Returns duration string in the VAST tag specific format "hours:minutes:seconds"
func getVideoDuration(path string) string {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fileReader, err := os.Open(path)
	if err != nil {
		log.Panicf("Error opening test file: %v", err)
	}

	data, err := ffprobe.ProbeReader(ctx, fileReader)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}
	durationInSeconds := int(data.Format.Duration().Seconds())

	hours := durationInSeconds / 3600
	minutes := (durationInSeconds % 3600) / 60
	seconds := durationInSeconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// Returns video format string in the VAST tag specific format eg. "video/mp4"
func getVideoFormat(path string) string {
	ext := filepath.Ext(path)
	if ext == ".mp4" {
		return "video/mp4"
	} else if ext == ".mkv" {
		return "video/mkv"
	} else {
		return ""
	}
}

func main() {
	FILENAME := "./videos/video_1.mp4"

	creative := Creative{}
	creative.path = FILENAME
	creative.duration = getVideoDuration(creative.path)
	creative.format = getVideoFormat(creative.path)
	fmt.Println(creative)

	generate_vast(creative)
}
