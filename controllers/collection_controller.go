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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "collections")
var validate = validator.New()

func CreateCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Collection
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CollectionResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newCollection := models.Collection{
		Id:          user.Id,
		StacVersion: user.StacVersion,
		Description: user.Description,
		Title:       user.Title,
		Links:       user.Links,
		Providers:   user.Providers,
	}

	result, err := userCollection.InsertOne(ctx, newCollection)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.CollectionResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("collectionId")
	var collection models.Collection
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&collection)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.CollectionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": collection}})
}

func EditCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collectionId := c.Params("collectionId")
	var collection models.Collection
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(collectionId)

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
		"providers":    collection.Providers,
	}

	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//get updated user details
	var updatedCollection models.Collection
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedCollection)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.CollectionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.CollectionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedCollection}})
}

func DeleteCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collectionId := c.Params("collectionId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(collectionId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
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

func GetCollections(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.Collection
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

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

		users = append(users, singleCollection)
	}

	return c.Status(http.StatusOK).JSON(
		responses.CollectionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}},
	)
}
