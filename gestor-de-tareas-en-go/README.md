Como Usarlo
Todos los comandos se ejecutan usando go run . seguido de la bandera correspondiente.

Listar todas las tareas

Muestra una tabla con todas tus tareas, sus estados y fechas.



go run . -listar
Agregar una nueva tarea

Las nuevas tareas se agregan por defecto en estado "Pendiente".



go run . -agregar "Hacer la compra semanal"
Editar el titulo de una tarea

Usa el formato "ID:Nuevo Titulo". Esto tambien actualizara la fecha de edicion.



go run . -editar "1:Preparar la cena para la familia"
Eliminar una tarea



Elimina una tarea permanentemente usando su ID.


go run . -eliminar 0
Alternar el estado de una tarea

Cambia el estado de una tarea en un ciclo (Pendiente -> En Progreso -> Completada -> Pendiente).



go run . -alternar 1
Establecer un estado especifico

Asigna directamente un estado a una tarea.



go run . -establecer-estado "2:Completada"


Obtener ayuda
Muestra la lista de todos los comandos disponibles.

go run . -ayuda



Ejecutar las Pruebas
Para verificar que todo funcione como se espera, puedes correr las pruebas unitarias desde la raiz del proyecto.

go test ./...