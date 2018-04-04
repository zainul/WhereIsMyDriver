package controllers

import (
	"WhereIsMyDriver/models"
	"WhereIsMyDriver/structs"
	"WhereIsMyDriver/structs/api"
	"log"

	"github.com/kataras/iris"

	"github.com/kataras/iris/context"
)

var user = new(models.User)

// UpdateLocation use for update the driver location
func UpdateLocation(c context.Context) {
	var res structs.Response

	loc := &api.UpdateLocation{}
	if err := c.ReadJSON(loc); err != nil {
		log.Println(err)

		return
	}

	updateLocationData := models.HistoryPosition{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Accuracy:  loc.Accuracy,
		UserID:    c.Params().Get("id"),
	}

	user.UpdateNewPositionDriver(
		updateLocationData,
		&res.Errors,
	)

	if len(res.Errors) > 0 {
		c.StatusCode(iris.StatusBadRequest)
	}

	c.JSON(res)
}
