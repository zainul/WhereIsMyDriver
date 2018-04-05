package databases

import (
	"WhereIsMyDriver/adapters"
	"WhereIsMyDriver/helper"
	"WhereIsMyDriver/models"
	"log"
	"strconv"
	"time"

	"github.com/icrowley/fake"
	"github.com/ivpusic/grpool"
	"golang.org/x/crypto/bcrypt"
)

// TotalUserSeed is total user seed data
var TotalUserSeed = (50 * 1000)

// SeedData ...
func SeedData() {
	start := time.Now()
	// number of workers, and size of job queue
	pool := grpool.NewPool(100, 500)

	// release resources used by pool
	defer pool.Release()

	if !IsCompletedSeedUserData() {

		// submit one or more jobs to pool
		for i := 0; i < TotalUserSeed; i++ {
			count := i
			log.Println(count)
			pool.JobQueue <- func() {
				SeedUserData(count)
			}
		}
	}

	log.Println("time for insert ", time.Since(start))
}

// SeedUserData ...
func SeedUserData(i int) {
	var user = models.User{}

	var errors = []string{}
	password, err := bcrypt.GenerateFromPassword(
		[]byte(fake.SimplePassword()),
		bcrypt.DefaultCost,
	)

	helper.CheckError("failed make bcrypt password", err)
	idxStr := strconv.Itoa(i)
	userData := models.User{
		Username:         fake.UserName() + idxStr,
		Phone:            idxStr + fake.Phone(),
		Email:            idxStr + fake.EmailAddress(),
		FirstName:        fake.FirstName(),
		FullName:         fake.FullName(),
		IdentifiedNumber: fake.CharactersN(10),
		LastName:         fake.LastName(),
		Password:         string(password),
		CurrentLatitude:  fake.Latitude(),
		CurrentLongitude: fake.Longitude(),
		CurrentAccuracy:  0.7,
	}
	userData.SetDefault()

	go user.AddUser(&userData, &errors)
}

// IsCompletedSeedUserData ...
func IsCompletedSeedUserData() bool {
	var user = models.User{}
	var count int
	db, err := adapters.ConnectDB()
	helper.CheckError("failed connect to database", err)
	db.Table(user.TableName()).Count(&count)
	log.Println(count)
	return (count >= TotalUserSeed)
}
