<?php
$requisitos = [
    'php' => '7.4',
    'mysql' => '5.6',
    'mariadb' => '10.1'
];

// Chequear la versión de PHP
if (version_compare(PHP_VERSION, $requisitos['php'], '>=')) {
    echo "Versión de PHP: " . PHP_VERSION . " ✔️\n";
} else {
    echo "Versión de PHP: " . PHP_VERSION . " ❌ (Requiere {$requisitos['php']} o superior)\n";
}

// Chequear la extensión de MySQL
if (extension_loaded('mysqli')) {
    echo "Extensión de MySQL: Cargada ✔️\n";
    
    // Chequear la versión de MySQL
    $conexion = mysqli_connect('localhost', 'root', 'pass');
    if ($conexion) {
        $version_mysql = mysqli_get_server_info($conexion);
        if (version_compare($version_mysql, $requisitos['mysql'], '>=')) {
            echo "Versión de MySQL: " . $version_mysql . " ✔️\n";
        } else {
            echo "Versión de MySQL: " . $version_mysql . " ❌ (Requiere {$requisitos['mysql']} o superior)\n";
        }
    } else {
        echo "No se pudo conectar a MySQL ❌\n";
    }
} else {
    echo "Extensión de MySQL: No cargada ❌\n";
}

// Chequear la extensión de MariaDB si es necesario
// ...

// Chequear HTTPS
if (!empty($_SERVER['HTTPS']) && $_SERVER['HTTPS'] !== 'off') {
    echo "HTTPS: Habilitado ✔️\n";
} else {
    echo "HTTPS: No habilitado ❌ (Recomendado para una instalación segura)\n";
}

// Asegúrate de reemplazar 'tu_usuario' y 'tu_contraseña' con tus credenciales de MySQL.
?>
