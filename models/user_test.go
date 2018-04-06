package models

import (
	"WhereIsMyDriver/helper"
	"WhereIsMyDriver/structs/api"
	"strconv"
	"testing"

	"github.com/icrowley/fake"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
)

var user = new(User)

func init() {
	for index := 0; index < 500; index++ {
		var user = User{}

		var errors = []string{}
		password, err := bcrypt.GenerateFromPassword(
			[]byte(fake.SimplePassword()),
			bcrypt.DefaultCost,
		)

		helper.CheckError("failed make bcrypt password", err)
		idxStr := strconv.Itoa(index)
		userData := User{
			Username:         fake.UserName() + idxStr,
			Phone:            idxStr + fake.Phone(),
			Email:            idxStr + fake.EmailAddress(),
			FirstName:        fake.FirstName(),
			FullName:         fake.FullName(),
			IdentifiedNumber: idxStr + fake.CharactersN(10),
			LastName:         fake.LastName(),
			Password:         string(password),
			CurrentLatitude:  fake.Latitude(),
			CurrentLongitude: fake.Longitude(),
			CurrentAccuracy:  0.7,
		}
		userData.SetDefault()

		go user.AddUser(&userData, &errors)
	}
}

func TestUserGetData(t *testing.T) {
	drivers := user.GetDriver(57, -122, 1000, 10)
	var driverType api.GetDriver

	Convey("Valid get request \n", t, func() {
		Convey("Should have driver records", func() {
			So(len(drivers), ShouldBeGreaterThan, 0)
		})
		Convey("Should have structure", func() {
			So(drivers[0].Distance, ShouldHaveSameTypeAs, driverType.Distance)
			So(drivers[0].ID, ShouldHaveSameTypeAs, driverType.ID)
			So(drivers[0].Latitude, ShouldHaveSameTypeAs, driverType.Latitude)
			So(drivers[0].Longitude, ShouldHaveSameTypeAs, driverType.Longitude)
		})
	})
}

func TestUserUpdatePosition(t *testing.T) {
	var err []string
	user.UpdateNewPositionDriver(
		HistoryPosition{
			Latitude:  37,
			Longitude: 55,
			UserID:    1,
		},
		&err,
	)

	Convey("Success update driver position \n", t, func() {
		Convey("Should have no error", func() {
			So(len(err), ShouldEqual, 0)
		})
	})
}
