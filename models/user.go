package models

import (
	"WhereIsMyDriver/adapters"
	"WhereIsMyDriver/helper"
	"WhereIsMyDriver/structs"
	"WhereIsMyDriver/structs/api"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/context"

	"github.com/go-sql-driver/mysql"
	validator "gopkg.in/go-playground/validator.v9"
)

// UserTableName ...
const UserTableName = "users"

// User ...
type User struct {
	ID                uint    `json:"id" gorm:"primary_key"`
	CurrentLatitude   float32 `gorm:"not null;primary_key" sql:"type:decimal(9,6);index"`
	CurrentLongitude  float32 `gorm:"not null;primary_key" sql:"type:decimal(9,6);index"`
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
	CurrentAccuracy   float32    `gorm:"not null"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at" sql:"index"`
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
	var base = new(Base)
	errDB := base.Create(v)
	helper.CheckError("failed save to database", errDB)

	if errDB != nil {
		(*errors) = append(*errors, "Failed save user to database")
	}
}

// SetDefault use for set default value if created and updated time
func (u *User) SetDefault() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// UpdateNewPositionDriver use for update the driver position
func (u *User) UpdateNewPositionDriver(
	position HistoryPosition,
	errStr *[]string,
) {
	position.SetDefault()

	db, err := adapters.ConnectDB()
	defer closeDB(db)
	helper.CheckError("error connect to database", err)

	tx := db.Begin()
	var user = User{
		CurrentLatitude:  position.Latitude,
		CurrentLongitude: position.Longitude,
		CurrentAccuracy:  position.Accuracy,
	}

	if position.UserID < 0 || position.UserID > (50*1000) {
		tx.Rollback()
		(*errStr) = append(*errStr, "The driver ID is invalid")
		return
	}

	if err := tx.Table(u.TableName()).
		Where("id = ?", position.UserID).
		Update(user).Error; err != nil {
		tx.Rollback()
		(*errStr) = append(*errStr, "failed update current location driver")
		return
	}

	if err := tx.Create(&position).Error; err != nil {
		tx.Rollback()
		(*errStr) = append(*errStr, "failed saved new position")
		errClose := db.Close()
		helper.CheckError("error close db", errClose)
		return
	}
	log.Println(errStr)

	tx.Commit()

}

// GetDriver is use for get driver location base on user location and filter
// by radius
func (u *User) GetDriver(
	latitude float32,
	longitude float32,
	radius int,
	limit int,
) []api.GetDriver {
	var count int
	db, err := adapters.ConnectDB()
	helper.CheckError("failed connect to database", err)
	db.Table(u.TableName()).Count(&count)
	errClose := db.Close()
	helper.CheckError("failed close db", errClose)

	diff := float64(count) / float64(10)
	sizeDiff := int(math.Ceil(diff))

	drivers := splitGetData(
		latitude,
		longitude,
		radius,
		limit,
		count,
		sizeDiff,
		u.TableName(),
	)

	return drivers
}

func splitGetData(
	latitude float32,
	longitude float32,
	radius int,
	limit int,
	totalRecord int,
	size int,
	tableName string,
) []api.GetDriver {

	var res []api.GetDriver

	diff := float64(totalRecord) / float64(size)
	page := int(math.Ceil(diff))
	var startID = 0
	var endID = size

	c := make(chan []api.GetDriver)

	for index := 1; index <= page; index++ {
		go func(
			latitude float32,
			longitude float32,
			radius int,
			limit int,
			tableName string,
			startID int,
			endID int,
			co chan<- []api.GetDriver,
		) {
			co <- getDriver(
				latitude,
				longitude,
				radius,
				limit,
				tableName,
				startID,
				endID,
			)
		}(latitude, longitude, radius, limit, tableName, startID, endID, c)

		startID = endID + 1
		endID = ((index + 1) * size)
	}

	for index := 1; index <= page; index++ {
		var drivers []api.GetDriver
		drivers = <-c
		(res) = append(res, drivers...)
	}

	sortByDistance(&res)
	recordDriverWithLimit := make([]api.GetDriver, 0)

	if len(res) >= limit {
		for j := 0; j < limit; j++ {
			recordDriverWithLimit = append(recordDriverWithLimit, res[j])
		}
	} else {
		recordDriverWithLimit = append(recordDriverWithLimit, res...)
	}

	return recordDriverWithLimit
}

func sortByDistance(res *[]api.GetDriver) {
	sort.Slice(*res, func(i, j int) bool {
		return (*res)[i].Distance < (*res)[i].Distance
	})
}

func getDriver(
	latitude float32,
	longitude float32,
	radius int,
	limit int,
	tableName string,
	startID int,
	endID int,
) []api.GetDriver {
	var drivers []api.GetDriver
	db, err := adapters.ConnectDB()
	defer closeDB(db)
	helper.CheckError("failed to connect database", err)

	qb := []string{
		`
		SELECT 
			id, 
			( 
				63710 * acos(
						cos( radians( ? ) ) * 
						cos( radians( current_latitude ) ) * 
						cos( radians( current_longitude ) - radians(?) ) + 
						sin(radians(?)) * sin(radians(current_latitude)) 
					)
			) AS distance,
			current_latitude,
			current_longitude 
		FROM `, tableName, `
		WHERE id > ? AND id <= ?
		HAVING distance < ?
		ORDER BY distance 
		LIMIT 0 , ? ;`,
	}

	sql := strings.Join(qb, " ")
	db.Raw(
		sql,
		latitude,
		longitude,
		latitude,
		startID,
		endID,
		radius,
		limit,
	).Scan(&drivers)

	return drivers
}

// ValidateQueryStrDriver ...
func (u *User) ValidateQueryStrDriver(
	ctx context.Context,
	errors *[]string,
) (
	paramLimit int,
	paramRadius int,
	paramLong float32,
	paramLat float32,
) {
	query := ctx.Request().URL.Query()

	if limit := query.Get("limit"); limit == "" {
		paramLimit = 10
	} else {
		paramLimit, _ = strconv.Atoi(limit)
	}

	if radius := query.Get("radius"); radius == "" {
		paramRadius = 500
	} else {
		paramRadius, _ = strconv.Atoi(radius)
	}
	if lat := query.Get("latitude"); lat == "" {
		(*errors) = append(*errors, "latitude is mandatory")
	} else {
		paramLat64, _ := strconv.ParseFloat(lat, 32)
		paramLat = float32(paramLat64)
		validateRangeLatitude(paramLat, errors)
	}

	if long := query.Get("longitude"); long == "" {
		(*errors) = append(*errors, "longitude is mandatory")
	} else {
		paramLong64, _ := strconv.ParseFloat(long, 32)
		paramLong = float32(paramLong64)
	}

	return
}

func validateRangeLatitude(paramLat float32, errors *[]string) {
	if paramLat < -90 || paramLat > 90 {
		(*errors) = append(*errors, "Latitude should be between +/- 90")
	}
}
