package internal

// PlanRodaje representa un plan de rodaje que contiene una serie de planos
type PlanRodaje struct {
	codigo     string  // Código único del plan de rodaje
	planos     []Plano // Lista de planos que forman parte del plan de rodaje
	pelicula   string  // Nombre de la película a la que se le asigna el plan
	direccion  []string  // Nombre del director o directora
	produccion []string  // Nombre del productor o productora
}

