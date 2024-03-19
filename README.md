## Pre-requisitos
Antes de comenzar, asegúrate de tener instalado lo siguiente en tu sistema:

- Node.js: Necesario para ejecutar el servidor auxiliar gRPC-front. [Descargar Node.js](https://nodejs.org/)
- Go: Requerido para el backend y la API de integración gRPC. [Descargar Go](https://go.dev/doc/install)
- MariaDB: Sistema de gestión de base de datos para el repositorio. Instalar MariaDB con XAMPP [Descargar XAMPP](https://www.apachefriends.org/es/index.html)
- gRPC: Necesario para la comunicación entre servicios. Guía de Instalación de gRPC [Instalación de gRPC](https://grpc.io/docs/languages/go/quickstart/)

## Configuración de la base de datos:
- Importar las sentencias SQL para la base de datos en MariaDB

## Guia de compilación

Dentro del código implementado vemos varios archivos, tanto de front, como de backend. El archivo del front se encuentra en la carpeta llamada **preactexample/portafolio** y tenemos un servidor proxy en Node.js dentro de la carpeta **servidornodejs**. Las demás carpetas estan relacionadas con el Backend en Golang. Para correr el código realizar lo siguiente:

- Instalar las siguientes librerias de gRpc **go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28** y **go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2**


