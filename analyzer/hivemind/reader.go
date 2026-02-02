package hivemind

import (
	"archive/zip"
	"fmt"
	"log"
	"strings"
)

func OpenAndParseZip(zipPath string) ([]HivemindEvent, []TourneyMatch, error) {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	hivemindEvents := []HivemindEvent{}
	matches := []TourneyMatch{}

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		// 3. Open individual file
		rc, err := file.Open()
		if err != nil {
			return hivemindEvents, matches, err
		}

		if strings.HasSuffix(file.Name, "/gameevent.csv") {
			hivemindEvents, err = parseGameEventCsv(rc)
			if err != nil {
				defer rc.Close()
				return hivemindEvents, matches, fmt.Errorf("error in file %s: %w", file.Name, err)
			}
		} else if strings.HasSuffix(file.Name, "/usergame.csv") {
			// Expected, so ignoring and closing
		} else if strings.HasSuffix(file.Name, "/game.csv") {
			matches, err = parseGameCsv(rc)
			if err != nil {
				defer rc.Close()
				return hivemindEvents, matches, fmt.Errorf("error in file %s: %w", file.Name, err)
			}
		} else {
			log.Printf("Unexpected file name: %s\n", file.Name)
		}

		err = rc.Close()
		if err != nil {
			return hivemindEvents, matches, err
		}
	}

	return hivemindEvents, matches, nil
}
