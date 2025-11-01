use std::env;
use std::path::Path;

mod errores;
mod tarea;
mod almacenamiento;
mod comandos;

use errores::ErrorApp;
use almacenamiento::{cargar_gestor, guardar_gestor};
use comandos::{
    agregar::agregar_tarea,
    listar::listar_tareas,
    cambiar_estado::cambiar_estado_tarea,
    eliminar::eliminar_tarea,
    editar::editar_tarea,
};

fn imprimir_ayuda() {
    println!("\nGestor de Tareas CLI en Rust");
    println!("Uso: cargo run <comando> [argumentos]\n");
    println!("Comandos:");
    println!("  agregar   \"<titulo>\" \"<descripcion>\"       - Agrega una nueva tarea.");
    println!("  listar                                    - Muestra todas las tareas.");
    println!("  estado    <ID> <estado>                   - Cambia el estado de una tarea.");
    println!("                                              (Estados: pendiente, enprogreso, completada)");
    println!("  editar    <ID> \"<titulo>\" \"<desc>\"        - Edita el titulo y descripcion de una tarea.");
    println!("  eliminar  <ID>                            - Elimina una tarea.");
}

fn ejecutar() -> Result<(), ErrorApp> {
    let ruta_archivo = Path::new("tareas.json");
    let mut gestor = cargar_gestor(ruta_archivo)?;
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        imprimir_ayuda();
        return Ok(());
    }

    let comando = &args[1];

    match comando.as_str() {
        "agregar" => {
            let titulo = args.get(2).cloned().ok_or_else(|| ErrorApp::EntradaInvalida("Falta el titulo.".to_string()))?;
            let descripcion = args.get(3).cloned().ok_or_else(|| ErrorApp::EntradaInvalida("Falta la descripcion.".to_string()))?;
            agregar_tarea(&mut gestor, titulo.clone(), descripcion);
            guardar_gestor(&gestor, ruta_archivo)?;
            println!("Tarea '{}' agregada con exito.", titulo);
        }
        "listar" => {
            listar_tareas(&gestor);
        }
        "estado" => {
            let id_str = args.get(2).ok_or_else(|| ErrorApp::EntradaInvalida("Falta el ID de la tarea.".to_string()))?;
            let id = id_str.parse::<u32>().map_err(|_| ErrorApp::EntradaInvalida("El ID debe ser un numero.".to_string()))?;
            let nuevo_estado = args.get(3).ok_or_else(|| ErrorApp::EntradaInvalida("Falta el nuevo estado.".to_string()))?;
            
            cambiar_estado_tarea(&mut gestor, id, nuevo_estado)?;
            guardar_gestor(&gestor, ruta_archivo)?;
            println!("Estado de la tarea con ID {} actualizado a '{}'.", id, nuevo_estado);
        }
        "editar" => {
            let id_str = args.get(2).ok_or_else(|| ErrorApp::EntradaInvalida("Falta el ID de la tarea.".to_string()))?;
            let id = id_str.parse::<u32>().map_err(|_| ErrorApp::EntradaInvalida("El ID debe ser un numero.".to_string()))?;
            let nuevo_titulo = args.get(3).cloned().ok_or_else(|| ErrorApp::EntradaInvalida("Falta el nuevo titulo.".to_string()))?;
            let nueva_descripcion = args.get(4).cloned().ok_or_else(|| ErrorApp::EntradaInvalida("Falta la nueva descripcion.".to_string()))?;

            editar_tarea(&mut gestor, id, nuevo_titulo, nueva_descripcion)?;
            guardar_gestor(&gestor, ruta_archivo)?;
            println!("Tarea con ID {} editada con exito.", id);
        }
        "eliminar" => {
            let id_str = args.get(2).ok_or_else(|| ErrorApp::EntradaInvalida("Falta el ID de la tarea.".to_string()))?;
            let id = id_str.parse::<u32>().map_err(|_| ErrorApp::EntradaInvalida("El ID debe ser un numero.".to_string()))?;
            eliminar_tarea(&mut gestor, id)?;
            guardar_gestor(&gestor, ruta_archivo)?;
            println!("Tarea con ID {} eliminada.", id);
        }
        _ => {
            println!("Comando no reconocido.");
            imprimir_ayuda();
        }
    }
    Ok(())
}

fn main() {
    if let Err(e) = ejecutar() {
        eprintln!("\nError en la aplicacion: {}", e);
        std::process::exit(1);
    }
}