GoPhpServer
GoPhpServer es un servidor web escrito en Go que permite servir scripts PHP. Básicamente, funciona como un intermediario entre las solicitudes HTTP entrantes y la ejecución de scripts PHP en el servidor.

Descripción
El proyecto consta de tres archivos principales:

main.go: Este archivo contiene el código principal en Go que configura el servidor web y maneja las solicitudes HTTP entrantes. Utiliza el paquete net/http de Go para iniciar el servidor en 127.0.0.1:8080. Cuando se recibe una solicitud, el servidor determina si se trata de la ruta raíz (/) o una ruta diferente, y luego llama a la función serveScript para ejecutar el script PHP correspondiente.
index.php: Este es un archivo PHP simple que se sirve cuando se accede a la ruta raíz (/). Actualmente, solo muestra un saludo "¡Hola Mundo!" y un mensaje indicando que el servidor funciona, pero que aún falta trabajo por hacer.
util.php: Este archivo PHP contiene funciones de utilidad para verificar los requisitos del sistema, como la versión de PHP, la extensión de MySQL y la disponibilidad de HTTPS. Actualmente, está configurado para verificar las versiones de PHP 7.4, MySQL 5.6 y MariaDB 10.1, pero estas versiones se pueden ajustar según las necesidades del proyecto.
Funcionamiento
Cuando se inicia el servidor Go (go run main.go), comienza a escuchar en 127.0.0.1:8080. Cuando se recibe una solicitud HTTP, el servidor determina la ruta solicitada y ejecuta el script PHP correspondiente utilizando el binario php-cgi.exe. La salida del script PHP se envía de vuelta al cliente como respuesta HTTP.

Requisitos
Go (versión compatible con el código)
PHP (versión 7.4 o superior)
MySQL (versión 5.6 o superior) o MariaDB (versión 10.1 o superior)
Binario php-cgi.exe (generalmente ubicado en la carpeta de instalación de PHP)
Instalación y uso
Clona el repositorio o descarga los archivos del proyecto.
Asegúrate de tener instalados Go, PHP y MySQL/MariaDB según los requisitos.
Actualiza la ruta del binario php-cgi.exe en main.go según la ubicación de tu instalación de PHP.
Abre una terminal, navega hasta el directorio del proyecto y ejecuta go run main.go para iniciar el servidor.
Abre un navegador web y visita http://127.0.0.1:8080 para ver la página de inicio.
Contribución
Si deseas contribuir a este proyecto, puedes abrir un issue o enviar una pull request en el repositorio de GitHub. Se aprecian todas las contribuciones, ya sean correcciones de errores, mejoras de características o documentación.

Licencia
Este proyecto está licenciado bajo la Licencia MIT.
