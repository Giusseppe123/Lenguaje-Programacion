pub mod agregar;
pub mod cambiar_estado;
pub mod eliminar;
pub mod listar;
pub mod editar;

#[cfg(test)]
mod tests {
    use super::*;
    use crate::tarea::{Tarea, Estado};
    use std::fs;
    use std::path::Path;
    use std::thread;
    use std::time::Duration;

    fn setup_test(test_name: &str) -> (Vec<Tarea>, &Path) {
        let ruta = Path::new(test_name);
        if ruta.exists() {
            fs::remove_file(ruta).unwrap();
        }
        (Vec::new(), ruta)
    }

    fn cleanup(ruta: &Path) {
        if ruta.exists() {
            fs::remove_file(ruta).unwrap();
        }
    }

    #[test]
    fn test_agregar_tarea() {
        let (mut tareas, ruta) = setup_test("test_agregar.json");
        
        agregar::agregar_tarea(&mut tareas, "Test 1".to_string(), "Desc 1".to_string());
        
        assert_eq!(tareas.len(), 1);
        assert_eq!(tareas[0].id, 1);
        assert_eq!(tareas[0].estado, Estado::Pendiente);

        cleanup(ruta);
    }
    
    #[test]
    fn test_editar_tarea() {
        let (mut tareas, ruta) = setup_test("test_editar.json");
        agregar::agregar_tarea(&mut tareas, "Titulo viejo".to_string(), "Desc vieja".to_string());
        
        thread::sleep(Duration::from_secs(1));

        let resultado = editar::editar_tarea(&mut tareas, 1, "Titulo nuevo".to_string(), "Desc nueva".to_string());
        assert!(resultado.is_ok());

        assert_eq!(tareas[0].titulo, "Titulo nuevo");
        assert_eq!(tareas[0].descripcion, "Desc nueva");
        assert!(tareas[0].fecha_edicion.is_some());
        assert!(tareas[0].fecha_edicion.unwrap() > tareas[0].fecha_creacion);

        cleanup(ruta);
    }

    #[test]
    fn test_cambiar_estado_tarea() {
        let (mut tareas, ruta) = setup_test("test_cambiar_estado.json");
        agregar::agregar_tarea(&mut tareas, "Tarea a modificar".to_string(), "desc".to_string());
        
        cambiar_estado::cambiar_estado_tarea(&mut tareas, 1, "completada").unwrap();
        assert_eq!(tareas[0].estado, Estado::Completada);

        cleanup(ruta);
    }

    #[test]
    fn test_eliminar_tarea() {
        let (mut tareas, ruta) = setup_test("test_eliminar.json");
        agregar::agregar_tarea(&mut tareas, "Tarea A".to_string(), "desc A".to_string());
        agregar::agregar_tarea(&mut tareas, "Tarea B".to_string(), "desc B".to_string());
        
        eliminar::eliminar_tarea(&mut tareas, 1).unwrap();
        
        assert_eq!(tareas.len(), 1);
        assert_eq!(tareas[0].id, 2);

        cleanup(ruta);
    }
}