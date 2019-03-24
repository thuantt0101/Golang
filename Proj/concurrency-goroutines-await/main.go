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
	noteID    int
}

type User struct {
	gorm.Model
	Id   int
	Name string
}

//Trả về con trỏ với kiểu dữ liệu là Struct Note và biến error
func findNote(db *gorm.DB, id int) (*Note, error) {
	//new ở đây là new một con trỏ
	note := new(Note)
	//Toán tử & được sử dụng để lấy địa chỉ của một biến
	//.First: find note
	//tìm note trong db với id và gán cho con trỏ note
	// Get first matched record

	err := db.Where("noteID=?", id).First(&note).Error
	//// SELECT * FROM note WHERE noteID = id limit 1;

	return note, err
}

func findCreator(db *gorm.DB, id int) (*User, error) {
	creator := new(User)
	err := db.Where("Id = ?", id).First(&creator).Error
	return creator, err
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/notes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&Note{}, &User{})

	wg := new(sync.WaitGroup)
	//noteID := 1
	creatorID := 1

	//note := new(Note)
	//wg.Add(1)

	//go rountine
	//go func() {
	//	defer wg.Done()
	//	note, _ = findNote(db, noteID)
	//}()

	creator := new(User)

	wg.Add(1)
	go func() {
		defer wg.Done()
		creator, _ = findCreator(db, creatorID)
	}()

	wg.Wait()

	//fmt.Println("note:", note)

	fmt.Println("creator:", creator)

	if creator == nil {
		fmt.Println("creator is nil")
	}

	fmt.Println("name:", creator.Name)

}
