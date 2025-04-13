package handlers

import (
	"github.com/gin-gonic/gin" //сервер
	"log"
	"net/http"
	"plane_rent_backend/db"     //бд
	"plane_rent_backend/models" //структуры-модели
)

func GetTypes(context *gin.Context) {
	var types []models.Type      //слайс пустой
	result := db.Db.Find(&types) //заполняем слайс кортежами из бд
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	context.JSON(http.StatusOK, types)
}
