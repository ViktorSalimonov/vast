package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

func get_video_duration(filename string) int {
	// return video duration in seconds
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fileReader, err := os.Open(filename)
	if err != nil {
		log.Panicf("Error opening test file: %v", err)
	}

	data, err := ffprobe.ProbeReader(ctx, fileReader)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}
	return int(data.Format.Duration().Seconds())
}

func main() {
	FILENAME := "./videos/video_1.mp4"

	duration := get_video_duration(FILENAME)
	fmt.Println(duration)
}
