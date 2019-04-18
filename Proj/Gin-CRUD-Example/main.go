package main

import (
	"net/http"

	model "github.com/Golang/Proj/Gin-CRUD-Example/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/k0kubun/pp"
)

//Xay dung ung dung truyen client server CRUD thong qua json

func main() {
	pp.Println("Start main form")

	db, err := gorm.Open("mysql", "root:123456@/notes?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Note1{})

	//create router
	r := gin.Default()

	//Return a string drive into respond body
	r.GET("/ping", func(c *gin.Context) {

		//String write the driven string to the respond body(from server to client)
		c.String(201, "Pong\n")
	})

	//Create
	//1.1 Create a note
	note1 := model.Note1{
		Title:     "Today is the second day of my life in hcm city",
		Completed: true,
	}
	//2.Create
	r.GET("/Create", func(c *gin.Context) {
		//Insert data into DB
		db.Save(&note1)
		//gin.H : Create a map
		//c.JSON: create a JSON with driven struct
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
			"Message": "Todo item created successfully!",
			"ID":      note1.ID})
	})

	//Read all notes
	r.GET("/ReadAll", func(c *gin.Context) {
		//create 1 slice note
		var notes []model.Note1
		var _notes []model.TranformNote1

		db.Find(&notes)
		if len(notes) <= 0 {
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusNotFound,
				"message": "no row found"})
			return
		}

		//transforms the notes for building a good response

		for _, item := range notes {
			completed := false
			if item.Completed == true {
				completed = true
			}
			_notes = append(_notes, model.TranformNote1{ID: item.ID, Title: item.Title, Completed: completed})
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _notes})
	})

	//Read one Note
	r.GET("/ReadOne/:id", func(c *gin.Context) {
		var note2 model.Note1
		noteID := c.Param("id")

		//find along with ID
		db.First(&note2, noteID)
		if note2.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound,
				"messgae": "no note found!"})
			return
		}
		completed := false
		if note2.Completed == true {
			completed = true
		}

		_note2 := model.TranformNote1{ID: note2.ID, Title: note2.Title, Completed: completed}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _note2})
	})

	//Update
	r.PUT("/Update/:id", func(c *gin.Context) {
		var note3 model.Note1
		noteID := c.Param("id")

		db.First(&note3, noteID)
		if note3.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no note found"})
			return
		}

		//Update
		//update notes set title  = "Hello my first updated"
		db.Model(&note3).Update("title", "Hello my first updated")
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "note update successfully"})
	})

	//Delete
	r.DELETE("/Delete/:id", func(c *gin.Context) {
		var note4 model.Note1
		noteid4 := c.Param("id")

		db.First(&note4, noteid4)

		if note4.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "note not found"})
			return
		}

		db.Delete(&note4)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "delete successfully"})
	})

	r.Run(":8081")
}
