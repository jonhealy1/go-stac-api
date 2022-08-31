package controllers

import (
	"context"
	"fmt"
	"go-stac-api/models"
	"go-stac-api/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PostSearch godoc
// @Summary POST Search request
// @Description Search for STAC items via the Search endpoint
// @Tags Search
// @ID post-search
// @Accept  json
// @Produce  json
// @Param search body models.Search true "Search body json"
// @Router /search [post]
func PostSearch(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var search models.Search
	var items []models.Item
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&search); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate_item.Struct(&search); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	filter := bson.M{}
	if search.Geometry.Type != "" {
		if search.Geometry.Type == "Point" {
			long := search.Geometry.Coordinates[0]
			lat := search.Geometry.Coordinates[1]
			filter["geometry"] = bson.M{
				"$geoIntersects": bson.M{
					"$geometry": bson.M{
						"type":        "Point",
						"coordinates": []float64{long, lat},
					},
				},
			}
		}
	}
	if len(search.Collections) > 0 {
		filter["collection"] = bson.M{"$in": search.Collections}
	}
	if len(search.Ids) > 0 {
		filter["id"] = bson.M{"$in": search.Ids}
	}
	// if len(search.Bbox) > 0 {
	// 	filter["geometry"] = bson.M{"$geoIntersects": {"$geometry": geom}}
	// }

	fmt.Println(filter)

	limit := 0
	if search.Limit > 0 {
		limit = search.Limit
	} else {
		limit = 100
	}

	opts := options.Find().SetLimit(int64(limit))
	results, err := stacItem.Find(ctx, filter, opts)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	defer results.Close(ctx)
	count := 0
	for results.Next(ctx) {
		var singleItem models.Item
		if err = results.Decode(&singleItem); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		items = append(items, singleItem)
		count = count + 1
	}

	context := models.Context{
		Returned: count,
		Limit:    limit,
	}

	itemCollection := models.ItemCollection{
		Type:     "FeatureCollection",
		Context:  context,
		Features: items,
	}

	return c.Status(http.StatusOK).JSON(
		responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"results": itemCollection}},
	)
}
