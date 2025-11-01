use crate::tarea::{Estado, GestorTareas};
use crate::errores::ErrorApp;

fn parsear_estado(s: &str) -> Result<Estado, ErrorApp> {
    match s.to_lowercase().as_str() {
        "pendiente" => Ok(Estado::Pendiente),
        "enprogreso" => Ok(Estado::EnProgreso),
        "completada" => Ok(Estado::Completada),
        _ => Err(ErrorApp::EntradaInvalida(format!("Estado '{}' no valido. Opciones: pendiente, enprogreso, completada.", s))),
    }
}

pub fn cambiar_estado_tarea(gestor: &mut GestorTareas, id: u32, nuevo_estado_str: &str) -> Result<(), ErrorApp> {
    let nuevo_estado = parsear_estado(nuevo_estado_str)?;
    
    if let Some(tarea) = gestor.tareas.iter_mut().find(|t| t.id == id) {
        tarea.estado = nuevo_estado;
        Ok(())
    } else {
        Err(ErrorApp::EntradaInvalida(format!("No se encontro una tarea con ID {}.", id)))
    }
}