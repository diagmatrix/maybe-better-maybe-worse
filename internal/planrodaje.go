package internal

import "time"

type AsignacionPlano struct {
	PlanoAAsignar	Plano
	FechaInicio	time.Time
}

type PlanRodaje struct {
	planosAsignados	[]AsignacionPlano
}
