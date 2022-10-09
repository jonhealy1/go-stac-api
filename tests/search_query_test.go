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

func TestSearchQueryEq(t *testing.T) {

	// Setup the app as it is done in the main function
	app := Setup()

	body := []byte(`{"query": {"eo:cloud_cover": {"eq": 96.91}}}`)

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
	assert.Nilf(t, err, "query eq")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.Equal(t, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"], 96.91)
}

func TestSearchQueryLte(t *testing.T) {
	app := Setup()

	body := []byte(`{"query": {"eo:cloud_cover": {"lte": 6.91}}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "sort asc test")

	assert.Equalf(t, 200, res.StatusCode, "sort asc test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query eq")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"], 6.91)
}

func TestSearchQueryGte(t *testing.T) {
	app := Setup()

	body := []byte(`{"query": {"eo:cloud_cover": {"gte": 50.1}}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "sort asc test")

	assert.Equalf(t, 200, res.StatusCode, "sort asc test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query eq")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, 50.1, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"])
}

func TestSearchQueryLt(t *testing.T) {
	app := Setup()

	body := []byte(`{"query": {"eo:cloud_cover": {"lt": 3.01}}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "sort asc test")

	assert.Equalf(t, 200, res.StatusCode, "sort asc test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query eq")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"], 3.00)
}

func TestSearchQueryGt(t *testing.T) {
	app := Setup()

	body := []byte(`{"query": {"eo:cloud_cover": {"gt": 99.91}}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "sort asc test")

	assert.Equalf(t, 200, res.StatusCode, "sort asc test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query eq")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, 99.91, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"])
}

func TestSearchQueryMulti(t *testing.T) {
	app := Setup()

	body := []byte(`{"query": {"data_coverage": {"gte": 20.18}, "eo:cloud_cover": {"lte": 17.19}}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "sort asc test")

	assert.Equalf(t, 200, res.StatusCode, "sort asc test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query eq")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Properties["eo:cloud_cover"], 17.19)
	assert.LessOrEqual(t, 20.18, bodyResponse.Data.Features[0].Properties["data_coverage"])
}
