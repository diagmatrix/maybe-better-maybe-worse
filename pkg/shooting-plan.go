package ppf

import "time"

// PlannedShot represents a shot that is planned in the shooting plan
type PlannedShot struct {
	shot         Shot
	startingDate time.Time
}

// ShootingPlan represents the shooting plan
type ShootingPlan struct {
	plannedShots []PlannedShot
}
