package main

import (
	"fmt"
	"os"
)

func main() {
	almacenamientoTareas := NuevoAlmacenamiento[GestorDeTareas]("tareas.json")

	var gestor GestorDeTareas
	err := almacenamientoTareas.Cargar(&gestor)
	if err != nil {
		fmt.Printf("Error fatal al cargar tareas: %v\n", err)
		os.Exit(1)
	}
	
	if gestor.SiguienteID == 0 && len(gestor.Tareas) > 0 {
		maxID := -1
		for _, tarea := range gestor.Tareas {
			if tarea.ID > maxID {
				maxID = tarea.ID
			}
		}
		gestor.SiguienteID = maxID + 1
	}

	indicadoresCmd := NuevosIndicadoresComando()

	indicadoresCmd.Ejecutar(&gestor, almacenamientoTareas)
}