use crate::tarea::{Tarea, Estado};
use chrono::Local;

pub fn agregar_tarea(tareas: &mut Vec<Tarea>, titulo: String, descripcion: String) {
    let nuevo_id = tareas.iter().map(|t| t.id).max().unwrap_or(0) + 1;

    let nueva_tarea = Tarea {
        id: nuevo_id,
        titulo,
        descripcion,
        estado: Estado::Pendiente,
        fecha_creacion: Local::now(),
        fecha_edicion: None,
    };
    tareas.push(nueva_tarea);
}