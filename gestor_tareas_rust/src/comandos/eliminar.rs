use crate::tarea::GestorTareas;
use crate::errores::ErrorApp;

pub fn eliminar_tarea(gestor: &mut GestorTareas, id: u32) -> Result<(), ErrorApp> {
    let tamano_original = gestor.tareas.len();
    gestor.tareas.retain(|t| t.id != id);

    if gestor.tareas.len() < tamano_original {
        Ok(())
    } else {
        Err(ErrorApp::EntradaInvalida(format!("No se encontro una tarea con ID {}.", id)))
    }
}