<?php
$servername = "127.0.0.1";
$username = "root";
$password = "pass";
$database = "nombre_de_tu_base_de_datos"; // Reemplaza con el nombre de tu base de datos

// Crear conexión
$conn = new mysqli($servername, $username, $password, $database);

// Verificar conexión
if ($conn->connect_error) {
    die("Conexión fallida: " . $conn->connect_error);
}
echo "Conexión exitosa";
$conn->close();
?>