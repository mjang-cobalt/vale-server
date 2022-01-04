package atom

import (
	"testing"
)

var feed = "https://github.com/errata-ai/Joblint/releases.atom"

func TestConvert(t *testing.T) {
	parsed, err := Parse(feed)
	if err != nil {
		t.Error(err)
	}

	dt1, err := ToTime(parsed.Updated)
	if err != nil {
		t.Error(err)
	}

	dt2, err := ToTime("2021-03-07T12:21:29-07:00")
	if err != nil {
		t.Error(err)
	}

	if !dt2.UTC().Before(dt1.UTC()) {
		t.Error("bad date")
	}
}
