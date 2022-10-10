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

func TestSearchDatetime(t *testing.T) {
	app := Setup()

	body := []byte(`{"datetime": "2018-10-04T21:05:21Z/2018-10-05T21:05:21Z"}`)

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
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Properties["datetime"], "2018-10-05T21:05:21Z")
	assert.LessOrEqual(t, "2018-10-04T21:05:21Z", bodyResponse.Data.Features[0].Properties["datetime"])
}

func TestSearchDatetimeOpenBefore(t *testing.T) {
	app := Setup()

	body := []byte(`{"datetime": "../2018-10-05T21:05:21Z"}`)

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
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Properties["datetime"], "2018-10-05T21:05:21Z")
}

func TestSearchDatetimeOpenAfter(t *testing.T) {
	app := Setup()

	body := []byte(`{"datetime": "2018-10-05T21:05:21Z/.."}`)

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
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, "2018-10-05T21:05:21Z", bodyResponse.Data.Features[0].Properties["datetime"])
}

func TestSearchDatetimeOpenEnded(t *testing.T) {
	app := Setup()

	body := []byte(`{"datetime": "../.."}`)

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
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, "1900-10-05T21:05:21Z", bodyResponse.Data.Features[0].Properties["datetime"])
	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Properties["datetime"], "2100-10-05T21:05:21Z")
}
