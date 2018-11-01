package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hill-daniel/alexa-rss-flashbriefing/aws"
	"github.com/hill-daniel/alexa-rss-flashbriefing/rss"
	"net/http"
	"os"
	"time"
)

const envAtomFeedURL = "ATOM_FEED_URL"

func main() {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	mapper := rss.CodecentricMapper{}
	url := os.Getenv(envAtomFeedURL)
	feedReader := rss.NewAtomFeedReader(client, mapper, url)
	handler := aws.NewFlashBriefingHandler(feedReader)

	lambda.Start(handler.Handle)
}
