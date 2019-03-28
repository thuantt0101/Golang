package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/k0kubun/pp"
)

//Base model definition gorm.Model
//including fields ID, CreatedAt,
//including fields ID, CreatedAt,
//you could embed it in your model,
// or only write those fields you want
// Base Model's definition
//type Model struct {
//	ID        uint `gorm:"primary_key"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time
// }
//Table name is the pluralized(số nhiều) version of struct name
type User struct { // default table name is `users`
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` //default size for string is 255
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard //One-To-One relationship	(has one)
	Emails     []Email    //One-To-Many relationship(has many)

	BillingAddress   Address //One-To-One relationship
	BillingAddressID sql.NullInt64

	ShippingAddress   Address //One-To-One relationship
	ShippingAddressID int

	IgnoreMe int        `gorm:"-"` //Ignore this field
	Language []Language `gorm:"many2many:user_languages"`
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // Foreign key(belong to),tag `index`
	Email      string `gorm:"type:varchar(100);unique_index"` //`type` set sc
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` //set field as not null
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` //Create index with name
	Code string `gorm:"index:idx_name_code"` //unique_index also works
}

func main() {
	//
	pp.Println("Start Model Definition")
	db, err := gorm.Open("mysql", "root:123456@/db001?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&Language{}, &User{}, &Address{}, &Email{}, &CreditCard{})

	//check if exists table
	if db.HasTable(&User{}) {
		pp.Println("User is exists in table model")
	}

	//check if exists table in db
	if db.HasTable("users") {
		pp.Println("users table is exists in database")
	}

	pp.Println("start create record")
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	if db.NewRecord(user) { //=>returns 'true' as primary key is blank
		fmt.Println(user.Name)
	}

	db.Create(&user)
	db.NewRecord(user) // => return `false` after `user` created
}
