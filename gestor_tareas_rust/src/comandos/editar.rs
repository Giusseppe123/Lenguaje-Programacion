use crate::tarea::GestorTareas;
use crate::errores::ErrorApp;
use chrono::Local;

pub fn editar_tarea(gestor: &mut GestorTareas, id: u32, nuevo_titulo: String, nueva_descripcion: String) -> Result<(), ErrorApp> {
    if let Some(tarea) = gestor.tareas.iter_mut().find(|t| t.id == id) {
        tarea.titulo = nuevo_titulo;
        tarea.descripcion = nueva_descripcion;
        tarea.fecha_edicion = Some(Local::now());
        Ok(())
    } else {
        Err(ErrorApp::EntradaInvalida(format!("No se encontro una tarea con ID {}.", id)))
    }
}