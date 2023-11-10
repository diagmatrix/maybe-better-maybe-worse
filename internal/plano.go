package internal

import "time"

type Plano struct {
	tiempoEstimado time.Duration
	actores        []Actor
	lugarGrabacion string
	fechaInicio    time.Time
}

