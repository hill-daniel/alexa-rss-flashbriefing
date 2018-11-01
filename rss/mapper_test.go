package rss_test

import (
	"github.com/hill-daniel/alexa-rss-flashbriefing"
	"github.com/hill-daniel/alexa-rss-flashbriefing/rss"
	assert "github.com/matryer/is"
	"testing"
	"time"
)

func Test_should_map_atom_feed_to_alexa_feed(t *testing.T) {
	is := assert.New(t)
	channel := createChannel("This is the first news item", "This is some longer text, which describes nicely what should be reported")
	atomFeed := &feed.AtomFeed{Channel: &channel}
	mapper := &rss.CodecentricMapper{}

	responseFeed, err := mapper.Map(atomFeed)

	is.NoErr(err)
	is.Equal(len(responseFeed), 1)
	expectedResponseItem := feed.AlexaItem{UUID: "A259DD9F-B2A3-460E-B242-E606CF4FE02C",
		RedirectionURL: "https://my.awesome.site/rss/A259DD9F-B2A3-460E-B242-E606CF4FE02C",
		TitleText:      "This is the first news item",
		MainText:       "This is some longer text, which describes nicely what should be reported",
		UpdateDate:     feed.UpdateDate(time.Date(2018, 2, 13, 6, 0, 0, 0, time.UTC))}
	is.Equal(responseFeed[0], expectedResponseItem)
}

func Test_should_truncate_main_text_till_weiterlesen_part_and_remove_html_in_title_and_description(t *testing.T) {
	is := assert.New(t)
	titleWithTags := "This is the <strong>first</strong> news item"
	descriptionWithTags := `<![CDATA[<p>Mit einem guten Freund streite ich mich immer wieder über das Thema Tests in Web-Frontends. In der Regel vertauschen sich unsere Positionen &#8211; je nachdem, wie frustriert derjenige ist, der aktuell Tests schreibt. Wir sprechen heute meistens eher von Web-Anwendungen als von Web-Frontends. Der Code dazu ist entsprechend umfangreich, obwohl moderne Web-Frameworks wie Angular, React... <a class="view-article" href="https://blog.codecentric.de/2018/02/javascript-ui-tests-mit-struktur/">Weiterlesen</a></p>
<p>The post <a rel="nofollow" href="https://blog.codecentric.de/2018/02/javascript-ui-tests-mit-struktur/">JavaScript UI Tests mit Struktur</a> appeared first on <a rel="nofollow" href="https://blog.codecentric.de">codecentric AG Blog</a>.</p>
]]>`
	channel := createChannel(titleWithTags, descriptionWithTags)
	atomFeed := &feed.AtomFeed{Channel: &channel}
	mapper := &rss.CodecentricMapper{}

	responseFeed, err := mapper.Map(atomFeed)

	is.NoErr(err)
	is.Equal(len(responseFeed), 1)
	expectedResponseItem := feed.AlexaItem{UUID: "A259DD9F-B2A3-460E-B242-E606CF4FE02C",
		RedirectionURL: "https://my.awesome.site/rss/A259DD9F-B2A3-460E-B242-E606CF4FE02C",
		TitleText:      "This is the first news item",
		MainText:       "Mit einem guten Freund streite ich mich immer wieder über das Thema Tests in Web-Frontends. In der Regel vertauschen sich unsere Positionen – je nachdem, wie frustriert derjenige ist, der aktuell Tests schreibt. Wir sprechen heute meistens eher von Web-Anwendungen als von Web-Frontends. Der Code dazu ist entsprechend umfangreich, obwohl moderne Web-Frameworks wie Angular, React.",
		UpdateDate:     feed.UpdateDate(time.Date(2018, 2, 13, 6, 0, 0, 0, time.UTC))}
	is.Equal(responseFeed[0], expectedResponseItem)
}

func createChannel(title, description string) feed.Channel {
	item := feed.AtomItem{GUID: "A259DD9F-B2A3-460E-B242-E606CF4FE02C",
		Link:    "https://my.awesome.site/rss/A259DD9F-B2A3-460E-B242-E606CF4FE02C",
		Title:   title,
		Desc:    description,
		PubDate: "Tue, 13 Feb 2018 07:00:00 +0100"}
	channel := feed.Channel{Title: "myRssChannel",
		Desc:     "Awesome channel",
		Items:    []feed.AtomItem{item},
		Language: "de"}
	return channel
}
