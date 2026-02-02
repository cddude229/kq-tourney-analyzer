package hivemind

import (
	"cddude229/kq-tourney-analyzer/models"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

func parseGameCsv(reader io.ReadCloser) ([]TourneyMatch, error) {
	// TODO: UTs
	var matches []TourneyMatch

	csvReader := csv.NewReader(reader)
	headerRow, err := csvReader.Read() // Skip header row
	if err == io.EOF {
		return matches, nil
	} else if err != nil {
		return nil, err
	}

	if strings.Join(headerRow, ",") != "id,start_time,end_time,win_condition,winning_team,map_name,player_count,cabinet_id,cabinet_name,tournament_match_id,blue_team,gold_team" {
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

		event, err := parseTourneyMatchRecord(record)
		if err != nil {
			return nil, err
		}

		matches = append(matches, *event)
	}

	sort.Sort(TourneyMatchById(matches))

	return matches, nil
}

func parseTourneyMatchRecord(record []string) (*TourneyMatch, error) {
	id, err := strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id (\"%s\") failed: %w", record[0], err)
	}

	// Sample: 2025-09-27 20:06:32.021+00
	startTime, err := time.Parse("2006-01-02 15:04:05-07", record[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse start_time (\"%s\") failed: %w", record[1], err)
	}

	endTime, err := time.Parse("2006-01-02 15:04:05-07", record[2])
	if err != nil {
		return nil, fmt.Errorf("failed to parse end_time (\"%s\") failed: %w", record[1], err)
	}

	winCondition := models.WinCondition(strings.ToUpper(record[3][:1]) + record[3][1:])

	playerCount, err := strconv.ParseInt(record[6], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse player_count (\"%s\") failed: %w", record[6], err)
	}

	cabinetId, err := strconv.ParseInt(record[7], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse cabinet_id (\"%s\") failed: %w", record[6], err)
	}

	matchId, err := strconv.ParseInt(record[9], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse match_id (\"%s\") failed: %w", record[6], err)
	}

	return &TourneyMatch{
		Id:                id,
		StartTime:         startTime,
		EndTime:           endTime,
		WinCondition:      winCondition,
		WinningTeam:       record[4],
		MapName:           record[5],
		PlayerCount:       playerCount,
		CabinetId:         cabinetId,
		CabinetName:       record[8],
		TournamentMatchId: matchId,
		BlueTeamName:      record[10],
		GoldTeamName:      record[11],
	}, nil
}
