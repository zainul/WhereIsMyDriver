package models

import (
	"WhereIsMyDriver/adapters"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Base ...
type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	ID        uint       `json:"id" gorm:"primary_key"`
}

var information = "DB"

// Connect base for make connection to database
func (b *Base) Connect() *gorm.DB {
	db, err := adapters.ConnectDB()
	CheckErr("error connect to database ", err)
	return db
}

// Find for find the record in database with some parameter as interface
func (b *Base) Find(v interface{}) (errDB error) {
	start := time.Now()
	db := b.Connect()
	defer db.Close()
	errDB = db.Find(v).Error

	b.LogTime(start, information)

	return
}

// FindOne for find the record in database with
// some parameter as interface and id
func (b *Base) FindOne(v interface{}, id int64) (notFound bool) {
	start := time.Now()
	db := b.Connect()
	defer db.Close()
	notFound = db.Find(v, id).RecordNotFound()
	b.LogTime(start, information)

	return
}

// Update for update some value to database with some parameter interface and id
func (b *Base) Update(name string, v interface{}, id int64) (errDB error) {
	db := b.Connect()
	defer db.Close()
	errDB = db.Table(name).Where("id = ?", id).Update(v).Error

	return
}

// Create for create record to database
func (b *Base) Create(v interface{}) (errDB error) {
	db := b.Connect()
	defer db.Close()
	errDB = db.Create(v).Error

	return
}

// Delete for deleting record to database
func (b *Base) Delete(v interface{}, id int64) (errDB error) {
	start := time.Now()
	db := b.Connect()
	defer db.Close()
	errDB = db.Delete(v, id).Error
	b.LogTime(start, information)

	return
}

// ToStruct use for assign interface to struct
func (b *Base) ToStruct(v interface{}, vDest interface{}) interface{} {
	byteData, err := json.Marshal(v)
	CheckErr("error while unmarshal struct ", err)
	json.Unmarshal(byteData, vDest)

	return vDest
}

// LogTime ...
func (b *Base) LogTime(start time.Time, information string) {
	elapsed := time.Since(start)
	fmt.Println("processing " + information + " time :=> " + elapsed.String())
}

//CheckErr ...
func CheckErr(msg string, err error) {
	if err != nil {
		log.Println(err)
	}
}
