use serde::{Serialize, Deserialize};
use chrono::{DateTime, Local};
use std::fmt;

#[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
pub enum Estado {
    Pendiente,
    EnProgreso,
    Completada,
}

impl fmt::Display for Estado {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Estado::Pendiente => write!(f, "Pendiente"),
            Estado::EnProgreso => write!(f, "En Progreso"),
            Estado::Completada => write!(f, "Completada"),
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
pub struct Tarea {
    pub id: u32,
    pub titulo: String,
    pub descripcion: String,
    pub estado: Estado,
    pub fecha_creacion: DateTime<Local>,
    pub fecha_edicion: Option<DateTime<Local>>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct GestorTareas {
    pub tareas: Vec<Tarea>,
    pub siguiente_id: u32,
}