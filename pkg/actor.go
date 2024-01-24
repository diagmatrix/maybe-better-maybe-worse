package ppf

import "time"

type TipoMiembro int

const (
	Actor TipoMiembro = iota
	Tecnico
)

type Miembro struct {
	horariosDisponibilidad []time.Time
	planosARodar           []Plano
	nombre                 string
	Tipo                   TipoMiembro
}
