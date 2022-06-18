// ./go/testing.go

package routes

import (
	"context"
	"go-stac-api/configs"
	"go-stac-api/models"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert" // add Testify package
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCollectionsRoute(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/collections",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("/collections", func(c *fiber.Ctx) error {
		// Return simple string as response
		return c.SendString("Hello, World!")
	})

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetCollection(t *testing.T) {
	var stacCollection *mongo.Collection = configs.GetCollection(configs.DB, "collections")
	app := fiber.New()
	req := httptest.NewRequest("GET", "/collections/sentinel-s2-l2a-cogs", nil)
	resp, _ := app.Test(req, 1)
	assert.Equal(t, 404, resp.StatusCode, "status ok")

	newCollection := models.Collection{
		Id:          "sentinel-s2-l2a-cogs",
		StacVersion: "1.0.0",
		Description: "Sentinel-2a and Sentinel-2b imagery, processed to Level 2A (Surface Reflectance) and converted to Cloud-Optimized GeoTIFFs",
		Title:       "Sentinel 2 L2A COGs",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stacCollection.InsertOne(ctx, newCollection)

	req = httptest.NewRequest("GET", "/collections/sentinel-s2-l2a-cogs", nil)
	resp, _ = app.Test(req)
	//body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode, "status ok")
}
