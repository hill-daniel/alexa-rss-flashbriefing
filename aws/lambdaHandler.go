// Package aws provides functionality for AWS abstractions.
package aws

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/hill-daniel/alexa-rss-flashbriefing"
	"github.com/pkg/errors"
)

// FlashBriefingHandler is the entry point for the lambda function.
// Returns Alexa Flash Briefing friendly JSON.
// See https://developer.amazon.com/docs/flashbriefing/flash-briefing-skill-api-feed-reference.html
type FlashBriefingHandler struct {
	feedReader feed.AlexaFeedReader
}

// NewFlashBriefingHandler creates a FlashBriefingHandler.
func NewFlashBriefingHandler(alexaFeedReader feed.AlexaFeedReader) *FlashBriefingHandler {
	return &FlashBriefingHandler{
		feedReader: alexaFeedReader,
	}
}

// Handle returns the most recent RSS feed as JSON, wrapped in an APIGatewayProxyResponse.
func (h FlashBriefingHandler) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	alexaFeed, err := h.feedReader.Read()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	jsonFeed, err := json.Marshal(alexaFeed)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.Wrapf(err, "failed to marshal alexa feed, %v", alexaFeed)
	}
	return events.APIGatewayProxyResponse{
		Body:       string(jsonFeed),
		StatusCode: 200,
	}, nil
}
