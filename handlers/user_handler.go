package handlers

import (
	"github.com/gin-gonic/gin" //сервер
	"log"
	"net/http"
	"plane_rent_backend/db"     //бд
	"plane_rent_backend/models" //структуры-модели
)

func GetUsers(context *gin.Context) {
	var users []models.User      //слайс пустой
	result := db.Db.Find(&users) //заполняем слайс кортежами из бд
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	context.JSON(http.StatusOK, users)
}
