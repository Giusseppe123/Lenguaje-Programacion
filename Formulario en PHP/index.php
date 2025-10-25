<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Registro de Usuario</title>
    <style>
        /* Estilos que ya tenías */
        body {
            font-family: Arial, sans-serif;
            background-color: #e9ecef;
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
            border: 1px solid #ced4da;
            border-radius: 5px;
            box-sizing: border-box;
            transition: border-color 0.3s;
        }
        input:focus {
            border-color: #007bff;
            outline: none;
        }
        /* Resaltar campos que fallen la validación de formato */
        input:invalid:focus {
            border-color: #dc3545; /* Rojo */
        }
        button {
            width: 100%;
            background-color: #007bff;
            color: white;
            padding: 12px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            margin-top: 10px;
            transition: background-color 0.3s;
        }
        button:hover {
            background-color: #0056b3;
        }
        /* Estilos para los mensajes de PHP */
        .message {
            padding: 10px;
            margin-bottom: 20px;
            border-radius: 5px;
            text-align: center;
            width: 100%;
            max-width: 350px;
            box-sizing: border-box;
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
    </style>
</head>
<body>
    
    <?php
    if (isset($_GET['success'])) {
        echo '<p class="message success">Registro completado con éxito.</p>';
    }
    if (isset($_GET['error'])) {
        echo '<p class="message error">ERROR: ' . htmlspecialchars($_GET['error']) . '</p>';
    }
    ?>

    <div class="container">
        <h2>Formulario de Registro de Usuario</h2>
        
        <form action="procesar.php" method="POST">
            
            <label for="nombre">Nombre:</label>
            <input type="text" id="nombre" name="nombre" required>

            <label for="correo">Correo Electrónico:</label>
            <input type="text" id="correo" name="correo" 
                   required
                   pattern="[a-zA-Z0-9._%+-]+@gmail\.com$"
                   title="El correo debe ser una dirección válida de Gmail (ej: nombre.apellido@gmail.com)">
            
            <label for="contraseña">Contraseña:</label>
            <input type="password" id="contraseña" name="contraseña" 
                   required
                   pattern="(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}"
                   title="La contraseña debe tener al menos 8 caracteres, incluyendo una mayúscula, una minúscula y un número.">
            
            <button type="submit">Registrar</button>
        </form>
    </div>
</body>
</html>