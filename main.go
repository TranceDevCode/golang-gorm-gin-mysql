package main

import (
	"log"
	"os"

	"gotenancy/routes"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	//Inicializacion de conexion a la base de datos
	//db := config.InitializeDB()
	//db.AutoMigrate(&CustomerModel{})
	//Creacion de la instancia de Gin Gonic
	router := gin.Default()

	//cargas las rutas definidas en el archivo routes/routes.go
	routes.InitializeRoutes(router)

	//Inicializamos el servidor en el puerto que definimos en el .env
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		return
	}
	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
