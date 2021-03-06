package controllers

import (
	"WhereIsMyDriver/helper"
	"WhereIsMyDriver/models"
	"WhereIsMyDriver/structs"
	"WhereIsMyDriver/structs/api"
	"log"
	"strconv"

	"github.com/kataras/iris"

	"github.com/kataras/iris/context"
)

var user = new(models.User)

// UpdateLocation use for update the driver location
func UpdateLocation(c context.Context) {
	log.Println(c.Request().URL)
	// log.Println(c.CO)
	var res structs.Response

	loc := &api.UpdateLocation{}
	if err := c.ReadJSON(loc); err != nil {
		log.Println(err)

		return
	}
	log.Println(loc)
	userIDint, errStrConv := strconv.Atoi(c.Params().Get("id"))
	helper.CheckError("failed convert user id", errStrConv)

	updateLocationData := models.HistoryPosition{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Accuracy:  loc.Accuracy,
		UserID:    userIDint,
	}

	user.UpdateNewPositionDriver(
		updateLocationData,
		&res.Errors,
	)

	if len(res.Errors) > 0 {
		c.StatusCode(iris.StatusBadRequest)
	}

	_, errWrite := c.JSON(res)
	helper.CheckError("Failed write response json ", errWrite)
}

// GetDrivers is use for get driver by user location
// Parameters:
// "latitude" — mandatory
// "longitude" — mandatory
// "radius" — optional defaults to 500 meters
// "limit" — optional defaults to 10
func GetDrivers(c context.Context) {
	log.Println(c.Request().URL)
	var (
		res     structs.Response
		drivers []api.GetDriver
	)

	limit, radius, long, lat := user.ValidateQueryStrDriver(c, &res.Errors)

	if len(res.Errors) == 0 {
		drivers = user.GetDriver(lat, long, radius, limit)
	}

	if len(res.Errors) > 0 {
		c.StatusCode(iris.StatusUnprocessableEntity)
	} else {
		res.Data = drivers
	}

	_, errWrite := c.JSON(res)
	helper.CheckError("Failed write response json ", errWrite)
}
