package feed

import "encoding/xml"

// AtomFeed is a struct representation of an Atom XML RSS feed.
type AtomFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

// Channel is the struct representation of XML element channel.
type Channel struct {
	Title    string     `xml:"title"`
	Link     string     `xml:"link"`
	Desc     string     `xml:"description"`
	Language string     `xml:"language"`
	Items    []AtomItem `xml:"item"`
}

// AtomItem is the struct representation of XML element item.
type AtomItem struct {
	GUID    string `xml:"guid"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Desc    string `xml:"description"`
	PubDate string `xml:"pubDate"`
}
