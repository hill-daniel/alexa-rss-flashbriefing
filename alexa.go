package feed

import (
	"fmt"
	"strings"
	"time"
)

// UpdateDate is a custom time type for JSON un/marshalling
type UpdateDate time.Time

func (d UpdateDate) String() string {
	return time.Time(d).UTC().Format(time.RFC3339)
}

// UnmarshalJSON unmarshals a given byte slice to the custom UpdateDate type from the following format: 2006-01-02T15:04:05Z
func (d *UpdateDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), "\"")
	parse, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return err
	}
	*d = UpdateDate(parse)
	return nil
}

// MarshalJSON marshals the custom UpdateDate type to the following format: 2006-01-02T15:04:05Z
func (d UpdateDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(d).UTC().Format(time.RFC3339))
	return []byte(stamp), nil
}

// AlexaFeed is a custom type to provide sorting for a AlexaItem slice.
type AlexaFeed []AlexaItem

func (f AlexaFeed) Len() int {
	return len(f)
}

func (f AlexaFeed) Less(i, j int) bool {
	return time.Time(f[i].UpdateDate).After(time.Time(f[j].UpdateDate))
}

func (f AlexaFeed) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// AlexaItem represents an Alexa Flash Briefing response.
type AlexaItem struct {
	UUID           string     `json:"uid"`
	UpdateDate     UpdateDate `json:"updateDate"`
	TitleText      string     `json:"titleText"`
	MainText       string     `json:"mainText"`
	RedirectionURL string     `json:"redirectionUrl"`
}

// AlexaFeedReader reads an AlexaFeed.
type AlexaFeedReader interface {
	Read() (AlexaFeed, error)
}

// AlexaMapper maps an Atom RSS feed to Alexa feed.
type AlexaMapper interface {
	Map(atomFeed *AtomFeed) (AlexaFeed, error)
}
