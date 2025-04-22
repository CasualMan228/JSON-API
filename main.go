package main

import (
	"github.com/gin-gonic/gin"    //сервер
	"plane_rent_backend/db"       //бд
	"plane_rent_backend/handlers" //обработчики (пока один обработчик JSON)
)

func main() {
	db.Connect()                                  //создаем соединение с бд
	server := gin.Default()                       //создаем сервер
	server.GET("/api/planes", handlers.GetPlanes) //при обращении по адресу, триггернуть обработчик
	server.GET("/api/types", handlers.GetTypes)
	server.GET("/api/categories", handlers.GetCategories)
	server.GET("/api/photos", handlers.GetPhotos)
	server.GET("/api/users", handlers.GetUsers)
	server.GET("/api/roles", handlers.GetRoles)
	server.GET("/api/bills", handlers.GetBills)
	server.Run(":8080") //запуск сервера
}
