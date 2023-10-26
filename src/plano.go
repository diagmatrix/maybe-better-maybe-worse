package src

import "time"

// Plano representa un plano de rodaje con los detalles más relevantes
type Plano struct {
	tiempoEstimado time.Duration // Tiempo estimado de grabación del plano
	actores        []Actor       // Lista de actores que participarán en el plano
	lugar          string        // Lugar donde se llevará a cabo el rodaje del plano
	fecha          time.Time     // Fecha en la que se rodará el plano
}
