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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userItem *mongo.Collection = configs.GetItem(configs.DB, "items")
var validate_item = validator.New()

// CreateItem godoc
// @Summary Create a STAC item
// @Description Create an item with an ID
// @Tags Items
// @ID post-item
// @Accept  json
// @Produce  json
// @Param collectionId path string true "Collection ID"
// @Router /collections/{collectionId}/items/ [post]
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
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ItemResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// GetItem godoc
// @Summary Get an item
// @Description Get an item by its ID
// @Tags Items
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param itemId path string true "Item ID"
// @Param collectionId path string true "Collection ID"
// @Router /collections/{collectionId}/items/{itemId} [get]
func GetItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	itemId := c.Params("itemId")
	collectionId := c.Params("collectionId")
	var item models.Item
	defer cancel()

	err := userItem.FindOne(ctx, bson.M{"id": itemId, "collection": collectionId}).Decode(&item)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": item}})
}
