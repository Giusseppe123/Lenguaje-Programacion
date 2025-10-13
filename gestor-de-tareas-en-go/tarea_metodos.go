package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func (g *GestorDeTareas) Agregar(titulo string) {
	nuevoID := g.SiguienteID
	g.SiguienteID++

	tarea := Tarea{
		ID:            nuevoID,
		Titulo:        titulo,
		Estado:        Pendiente,
		FechaCreacion: time.Now(),
		FechaEdicion:  nil,
	}
	g.Tareas = append(g.Tareas, tarea)
}

func (g *GestorDeTareas) validarID(id int) (int, error) {
	for i, tarea := range g.Tareas {
		if tarea.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("ID de tarea no encontrado")
}

func (g *GestorDeTareas) Eliminar(id int) error {
	idx, err := g.validarID(id)
	if err != nil {
		return err
	}
	g.Tareas = append((g.Tareas)[:idx], (g.Tareas)[idx+1:]...)
	return nil
}

func (g *GestorDeTareas) AlternarEstado(id int) error {
	idx, err := g.validarID(id)
	if err != nil {
		return err
	}
	switch g.Tareas[idx].Estado {
	case Pendiente:
		g.Tareas[idx].Estado = EnProgreso
	case EnProgreso:
		g.Tareas[idx].Estado = Completada
	case Completada:
		g.Tareas[idx].Estado = Pendiente
	}
	return nil
}

func (g *GestorDeTareas) Editar(id int, nuevoTitulo string) error {
	idx, err := g.validarID(id)
	if err != nil {
		return err
	}
	g.Tareas[idx].Titulo = nuevoTitulo
	ahora := time.Now()
	g.Tareas[idx].FechaEdicion = &ahora
	return nil
}

func (g *GestorDeTareas) Imprimir() {
	if len(g.Tareas) == 0 {
		fmt.Println("No hay tareas para mostrar. ¡Agrega alguna!")
		return
	}

	sort.SliceStable(g.Tareas, func(i, j int) bool {
		return g.Tareas[i].ID < g.Tareas[j].ID
	})

	fmt.Printf("%-5s %-30s %-15s %-20s %-20s\n", "ID", "Titulo", "Estado", "Creada el", "Editada el")
	fmt.Println("----- ------------------------------ --------------- -------------------- --------------------")
	for _, tarea := range g.Tareas {
		fechaEdicionStr := "-"
		if tarea.FechaEdicion != nil {
			fechaEdicionStr = tarea.FechaEdicion.Format("2006-01-02 15:04:05")
		}
		fmt.Printf("%-5d %-30s %-15s %-20s %-20s\n",
			tarea.ID,
			tarea.Titulo,
			tarea.Estado.String(),
			tarea.FechaCreacion.Format("2006-01-02 15:04:05"),
			fechaEdicionStr,
		)
	}
}