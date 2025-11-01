use std::fs;
use std::io;
use std::path::Path;
use crate::tarea::GestorTareas;
use crate::errores::ErrorApp;

pub fn cargar_gestor(ruta: &Path) -> Result<GestorTareas, ErrorApp> {
    match fs::read_to_string(ruta) {
        Ok(contenido) => {
            let gestor: GestorTareas = serde_json::from_str(&contenido)?;
            Ok(gestor)
        }
        Err(e) if e.kind() == io::ErrorKind::NotFound => {
            Ok(GestorTareas { tareas: vec![], siguiente_id: 1 })
        }
        Err(e) => Err(ErrorApp::from(e)),
    }
}

pub fn guardar_gestor(gestor: &GestorTareas, ruta: &Path) -> Result<(), ErrorApp> {
    let json = serde_json::to_string_pretty(gestor)?;
    fs::write(ruta, json)?;
    Ok(())
}