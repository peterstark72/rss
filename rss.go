package rss

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

//MediaContent is the Yahoo Media RSS
type MediaContent struct {
	URL string `xml:"url,attr"`
}

//Enclosure is an RSS 2.0 Enclosure
type Enclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

//Item is an RSS 2.0 Item
type Item struct {
	PubDate     string       `xml:"pubDate"`
	Description string       `xml:"description"`
	GUID        string       `xml:"guid"`
	Link        string       `xml:"link"`
	Creator     string       `xml:"creator"`
	Category    []string     `xml:"category"`
	Content     MediaContent `xml:"media:content"`
	Enclosure   Enclosure    `xml:"enclosure"`
	Title       string       `xml:"title"`
}

//Channel is an RSS 2.0 Channel
type Channel struct {
	Items       []Item `xml:"item"`
	Description string `xml:"description"`
	Title       string `xml:"title"`
}

//Feed is the RSS 2.0 root
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"xmlns:media,attr"`
	Channel Channel  `xml:"channel"`
}

//DateLayout is the date format often used by RSS
const DateLayout = "Mon, 2 Jan 2006 15:04:05 -0700"

//Version is the supported RSS version
const Version = "2.0"

//YahooMediaNamespace is the Yahoo media namespace URI
const YahooMediaNamespace = "http://search.yahoo.com/mrss/"

//ReadAll returns as channel with items
func ReadAll(r io.Reader) (fd Feed) {
	data, _ := ioutil.ReadAll(r)
	xml.Unmarshal(data, &fd)

	return
}

//NewFeed created an empty RSS Feed
func NewFeed(title string) (fd Feed) {
	fd.Channel.Title = title
	fd.Version = Version
	fd.Media = YahooMediaNamespace

	return
}
