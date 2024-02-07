package utils

import (
	"errors"
	"fmt"
	"time"
)

const (
	TIMETABLEHEADER      = "Day;Start;End"
	TECHNICALGUIDEHEADER = "Scene;Shot;Duration;Setting;Cast;Crew;Details"
)

// InvalidHeader is returned when the header of a file is invalid
var ErrInvalidHeader = errors.New("invalid header")

// InvalidFileType is returned when the file type is invalid
var ErrInvalidFileType = errors.New("invalid file type")

// InvalidAvailability is returned when the availability of a crew member is invalid
var ErrInvalidAvailability = errors.New("invalid availability")

// InvalidCast is returned when the cast of a shot is invalid
var ErrInvalidCast = errors.New("invalid cast")

// InvalidCrew is returned when the crew of a shot is invalid
var ErrInvalidCrew = errors.New("invalid crew")

// formatTime parses a string pair of date and time into a time.Time object with layout dd-mm-yyyy HH:MM
// It takes two strings as arguments, the first one is the date in format dd-mm-yyyy and the second one is the time in format HH:MM
// It returns a time.Time object or an error if the parsing fails
func FormatTime(dateString string, timeString string) (time.Time, error) {
	layout := "02/01/2006 15:04"
	dateTimeString := fmt.Sprintf("%s %s", dateString, timeString)
	return time.Parse(layout, dateTimeString)
}
