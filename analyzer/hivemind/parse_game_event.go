package hivemind

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type HivemindEvent struct {
	Id        int64
	Timestamp time.Time
	EventType string
	Values    string
	GameId    int64
}

func parseGameEventCsv(reader io.ReadCloser) ([]HivemindEvent, error) {
	var events []HivemindEvent

	csvReader := csv.NewReader(reader)
	headerRow, err := csvReader.Read() // Skip header row
	if err != nil {
		return nil, err
	}

	if strings.Join(headerRow, ",") != "id,timestamp,event_type,values,game_id" {
		return nil, fmt.Errorf("New CSV header format detected.  Code needs updating.")
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse id (\"%s\") failed: %w", record[0], err)
		}

		// Sample: 2025-09-27 20:06:32.021+00
		timestamp, err := time.Parse("2006-01-02 15:04:05-07", record[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse timestamp (\"%s\") failed: %w", record[1], err)
		}

		gameId, err := strconv.ParseInt(record[4], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse gameId (\"%s\") failed: %w", record[4], err)
		}

		events = append(events, HivemindEvent{
			Id:        id,
			Timestamp: timestamp,
			EventType: record[2],
			Values:    record[3],
			GameId:    gameId,
		})

	}

	return events, nil
}
