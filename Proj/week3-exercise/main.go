package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"

	"github.com/Golang/Proj/week3-exercise/model"
)

func main() {

	//0.Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//1. Lien quan tới Database
	db, err := gorm.Open("mysql", "root:123456@/notes1?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&model.Note{})

	//2. Write access log ra file & de giu lai cai Println -->stdout

	fileWrite, err := os.Create("access.log")

	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = fileWrite

	//3. Tạo ra router
	///r := gin.Default()
	//handler

}
