package adapters

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// only for dialect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBCredential ...
type DBCredential struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
}

var dbCredential DBCredential

// ConnectDB use for connect to db with credential and return db gorm
func ConnectDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql",
		dbCredential.User+":"+dbCredential.Password+"@tcp("+
			dbCredential.Host+":"+dbCredential.Port+")/"+dbCredential.DBName+
			"?charset=utf8&parseTime=True&loc=Local")
	CheckErr("Error while connect to database ", err)
	db = db.Debug()

	return
}

func init() {
	dbCredential.DBName = os.Getenv("DB_NAME")
	dbCredential.Host = os.Getenv("DB_HOST")
	dbCredential.Password = os.Getenv("DB_PASSWORD")
	dbCredential.Port = os.Getenv("DB_PORT")
	dbCredential.User = os.Getenv("DB_USER")
}

// CheckErr is use for checking general error 
func CheckErr(msg string, err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
	}
}
