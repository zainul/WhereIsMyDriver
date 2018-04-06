package models

import (
	"WhereIsMyDriver/structs/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var user = new(User)

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
