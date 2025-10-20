use crate::tarea::Tarea;
use crate::errores::ErrorApp;

pub fn eliminar_tarea(tareas: &mut Vec<Tarea>, id: u32) -> Result<(), ErrorApp> {
    let tamano_original = tareas.len();
    tareas.retain(|t| t.id != id);

    if tareas.len() < tamano_original {
        Ok(())
    } else {
        Err(ErrorApp::EntradaInvalida(format!("No se encontro una tarea con ID {}.", id)))
    }
}