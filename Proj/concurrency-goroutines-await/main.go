package main

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type (
	Note struct {
		gorm.Model
		Id        int
		Name      string
		IdCreater string
	}
	
	User struct {
		gorm.Model
		Id   int
		Name string
	}
)

func getNote(db *gorm.DB, id int) (*Note, error) {
	note := new(Note)
	err := db.Where("id = ?", id).First(&note).Error

	return note, err
}

func getCreator(db *gorm.DB, id int) (*User, error) {
	user := new(User)
	err := db.Where("id = ?", id).First(&user).Error

	return user, err
}

func init() {
}

func main() {
	fmt.Println("Trần thanh thuận")
	db, err := gorm.Open("mysql", "root:123456@/notes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&Note{}, &User{})

	//noteID := 1
	creatorID := 1

	note := &Note{}
	user := &User{}

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		note, _ = getNote(db, creatorID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user, _ = getCreator(db, creatorID)
	}()

	wg.Wait()

	fmt.Println("note", note)
	fmt.Println("user", user)
	fmt.Println("abc", note.Name)
}
