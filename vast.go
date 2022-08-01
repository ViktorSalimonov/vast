package main

import (
	"os"
	"strconv"

	"github.com/beevik/etree"
)

const Skipoffset = "00:00:05"

func generate_vast(creative_data Creative) {
	doc := etree.NewDocument()
	doc.CreateProcInst("VAST", `version="3.0"`)

	ad := doc.CreateElement("Ad")
	ad.CreateAttr("id", "AdId")

	inline := ad.CreateElement("InLine")

	adsystem := inline.CreateElement("AdSystem")
	adsystem.CreateText("Test AdSystem")

	adtitle := inline.CreateElement("AdTitle")
	adtitle.CreateText("Test video ad")

	impression := inline.CreateElement("Impression")
	impression.CreateCData("https://adserver.com/track/impression")

	creatives := inline.CreateElement("Creatives")
	creative := creatives.CreateElement("Creative")

	linear := creative.CreateElement("Linear")
	linear.CreateAttr("skipoffset", Skipoffset)

	duration := linear.CreateElement("Duration")
	duration.CreateText(creative_data.duration)

	mediafiles := linear.CreateElement("MediaFiles")
	mediafile := mediafiles.CreateElement("MediaFile")
	mediafile.CreateAttr("delivery", "progressive")
	mediafile.CreateAttr("type", creative_data.format)
	mediafile.CreateAttr("width", strconv.Itoa(creative_data.width))
	mediafile.CreateAttr("height", strconv.Itoa(creative_data.heignt))
	mediafile.CreateCData("htttp://example.com/") // TODO: upload the file to s3 and use its s3 url here

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}
