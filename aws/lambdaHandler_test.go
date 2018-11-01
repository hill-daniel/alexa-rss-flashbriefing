package aws_test

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/hill-daniel/alexa-rss-flashbriefing"
	"github.com/hill-daniel/alexa-rss-flashbriefing/aws"
	assert "github.com/matryer/is"
	"testing"
	"time"
)

func Test_should_return_http_200_and_feed_as_json_in_body(t *testing.T) {
	is := assert.New(t)
	feedResponse := createResponse()
	feedReader := testFeedReader{response: feedResponse}
	handler := aws.NewFlashBriefingHandler(feedReader)

	response, err := handler.Handle(events.APIGatewayProxyRequest{})

	is.NoErr(err)
	body, _ := json.Marshal(feedResponse)
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}
	is.Equal(response, expectedResponse)
}

func createResponse() []feed.AlexaItem {
	feedResponse := []feed.AlexaItem{{UUID: "CAF47550-B2E4-4294-AD88-FE3332E82997",
		RedirectionURL: "https://blog.codecentric.de/en/2018/02/continuous-integration-drone-aws/",
		TitleText:      "Continuous Integration with Drone on AWS",
		MainText:       "Drone is a Continuous Delivery system built on container technology and written in Go. It uses a simple configuration yaml, with syntax similar to a docker-compose file, to define and execute pipelines inside Docker containers.",
		UpdateDate:     feed.UpdateDate(time.Date(2018, 2, 14, 13, 37, 0, 0, time.UTC))}}
	return feedResponse
}

type testFeedReader struct {
	response feed.AlexaFeed
}

func (tr testFeedReader) Read() (feed.AlexaFeed, error) {
	return tr.response, nil
}
