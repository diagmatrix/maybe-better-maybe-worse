package ppf

import "time"

type Plano struct {
	tiempoEstimado time.Duration
	actores        []Miembro
	equipoTecnico  []Miembro
	lugarGrabacion string
	escena         string
	identificador  string
}
