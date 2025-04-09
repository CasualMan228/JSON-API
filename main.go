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
	server.Run(":8080")                           //запуск сервера
}
