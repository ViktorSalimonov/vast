package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

const ResultXmlFolder string = "results"

type Creative struct {
	path         string
	duration     string
	format       string
	width        int
	heignt       int
	clickthrough string
}

func getVideoData(path string) *ffprobe.ProbeData {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fileReader, err := os.Open(path)
	if err != nil {
		log.Panicf("Error opening the file: %v", err)
	}

	data, err := ffprobe.ProbeReader(ctx, fileReader)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}
	return data
}

func newCreative(path string) *Creative {
	videoData := getVideoData(path)

	width, height := videoData.FirstVideoStream().Width, videoData.FirstVideoStream().Height
	durationInSeconds := int(videoData.Format.DurationSeconds)
	durationFormatted := fmt.Sprintf("%02d:%02d:%02d", durationInSeconds/3600, (durationInSeconds%3600)/60, durationInSeconds%60)
	videoFormatMap := map[string]string{".mp4": "video/mp4", ".avi": "video/api"}

	creative := Creative{path: path}
	creative.width, creative.heignt = width, height
	creative.duration = durationFormatted
	creative.format = videoFormatMap[filepath.Ext(path)]
	return &creative
}

func main() {
	var videoPath string
	fmt.Print("Enter the creative file path: ")
	fmt.Scanln(&videoPath)

	var landingPage string
	fmt.Print("Enter the landing page: ")
	fmt.Scanln(&landingPage)

	creative := newCreative(videoPath)
	creative.clickthrough = landingPage

	vastTree := generateVastTree(*creative)

	// Save the VAST tag to the file .xml file
	xmlFileName := strings.TrimSuffix(filepath.Base(videoPath), filepath.Ext(filepath.Base(videoPath)))
	xmlFileNamePath := fmt.Sprintf("%s/%s.xml", ResultXmlFolder, xmlFileName)
	vastTree.WriteToFile(xmlFileNamePath)

	// TODO: Save the vast tag in the sql DB
}
