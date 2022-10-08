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

func TestSearchFields(t *testing.T) {
	var expected_item models.Item
	var expected_items []models.Item
	var expected_fc models.ItemCollection
	jsonFile, _ := os.Open("setup_data/S2B_1CCV_20181004_0_L2A-test.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &expected_item)

	expected_items = append(expected_items, expected_item)
	expected_fc.Features = expected_items

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  models.ItemCollection
	}{
		{
			description:   "POST search fields test route",
			route:         "/search",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  expected_fc,
		},
	}

	// Setup the app as it is done in the main function
	app := Setup()

	var fields models.Fields
	fields.Include = append(fields.Include, "properties")

	body := []byte(`{"fields": {"include":["properties"]}}`)

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

		// // Read the response body
		// body, err := ioutil.ReadAll(res.Body)
		// assert.Nilf(t, err, "get item")

		// var stac_fc models.ItemCollection

		// json.Unmarshal(body, &stac_fc)

		// // Reading the response body should work everytime, such that
		// // the err variable should be nil
		// assert.Nilf(t, err, test.description)

		// // Verify, that the reponse body equals the expected body
		// assert.Equalf(t, test.expectedBody.Features, stac_fc.Features, test.description)
	}
}
