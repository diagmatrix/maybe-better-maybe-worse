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

// MemberTypeFromString converts a string to a MemberType
// It takes a string as an argument
// It returns a MemberType or an error if the conversion fails
func MemberTypeFromString(memberType string) (MemberType, error) {
	switch memberType {
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
	Name         string
	crewType     MemberType
}

// NewMember creates a new member of the film crew
// It takes the name of the member, the type of the member and the availability of the member as arguments
// It returns a Member
func NewMember(name string, crewType MemberType, availability [][]time.Time) Member {
	return Member{
		availability: availability,
		shots:        []Shot{},
		Name:         name,
		crewType:     crewType,
	}
}

// ParseTimeTable parses a time table into the data needed to create a crew member
// It takes the file path as an argument
// It returns the name of the crew member, the type of the crew member and the availability of the crew member
func ParseTimeTable(filePath string) (Member, error) {

	filePath = filepath.FromSlash(filePath)

	dir := filepath.Dir(filePath)
	crewTypeString := filepath.Base(dir)
	crewType, err := MemberTypeFromString(crewTypeString)
	if err != nil {
		return Member{}, err
	}

	fileName := filepath.Base(filePath)
	name := fileName[:len(fileName)-len(filepath.Ext(fileName))]

	file, err := os.Open(filePath)
	if err != nil {
		return Member{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return Member{}, err
	}

	if strings.Join(lines[0], ";") != utils.TIMETABLEHEADER {
		return Member{}, utils.ErrInvalidHeader
	}

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
		availability = append(availability, []time.Time{startDate, endDate})
	}

	return NewMember(name, crewType, availability), nil
}
