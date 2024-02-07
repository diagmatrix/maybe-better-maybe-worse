package ppf

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/diagmatrix/maybe-better-maybe-worse/pkg/utils"
)

// MemberType represents the different kinds of crew members of the shooting plan
type MemberType int

// ErrInvalidMemberType is returned when the member type is invalid
var ErrInvalidMemberType = errors.New("invalid member type")

const (
	Actor MemberType = iota
	Crew
)

// MemberTypeFromString converts stringMemberType to a MemberType and returns an error if the string is invalid
func MemberTypeFromString(stringMemberType string) (MemberType, error) {
	switch stringMemberType {
	case "Actor", "actor":
		return Actor, nil
	case "Crew", "crew":
		return Crew, nil
	default:
		return 0, ErrInvalidMemberType
	}
}

// Member represents a member of the film crew
type Member struct {
	availability [][]time.Time
	shots        []Shot
	name         string
	crewType     MemberType
}

// NewMember creates a new member of the film crew given its name, type and availability
func NewMember(name string, crewType MemberType, availability [][]time.Time) Member {
	return Member{
		availability: availability,
		shots:        []Shot{},
		name:         name,
		crewType:     crewType,
	}
}

// GetAvailability returns the availability of the crew member
func (m Member) GetAvailability() [][]time.Time {
	return m.availability
}

// GetName returns the name of the crew member
func (m Member) GetName() string {
	return m.name
}

// GetCrewType returns the type of the crew member
func (m Member) GetCrewType() MemberType {
	return m.crewType
}

// ParseTimeTable parses a time table into the data needed to create a crew member given its file path.
// filePath contains the filepath of the time table (a csv file).
// This function returns a Member object or an error if something goes wrong.
func ParseTimeTable(filePath string) (Member, error) {

	filePath = filepath.FromSlash(filePath)

	// Get the crew type from the file path
	dir := filepath.Dir(filePath)
	crewTypeString := filepath.Base(dir)
	crewType, err := MemberTypeFromString(crewTypeString)
	if err != nil {
		return Member{}, err
	}

	// Get the name of the crew member from the file name
	fileName := filepath.Base(filePath)
	name := fileName[:len(fileName)-len(filepath.Ext(fileName))]

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return Member{}, err
	}

	// Check file extension
	if filepath.Ext(filePath) != ".csv" {
		return Member{}, utils.ErrInvalidFileType
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return Member{}, err
	}
	defer file.Close()

	// Read the file
	reader := csv.NewReader(file)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return Member{}, err
	}

	// Check the header
	if strings.Join(lines[0], ";") != utils.TIMETABLEHEADER {
		return Member{}, utils.ErrInvalidHeader
	}

	// Parse the availability
	var availability [][]time.Time
	for _, line := range lines[1:] {
		startDate, err := utils.FormatTime(line[0], line[1])
		if err != nil {
			return Member{}, err
		}
		endDate, err := utils.FormatTime(line[0], line[2])
		if err != nil {
			return Member{}, err
		}

		if endDate.Before(startDate) {
			return Member{}, utils.ErrInvalidAvailability
		}

		availability = append(availability, []time.Time{startDate, endDate})
	}

	return NewMember(name, crewType, availability), nil
}
