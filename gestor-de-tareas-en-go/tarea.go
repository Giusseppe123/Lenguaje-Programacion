package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type EstadoTarea int

const (
	Pendiente EstadoTarea = iota
	EnProgreso
	Completada
)

func (et EstadoTarea) String() string {
	switch et {
	case Pendiente:
		return "Pendiente"
	case EnProgreso:
		return "En Progreso"
	case Completada:
		return "Completada"
	default:
		return "Desconocido"
	}
}

func (et EstadoTarea) MarshalJSON() ([]byte, error) {
	return json.Marshal(et.String())
}

func (et *EstadoTarea) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Pendiente":
		*et = Pendiente
	case "En Progreso":
		*et = EnProgreso
	case "Completada":
		*et = Completada
	default:
		return fmt.Errorf("estado de tarea desconocido: %s", s)
	}
	return nil
}

type Tarea struct {
	ID            int         `json:"id"`
	Titulo        string      `json:"titulo"`
	Estado        EstadoTarea `json:"estado"`
	FechaCreacion time.Time   `json:"fecha_creacion"`
	FechaEdicion  *time.Time  `json:"fecha_edicion,omitempty"`
}

type Tareas []Tarea