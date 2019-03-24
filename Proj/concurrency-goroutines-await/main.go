package main

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Note struct {
	gorm.Model
	Title     string
	Completed bool
	Creator   int
}

type User struct {
	gorm.Model
	Name string
}

func findNote(db *gorm.DB, id int) (*Note, error) {
	//new ở đây là new một con trỏ
	note := new(Note)
	err := db.Where("id =?", id).First(&note).Error
	return note, err
}

func findCreator(db *gorm.DB, id int) (*User, error) {
	creator := new(User)
	err := db.Where("id = ?", id).First(&creator).Error
	return creator, err
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/notes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Note{}, &User{})
	wg := new(sync.WaitGroup)
	noteID := 1
	creatorID := 1

	note := new(Note)
	wg.Add(1)
	go func() {
		defer wg.Done()
		note, _ = findNote(db, noteID)
	}()

	creator := new(User)

	wg.Add(1)
	go func() {
		defer wg.Done()
		creator, _ = findCreator(db, creatorID)
	}()

	wg.Wait()
	fmt.Println("note:", note)
	fmt.Println("creator:", creator)
}
