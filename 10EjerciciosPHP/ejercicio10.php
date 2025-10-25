<?php
$numero = 7;
$contador = 1;

echo "<h3>Tabla de Multiplicar del " . $numero . "</h3>";

while ($contador <= 10) {
    $resultado = $numero * $contador;
    echo $numero . " X " . $contador . " = " . $resultado . "<br>";
    $contador++;
}
?>