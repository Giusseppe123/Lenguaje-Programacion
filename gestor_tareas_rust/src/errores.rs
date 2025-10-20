use std::fmt;

#[derive(Debug)]
pub enum ErrorApp {
    ErrorIO(std::io::Error),
    ErrorJSON(serde_json::Error),
    EntradaInvalida(String),
}

impl fmt::Display for ErrorApp {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            ErrorApp::ErrorIO(e) => write!(f, "Error de archivo: {}", e),
            ErrorApp::ErrorJSON(e) => write!(f, "Error de formato JSON: {}", e),
            ErrorApp::EntradaInvalida(msg) => write!(f, "Error de entrada: {}", msg),
        }
    }
}

impl std::error::Error for ErrorApp {}

impl From<std::io::Error> for ErrorApp {
    fn from(err: std::io::Error) -> ErrorApp {
        ErrorApp::ErrorIO(err)
    }
}

impl From<serde_json::Error> for ErrorApp {
    fn from(err: serde_json::Error) -> ErrorApp {
        ErrorApp::ErrorJSON(err)
    }
}