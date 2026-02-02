package aggregation

import (
	"cddude229/kq-tourney-analyzer/models"
	"encoding/csv"
	"io"
	"os"
)

// TODO: Extend this to be assymetric

type directTeamMapGenerator struct {
	TeamName string

	BStripesName string
	BAbsName     string
	BQueenName   string
	BSkullsName  string
	BChexName    string

	GStripesName string
	GAbsName     string
	GQueenName   string
	GSkullsName  string
	GChexName    string
}

func (e directTeamMapGenerator) NamePlayer(id models.PlayerId, gameId int64, matchId int64, teamName string) *string {
	if teamName == e.TeamName {
		switch id {
		case models.BlueStripes:
			return &e.BStripesName
		case models.BlueAbs:
			return &e.BAbsName
		case models.BlueQueen:
			return &e.BQueenName
		case models.BlueSkulls:
			return &e.BSkullsName
		case models.BlueChex:
			return &e.BChexName

		case models.GoldChex:
			return &e.GChexName
		case models.GoldSkulls:
			return &e.GSkullsName
		case models.GoldQueen:
			return &e.GQueenName
		case models.GoldAbs:
			return &e.GAbsName
		case models.GoldStripes:
			return &e.GStripesName

		default:
			return nil
		}
	}

	return nil
}

func TeamsFromCsvFile(filePath string) ([]PlayerNameGenerator, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	_, err = csvReader.Read() // Skip header row
	if err != nil {
		return nil, err
	}

	results := make([]PlayerNameGenerator, 0)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		gen := directTeamMapGenerator{
			TeamName: record[0],

			BStripesName: record[1],
			BAbsName:     record[2],
			BQueenName:   record[3],
			BSkullsName:  record[4],
			BChexName:    record[5],

			GStripesName: record[6],
			GAbsName:     record[7],
			GQueenName:   record[8],
			GSkullsName:  record[9],
			GChexName:    record[10],
		}

		results = append(results, gen)
	}

	return results, nil

}
