package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestAgregarTarea(t *testing.T) {
	gestor := GestorDeTareas{}
	gestor.Agregar("Comprar leche")

	if len(gestor.Tareas) != 1 {
		t.Fatalf("Fallo AgregarTarea: se esperaba 1 tarea, pero hay %d", len(gestor.Tareas))
	}
	if gestor.Tareas[0].ID != 0 {
		t.Errorf("Fallo AgregarTarea: se esperaba ID 0, pero se obtuvo %d", gestor.Tareas[0].ID)
	}
	if gestor.SiguienteID != 1 {
		t.Errorf("Fallo AgregarTarea: se esperaba SiguienteID 1, pero se obtuvo %d", gestor.SiguienteID)
	}
}

func TestEliminarTarea(t *testing.T) {
	gestor := GestorDeTareas{
		Tareas: Tareas{
			{ID: 0, Titulo: "Tarea 0"},
			{ID: 1, Titulo: "Tarea 1"},
		},
		SiguienteID: 2,
	}

	err := gestor.Eliminar(0)
	if err != nil {
		t.Errorf("Fallo EliminarTarea para ID 0: %v", err)
	}
	if len(gestor.Tareas) != 1 || gestor.Tareas[0].ID != 1 {
		t.Errorf("Fallo EliminarTarea: se esperaba que quedara 1 tarea con ID 1, pero se obtuvo %+v", gestor.Tareas)
	}
}

func TestIDsUnicosDespuesDeEliminar(t *testing.T) {
	gestor := GestorDeTareas{}
	gestor.Agregar("Tarea A") 
	gestor.Agregar("Tarea B") 
	gestor.Agregar("Tarea C") 

	
	if gestor.SiguienteID != 3 {
		t.Fatalf("Fallo preparacion: se esperaba SiguienteID 3, pero es %d", gestor.SiguienteID)
	}

	
	err := gestor.Eliminar(1)
	if err != nil {
		t.Fatalf("Fallo al eliminar tarea para la prueba: %v", err)
	}

	
	gestor.Agregar("Tarea D")

	
	nuevaTarea := gestor.Tareas[len(gestor.Tareas)-1]
	if nuevaTarea.ID != 3 {
		t.Errorf("Fallo ID unico: el nuevo ID deberia ser 3, pero es %d", nuevaTarea.ID)
	}
	if gestor.SiguienteID != 4 {
		t.Errorf("Fallo ID unico: SiguienteID deberia ser 4, pero es %d", gestor.SiguienteID)
	}
}

func TestAlternarEstado(t *testing.T) {
	gestor := GestorDeTareas{
		Tareas:      Tareas{{ID: 0, Titulo: "Tarea 0", Estado: Pendiente}},
		SiguienteID: 1,
	}

	gestor.AlternarEstado(0)
	if gestor.Tareas[0].Estado != EnProgreso {
		t.Error("Fallo AlternarEstado: se esperaba EnProgreso")
	}

	gestor.AlternarEstado(0)
	if gestor.Tareas[0].Estado != Completada {
		t.Error("Fallo AlternarEstado: se esperaba Completada")
	}
}

func TestEditarTarea(t *testing.T) {
	gestor := GestorDeTareas{
		Tareas:      Tareas{{ID: 0, Titulo: "Titulo antiguo"}},
		SiguienteID: 1,
	}

	err := gestor.Editar(0, "Titulo nuevo")
	if err != nil {
		t.Errorf("Fallo EditarTarea: %v", err)
	}
	if gestor.Tareas[0].Titulo != "Titulo nuevo" {
		t.Errorf("Fallo EditarTarea: se esperaba 'Titulo nuevo', pero se obtuvo '%s'", gestor.Tareas[0].Titulo)
	}
	if gestor.Tareas[0].FechaEdicion == nil {
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
}

func TestAlmacenamiento(t *testing.T) {
	nombreArchivoTest := "test_tareas.json"
	almacenamiento := NuevoAlmacenamiento[GestorDeTareas](nombreArchivoTest)
	defer os.Remove(nombreArchivoTest)

	gestorInicial := GestorDeTareas{
		Tareas: Tareas{
			{ID: 0, Titulo: "Prueba", Estado: Pendiente, FechaCreacion: time.Now()},
		},
		SiguienteID: 1,
	}

	err := almacenamiento.Guardar(gestorInicial)
	if err != nil {
		t.Fatalf("Fallo al guardar: %v", err)
	}

	var gestorLeido GestorDeTareas
	err = almacenamiento.Cargar(&gestorLeido)
	if err != nil {
		t.Fatalf("Fallo al cargar: %v", err)
	}

	
	gestorInicial.Tareas[0].FechaCreacion = gestorInicial.Tareas[0].FechaCreacion.Truncate(time.Second)
	gestorLeido.Tareas[0].FechaCreacion = gestorLeido.Tareas[0].FechaCreacion.Truncate(time.Second)

	if !reflect.DeepEqual(gestorInicial, gestorLeido) {
		t.Error("El gestor leido del archivo no es igual al inicial.")
	}
}