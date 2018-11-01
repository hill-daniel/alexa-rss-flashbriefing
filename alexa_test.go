package feed_test

import (
	"encoding/json"
	"github.com/hill-daniel/alexa-rss-flashbriefing"
	assert "github.com/matryer/is"
	"testing"
	"time"
)

func Test_should_marshal_update_time_to_expected_format(t *testing.T) {
	is := assert.New(t)
	date := feed.UpdateDate(time.Date(2018, 2, 13, 14, 12, 49, 0, time.UTC))

	bytes, err := json.Marshal(date)

	is.NoErr(err)
	is.Equal(string(bytes), "\"2018-02-13T14:12:49Z\"")
}

func Test_should_unmarshal_JSON_to_UpdateDate(t *testing.T) {
	is := assert.New(t)
	dateString := "\"2018-02-13T14:12:49Z\""
	updateDate := &feed.UpdateDate{}

	err := json.Unmarshal([]byte(dateString), updateDate)

	is.NoErr(err)
	is.Equal(time.Time(*updateDate), time.Date(2018, 2, 13, 14, 12, 49, 0, time.UTC))
}
