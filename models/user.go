package models

import (
	"WhereIsMyDriver/helper"
	"WhereIsMyDriver/structs"
	"time"

	"github.com/go-sql-driver/mysql"
	validator "gopkg.in/go-playground/validator.v9"
)

// UserTableName ...
const UserTableName = "users"

// User ...
type User struct {
	Email             string  `gorm:"not null;unique_index" validate:"required,email"`
	Username          string  `gorm:"not null;unique" validate:"required,gte=6"`
	Password          string  `gorm:"not null;type:varchar(1000)" validate:"required,gte=6"`
	Phone             string  `gorm:"not null;unique" validate:"required,gte=6"`
	Photo             *string `gorm:"type:varchar(1000)"`
	IdentifiedNumber  string
	Token             *string
	FullName          string
	FirstName         string
	LastName          string
	TokenConfirmation *string
	ConfirmationAt    mysql.NullTime
	CurrentLatitude   float32 `gorm:"not null" sql:"type:decimal(9,6);"`
	CurrentLongitude  float32 `gorm:"not null" sql:"type:decimal(9,6);"`
	CurrentAccuracy   float32 `gorm:"not null"`
	Base
}

// TableName ...
func (u *User) TableName() string {
	return UserTableName
}

// RuleValidation ...
func (u *User) RuleValidation() (errors []string) {
	validate := validator.New()
	err := validate.Struct(u)
	errors = structs.MapValidation(err)
	return
}

// AddUser use for saving user
func (u *User) AddUser(v interface{}, errors *[]string) {
	errDB := u.Create(v)
	helper.CheckError("failed save to database", errDB)

	if errDB != nil {
		(*errors) = append(*errors, "Failed save user to database")
	}
}

func (u *User) SetDefault() {
	u.Base.CreatedAt = time.Now()
	u.Base.UpdatedAt = time.Now()
}
