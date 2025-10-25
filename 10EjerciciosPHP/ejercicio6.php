<?php
$tamano = 5;

echo "Cuadrado de " . $tamano . "x" . $tamano . ":<br>";

for ($fila = 1; $fila <= $tamano; $fila++) {
    for ($columna = 1; $columna <= $tamano; $columna++) {
        echo "* ";
    }
    echo "<br>";
}
?>