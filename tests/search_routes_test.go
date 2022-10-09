package tests

import (
	"bytes"
	"encoding/json"
	"go-stac-api/models"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchFieldsExclude(t *testing.T) {
	var expected_item models.Item

	jsonFile, _ := os.Open("setup_data/S2B_1CCV_20181004_0_L2A-test.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &expected_item)

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  models.Item
	}{
		{
			description:   "POST search fields test route",
			route:         "/search",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  expected_item,
		},
	}

	// Setup the app as it is done in the main function
	app := Setup()

	body := []byte(`{"fields": {"exclude":["properties.created"]}}`)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"POST",
			test.route,
			bytes.NewBuffer(body),
		)
		req.Header.Add("Content-Type", "application/json")

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
		assert.Nilf(t, err, "get item")

		var bodyResponse models.ItemResponse

		json.Unmarshal(body, &bodyResponse)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		assert.Nil(t, bodyResponse.Data.Features[0].Properties["created"])

		delete(test.expectedBody.Properties, "created")

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody.Properties, bodyResponse.Data.Features[0].Properties, test.description)
	}
}

func TestSearchFieldsInclude(t *testing.T) {
	var expected_item models.Item

	jsonFile, _ := os.Open("setup_data/S2B_1CCV_20181004_0_L2A-test.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &expected_item)

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  models.Item
	}{
		{
			description:   "POST search fields test route",
			route:         "/search",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  expected_item,
		},
	}

	app := Setup()

	body := []byte(`{"fields": {"include":["properties.created"]}}`)

	for _, test := range tests {
		req, _ := http.NewRequest(
			"POST",
			test.route,
			bytes.NewBuffer(body),
		)
		req.Header.Add("Content-Type", "application/json")

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "get item")

		var bodyResponse models.ItemResponse

		json.Unmarshal(body, &bodyResponse)

		assert.Nilf(t, err, test.description)

		assert.Nil(t, bodyResponse.Data.Features[0].Assets)
		assert.Nil(t, bodyResponse.Data.Features[0].Links)
		assert.Nil(t, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"])

		assert.Equalf(t, test.expectedBody.Properties["created"], bodyResponse.Data.Features[0].Properties["created"], test.description)
	}
}
