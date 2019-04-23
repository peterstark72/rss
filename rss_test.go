package rss_test

import (
	"strings"
	"testing"
	"time"

	"github.com/peterstark72/rss"
)

const (
	PolisenRSS     = "https://polisen.se/aktuellt/rss/skane/handelser-rss---skane/"
	SydsvenskanRSS = "https://tygelsjo-20030801.appspot.com/?q=Tygelsj%C3%B6"
	NYT            = "http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml"
)

func TestNewFeed(t *testing.T) {

	var r = `<rss version="2.0"><channel><item><pubDate>2019-03-04T17:41:04+01:00</pubDate><title>Test</title></item></channel></rss>`

	f, err := rss.NewFeed(strings.NewReader(r))
	if err != nil {
		t.Error("Could not read string", err)
	}

	d, _ := f.Channel.Items[0].ParsePubDate()
	if d.Format(time.RFC3339) != "2019-03-04T17:41:04+01:00" {
		t.Error("Wrong date")
	}

}
