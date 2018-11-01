// Package rss provides functionality to handle rss feeds.
package rss

import (
	"encoding/xml"
	"github.com/hill-daniel/alexa-rss-flashbriefing"
	"github.com/hill-daniel/alexa-rss-flashbriefing/http"
	"github.com/pkg/errors"
	"io"
)

// AtomFeedReader reads a RSS feed and returns an AlexaFeed.
type AtomFeedReader struct {
	client http.Getter
	mapper feed.AlexaMapper
	url    string
}

// NewAtomFeedReader creates a AtomFeedReader.
func NewAtomFeedReader(client http.Getter, mapper feed.AlexaMapper, url string) *AtomFeedReader {
	return &AtomFeedReader{
		client: client,
		mapper: mapper,
		url:    url,
	}
}

func (r AtomFeedReader) Read() (feed.AlexaFeed, error) {
	response, err := r.client.Get(r.url)
	if err != nil {
		return nil, err
	}
	defer safeClose(response.Body, &err)
	decoder := xml.NewDecoder(response.Body)
	atomFeed := &feed.AtomFeed{}
	err = decoder.Decode(atomFeed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode xml")
	}

	alexaFeed, err := r.mapper.Map(atomFeed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to map atom feed to alexa feed")
	}
	return alexaFeed, err
}

func safeClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && *err == nil {
		*err = cerr
	}
}
