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

var stacCollection *mongo.Collection = configs.GetCollection(configs.DB, "collections")
var validate = validator.New()

func Root(c *fiber.Ctx) error {
	links := []models.Link{
		{
			Rel:   "self",
			Type:  "application/json",
			Href:  "/",
			Title: "root catalog",
		},
		{
			Rel:   "children",
			Type:  "application/json",
			Href:  "/collections",
			Title: "stac child collections",
		},
	}

	rootCatalog := models.Root{
		Id:          "test-catalog",
		StacVersion: "1.0.0",
		Description: "test catalog for go-stac-api, please edit",
		Title:       "go-stac-api",
		Links:       links,
	}

	return c.Status(http.StatusOK).JSON(rootCatalog)
}

func Conformance(c *fiber.Ctx) error {
	conformsTo := []string{
		"http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/core",
		"http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/oas30",
		"http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/geojson",
	}
	conformance := bson.M{
		"conformsTo": conformsTo,
	}

	return c.Status(http.StatusOK).JSON(conformance)
}

// CreateCollection godoc
// @Summary Create a STAC collection
// @Description Create a collection with a unique ID
// @Tags Collections
// @ID post-collection
// @Accept  json
// @Produce  json
// @Param collection body models.Collection true "STAC Collection json"
// @Router /collections [post]
func CreateCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var collection models.Collection
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&collection); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&collection); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newCollection := models.Collection{
		Id:          collection.Id,
		StacVersion: collection.StacVersion,
		Description: collection.Description,
		Title:       collection.Title,
		Links:       collection.Links,
		Extent:      collection.Extent,
		Providers:   collection.Providers,
	}

	result, err := stacCollection.InsertOne(ctx, newCollection)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.CollectionResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// GetCollection godoc
// @Summary Get a Collection
// @Description Get a collection by ID
// @Tags Collections
// @ID get-collection-by-id
// @Accept  json
// @Produce  json
// @Param collectionId path string true "Collection ID"
// @Router /collections/{collectionId} [get]
// @Success 200 {object} models.Collection
func GetCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collectionId := c.Params("collectionId")
	var collection models.Collection
	defer cancel()

	err := stacCollection.FindOne(ctx, bson.M{"id": collectionId}).Decode(&collection)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.CollectionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": collection}})
}

// EditCollection godoc
// @Summary Edit a Collection
// @Description Edit a collection by ID
// @Tags Collections
// @ID edit-collection
// @Accept  json
// @Produce  json
// @Param collectionId path string true "Collection ID"
// @Param collection body models.Collection true "STAC Collection json"
// @Router /collections/{collectionId} [put]
// @Success 200 {object} models.Collection
func EditCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collectionId := c.Params("collectionId")
	var collection models.Collection
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&collection); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&collection); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"id":           collection.Id,
		"stac_version": collection.StacVersion,
		"description":  collection.Description,
		"title":        collection.Title,
		"links":        collection.Links,
		"extent":       collection.Extent,
		"providers":    collection.Providers,
	}

	result, err := stacCollection.UpdateOne(ctx, bson.M{"id": collectionId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//get updated collection details
	var updatedCollection models.Collection
	if result.MatchedCount == 1 {
		err := stacCollection.FindOne(ctx, bson.M{"id": collectionId}).Decode(&updatedCollection)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.CollectionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedCollection}})
}

// DeleteCollection godoc
// @Summary Delete a Collection
// @Description Delete a collection by ID
// @Tags Collections
// @ID delete-collection-by-id
// @Accept  json
// @Produce  json
// @Param collectionId path string true "Collection ID"
// @Router /collections/{collectionId} [delete]
func DeleteCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collectionId := c.Params("collectionId")
	defer cancel()

	result, err := stacCollection.DeleteOne(ctx, bson.M{"id": collectionId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.CollectionResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Collection with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.CollectionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Collection successfully deleted!"}},
	)
}

// GetCollections godoc
// @Summary Get all Collections
// @Description Get all Collections
// @Tags Collections
// @ID get-all-collections
// @Accept  json
// @Produce  json
// @Router /collections [get]
// @Success 200 {object} []models.Collection
func GetCollections(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var collections []models.Collection
	defer cancel()

	results, err := stacCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleCollection models.Collection
		if err = results.Decode(&singleCollection); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		collections = append(collections, singleCollection)
	}

	return c.Status(http.StatusOK).JSON(
		responses.CollectionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": collections}},
	)
}
