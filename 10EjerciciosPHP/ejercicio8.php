<?php
$inventario = [
    "Leche" => 15,
    "Pan" => 30,
    "Cafe" => 8,
    "Harina Pan" => 10,
    "Chocolate" => 4
];

echo "<h3>Reporte de Inventario:</h3>";

foreach ($inventario as $producto => $stock) {
    echo "Producto: " . $producto . " - Stock: " . $stock;

    if ($stock <= 10) {
        echo " ALERTA: Stock bajo";
    }

    echo "<br>";
}
?>