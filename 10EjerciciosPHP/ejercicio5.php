<?php
$nota1 = 9.5;
$nota2 = 12.0;
$nota3 = 10.0;

$promedio = ($nota1 + $nota2 + $nota3) / 3;

echo "Promedio Obtenido: " . round($promedio, 2) . "<br>";

if ($promedio >= 9) {
    echo "Aprobado con Distincion";
} elseif ($promedio >= 6) {
    echo "Aprobado";
} else {
    echo "Reprobado";
}
?>