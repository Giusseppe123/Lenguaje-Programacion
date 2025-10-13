package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IndicadoresComando struct {
	Agregar         string
	Eliminar        int
	Editar          string
	Alternar        int
	Listar          bool
	EstablecerEstado string
	Ayuda           bool
}

func NuevosIndicadoresComando() *IndicadoresComando {
	ic := &IndicadoresComando{}

	flag.StringVar(&ic.Agregar, "agregar", "", "agrega una nueva tarea ejemplo: -agregar \"Título\"")
	flag.IntVar(&ic.Eliminar, "eliminar", -1, "elimina una tarea por su ID por ejemplo: -eliminar <ID>")
	flag.StringVar(&ic.Editar, "editar", "", "edita el título de una tarea por ejemplo: -editar \"<ID>:Nuevo Título\"")
	flag.IntVar(&ic.Alternar, "alternar", -1, "cambia el estado de una tarea (ciclo). por ejemplo: -alternar <ID>")
	flag.BoolVar(&ic.Listar, "listar", false, "muestra todas las tareas.")
	flag.StringVar(&ic.EstablecerEstado, "establecer-estado", "", "establece un estado específico. por ejemplo: -establecer-estado \"<ID>:Estado\"")
	flag.BoolVar(&ic.Ayuda, "ayuda", false, "muestra los comandos disponibles.")

	flag.Parse()
	return ic
}

func (ic *IndicadoresComando) Ejecutar(tareas *Tareas, almacenamiento *Almacenamiento[Tareas]) {
	if ic.Ayuda || (ic.Agregar == "" && ic.Eliminar == -1 && ic.Editar == "" && ic.Alternar == -1 && !ic.Listar && ic.EstablecerEstado == "") {
		fmt.Println("")
		fmt.Println("Gestor de Tareas en Go")
		fmt.Println("\ncomandos disponibles:")
		flag.PrintDefaults()
		return
	}

	var err error
	ejecutado := false

	switch {
	case ic.Agregar != "":
		tareas.Agregar(ic.Agregar)
		fmt.Printf(" Tarea \"%s\" agregada.\n", ic.Agregar)
		ejecutado = true
	case ic.Eliminar != -1:
		err = tareas.Eliminar(ic.Eliminar)
		if err != nil {
			fmt.Printf(" Error al eliminar tarea (ID %d): %v\n", ic.Eliminar, err)
		} else {
			fmt.Printf(" Tarea con ID %d eliminada.\n", ic.Eliminar)
		}
		ejecutado = true
	case ic.Editar != "":
		partes := strings.SplitN(ic.Editar, ":", 2)
		if len(partes) != 2 {
			fmt.Println(" formato invalido. Su forma es: -editar \"<ID>:Nuevo Título\"")
			os.Exit(1)
		}
		id, errParse := strconv.Atoi(partes[0])
		if errParse != nil {
			fmt.Printf(" ID de tarea invalido: %v\n", errParse)
			os.Exit(1)
		}
		nuevoTitulo := partes[1]
		err = tareas.Editar(id, nuevoTitulo)
		if err != nil {
			fmt.Printf(" Error al editar tarea (ID %d): %v\n", id, err)
		} else {
			fmt.Printf("Tarea con ID %d actualizada.\n", id)
		}
		ejecutado = true
	case ic.Alternar != -1:
		err = tareas.AlternarEstado(ic.Alternar)
		if err != nil {
			fmt.Printf(" Error al cambiar estado (ID %d): %v\n", ic.Alternar, err)
		} else {
			fmt.Printf(" Estado de la tarea con ID %d cambiado.\n", ic.Alternar)
		}
		ejecutado = true
	case ic.EstablecerEstado != "":
		partes := strings.SplitN(ic.EstablecerEstado, ":", 2)
		if len(partes) != 2 {
			fmt.Println(" Formato invalido. Su forma es: -establecer-estado \"<ID>:Estado\"")
			os.Exit(1)
		}
		id, errParse := strconv.Atoi(partes[0])
		if errParse != nil {
			fmt.Printf(" ID de tarea invalido: %v\n", errParse)
			os.Exit(1)
		}
		estadoStr := strings.TrimSpace(partes[1])
		var nuevoEstado EstadoTarea
		errUnmarshal := nuevoEstado.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", estadoStr)))
		if errUnmarshal != nil {
			fmt.Printf(" Estado inválido: \"%s\". Usa: Pendiente, En Progreso, Completada.\n", estadoStr)
			os.Exit(1)
		}

		idx, errValidar := tareas.validarID(id)
		if errValidar != nil {
			fmt.Printf(" Error al establecer estado (ID %d): %v\n", id, errValidar)
		} else {
			(*tareas)[idx].Estado = nuevoEstado
			fmt.Printf(" Estado de la tarea con ID %d establecido a \"%s\".\n", id, nuevoEstado.String())
		}
		ejecutado = true
	case ic.Listar:
		tareas.Imprimir()
		ejecutado = true
	}

	if ejecutado {
		errGuardar := almacenamiento.Guardar(*tareas)
		if errGuardar != nil {
			fmt.Printf(" Advertencia: Error al guardar tareas: %v\n", errGuardar)
		}
	}
}