package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestAgregarTarea(t *testing.T) {
	tareas := Tareas{}
	tareas.Agregar("Comprar leche")

	if len(tareas) != 1 || tareas[0].Titulo != "Comprar leche" || tareas[0].Estado != Pendiente {
		t.Errorf("Fallo AgregarTarea: se esperaba 1 tarea 'Comprar leche' en estado Pendiente, pero se obtuvo %+v", tareas)
	}
}

func TestEliminarTarea(t *testing.T) {
	tareas := Tareas{
		{ID: 0, Titulo: "Tarea 0"},
		{ID: 1, Titulo: "Tarea 1"},
	}

	err := tareas.Eliminar(0)
	if err != nil {
		t.Errorf("Fallo EliminarTarea para ID 0: %v", err)
	}
	if len(tareas) != 1 || tareas[0].ID != 1 {
		t.Errorf("Fallo EliminarTarea: se esperaba que quedara 1 tarea con ID 1, pero se obtuvo %+v", tareas)
	}

	err = tareas.Eliminar(99)
	if err == nil || err.Error() != "ID de tarea no encontrado" {
		t.Errorf("Fallo EliminarTarea para un ID que no existe: se esperaba un error, pero se obtuvo '%v'", err)
	}
}

func TestAlternarEstado(t *testing.T) {
	tareas := Tareas{{ID: 0, Titulo: "Tarea 0", Estado: Pendiente}}

	tareas.AlternarEstado(0)
	if tareas[0].Estado != EnProgreso {
		t.Error("Fallo AlternarEstado: se esperaba EnProgreso")
	}

	tareas.AlternarEstado(0)
	if tareas[0].Estado != Completada {
		t.Error("Fallo AlternarEstado: se esperaba Completada")
	}
}

func TestEditarTarea(t *testing.T) {
	tareas := Tareas{{ID: 0, Titulo: "Titulo antiguo"}}

	err := tareas.Editar(0, "Titulo nuevo")
	if err != nil {
		t.Errorf("Fallo EditarTarea: %v", err)
	}
	if tareas[0].Titulo != "Titulo nuevo" {
		t.Errorf("Fallo EditarTarea: se esperaba 'Titulo nuevo', pero se obtuvo '%s'", tareas[0].Titulo)
	}
	if tareas[0].FechaEdicion == nil {
		t.Error("Fallo EditarTarea: la fecha de edicion no se actualizo y es nil")
	}
}

func TestEstadoTareaJSON(t *testing.T) {
	var estado EstadoTarea = Completada
	bytes, err := json.Marshal(estado)
	if err != nil {
		t.Fatalf("Fallo Marshal: %v", err)
	}
	if string(bytes) != "\"Completada\"" {
		t.Errorf("Fallo Marshal: se esperaba '\"Completada\"', pero se obtuvo '%s'", string(bytes))
	}

	var nuevoEstado EstadoTarea
	err = json.Unmarshal([]byte("\"En Progreso\""), &nuevoEstado)
	if err != nil {
		t.Fatalf("Fallo Unmarshal: %v", err)
	}
	if nuevoEstado != EnProgreso {
		t.Errorf("Fallo Unmarshal: se esperaba EnProgreso, pero se obtuvo %s", nuevoEstado.String())
	}
}

func TestAlmacenamiento(t *testing.T) {
	nombreArchivoTest := "test_tareas.json"
	almacenamiento := NuevoAlmacenamiento[Tareas](nombreArchivoTest)
	defer os.Remove(nombreArchivoTest)

	tareasIniciales := Tareas{{ID: 0, Titulo: "Prueba", Estado: Pendiente, FechaCreacion: time.Now()}}

	err := almacenamiento.Guardar(tareasIniciales)
	if err != nil {
		t.Fatalf("Fallo al guardar: %v", err)
	}

	var tareasLeidas Tareas
	err = almacenamiento.Cargar(&tareasLeidas)
	if err != nil {
		t.Fatalf("Fallo al cargar: %v", err)
	}

	// Truncamos para evitar problemas de precision con nanosegundos
	tareasIniciales[0].FechaCreacion = tareasIniciales[0].FechaCreacion.Truncate(time.Second)
	tareasLeidas[0].FechaCreacion = tareasLeidas[0].FechaCreacion.Truncate(time.Second)

	if !reflect.DeepEqual(tareasIniciales, tareasLeidas) {
		t.Error("Las tareas leidas del archivo no son iguales a las iniciales.")
	}
}