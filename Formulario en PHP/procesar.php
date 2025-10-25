<?php

$registro_archivo = 'usuarios.json';

if (empty($_POST['nombre']) || empty($_POST['correo']) || empty($_POST['contraseña']) || empty($_POST['confirmar_contraseña'])) {
    header("Location: index.php?error=Datos incompletos. Llena todos los campos.");
    exit;
}

$nombre_usuario = $_POST['nombre'];
$email_usuario = $_POST['correo'];
$contraseña = $_POST['contraseña'];
$confirmacion = $_POST['confirmar_contraseña'];

if ($contraseña !== $confirmacion) {
    header("Location: index.php?error=Las contraseñas no coinciden.");
    exit;
}

$lista_usuarios = [];

if (file_exists($registro_archivo)) {
    $json_contenido = file_get_contents($registro_archivo);
    $usuarios_existentes = json_decode($json_contenido, true);

    if (is_array($usuarios_existentes)) {
        $lista_usuarios = $usuarios_existentes;
    }
}

foreach ($lista_usuarios as $usuario_existente) {
    if ($usuario_existente['email'] === $email_usuario) {
        header("Location: index.php?error=Ya existe un usuario registrado con este correo electrónico.");
        exit;
    }
}

$clave_hasheada = password_hash($contraseña, PASSWORD_DEFAULT);

$nuevo_registro = [
    'nombre' => $nombre_usuario,
    'email' => $email_usuario,
    'password_hash' => $clave_hasheada,
];

$lista_usuarios[] = $nuevo_registro;

$json_a_guardar = json_encode($lista_usuarios, JSON_PRETTY_PRINT);

if (file_put_contents($registro_archivo, $json_a_guardar) === false) {
    header("Location: index.php?error=error al guardar el archivo");
    exit;
}

header("Location: login.php?registered=true");
exit;