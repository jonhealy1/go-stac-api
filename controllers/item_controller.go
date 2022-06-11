package controllers

import (
	"context"
	"go-stac-api/configs"
	"go-stac-api/models"
	"go-stac-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var userItem *mongo.Collection = configs.GetItem(configs.DB, "items")
var validate_item = validator.New()

func CreateItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Item
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate_item.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newItem := models.Item{
		Id:             user.Id,
		Type:           user.Type,
		StacVersion:    user.StacVersion,
		Collection:     user.Collection,
		StacExtensions: user.StacExtensions,
		Bbox:           user.Bbox,
		Geometry:       user.Geometry,
		Properties:     user.Properties,
		Assets:         user.Assets,
		Links:          user.Links,
	}

	result, err := userItem.InsertOne(ctx, newItem)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ItemResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
