// Package rss provides functionality to handle rss feeds.
package rss

import (
	"github.com/hill-daniel/alexa-rss-flashbriefing"
	"github.com/hill-daniel/alexa-rss-flashbriefing/html"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const codecentricFeedDateFormat = "Mon, 2 Jan 2006 15:04:05 -0700"

// CodecentricMapper maps the codecentric RSS feed to an AlexaFeed.
type CodecentricMapper struct {
}

// Map converts a given rss feed to an alexa feed.
func (CodecentricMapper) Map(atomFeed *feed.AtomFeed) (feed.AlexaFeed, error) {
	alexaFeed := feed.AlexaFeed{}
	for _, item := range atomFeed.Channel.Items {
		updateTime, err := time.Parse(codecentricFeedDateFormat, item.PubDate)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse date from item %v", item)
		}
		alexaItem := createAlexaItem(item, updateTime)
		alexaFeed = append(alexaFeed, alexaItem)
	}
	return alexaFeed, nil
}

func createAlexaItem(item feed.AtomItem, updateTime time.Time) feed.AlexaItem {
	feedItem := feed.AlexaItem{UUID: item.GUID,
		RedirectionURL: item.Link,
		TitleText:      html.RemoveTags(item.Title),
		MainText:       truncate(html.RemoveTags(item.Desc)),
		UpdateDate:     feed.UpdateDate(updateTime.UTC())}
	return feedItem
}

func truncate(value string) string {
	index := strings.Index(value, ".. Weiterlesen")
	if index > 0 {
		return strings.TrimSpace(value[0:index])
	}
	return value
}
