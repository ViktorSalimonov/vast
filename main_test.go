package main

import (
	"testing"
)

func TestGenerateVastTag(t *testing.T) {
	path := "./videos/sample.mp4"
	landingPage := "https://sample.com/"

	creative := NewCreative(path, landingPage)
	creative.generateVastTree()

	expectedVastString := `<?VAST version="3.0"?>
<Ad id="AdId">
  <InLine>
    <AdSystem>Test AdSystem</AdSystem>
    <AdTitle>Test video ad</AdTitle>
    <Impression><![CDATA[https://adserver.com/track/impression]]></Impression>
    <TrackingEvents>
      <Tracking type="start"><![CDATA[https://adserver.com/track/start]]></Tracking>
      <Tracking type="midpoint"><![CDATA[https://adserver.com/track/midpoint]]></Tracking>
      <Tracking type="complete"><![CDATA[https://adserver.com/track/complete]]></Tracking>
    </TrackingEvents>
    <VideoClicks>
      <ClickThrough><![CDATA[https://sample.com/]]></ClickThrough>
    </VideoClicks>
    <Creatives>
      <Creative>
        <Linear skipoffset="00:00:05">
          <Duration>00:00:34</Duration>
          <MediaFiles>
            <MediaFile delivery="progressive" type="video/mp4" width="464" height="848"><![CDATA[htttp://example.com/]]></MediaFile>
          </MediaFiles>
        </Linear>
      </Creative>
    </Creatives>
  </InLine>
</Ad>
`

	vastString, err := creative.vastTree.WriteToString()
	if err != nil {
		t.Errorf("Can't write to string %v", err)
	}

	if vastString != expectedVastString {
		t.Errorf("Got %s, expected %s", vastString, expectedVastString)
	}
}
