package main

import (
	"ZeroDay_Go/src/estudiantes/infraestructure"
	maestrosInfra "ZeroDay_Go/src/maestros/infraestructure"
	materias "ZeroDay_Go/src/materias/infraestructure"
	serverZeroday "ZeroDay_Go/src/server_ZeroDay"
	server "ZeroDay_Go/src/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear el router principal de Gin
	r := gin.Default()

	// Inicializar rutas del polling ZeroDay (ya no crea su propio router)
	serverZeroday.Run(nil)

	// Inicializar rutas del polling maestro/alumno
	server.Init(r)

	// Inicializar rutas de otras áreas
	infraestructure.Init(r)
	maestrosInfra.Init(r)
	materias.Init(r)

	// Ejecutar el servidor en el puerto 8080
	r.Run(":8080")
}
