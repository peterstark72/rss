// Package rss provides functions to parse RSS 2.0 feeds.
package rss

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"time"
)

// MediaContent is the Yahoo Media RSS
type MediaContent struct {
	URL    string `xml:"url,attr" json:"URL,omitempty"`
	Width  string `xml:"width,attr" json:"Width,omitempty"`
	Height string `xml:"height,attr" json:"Height,omitempty"`
	Medium string `xml:"medium,attr" json:"Medium,omitempty"`
}

// Enclosure is an RSS 2.0 Enclosure
type Enclosure struct {
	URL  string `xml:"url,attr" json:"URL"`
	Type string `xml:"type,attr" json:"Type,omitempty"`
}

// Source is the RSS channel that the item came from.
type Source struct {
	URL  string `xml:"url,attr" json:"URL,omitempty"`
	Name string `xml:",chardata" json:"Name,omitempty"`
}

type PubDate struct {
	T time.Time
}

// CommonDateLayouts is an array of commonly used date formats
var CommonDateLayouts = []string{time.RFC1123, time.RFC1123Z, time.RFC3339}

// UnmarshalXML unmarshals RSS dates
func (p *PubDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	var err error
	for _, layout := range CommonDateLayouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			p.T = t
			return nil
		}
	}

	return err
}

// MarshalJSON marshals RSS Pubdate to JSON
func (p PubDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.T.Format("2006-01-02"))
}

// Item is an RSS 2.0 Item
type Item struct {
	PubDate     PubDate      `xml:"pubDate" json:"PubDate"`
	Description string       `xml:"description" json:"Description"`
	GUID        string       `xml:"guid" json:"GUID"`
	Link        string       `xml:"link" json:"Link"`
	Creator     string       `xml:"creator" json:"Creator"`
	Category    []string     `xml:"category" json:"Category"`
	Content     MediaContent `xml:"content" json:"Content"`
	Enclosure   Enclosure    `xml:"enclosure" json:"Enclosure"`
	Title       string       `xml:"title" json:"Title"`
	Source      Source       `xml:"source" json:"Source"`
}

// DateSorter sorts a list of items by date
type DateSorter []Item

func (s DateSorter) Len() int      { return len(s) }
func (s DateSorter) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s DateSorter) Less(i, j int) bool {
	return s[i].PubDate.T.After(s[j].PubDate.T)
}

// Channel is an RSS 2.0 Channel
type Channel struct {
	Items       []Item `xml:"item"`
	Description string `xml:"description"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Language    string `xml:"language"`
	Copyright   string `xml:"copyright"`
}

// Feed is the RSS 2.0 root
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"xmlns:media,attr"`
	Channel Channel  `xml:"channel"`
}

// NewFeed creates a new Feed from r
func NewFeed(r io.Reader) (f *Feed, err error) {
	f = &Feed{}
	data, err := io.ReadAll(r)
	if err != nil {
		return
	}
	xml.Unmarshal(data, &f)

	return
}
