package main

import (
	"WhereIsMyDriver/helper"
	"WhereIsMyDriver/structs"
	"WhereIsMyDriver/structs/api"
	"encoding/json"
	"testing"
	// "github.com/buger/jsonparser"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	. "github.com/smartystreets/goconvey/convey"
)

// $ go test -v
func TestValidGetDriver(t *testing.T) {
	var res structs.Response
	var getDrivers []api.GetDriver

	app := irisApp()
	e := httptest.New(t, app)
	response := e.GET("/drivers").
		WithQuery("latitude", 37).
		WithQuery("longitude", -122).
		WithQuery("radius", 1000*1000).
		WithQuery("limit", 10).
		Expect()

	bodyByte := []byte(response.Body().Raw())
	json.Unmarshal(bodyByte, &res)

	helper.ToStruct(res.Data, &getDrivers)

	Convey("Valid get request \n", t, func() {
		Convey("Should have status 200", func() {
			So(response.Raw().StatusCode, ShouldEqual, iris.StatusOK)
		})
		Convey("Should have no error", func() {
			So(res.Errors, ShouldEqual, nil)
		})
		Convey("should have data driver location", func() {
			So(len(getDrivers), ShouldBeGreaterThan, 0)
		})
	})
}

func TestWrongLatitude(t *testing.T) {
	var res structs.Response

	app := irisApp()
	e := httptest.New(t, app)
	response := e.GET("/drivers").
		WithQuery("latitude", 92.00).
		WithQuery("longitude", -122).
		WithQuery("radius", 10).
		WithQuery("limit", 10).
		Expect()

	bodyByte := []byte(response.Body().Raw())
	json.Unmarshal(bodyByte, &res)
	Convey("Valid get request \n", t, func() {
		Convey("Should have status 422", func() {
			So(response.Raw().StatusCode,
				ShouldEqual,
				iris.StatusUnprocessableEntity)
		})
		Convey("Should have error", func() {
			So(res.Errors[0], ShouldEqual, "Latitude should be between +/- 90")
		})
	})
}
