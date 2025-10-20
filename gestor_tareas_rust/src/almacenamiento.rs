use std::fs;
use std::io;
use std::path::Path;
use crate::tarea::Tarea;
use crate::errores::ErrorApp;

pub fn cargar_tareas(ruta: &Path) -> Result<Vec<Tarea>, ErrorApp> {
    match fs::read_to_string(ruta) {
        Ok(contenido) => {
            let tareas: Vec<Tarea> = serde_json::from_str(&contenido)?;
            Ok(tareas)
        }
        Err(e) if e.kind() == io::ErrorKind::NotFound => {
            Ok(vec![])
        }
        Err(e) => Err(ErrorApp::from(e)),
    }
}

pub fn guardar_tareas(tareas: &Vec<Tarea>, ruta: &Path) -> Result<(), ErrorApp> {
    let json = serde_json::to_string_pretty(tareas)?;
    fs::write(ruta, json)?;
    Ok(())
}