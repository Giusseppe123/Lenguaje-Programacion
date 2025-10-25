<?php
session_start();

$archivo_json = 'usuarios.json';

if (empty($_POST['email']) || empty($_POST['contraseña'])) {
    $_SESSION['error_login'] = "Ingresa tu correo y contraseña.";
    header("Location: login.php");
    exit;
}

$email_ingresado = $_POST['email'];
$contraseña_ingresada = $_POST['contraseña'];
$usuario_encontrado = false;

if (file_exists($archivo_json)) {
    $datos_json = file_get_contents($archivo_json);
    $lista_usuarios = json_decode($datos_json, true);

    if (is_array($lista_usuarios)) {
        foreach ($lista_usuarios as $usuario) {
            if ($usuario['email'] === $email_ingresado) {
                if (password_verify($contraseña_ingresada, $usuario['password_hash'])) {
                    $usuario_encontrado = true;
                    
                    $_SESSION['autenticado'] = true;
                    $_SESSION['usuario_nombre'] = $usuario['nombre'];
                    
                    header("Location: omar.php");
                    exit;
                }
            }
        }
    }
}

$_SESSION['error_login'] = "Correo o contraseña incorrectos.";
header("Location: login.php");
exit;