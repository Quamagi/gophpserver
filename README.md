![logo](https://i3.wp.com/raw.githubusercontent.com/Quamagi/gophpserver/main/logo.jpg)
GoPhpServer
===========

# GoPhpServer

GoPhpServer es un servidor web escrito en Go que permite servir scripts PHP. Además de ejecutar los scripts PHP, también proporciona soporte para CORS (Cross-Origin Resource Sharing) y maneja diferentes métodos HTTP como GET, POST, PUT y DELETE.

## Descripción

El proyecto consta de varios archivos principales:

1. **main.go**: Este archivo contiene el código principal en Go que configura el servidor web y maneja las solicitudes HTTP entrantes. Utiliza el paquete `net/http` de Go para iniciar el servidor en `127.0.0.1:8080`. Cuando se recibe una solicitud, el servidor determina si se trata de la ruta raíz (`/`), la ruta `/check_mysql.php` o una ruta diferente, y luego llama a la función `serveScript` para ejecutar el script PHP correspondiente.

2. **index.php**: Este es un archivo PHP simple que se sirve cuando se accede a la ruta raíz (`/`). Actualmente, solo muestra un saludo "¡Hola Mundo!" y un mensaje indicando que el servidor funciona, pero que aún falta trabajo por hacer.

3. **check_mysql.php**: Este archivo PHP verifica si la extensión de MySQL está cargada en la instalación de PHP.

4. **util.php**: Este archivo PHP contiene funciones de utilidad para verificar los requisitos del sistema, como la versión de PHP, la extensión de MySQL y la disponibilidad de HTTPS.

## Funcionalidades

- **Ejecución de scripts PHP**: El servidor puede ejecutar scripts PHP utilizando el binario `php-cgi.exe`.
- **Soporte para CORS**: El servidor configura los encabezados necesarios para permitir solicitudes CORS desde cualquier origen.
- **Manejo de métodos HTTP**: El servidor admite los métodos HTTP GET, POST, PUT y DELETE. Para los métodos POST, PUT y DELETE, el cuerpo de la solicitud se pasa al script PHP a través de `cmd.Stdin`.

## Requisitos

- Go (versión compatible con el código)
- PHP (versión 7.4 o superior)
- MySQL (versión 5.6 o superior) o MariaDB (versión 10.1 o superior)
- Binario `php-cgi.exe` (generalmente ubicado en la carpeta de instalación de PHP)

## Instalación y uso

1. Clona el repositorio o descarga los archivos del proyecto.
2. Asegúrate de tener instalados Go, PHP y MySQL/MariaDB según los requisitos.
3. Actualiza la ruta del binario `php-cgi.exe` en `main.go` según la ubicación de tu instalación de PHP.
4. Abre una terminal, navega hasta el directorio del proyecto y ejecuta `go run main.go` para iniciar el servidor.
5. Abre un navegador web y visita `http://127.0.0.1:8080` para ver la página de inicio.
6. Para verificar si la extensión de MySQL está cargada, visita `http://127.0.0.1:8080/check_mysql.php`.

## Contribución

Si deseas contribuir a este proyecto, puedes abrir un issue o enviar una pull request en el repositorio de GitHub. Se aprecian todas las contribuciones, ya sean correcciones de errores, mejoras de características o documentación.

## Licencia

Este proyecto está licenciado bajo la [Licencia MIT](LICENSE).
