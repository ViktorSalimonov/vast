package main

import (
	"strconv"

	"github.com/beevik/etree"
)

const SkipOffset = "00:00:05"

func (c *Creative) generateVastTree() {
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

	trackingEvents := inline.CreateElement("TrackingEvents")
	trackingEventStart := trackingEvents.CreateElement("Tracking")
	trackingEventStart.CreateAttr("type", "start")
	trackingEventStart.CreateCData("https://adserver.com/track/start")
	trackingEventMid := trackingEvents.CreateElement("Tracking")
	trackingEventMid.CreateAttr("type", "midpoint")
	trackingEventMid.CreateCData("https://adserver.com/track/midpoint")
	trackingEventComplete := trackingEvents.CreateElement("Tracking")
	trackingEventComplete.CreateAttr("type", "complete")
	trackingEventComplete.CreateCData("https://adserver.com/track/complete")

	videoClicks := inline.CreateElement("VideoClicks")
	clickThrough := videoClicks.CreateElement("ClickThrough")
	clickThrough.CreateCData(c.clickthrough)

	creatives := inline.CreateElement("Creatives")
	creativeElem := creatives.CreateElement("Creative")

	linear := creativeElem.CreateElement("Linear")
	linear.CreateAttr("skipoffset", SkipOffset)

	duration := linear.CreateElement("Duration")
	duration.CreateText(c.duration)

	mediafiles := linear.CreateElement("MediaFiles")
	mediafile := mediafiles.CreateElement("MediaFile")
	mediafile.CreateAttr("delivery", "progressive")
	mediafile.CreateAttr("type", c.format)
	mediafile.CreateAttr("width", strconv.Itoa(c.width))
	mediafile.CreateAttr("height", strconv.Itoa(c.heignt))
	mediafile.CreateCData("htttp://example.com/") // TODO: upload the file to s3 and use its s3 url here

	doc.Indent(2)

	c.vastTree = *doc
}
