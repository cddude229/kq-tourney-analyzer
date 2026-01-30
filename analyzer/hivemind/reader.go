package hivemind

import (
	"archive/zip"
	"log"
	"strings"
)

func OpenAndParseZip(zipPath string) ([]HivemindEvent, error) {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		// 3. Open individual file
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		if strings.HasSuffix(file.Name, "gameevent.csv") {
			defer rc.Close()
			return parseGameEventCsv(rc)
		} else if strings.HasSuffix(file.Name, "game.csv") || strings.HasSuffix(file.Name, "usergame.csv") {
			// Expected, so ignoring and closing
		} else {
			log.Printf("Unexpected file name: %s\n", file.Name)
		}

		err = rc.Close()
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
