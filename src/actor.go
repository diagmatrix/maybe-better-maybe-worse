package src

import "time"

// Actor representa un actor con información básica y su disponibilidad horaria.
type Actor struct {
	DNI            string      // Número de identificación del actor
	nombre         string      // Nombre del actor
	edad           int         // Edad del actor
	disponibilidad []time.Time // Horarios de disponibilidad del actor
}
