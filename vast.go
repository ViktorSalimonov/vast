package main

import (
	"os"
	"strconv"

	"github.com/beevik/etree"
)

const SkipOffset = "00:00:05"

func generate_vast(creative Creative) {
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
	creative_elem := creatives.CreateElement("Creative")

	linear := creative_elem.CreateElement("Linear")
	linear.CreateAttr("skipoffset", SkipOffset)

	duration := linear.CreateElement("Duration")
	duration.CreateText(creative.duration)

	mediafiles := linear.CreateElement("MediaFiles")
	mediafile := mediafiles.CreateElement("MediaFile")
	mediafile.CreateAttr("delivery", "progressive")
	mediafile.CreateAttr("type", creative.format)
	mediafile.CreateAttr("width", strconv.Itoa(creative.width))
	mediafile.CreateAttr("height", strconv.Itoa(creative.heignt))
	mediafile.CreateCData("htttp://example.com/") // TODO: upload the file to s3 and use its s3 url here

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}
