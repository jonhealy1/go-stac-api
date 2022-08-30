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

var stacItem *mongo.Collection = configs.GetItem(configs.DB, "items")
var validate_item = validator.New()

// CreateItem godoc
// @Summary Create a STAC item
// @Description Create an item with an ID
// @Tags Items
// @ID post-item
// @Accept  json
// @Produce  json
// @Param collectionId path string true "Collection ID"
// @Param item body models.Item true "STAC Item json"
// @Router /collections/{collectionId}/items [post]
func CreateItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var item models.Item
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate_item.Struct(&item); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newItem := models.Item{
		Id:             item.Id,
		Type:           item.Type,
		StacVersion:    item.StacVersion,
		Collection:     item.Collection,
		StacExtensions: item.StacExtensions,
		Bbox:           item.Bbox,
		Geometry:       item.Geometry,
		Properties:     item.Properties,
		Assets:         item.Assets,
		Links:          item.Links,
	}

	result, err := stacItem.InsertOne(ctx, newItem)
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
// @Success 200 {object} models.Item
func GetItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	itemId := c.Params("itemId")
	collectionId := c.Params("collectionId")
	var item models.Item
	defer cancel()

	err := stacItem.FindOne(ctx, bson.M{"id": itemId, "collection": collectionId}).Decode(&item)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": item}})
}

// GetItemCollection godoc
// @Summary Get all Items from a Collection
// @Description Get all Items with a Collection ID
// @Tags ItemCollection
// @ID get-item-collection
// @Accept  json
// @Produce  json
// @Param collectionId path string true "Collection ID"
// @Router /collections/{collectionId}/items [get]
// @Success 200 {object} models.ItemCollection
func GetItemCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var items []models.Item
	defer cancel()

	results, err := stacItem.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleItem models.Item
		if err = results.Decode(&singleItem); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		items = append(items, singleItem)
	}
	itemCollection := models.ItemCollection{
		Type:     "FeatureCollection",
		Features: items,
	}

	return c.Status(http.StatusOK).JSON(
		responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": itemCollection}},
	)
}

// DeleteItem godoc
// @Summary Delete an Item
// @Description Delete an Item by ID is a specified collection
// @Tags Items
// @ID delete-item-by-id
// @Accept  json
// @Produce  json
// @Param itemId path string true "Item ID"
// @Param collectionId path string true "Collection ID"
// @Router /collections/{collectionId}/items/{itemId} [delete]
func DeleteItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	itemId := c.Params("itemId")
	collectionId := c.Params("collectionId")
	defer cancel()

	result, err := stacItem.DeleteOne(ctx, bson.M{"collection": collectionId, "id": itemId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.ItemResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Item with specified ID in collection not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Item successfully deleted!"}},
	)
}

// EditItem godoc
// @Summary Edit an Item
// @Description Edit a stac item by ID
// @Tags Collections
// @ID edit-item
// @Accept  json
// @Produce  json
// @Param collectionId path string true "Collection ID"
// @Param itemId path string true "Item ID"
// @Param item body models.Item true "STAC Collection json"
// @Router /collections/{collectionId}/items/{itemId} [put]
// @Success 200 {object} models.Item
func EditItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collectionId := c.Params("collectionId")
	itemId := c.Params("itemId")
	var item models.Item
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&item); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"id":              item.Id,
		"type":            item.Type,
		"stac_version":    item.StacVersion,
		"collection":      item.Collection,
		"stac_extensions": item.StacExtensions,
		"bbox":            item.Bbox,
		"geometry":        item.Geometry,
		"properties":      item.Properties,
		"assets":          item.Assets,
		"links":           item.Links,
	}

	result, err := stacItem.UpdateOne(ctx, bson.M{"collection": collectionId, "id": itemId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	var updatedItem models.Item
	if result.MatchedCount == 1 {
		err := stacItem.FindOne(ctx, bson.M{"collection": collectionId, "id": itemId}).Decode(&updatedItem)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedItem}})
}

// PostSearch godoc
// @Summary POST Search request
// @Description Search for STAC items via the Search endpoint
// @Tags Search
// @ID post-search
// @Accept  json
// @Produce  json
// @Param search body models.Search true "Search body json"
// @Router /search
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
