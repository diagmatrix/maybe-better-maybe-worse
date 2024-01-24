package ppf

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/diagmatrix/maybe-better-maybe-worse/pkg/utils"
)

// Shot represents a shot in a the shooting plan
type Shot struct {
	estimatedDuration time.Duration
	cast              []Member
	crew              []Member
	setting           string
	scene             string
	id                string
	details           string
}

// NewShot creates a new shot
// It takes the estimated duration of the shot, the cast members, the crew members, the setting, the scene, the id and the details as arguments
// It returns a Shot
func NewShot(estimatedDuration time.Duration, cast []Member, crew []Member, setting, scene, id, details string) Shot {
	return Shot{
		estimatedDuration: estimatedDuration,
		cast:              cast,
		crew:              crew,
		setting:           setting,
		scene:             scene,
		id:                id,
		details:           details,
	}
}

// ParseTechnicalGuide parses a technical guide into the data needed to create a shot
// It takes the file path as an argument and the list of crew members
// It returns a list of shots or an error if the parsing fails
func ParseTechnicalGuide(filePath string, members []Member) ([]Shot, error) {
	filePath = filepath.FromSlash(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return []Shot{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return []Shot{}, err
	}

	if strings.Join(lines[0], ";") != utils.TECHNICALGUIDEHEADER {
		return []Shot{}, utils.ErrInvalidHeader
	}

	var shots []Shot
	for _, line := range lines[1:] {
		scene := line[0]
		shotID := line[1]
		setting := line[3]
		details := ""
		if len(line) > 5 {
			details = line[6]
		}

		estimatedDuration, err := time.ParseDuration(line[2])
		if err != nil {
			return []Shot{}, err
		}

		var cast []Member
		if len(line[4]) == 0 {
			cast = []Member{}
		} else {
			cast_names := strings.Split(line[4], ",")
			for _, name := range cast_names {
				for _, member := range members {
					if member.Name == name {
						cast = append(cast, member)
					}
				}
			}
		}

		var crew []Member
		if len(line[5]) == 0 {
			crew = []Member{}
		} else {
			crew_names := strings.Split(line[5], ",")
			for _, name := range crew_names {
				for _, member := range members {
					if member.Name == name {
						crew = append(crew, member)
					}
				}
			}
		}
		shots = append(shots, NewShot(estimatedDuration, cast, crew, setting, scene, shotID, details))
	}

	return shots, nil
}
