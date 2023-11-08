package internal

import "time"

// Actor representa un actor con información básica y su disponibilidad horaria.
type Actor struct {
	dni            string      // Número de identificación del actor
	nombre         string      // Nombre del actor
	disponibilidad []time.Time // Horarios de disponibilidad del actor
	planosARodar   []Plano     // Planos en los que participa
}

// NewActor crea una nueva instancia de la estructura Actor con los datos proporcionados y la devuelve como un puntero
//
// Parámetros:
//
//	dni (string): el número de identificación del actor
//	nombre (string): el nombre del actor
//	edad (int): la edad del actor
//	disponibilidad ([]time.Time): un slice de objetos time.Time que representa los horarios de disponibilidad del actor
//
// Valor de retorno:
//
//	*Actor: un puntero a la nueva instancia de Actor creada con los datos proporcionados
//
// Ejemplo de uso:
//
//	disponibilidad := []time.Time{time.Date(2023, time.October, 26, 9, 0, 0, 0, time.UTC), time.Date(2023, time.October, 27, 14, 30, 0, 0, time.UTC),}
//	actor := NewActor("123456789", "Juana Jiménez", 30, disponibilidad)
func NewActor(DNI, nombre string, edad int, disponibilidad []time.Time) *Actor {
	return &Actor{
		dni:            DNI,
		nombre:         nombre,
		edad:           edad,
		disponibilidad: disponibilidad,
	}
}
