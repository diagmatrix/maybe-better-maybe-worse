package ppf

import "time"

type Actor struct {
	horariosDisponibilidad []time.Time
	planosARodar           []Plano
}
