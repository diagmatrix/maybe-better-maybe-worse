package tests

import (
	"testing"
	"time"

	"github.com/diagmatrix/maybe-better-maybe-worse/pkg/ppf"
	"github.com/stretchr/testify/assert"
)

// Test to check if ParseTechnicalGuide	function returns the correct data for a valid technical guide
func TestParseTechnicalGuide(t *testing.T) {
	// Create the Member list
	actor := ppf.NewMember("actor", ppf.Actor, [][]time.Time{{time.Now(), time.Now()}})
	actor2 := ppf.NewMember("actor2", ppf.Actor, [][]time.Time{{time.Now(), time.Now()}})
	crew := ppf.NewMember("crew", ppf.Crew, [][]time.Time{{time.Now(), time.Now()}})

	// Get times (10 minutes)
	tenMinutes := time.Duration(600000000000)

	// Test for a valid technical guide
	shots, err := ppf.ParseTechnicalGuide("data/valid-tech-guide.csv", []ppf.Member{actor, actor2, crew})
	// Basic checks
	assert.Nil(t, err)
	assert.Equal(t, 2, len(shots))
	// Check the first shot
	firstShot := shots[0]
	assert.Equal(t, tenMinutes, firstShot.EstimatedDuration)
	assert.ElementsMatch(t, []ppf.Member{actor}, firstShot.Cast)
	assert.ElementsMatch(t, []ppf.Member{crew}, firstShot.Crew)
	assert.Equal(t, "Test Location", firstShot.Setting)
	assert.Equal(t, "INTRO", firstShot.Scene)
	assert.Equal(t, "P1", firstShot.Id)
	assert.Equal(t, "Test Details", firstShot.Details)
}

// Test to check if ParseTechnicalGuide function returns an error when the file does not exist or the file type is invalid
func TestParseTechnicalGuideInvalidFile(t *testing.T) {
	_, errFilePath := ppf.ParseTechnicalGuide("data/invalid-tech-guide.csv", []ppf.Member{})
	assert.NotNil(t, errFilePath)

	_, errFileType := ppf.ParseTechnicalGuide("data/invalid-tech-guide.txt", []ppf.Member{})
	assert.NotNil(t, errFileType)
}

// Test to check if ParseTechnicalGuide function returns an error when the header of the file is invalid
func TestParseTechnicalGuideInvalidHeader(t *testing.T) {
	_, err := ppf.ParseTechnicalGuide("data/invalid-tech-guide-header.csv", []ppf.Member{})
	assert.NotNil(t, err)
}

// Test to check if ParseTechnicalGuide function returns an error when the cast or crew of a shot is invalid
func TestParseTechnicalGuideInvalidCastCrew(t *testing.T) {
	// Create the Member list
	actor := ppf.NewMember("actor", ppf.Actor, [][]time.Time{{time.Now(), time.Now()}})
	crew := ppf.NewMember("crew", ppf.Crew, [][]time.Time{{time.Now(), time.Now()}})

	// Test for a valid technical guide
	_, err := ppf.ParseTechnicalGuide("data/valid-tech-guide.csv", []ppf.Member{actor, crew})
	assert.NotNil(t, err)
}
