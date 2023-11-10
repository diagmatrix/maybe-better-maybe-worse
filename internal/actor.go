package internal

import "time"

// Actor representa un actor con información básica y su disponibilidad horaria.
type Actor struct {
	horariosDisponibilidad []time.Time
	planosARodar   []Plano
}

