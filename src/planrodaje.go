package src

// PlanRodaje representa un plan de rodaje que contiene una serie de planos
type PlanRodaje struct {
	codigo     string  // Código único del plan de rodaje
	planos     []Plano // Lista de planos que forman parte del plan de rodaje
	pelicula   string  // Nombre de la película a la que se le asigna el plan
	direccion  []string  // Nombre del director o directora
	produccion []string  // Nombre del productor o productora
}

// NewPlanRodaje crea una nueva instancia de la estructura PlanRodaje con los datos proporcionados y la devuelve como un puntero
//
// Parámetros:
//
//	codigo (string): el código único del plan de rodaje
//	planos ([]Plano): la lista de planos que forman parte del plan de rodaje
//	pelicula (string): el nombre de la película a la que se le asigna el plan
//	direccion (string): el nombre del director o directora
//	produccion (string): el nombre del productor o productora
//
// Valor de retorno:
//
//	*PlanRodaje: un puntero a la nueva instancia de PlanRodaje creada con los datos proporcionados
//
// Ejemplo de uso:
//
//	plano1 := plano{Codigo: "P1", Descripcion: "Escena de apertura"}
//	plano2 := plano{Codigo: "P2", Descripcion: "Escena final"}
//	planos := []Plano{plano1, plano2}
//	planRodaje := NewPlanRodaje("PR123", planos, "Mi Película", "Director A", "Productor B")
func NewPlanRodaje(codigo string, planos []Plano, pelicula string, direccion string, produccion string) *PlanRodaje {
	return &PlanRodaje{
		codigo:     codigo,
		planos:     planos,
		pelicula:   pelicula,
		direccion:  direccion,
		produccion: produccion,
	}
}
