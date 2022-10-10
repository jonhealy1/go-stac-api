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

func TestSearchBbox(t *testing.T) {
	app := Setup()

	body := []byte(`{"bbox": [147.795950,-78.076921,179.557669,-65.760298]}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "bbox test")

	assert.Equalf(t, 200, res.StatusCode, "bbox test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[0], 179.557669)
	assert.LessOrEqual(t, 147.795950, bodyResponse.Data.Features[0].Bbox[0])
	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[2], 179.557669)
	assert.LessOrEqual(t, 147.795950, bodyResponse.Data.Features[0].Bbox[2])

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[1], -65.760298)
	assert.LessOrEqual(t, -78.076921, bodyResponse.Data.Features[0].Bbox[1])
	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[3], -65.760298)
	assert.LessOrEqual(t, -78.076921, bodyResponse.Data.Features[0].Bbox[3])
}

func TestSearchPoint(t *testing.T) {
	app := Setup()

	body := []byte(`{"geometry": {"type": "Point", "coordinates": [178.01642, -72.31064]}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "bbox test")

	assert.Equalf(t, 200, res.StatusCode, "bbox test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[0], 178.01642)
	assert.LessOrEqual(t, 178.01642, bodyResponse.Data.Features[0].Bbox[2])

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[1], -72.31064)
	assert.LessOrEqual(t, -72.31064, bodyResponse.Data.Features[0].Bbox[3])
}

func TestSearchLineString(t *testing.T) {
	app := Setup()

	body := []byte(`{"geometry": {"type": "LineString", "coordinates": [[177.85156249999997,-72.554563528593656],[177.101642,-72.690647]]}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "bbox test")

	assert.Equalf(t, 200, res.StatusCode, "bbox test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[0], 177.101642)
	assert.LessOrEqual(t, 177.85156249999997, bodyResponse.Data.Features[0].Bbox[2])

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[1], -72.690647)
	assert.LessOrEqual(t, -72.554563528593656, bodyResponse.Data.Features[0].Bbox[3])
}

func TestSearchPolygon(t *testing.T) {
	app := Setup()

	body := []byte(`{"geometry": {
		"type": "Polygon",
        "coordinates": [[
            [
              177.8515625,
              -74.14512718337613
            ],
            [
              178.35937499999999,
              -74.14512718337613
            ],
            [
              178.35937499999999,
              -72.15296965617042
            ],
            [
              177.8515625,
              -72.15296965617042
            ],
            [
              177.8515625,
              -74.14512718337613
            ]
        ]]
	}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "bbox test")

	assert.Equalf(t, 200, res.StatusCode, "bbox test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[0], 179.557669)
	assert.LessOrEqual(t, 147.795950, bodyResponse.Data.Features[0].Bbox[0])
	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[2], 179.557669)
	assert.LessOrEqual(t, 147.795950, bodyResponse.Data.Features[0].Bbox[2])

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[1], -65.760298)
	assert.LessOrEqual(t, -78.076921, bodyResponse.Data.Features[0].Bbox[1])
	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[3], -65.760298)
	assert.LessOrEqual(t, -78.076921, bodyResponse.Data.Features[0].Bbox[3])
}

func TestSearchGeometryCollection(t *testing.T) {
	app := Setup()

	body := []byte(`{"geometry": {
		"type": "GeometryCollection", 
		"geometries": [
			{
                "type": "Point",
                "coordinates": [178.01, -72.31]
            }
		]
	}}`)

	req, _ := http.NewRequest(
		"POST",
		"/search",
		bytes.NewBuffer(body),
	)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "bbox test")

	assert.Equalf(t, 200, res.StatusCode, "bbox test")

	resp_body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "query dateime")

	var bodyResponse models.ItemResponse

	json.Unmarshal(resp_body, &bodyResponse)

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[0], 178.01642)
	assert.LessOrEqual(t, 178.01642, bodyResponse.Data.Features[0].Bbox[2])

	assert.LessOrEqual(t, bodyResponse.Data.Features[0].Bbox[1], -72.31064)
	assert.LessOrEqual(t, -72.31064, bodyResponse.Data.Features[0].Bbox[3])
}
