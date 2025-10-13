package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func (t *Tareas) Agregar(titulo string) {
	nuevoID := 0
	if len(*t) > 0 {
		sort.Slice(*t, func(i, j int) bool {
			return (*t)[i].ID < (*t)[j].ID
		})
		nuevoID = (*t)[len(*t)-1].ID + 1
	}

	tarea := Tarea{
		ID:            nuevoID,
		Titulo:        titulo,
		Estado:        Pendiente,
		FechaCreacion: time.Now(),
	}
	*t = append(*t, tarea)
}

func (t *Tareas) validarID(id int) (int, error) {
	for i, tarea := range *t {
		if tarea.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("ID de tarea no encontrado")
}

func (t *Tareas) Eliminar(id int) error {
	idx, err := t.validarID(id)
	if err != nil {
		return err
	}
	*t = append((*t)[:idx], (*t)[idx+1:]...)
	return nil
}

func (t *Tareas) AlternarEstado(id int) error {
	idx, err := t.validarID(id)
	if err != nil {
		return err
	}

	switch (*t)[idx].Estado {
	case Pendiente:
		(*t)[idx].Estado = EnProgreso
	case EnProgreso:
		(*t)[idx].Estado = Completada
	case Completada:
		(*t)[idx].Estado = Pendiente
	}
	return nil
}

func (t *Tareas) Editar(id int, nuevoTitulo string) error {
	idx, err := t.validarID(id)
	if err != nil {
		return err
	}
	(*t)[idx].Titulo = nuevoTitulo
	

	ahora := time.Now()
	(*t)[idx].FechaEdicion = &ahora
	

	return nil
}

func (t *Tareas) Imprimir() {
	if len(*t) == 0 {
		fmt.Println("No hay tareas para mostrar. ¡Agrega alguna!")
		return
	}

	sort.SliceStable(*t, func(i, j int) bool {
		return (*t)[i].ID < (*t)[j].ID
	})

	
	fmt.Printf("%-5s %-30s %-15s %-20s %-20s\n", "ID", "Título", "Estado", "Creada el", "Editada el")
	fmt.Println("----- ------------------------------ --------------- -------------------- --------------------")
	for _, tarea := range *t {
		
		
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
