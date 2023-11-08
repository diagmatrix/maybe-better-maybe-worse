package internal

import "time"

// Actor representa un actor con información básica y su disponibilidad horaria.
type Actor struct {
	dni            string      // Número de identificación del actor
	nombre         string      // Nombre del actor
	disponibilidad []time.Time // Horarios de disponibilidad del actor
	planosARodar   []Plano     // Planos en los que participa
}

