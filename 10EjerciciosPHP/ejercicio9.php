<?php
$n1 = 0;
$n2 = 1;
$limite = 30;

echo "<h3>Serie de Fibonacci (Primeros 30 numeros):</h3>";
echo $n1 . ", " . $n2;

for ($i = 3; $i <= $limite; $i++) {
    $siguiente = $n1 + $n2;
    echo ", " . $siguiente;

    $n1 = $n2;
    $n2 = $siguiente;
}
?>