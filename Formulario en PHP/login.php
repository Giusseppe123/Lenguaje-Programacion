<?php
session_start();
?>
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Inicio de Sesion</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #d0608fff;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            padding: 20px;
        }
        .container {
            background-color: #ffffff;
            padding: 30px 40px;
            border-radius: 10px;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
            width: 100%;
            max-width: 350px;
        }
        h2 {
            text-align: center;
            color: #333;
            margin-bottom: 25px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            color: #555;
            font-weight: bold;
        }
        input {
            width: 100%;
            padding: 10px;
            margin-bottom: 15px;
            border: 1px solid #e10080ff;
            border-radius: 5px;
            box-sizing: border-box;
        }
        button {
            width: 100%;
            background-color: #ac398cff;
            color: white;
            padding: 12px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            margin-top: 10px;
        }
        button:hover {
            background-color: #bd5396ff;
        }
        .message {
            padding: 10px;
            margin-bottom: 20px;
            border-radius: 5px;
            text-align: center;
            width: 100%;
            max-width: 350px;
        }
        .success {
            background-color: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
        }
        .error {
            background-color: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
        }
        .registro-link {
            text-align: center;
            margin-top: 15px;
            font-size: 0.9em;
        }
        .registro-link a {
            color: #a94a7eff;
            text-decoration: none;
        }
    </style>
</head>
<body>
    
    <?php
    if (isset($_SESSION['error_login'])) {
        echo '<p class="message error">' . htmlspecialchars($_SESSION['error_login']) . '</p>';
        unset($_SESSION['error_login']);
    }
    if (isset($_GET['registered'])) {
        echo '<p class="message success">¡Registro completado! Por favor, inicia sesión.</p>';
    }
    ?>

    <div class="container">
        <h2>Iniciar Sesion</h2>
        
        <form action="autenticar.php" method="POST">
            
            <label for="email">Correo Electrónico:</label>
            <input type="text" id="email" name="email" required>
            
            <label for="contraseña">Contraseña:</label>
            <input type="password" id="contraseña" name="contraseña" required>
            
            <button type="submit">Entrar</button>
        </form>
        
        <p class="registro-link"><a href="index.php">¿No tienes cuenta? Registrate</a></p>
    </div>
</body>
</html>