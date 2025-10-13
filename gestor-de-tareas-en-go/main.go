package main

import (
	"fmt"
	"os"
)

func main() {
	almacenamientoTareas := NuevoAlmacenamiento[Tareas]("tareas.json")

	var tareas Tareas
	err := almacenamientoTareas.Cargar(&tareas)
	if err != nil {
		fmt.Printf("error al guardar las tareas: %v\n", err)
		os.Exit(1)
	}

	indicadoresCmd := NuevosIndicadoresComando()

	indicadoresCmd.Ejecutar(&tareas, almacenamientoTareas)
}