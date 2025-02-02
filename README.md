# WeatherAPI

WeatherAPI es un servicio que proporciona previsiones meteorológicas y de olas para diferentes ciudades. También incluira un sistema de notificaciones para alertar a los usuarios sobre las condiciones meteorológicas.

## Características

- Obtener previsiones meteorológicas para ciudades
- Obtener previsiones de olas para ciudades costeras
- Gestión de usuarios (crear, leer, actualizar)

## Endpoints

### Endpoints de Clima

- `GET /weather/{cityID}`: Obtener previsión meteorológica para una ciudad
- `GET /waves/{cityID}/{day}`: Obtener previsión de olas para una ciudad en un día específico, siendo 0 hoy, 1 mañana y 2 pasado mañana.

### Endpoints de Usuarios

- `GET /users`: Obtener todos los usuarios
- `POST /users`: Crear un nuevo usuario
- `GET /users/{id}`: Obtener un usuario por ID
- `PUT /users/{id}`: Actualizar un usuario por ID

## Configuración

La aplicación utiliza un archivo de configuración (`config.yaml`) para configurar el servidor y la conexión a la base de datos. Ejemplo:

```yaml
SERVER_ADDRESS: ":8080"
DATABASE_URL: "user:password@tcp(localhost:3306)/weather"
```

## Ejecutar la Aplicación

1. Cargar la configuración:
    ```sh
    cp config.example.yaml config.yaml
    ```

2. Compilar y ejecutar la aplicación:
    ```sh
    go run cmd/myapp/main.go
    ```

## Arquitectura

La aplicación está estructurada de la siguiente manera:

- `cmd/myapp`: Punto de entrada de la aplicación
- `internal/config`: Carga de configuración
- `internal/database`: Conexión y configuración de la base de datos
- `internal/handlers`: Manejadores HTTP para diferentes rutas
- `internal/services`: Lógica de negocio y servicios
- `pkg/model`: Modelos de datos y operaciones de base de datos

## Diagrama

```
+-------------------+
|      Cliente      |
+--------+----------+
         |
         v
+--------+----------+
|    Servidor HTTP  |
| (Gorilla Mux)     |
+--------+----------+
         |
         v
+--------+----------+
|    Manejadores    |
| (weather.go,      |
|  users.go)        |
+--------+----------+
         |
         v
+--------+----------+
|    Servicios      |
| (notification.go) |
+--------+----------+
         |
         v
+--------+----------+
|    Modelos        |
| (users.go,        | 
| weahter.go)       |
+--------+----------+
         |
         v
+--------+----------+
|    Base de Datos  |
| (MySQL)           |
+-------------------+
```

## Dependencias

- `github.com/gorilla/mux`: Router para manejar solicitudes HTTP
- `github.com/go-sql-driver/mysql`: Driver MySQL para la conexión a la base de datos
- `golang.org/x/net/html/charset`: Lector de charset para la decodificación de XML
