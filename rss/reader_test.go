package rss_test

import (
	"github.com/hill-daniel/alexa-rss-flashbriefing"
	"github.com/hill-daniel/alexa-rss-flashbriefing/rss"
	assert "github.com/matryer/is"
	"net/http"
	"os"
	"testing"
	"time"
)

func Test_should_read_rss_feed_and_return_alexa_feed(t *testing.T) {
	is := assert.New(t)
	client := &testClient{}
	mapper := rss.CodecentricMapper{}
	feedReader := rss.NewAtomFeedReader(client, mapper, "does not matter")

	alexaFeed, err := feedReader.Read()

	is.NoErr(err)
	expectedResponse := feed.AlexaItem{UUID: "https://blog.codecentric.de/?p=56032",
		TitleText:      "Application Lifecycle Intelligence: Analyse von Wertschöpfung in Entwicklungsprozessen",
		MainText:       "Wenn wir uns mit agiler Softwareentwicklung beschäftigen, sprechen wir grundsätzlich auch über Application Lifecycle Management (ALM). Ebenso treibt das Business, das hinter allen Anforderungen für die Entwicklung von Software steht, immer die Frage nach Wertschöpfung um. Damit wir euch Antworten auf eure Fragen geben können, müssen wir die Werte von ALM mit den Methodiken von.",
		RedirectionURL: "http://feedproxy.google.com/~r/CodecentricBlog/~3/4eL6TuQHvVs/",
		UpdateDate:     feed.UpdateDate(time.Date(2018, 9, 26, 6, 0, 56, 0, time.UTC))}
	is.Equal(alexaFeed[0], expectedResponse)
}

type testClient struct {
}

func (c testClient) Get(url string) (*http.Response, error) {
	file, err := os.Open("../resource/test_rss_feeed_response.xml")
	return &http.Response{Status: "200 OK",
		Body:       file,
		StatusCode: 200}, err
}
