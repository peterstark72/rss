//Package rss provides functions to parse RSS 2.0 feeds.
package rss

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"time"
)

//MediaContent is the Yahoo Media RSS
type MediaContent struct {
	URL    string `xml:"url,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Medium string `xml:"medium,attr"`
}

//Enclosure is an RSS 2.0 Enclosure
type Enclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

//Source is the RSS channel that the item came from.
type Source struct {
	URL  string `xml:"url,attr"`
	Name string `xml:",chardata"`
}

//Item is an RSS 2.0 Item
type Item struct {
	PubDate     string       `xml:"pubDate"`
	Description string       `xml:"description"`
	GUID        string       `xml:"guid"`
	Link        string       `xml:"link"`
	Creator     string       `xml:"creator"`
	Category    []string     `xml:"category"`
	Content     MediaContent `xml:"content"`
	Enclosure   Enclosure    `xml:"enclosure"`
	Title       string       `xml:"title"`
	Source      Source       `xml:"source"`
}

//CommonDateLayouts is an array of commonly used date formats
var CommonDateLayouts = []string{time.RFC1123, time.RFC1123Z, time.RFC3339}

//ParsePubDate attempts to parse PubDate using CommonDateLayouts
func (itm *Item) ParsePubDate() (t time.Time, err error) {
	for _, layout := range CommonDateLayouts {
		t, err := time.Parse(layout, itm.PubDate)
		if err == nil {
			return t, nil
		}
	}
	return
}

//Channel is an RSS 2.0 Channel
type Channel struct {
	Items       []Item `xml:"item"`
	Description string `xml:"description"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Language    string `xml:"language"`
	Copyright   string `xml:"copyright"`
}

//Feed is the RSS 2.0 root
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"xmlns:media,attr"`
	Channel Channel  `xml:"channel"`
}

//NewFeed creates a new Feed from r
func NewFeed(r io.Reader) (f *Feed, err error) {
	f = &Feed{}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	xml.Unmarshal(data, &f)

	return
}
