# BackEnd-Hotel-App-v2

Este es el back-end para la aplicación web del proyecto final de quinto semestre de la Fundación Universitaria Tecnológico Comfenalco.

## Descripción del Proyecto

Este proyecto es una aplicación de gestión de hotel desarrollada en Go. Proporciona una API RESTful para manejar operaciones relacionadas con usuarios, reservas, habitaciones y otros aspectos de la gestión hotelera.

## Características Principales

- Gestión de usuarios
- Manejo de reservas
- Administración de habitaciones
- Operaciones CRUD para varias entidades

## Tecnologías Utilizadas

- Go (Golang)
- GORM (ORM para Go)
- Base de datos (MySQL)
- Gorilla Mux (para enrutamiento HTTP)

## Estructura del Proyecto

El proyecto sigue una arquitectura de capas:

- `Pkg/Controllers`: Maneja las solicitudes HTTP y respuestas
- `Pkg/Services`: Contiene la lógica de negocio
- `Pkg/Services/Impl`: Implementaciones concretas de los servicios
- `Pkg/Services/Interfaces`: Interfaces de los servicios
- `Pkg/Entities`: Define las estructuras de datos principales
- `Pkg/Database`: Configuración y conexión a la base de datos

## Instalación y Configuración

#1. Clona el repositorio:
   ```
   git clone https://github.com/tu-usuario/BackEnd-Hotel-App-v2.git
   ```
#2. Navega al directorio del proyecto:
   ```
   cd BackEnd-Hotel-App-v2
   ```
#3. Instala las dependencias:
   ```
   go mod tidy
   ```
#4. Configura las variables de entorno (si es necesario)
#5. Ejecuta la aplicación:
   ```
   go run main.go
   ```

## Contacto

Eduardo Rodriguez - Solanoe934@gmail.com


Enlace del proyecto: [https://github.com/EduRoDev/BackEnd-Hotel-App-v2](https://github.com/EduRoDev/BackEnd-Hotel-App-v2)
