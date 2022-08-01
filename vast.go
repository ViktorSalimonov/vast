package main

import (
	"os"

	"github.com/beevik/etree"
)

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

	duration := linear.CreateElement("Duration")
	duration.CreateText(creative_data.duration)

	mediafiles := linear.CreateElement("MediaFiles")
	mediafile := mediafiles.CreateElement("MediaFile")
	mediafile.CreateAttr("delivery", "progressive")
	mediafile.CreateAttr("type", creative_data.format)

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}
