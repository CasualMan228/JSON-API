package db

import (
	"gorm.io/driver/sqlserver" //ssms
	"gorm.io/gorm"             //работа со структурами-моделями
	"log"                      //выводит ошибку И ЗАВЕРШАЕТ ПРОГРАММУ
)

var Db *gorm.DB

func Connect() {
	dbCurrent := "sqlserver://localhost:1433?database=PlaneRent" //УКАЗЫВАЙ СВОЮ БД!!!
	connection, err := gorm.Open(sqlserver.Open(dbCurrent), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db = connection //зафиксировать соединение для других пакетов
}
