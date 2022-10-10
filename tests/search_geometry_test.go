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
