package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-stac-api/configs"
	"go-stac-api/models"
	"go-stac-api/responses"
	"go-stac-api/routes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/assert"
)

func Setup() *fiber.App {
	configs.ConnectDB()
	app := fiber.New()

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(cache.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(recover.New())

	routes.CollectionRoute(app)
	routes.ItemRoute(app)

	return app
}

func TestCreateCollection(t *testing.T) {
	var expected_collection models.Collection
	jsonFile, err := os.Open("setup_data/collection.json")

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &expected_collection)
	responseBody := bytes.NewBuffer(byteValue)

	resp, err := http.Post("http://localhost:6001/collections", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	assert.Equalf(t, 201, resp.StatusCode, "create collection")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var collection_response responses.CollectionResponse
	json.Unmarshal(body, &collection_response)

	assert.Equal(t, "success", collection_response.Message, "create collection")
}
func TestGetCollection(t *testing.T) {
	LoadCollection()
	LoadItems()

	var expected_collection models.Collection
	jsonFile, _ := os.Open("setup_data/collection.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &expected_collection)

	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  models.Collection
	}{
		{
			description:   "GET collection route",
			route:         "/collections/sentinel-s2-l2a-cogs-test",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  expected_collection,
		},
		{
			description:   "GET collection route",
			route:         "/collections/sentinel-s2-l2a-cogs-test",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  expected_collection,
		},
	}

	// Setup the app as it is done in the main function
	app := Setup()

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// // verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		var stac_collection models.Collection

		json.Unmarshal(body, &stac_collection)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, stac_collection, test.description)
	}
}

func TestGetAllCollections(t *testing.T) {
	tests := []struct {
		description string
		route       string

		expectedError bool
		expectedCode  int
	}{
		{
			description:   "GET collections route",
			route:         "/collections",
			expectedError: false,
			expectedCode:  200,
		},
	}

	// Setup the app as it is done in the main function
	app := Setup()

	// Iterate through test single test cases
	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// // verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		var stac_collection []models.Collection

		json.Unmarshal(body, &stac_collection)

		// length := len(stac_collection)

		// assert.Equal(t, length, 1, test.description)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)
	}
}

func TestDeleteCollection(t *testing.T) {
	// Create client
	client := &http.Client{}

	// Creaate Request
	req, err := http.NewRequest("DELETE", "/collections/sentinel-s2-l2a-cogs-test", nil)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	assert.Equalf(t, 201, resp.Status, "create collection")

	// Read Response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var collection_response responses.CollectionResponse
	json.Unmarshal(body, &collection_response)

	assert.Equal(t, "success", collection_response.Message, "create collection")

	// // Display Results
	// fmt.Println("response Status : ", resp.Status)
	// fmt.Println("response Headers : ", resp.Header)
	// fmt.Println("response Body : ", string(respBody))
}
