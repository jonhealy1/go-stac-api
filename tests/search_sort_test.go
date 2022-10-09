package tests

import (
	"bytes"
	"encoding/json"
	"go-stac-api/models"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchSortAsc(t *testing.T) {

	// Setup the app as it is done in the main function
	app := Setup()

	body := []byte(`{"sort": [{"field": "eo:cloud_cover", "direction": "asc"}]}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	// Perform the request plain with the app.
	// The -1 disables request latency.
	res, err := app.Test(req, -1)

	// // verify that no error occured, that is not expected
	assert.Equalf(t, false, err != nil, "sort asc test")

	// Verify if the status code is as expected
	assert.Equalf(t, 200, res.StatusCode, "sort asc test")

	// Read the response body
	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "get item")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"], 1.1)
}

func TestSearchSortDesc(t *testing.T) {

	// Setup the app as it is done in the main function
	app := Setup()

	body := []byte(`{"sort": [{"field": "eo:cloud_cover", "direction": "desc"}]}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	// Perform the request plain with the app.
	// The -1 disables request latency.
	res, err := app.Test(req, -1)

	// // verify that no error occured, that is not expected
	assert.Equalf(t, false, err != nil, "sort asc test")

	// Verify if the status code is as expected
	assert.Equalf(t, 200, res.StatusCode, "sort asc test")

	// Read the response body
	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "get item")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, 99.9, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"])
}
