package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/beevik/etree"
	"gopkg.in/vansante/go-ffprobe.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ResultXmlFolder string = "results"

type Creative struct {
	path         string
	duration     string
	format       string
	width        int
	heignt       int
	clickthrough string
	vastTree     etree.Document
	xmlPath      string
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

func NewCreative(path string, landingPage string) *Creative {
	videoData := getVideoData(path)

	durationInSeconds := int(videoData.Format.DurationSeconds)
	durationFormatted := fmt.Sprintf("%02d:%02d:%02d", durationInSeconds/3600, (durationInSeconds%3600)/60, durationInSeconds%60)
	videoFormatMap := map[string]string{".mp4": "video/mp4", ".mov": "video/mov", ".wmv": "video/wmv"}

	return &Creative{
		path:         path,
		duration:     durationFormatted,
		format:       videoFormatMap[filepath.Ext(path)],
		width:        videoData.FirstVideoStream().Width,
		heignt:       videoData.FirstVideoStream().Height,
		clickthrough: landingPage,
	}
}

func (c *Creative) saveVastToFile() {
	xmlFileName := strings.TrimSuffix(filepath.Base(c.path), filepath.Ext(filepath.Base(c.path)))
	xmlFileNamePath := fmt.Sprintf("%s/%s.xml", ResultXmlFolder, xmlFileName)
	c.vastTree.WriteToFile(xmlFileNamePath)

	c.xmlPath = xmlFileNamePath
}

func (c *Creative) saveVastToDB() {
	var db *gorm.DB
	var err error

	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("NAME")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to DB!", db)
	}
}

func processVideo(videoPath string, landingPage string) *Creative {
	creative := NewCreative(videoPath, landingPage)

	creative.generateVastTree()

	creative.saveVastToFile()
	// creative.saveVastToDB()

	return creative
}

func GenerateVastHttpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// render input page
		t, err := template.ParseFiles("web/index.html")
		if err != nil {
			log.Fatal(err)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		// parse multipart/form-data input
		r.ParseMultipartForm(32 << 20)

		// retrieve file from posted form-data
		landingPage := r.FormValue("landingPage")
		file, handler, err := r.FormFile("creativeFile")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		log.Printf("Uploaded file: %+v\n", handler.Filename)
		log.Printf("File size: %+v\n", handler.Size)
		log.Printf("MIME header: %+v\n", handler.Header)

		// write tmp file on the server
		tempFile, err := ioutil.TempFile("videos", fmt.Sprintf("*-%s", handler.Filename))
		if err != nil {
			log.Fatal(err)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
			return
		}
		tempFile.Write(fileBytes)
		defer tempFile.Close()

		// process the file
		c := processVideo(tempFile.Name(), landingPage)
		vast, err := c.vastTree.WriteToString()
		if err != nil {
			log.Fatal(err)
			return
		}

		// render result page
		t, err := template.ParseFiles("web/result.html")
		if err != nil {
			log.Fatal(err)
			return
		}
		t.Execute(w, vast)
	}
}

func setupRoutes() {
	http.HandleFunc("/", GenerateVastHttpHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRoutes()
}
