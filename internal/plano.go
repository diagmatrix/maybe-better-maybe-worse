package internal

import "time"

// Plano representa un plano de rodaje con los detalles más relevantes
type Plano struct {
	tiempoEstimado time.Duration
	actores        []Actor
	lugarGrabacion string
	fechaInicio    time.Time
}

