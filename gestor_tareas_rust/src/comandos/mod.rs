pub mod agregar;
pub mod cambiar_estado;
pub mod eliminar;
pub mod listar;
pub mod editar;

#[cfg(test)]
mod tests {
    use super::*;
    use crate::tarea::{GestorTareas, Estado};
    use std::fs;
    use std::path::Path;
    use std::thread;
    use std::time::Duration;

    fn setup_test(test_name: &str) -> (GestorTareas, &Path) {
        let ruta = Path::new(test_name);
        if ruta.exists() {
            fs::remove_file(ruta).unwrap();
        }
        (GestorTareas { tareas: vec![], siguiente_id: 1 }, ruta)
    }

    fn cleanup(ruta: &Path) {
        if ruta.exists() {
            fs::remove_file(ruta).unwrap();
        }
    }

    #[test]
    fn test_agregar_tarea() {
        let (mut gestor, ruta) = setup_test("test_agregar.json");
        
        agregar::agregar_tarea(&mut gestor, "Test 1".to_string(), "Desc 1".to_string());
        
        assert_eq!(gestor.tareas.len(), 1);
        assert_eq!(gestor.tareas[0].id, 1);
        assert_eq!(gestor.siguiente_id, 2);
        assert_eq!(gestor.tareas[0].estado, Estado::Pendiente);

        cleanup(ruta);
    }

    #[test]
    fn test_ids_no_se_reutilizan() {
        let (mut gestor, ruta) = setup_test("test_ids_unicos.json");
        agregar::agregar_tarea(&mut gestor, "Tarea A".to_string(), "desc A".to_string());
        agregar::agregar_tarea(&mut gestor, "Tarea B".to_string(), "desc B".to_string());
        
        assert_eq!(gestor.tareas[1].id, 2);
        assert_eq!(gestor.siguiente_id, 3);

        eliminar::eliminar_tarea(&mut gestor, 1).unwrap();
        
        agregar::agregar_tarea(&mut gestor, "Tarea C".to_string(), "desc C".to_string());
        
        assert_eq!(gestor.tareas.len(), 2);
        assert_eq!(gestor.tareas[1].id, 3);
        assert_eq!(gestor.siguiente_id, 4);

        cleanup(ruta);
    }
    
    #[test]
    fn test_editar_tarea() {
        let (mut gestor, ruta) = setup_test("test_editar.json");
        agregar::agregar_tarea(&mut gestor, "Titulo viejo".to_string(), "Desc vieja".to_string());
        
        thread::sleep(Duration::from_secs(1));

        let resultado = editar::editar_tarea(&mut gestor, 1, "Titulo nuevo".to_string(), "Desc nueva".to_string());
        assert!(resultado.is_ok());

        assert_eq!(gestor.tareas[0].titulo, "Titulo nuevo");
        assert!(gestor.tareas[0].fecha_edicion.is_some());

        cleanup(ruta);
    }

    #[test]
    fn test_cambiar_estado_tarea() {
        let (mut gestor, ruta) = setup_test("test_cambiar_estado.json");
        agregar::agregar_tarea(&mut gestor, "Tarea a modificar".to_string(), "desc".to_string());
        
        cambiar_estado::cambiar_estado_tarea(&mut gestor, 1, "completada").unwrap();
        assert_eq!(gestor.tareas[0].estado, Estado::Completada);

        cleanup(ruta);
    }

    #[test]
    fn test_eliminar_tarea() {
        let (mut gestor, ruta) = setup_test("test_eliminar.json");
        agregar::agregar_tarea(&mut gestor, "Tarea A".to_string(), "desc A".to_string());
        agregar::agregar_tarea(&mut gestor, "Tarea B".to_string(), "desc B".to_string());
        
        eliminar::eliminar_tarea(&mut gestor, 1).unwrap();
        
        assert_eq!(gestor.tareas.len(), 1);
        assert_eq!(gestor.tareas[0].id, 2);

        cleanup(ruta);
    }
}