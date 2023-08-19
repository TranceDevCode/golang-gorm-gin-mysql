package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Esta funcion va a retornar un objeto db
var Database = func() (db *gorm.DB) {
	//Se valida la existencia de .env y variables de entorno
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		return
	}
	//creamos una variable para dsn
	//dsn := "root::@tcp(localhost:3306)/golang_mux_gorm?charset+utf8mb4&parseTime=true&loc=Local"
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset+utf8mb4&parseTime=true&loc=Local"

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Erro de conexion")
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}()
