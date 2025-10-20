use crate::tarea::Tarea;

pub fn listar_tareas(tareas: &Vec<Tarea>) {
    if tareas.is_empty() {
        println!("No hay tareas pendientes.");
        return;
    }
    
    println!("{:-<105}", "");
    println!("{:^5} | {:^20} | {:^15} | {:^25} | {:^25}", "ID", "Titulo", "Estado", "Fecha de Creacion", "Fecha de Edicion");
    println!("{:-<105}", "");
    for tarea in tareas {
        let fecha_edicion_str = match tarea.fecha_edicion {
            Some(fecha) => fecha.format("%Y-%m-%d %H:%M").to_string(),
            None => "-".to_string(),
        };

        println!(
            "{:^5} | {:<20} | {:^15} | {:<25} | {:<25}",
            tarea.id,
            tarea.titulo,
            tarea.estado,
            tarea.fecha_creacion.format("%Y-%m-%d %H:%M"),
            fecha_edicion_str,
        );
    }
    println!("{:-<105}", "");
}