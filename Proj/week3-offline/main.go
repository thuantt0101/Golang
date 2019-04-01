package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Note struct {
	gorm.Model
	Title     string
	Completed bool
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/notes1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Note{})

	//r stand for router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(201, "pong\n")
	})

	r.GET("/note", func(c *gin.Context) {
		note := Note{}
		c.BindJSON(&note)

		db.Create(&note)
		c.JSON(200, note)
	})

	r.GET("/note/:id", func(c *gin.Context) {
		id := c.Param("id")
		note := Note{}
		db.Where("id=?", id).First(&note)
		c.JSON(200, note)
	})

	r.Run(":8081")
}
