package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-stac-api/models"
	"go-stac-api/responses"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	var expected_collection models.Collection
	jsonFile, err := os.Open("setup_data/S2B_1CCV_20181004_0_L2A-test.json")

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &expected_collection)
	responseBody := bytes.NewBuffer(byteValue)

	resp, err := http.Post("http://localhost:6001/collections/sentinel-s2-l2a-cogs-test/items", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	assert.Equalf(t, 201, resp.StatusCode, "create item")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var collection_response responses.CollectionResponse
	json.Unmarshal(body, &collection_response)

	assert.Equalf(t, "success", collection_response.Message, "create collection")
}
