package internal

import "time"

// Plano representa un plano de rodaje con los detalles más relevantes
type Plano struct {
	tiempoEstimado time.Duration // Tiempo estimado de grabación del plano
	actores        []Actor       // Lista de actores que participarán en el plano
	lugar          string        // Lugar donde se llevará a cabo el rodaje del plano
	fechaInicio    time.Time     // Fecha en la que se rodará el plano
}

// NewPlano crea una nueva instancia de la estructura Plano con los datos proporcionados y la devuelve como un puntero
//
// Parámetros:
//
//	tiempoEstimado (time.Duration): El tiempo estimado de grabación del plano
//	actores ([]Actor): La lista de actores que participarán en el plano
//	lugar (string): El lugar donde se llevará a cabo el rodaje del plano
//	fechaInicio (time.Time): La fecha en la que se rodará el plano
//
// Valor de Retorno:
//
//	*Plano: Un puntero a la nueva instancia de Plano creada con los datos proporcionados
//
// Ejemplo de uso:
//
//	actores := []Actor{actor1, actor2}
//	fecha := time.Date(2023, time.October, 28, 10, 0, 0, 0, time.UTC)
//	plano := NewPlano(2*time.Hour, actores, "Locación A", fecha)
func NewPlano(tiempoEstimado time.Duration, actores []Actor, lugar string, fechaInicio time.Time) *Plano {
	return &Plano{
		tiempoEstimado: tiempoEstimado,
		actores:        actores,
		lugar:          lugar,
		fechaInicio:    fechaInicio,
	}
}
