<?php
session_start();


if (!isset($_SESSION['autenticado']) || $_SESSION['autenticado'] !== true) {
    $_SESSION['error_login'] = "Necesitas iniciar sesión para acceder a esta página.";
    header("Location: login.php");
    exit;
}

$nombre = $_SESSION['usuario_nombre'] ?? 'Fan';
?>
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Omar Courtz - Dashboard</title>
    <style>
        @import url('https://fonts.googleapis.com/css2?family=Oswald:wght@700&family=Pacifico&display=swap');
        
        :root {
            --color-pink-light: #ffebf0;
            --color-pink-dark: #ffc2d1;
            --color-pink-primary: #ff69b4;
            --color-text-white: white;
            --color-text-dark: #333;
        }

        body {
            font-family: 'Oswald', sans-serif;
            background: linear-gradient(180deg, var(--color-pink-light) 0%, var(--color-pink-dark) 100%);
            color: var(--color-text-dark);
            text-align: center;
            min-height: 100vh;
            margin: 0;
            padding: 0;
            display: flex;
            flex-direction: column;
            align-items: center;
        }
        
       

        .hero {
            width: 100%;
            background-color: var(--color-pink-primary);
            padding: 50px 0;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
            position: relative;
        }

        .container_nav { 
            width: 90%;
            max-width: 1200px;
            margin: 0 auto;
            display: flex;
            justify-content: space-between;
            align-items: center;
            color: var(--color-text-white);
            font-weight: 700;
            padding-bottom: 20px;
        }
        
        .div-block-37 { 
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
        }
        
        .card_component-copy-copy { 
            background: rgba(255, 255, 255, 0.1);
            border-radius: 10px;
            padding: 10px;
            margin-bottom: 20px;
        }

        .main_tour_section {
            padding: 40px 20px;
            width: 100%;
            max-width: 1000px;
            margin: 0 auto;
        }
        


        .cursive-title {
            font-family: 'Pacifico', cursive;
            font-size: 50px;
            line-height: 0.8;
            color: var(--color-text-white);
            text-shadow: 0 0 10px rgba(255, 255, 255, 0.8);
            margin: 0;
        }
        
        .tour-text {
            font-size: 30px;
            font-weight: 900;
            letter-spacing: 10px;
            color: var(--color-text-white);
            text-shadow: 0 0 8px var(--color-pink-primary);
            margin-top: 10px;
            text-transform: uppercase;
        }

        .artist-placeholder {
            width: 150px;
            height: 150px;
            background-color: var(--color-pink-light);

            background-image: url('omar.jpg'); 
            background-size: cover;
            background-position: center top; 
            border-radius: 50%;
            margin: 20px auto;
            box-shadow: 0 0 10px rgba(255, 255, 255, 0.8);
        }

        .logout a {
            color: var(--color-text-dark);
            background-color: var(--color-text-white);
            padding: 10px 20px;
            border-radius: 5px;
            text-decoration: none;
            font-weight: 700;
        }
        .logout a:hover {
             background-color: #ddd;
        }
    </style>
</head>
<body>
    
    <div class="hero">
        <div class="container_nav">
            <div style="font-size: 1.2em;">OMAR COURTS</div>
            <div>
                <a href="logout.php" style="color: var(--color-text-white); margin-left: 10px;">Salir</a>
            </div>
        </div>
        
        <div class="div-block-37">
            <div class="sub-title-main" style="color: white; font-size: 1.5em; letter-spacing: 5px;">
                OMAR COURTS
            </div>
            
            <div class="card_component-copy-copy">
                <h1 class="cursive-title">Primera Musa</h1>
                
                <div class="artist-placeholder"></div> 

                <p class="tour-text">WORLD TOUR</p>
            </div>
        </div>
    </div>
    
    <div class="main_tour_section">
        
        <div class="div-block-38">
            <h2>Bienvenido, <?php echo htmlspecialchars($nombre); ?></h2>
            <p>Tu acceso exclusivo a las fechas de la gira LATAM 2025.</p>
        </div>
        
        <div class="div-block-29">
            <div class="section-large">
                <p style="font-size: 1.2em; font-weight: 700;">Próxima Cita:</p>
                <p> Valencia Venezuela 17 de octubre</p>
                <p> Caracas Venezuela 18 de octubre</p>
            </div>
        </div>
    </div>

    <div class="logout" style="margin-top: 50px;">
        <a href="logout.php">Cerrar Sesión</a>
    </div>
    
</body>
</html>