package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Almacenamiento[T any] struct {
	NombreArchivo string
}

func NuevoAlmacenamiento[T any](nombreArchivo string) *Almacenamiento[T] {
	return &Almacenamiento[T]{
		NombreArchivo: nombreArchivo,
	}
}

func (a *Almacenamiento[T]) Guardar(datos T) error {
	datosArchivo, err := json.MarshalIndent(datos, "", "    ")
	if err != nil {
		return fmt.Errorf("error al serializar datos a JSON: %w", err)
	}

	err = ioutil.WriteFile(a.NombreArchivo, datosArchivo, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir datos en el archivo %s: %w", a.NombreArchivo, err)
	}
	return nil
}

func (a *Almacenamiento[T]) Cargar(datos *T) error {
	datosArchivo, err := ioutil.ReadFile(a.NombreArchivo)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("error al leer el archivo %s: %w", a.NombreArchivo, err)
	}

	err = json.Unmarshal(datosArchivo, datos)
	if err != nil {
		return fmt.Errorf("error al deserializar JSON del archivo %s: %w", a.NombreArchivo, err)
	}
	return nil
}