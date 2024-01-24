package ppf

import "time"

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
