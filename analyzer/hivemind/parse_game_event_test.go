package hivemind

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseRecord(t *testing.T) {
	event, err := parseRecord([]string{
		"175893643", "2025-09-27 20:06:32.021+00", "useMaiden", "{560,260,maiden_wings,8}", "1650408",
	})

	if assert.NoError(t, err) {
		// I hate that I'm parsing here, but I couldn't get Location to work right
		parsedTime, err := time.Parse("2006-01-02 15:04:05-07", "2025-09-27 20:06:32.021+00")
		assert.NoError(t, err)

		assert.EqualValues(t, &HivemindEvent{
			Id:        175893643,
			Timestamp: parsedTime,
			EventType: "useMaiden",
			Values:    "{560,260,maiden_wings,8}",
			GameId:    1650408,
		}, event)
	}

}
