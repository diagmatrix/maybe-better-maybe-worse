package tests

import (
	"testing"
	"time"

	"github.com/diagmatrix/maybe-better-maybe-worse/pkg/ppf"
	"github.com/stretchr/testify/assert"
)

// Test to check if the MemberTypeFromString function returns the correct MemberType
func TestMemberTypeFromString(t *testing.T) {
	// Test for the actor member type
	memberType, err := ppf.MemberTypeFromString("Actor")
	assert.Nil(t, err)
	assert.Equal(t, ppf.Actor, memberType)

	// Test for the crew member type
	memberType, err = ppf.MemberTypeFromString("Crew")
	assert.Nil(t, err)
	assert.Equal(t, ppf.Crew, memberType)
}

// Test to check if the MemberTypeFromString function returns an error for an invalid member type
func TestMemberTypeFromStringInvalid(t *testing.T) {
	memberType, err := ppf.MemberTypeFromString("Invalid")
	assert.NotNil(t, err)
	assert.Equal(t, ppf.MemberType(0), memberType)
}

// Test to check if the ParseTimeTable function returns the correct data for a valid time table
func TestParseTimeTable(t *testing.T) {
	// Create the available time list
	start, _ := time.Parse("02/01/2006 15:04", "23/11/2023 16:00")
	end, _ := time.Parse("02/01/2006 15:04", "23/11/2023 18:00")
	availableTime := [][]time.Time{{start, end}}

	// Test for a valid time table
	crewMember, err := ppf.ParseTimeTable("data/actor/valid-actor.csv")
	assert.Nil(t, err)
	assert.Equal(t, "valid-actor", crewMember.Name)
	assert.Equal(t, ppf.Actor, crewMember.CrewType)
	assert.ElementsMatch(t, availableTime, crewMember.Availability)
}

// Test to check if the ParseTimeTable function returns an error when the file does not exist or the file type is invalid
func TestParseTimeTableInvalidFile(t *testing.T) {
	_, errFilePath := ppf.ParseTimeTable("data/actor/invalid-file-path.csv")
	assert.NotNil(t, errFilePath)

	_, errFileType := ppf.ParseTimeTable("data/actor/invalid-file-type.txt")
	assert.NotNil(t, errFileType)
}

// Test to check if the ParseTimeTable function returns an error when the header of the file is invalid
func TestParseTimeTableInvalidHeader(t *testing.T) {
	_, err := ppf.ParseTimeTable("data/actor/invalid-header.csv")
	assert.NotNil(t, err)
}

// Test to check if the ParseTimeTable function returns an error when the time availability is invalid
func TestParseTimeTableInvalidAvailability(t *testing.T) {
	_, err := ppf.ParseTimeTable("data/actor/invalid-times.csv")
	assert.NotNil(t, err)
}
