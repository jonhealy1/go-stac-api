package controllers

import (
	"context"
	"go-stac-api/models"
	"go-stac-api/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

	count := 0
	if search.Collection != "" {
		results, err := stacItem.Find(ctx, bson.M{"id": bson.M{"$in": search.Ids}, "collection": search.Collection})
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleItem models.Item
			if err = results.Decode(&singleItem); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
			}
			count = count + 1
			items = append(items, singleItem)
		}
	} else {
		results, err := stacItem.Find(ctx, bson.M{"id": bson.M{"$in": search.Ids}})
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleItem models.Item
			if err = results.Decode(&singleItem); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
			}
			count = count + 1
			items = append(items, singleItem)
		}
	}

	itemCollection := models.ItemCollection{
		Type:     "FeatureCollection",
		Features: items,
	}

	return c.Status(http.StatusOK).JSON(
		responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"count": count, "data": itemCollection}},
	)

	//return c.Status(http.StatusOK).JSON(responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": results}})
}
