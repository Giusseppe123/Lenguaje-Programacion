use crate::tarea::{Tarea, Estado, GestorTareas};
use chrono::Local;

pub fn agregar_tarea(gestor: &mut GestorTareas, titulo: String, descripcion: String) {
    let nuevo_id = gestor.siguiente_id;
    gestor.siguiente_id += 1;

    let nueva_tarea = Tarea {
        id: nuevo_id,
        titulo,
        descripcion,
        estado: Estado::Pendiente,
        fecha_creacion: Local::now(),
        fecha_edicion: None,
    };
    gestor.tareas.push(nueva_tarea);
}