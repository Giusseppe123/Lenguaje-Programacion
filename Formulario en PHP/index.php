<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Registro de Usuario Completo</title>
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
            max-width: 380px;
        }
        h2 {
            text-align: center;
            color: #333;
            margin-bottom: 25px;
        }
        .form-group {
            margin-bottom: 15px;
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
            border: 1px solid #f383c4ff;
            border-radius: 5px;
            box-sizing: border-box;
            transition: border-color 0.3s;
        }
        input:focus {
            border-color: #007bff;
            outline: none;
        }
        input:invalid:focus {
            border-color: #dc3545;
        }
        .error-message {
            color: #dc3545;
            font-size: 0.8em;
            margin-top: 5px;
            display: block;
            min-height: 1.2em;
        }
        button {
            width: 100%;
            background-color: #ba5695ff;
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
            background-color: #9f4a8bff;
        }
        .message { padding: 10px; margin-bottom: 20px; border-radius: 5px; text-align: center; width: 100%; max-width: 380px; box-sizing: border-box; }
        .success { background-color: #d4edda; color: #155724; border: 1px solid #c3e6cb; }
        .error { background-color: #f8d7da; color: #721c24; border: 1px solid #f5c6cb; }
        .registro-link { text-align: center; margin-top: 15px; font-size: 0.9em; }
        .registro-link a { color: #ca5d92ff; text-decoration: none; }
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
        <h2>Registro de Usuario</h2>
        
        <form id="registroForm" action="procesar.php" method="POST">
            
            <div class="form-group">
                <label for="nombre">Nombre Completo:</label>
                <input type="text" id="nombre" name="nombre" 
                       required 
                       minlength="3" 
                       maxlength="100" 
                       pattern="[A-Za-zñÑáéíóúÁÉÍÓÚ\s]+"
                       title="Solo letras y espacios. Debe tener al menos 3 caracteres."
                       autocomplete="name"
                       aria-label="Nombre completo del usuario">
                <span class="error-message" aria-live="polite"></span>
            </div>

            <div class="form-group">
                <label for="correo">Correo Electrónico:</label>
                <input type="email" id="correo" name="correo" 
                    required
                    minlength="3" 
                    maxlength="255"
                    pattern="^.{3,}@gmail\.com$" 
                    title="Debe usar el dominio @gmail.com y tener al menos 3 caracteres antes del @ (ej: usu@gmail.com)"
                    autocomplete="email"
                    aria-describedby="correo-hint">
                <span id="correo-hint" class="error-message"></span>
            </div>
            
            <div class="form-group">
                <label for="contraseña">Contraseña:</label>
                <input type="password" id="contraseña" name="contraseña" 
                       required
                       minlength="8"
                       maxlength="50"
                       pattern="(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+}{:&quot;?/>.<,;]).{8,}"
                       title="Mínimo 8 caracteres, incluyendo una mayúscula, una minúscula, un número y un símbolo especial."
                       autocomplete="new-password"
                       aria-describedby="contraseña-hint">
                <span id="contraseña-hint" class="error-message">Debe tener 8+ caracteres, Mayús, Minús, Número, Símbolo.</span>
            </div>

            <div class="form-group">
                <label for="confirmar_contraseña">Confirmar Contraseña:</label>
                <input type="password" id="confirmar_contraseña" name="confirmar_contraseña" 
                       required
                       minlength="8"
                       maxlength="50"
                       autocomplete="new-password"
                       aria-describedby="confirmar-contraseña-hint">
                <span id="confirmar-contraseña-hint" class="error-message"></span>
            </div>
            
            <button type="submit">Registrar mi Cuenta</button>
        </form>
        
        <p class="registro-link"><a href="login.php">¿Ya tienes cuenta? Inicia Sesión</a></p>
    </div>

    <script>
        document.getElementById('registroForm').addEventListener('submit', function(event) {
            const password = document.getElementById('contraseña').value;
            const confirmPassword = document.getElementById('confirmar_contraseña').value;
            const confirmError = document.getElementById('confirmar-contraseña-hint');
            
            const emailInput = document.getElementById('correo');
            const emailHint = document.getElementById('correo-hint');
            const emailPattern = new RegExp(emailInput.pattern);

           
            if (password !== confirmPassword) {
                event.preventDefault(); 
                confirmError.textContent = 'Las contraseñas no coinciden.';
                document.getElementById('confirmar_contraseña').focus();
            } else {
                confirmError.textContent = '';
            }

            if (emailInput.value.length < 3) {

                emailHint.textContent = emailInput.title;
            } else if (!emailPattern.test(emailInput.value)) {
                event.preventDefault();
                emailHint.textContent = emailInput.title;
            } else {
                emailHint.textContent = '';
            }
        });

        
        document.getElementById('confirmar_contraseña').addEventListener('input', function() {
            document.getElementById('confirmar-contraseña-hint').textContent = '';
        });
    </script>
</body>
</html>