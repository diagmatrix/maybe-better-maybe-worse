package ppf

import "time"

// MemberType represents the different kinds of crew members of the shooting plan
type MemberType int

const (
	Actor MemberType = iota
	Crew
)

// Member represents a member of the film crew
type Member struct {
	availability []time.Time
	shots        []Shot
	name         string
	crewType     MemberType
}
