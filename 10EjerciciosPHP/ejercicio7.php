<?php
$texto = "Luis Chirivella no quiere estudiar PHP";
$contadorVocales = 0;
$longitud = strlen($texto);

for ($i = 0; $i < $longitud; $i++) {
    $caracter = strtolower($texto[$i]);

    switch ($caracter) {
        case 'a':
        case 'e':
        case 'i':
        case 'o':
        case 'u':
            $contadorVocales++;
            break;
    }
}

echo "Frase: '" . $texto . "'<br>";
echo "El numero total de vocales que tiene la frase es: " . $contadorVocales;
?>