package handlers

import (
	"github.com/gin-gonic/gin" //сервер
	"log"
	"net/http"
	"plane_rent_backend/db"     //бд
	"plane_rent_backend/models" //структуры-модели
)

func GetPlanes(context *gin.Context) {
	var planes []models.Plane     //слайс пустой
	result := db.Db.Find(&planes) //заполняем слайс кортежами из бд
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	context.JSON(http.StatusOK, planes)
}
