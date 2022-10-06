package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-stac-api/models"
	"go-stac-api/responses"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PostSearch godoc
// @Summary POST Search request
// @Description Search for STAC items via the Search endpoint
// @Tags Search
// @ID post-search
// @Accept  json
// @Produce  json
// @Param search body models.Search true "Search body json"
// @Router /search [post]
func PostSearch(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var search models.Search
	var items []models.Item
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&search); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate_item.Struct(&search); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
	}

	filter := bson.M{}

	queries := make([]string, 0, len(search.Query))

	for k := range search.Query {
		if val, ok := search.Query[k]["lte"]; ok {
			property := "properties." + k
			filter[property] = bson.M{"$lte": val}
		}
		queries = append(queries, k)
	}

	fmt.Println(queries)

	if strings.Contains(search.Datetime, "/") {
		parsed := returnDatetime(search.Datetime)
		filter["properties.datetime"] = bson.M{"$gte": parsed[0], "$lte": parsed[1]}
	}
	if len(search.Bbox) > 0 {
		geom := bbox2polygon(search.Bbox)

		filter["geometry"] = bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": geom,
				},
			},
		}
	}
	if search.Geometry.Type == "Point" {
		geom := models.GeoJSONPoint{}.Coordinates
		json.Unmarshal(search.Geometry.Coordinates, &geom)
		filter["geometry"] = bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        search.Geometry.Type,
					"coordinates": geom,
				},
			},
		}
	}

	if search.Geometry.Type == "Polygon" {
		geom := models.GeoJSONPolygon{}.Coordinates
		json.Unmarshal(search.Geometry.Coordinates, &geom)
		filter["geometry"] = bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        search.Geometry.Type,
					"coordinates": geom,
				},
			},
		}
	}

	if search.Geometry.Type == "LineString" {
		geom := models.GeoJSONLine{}.Coordinates
		json.Unmarshal(search.Geometry.Coordinates, &geom)
		filter["geometry"] = bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        search.Geometry.Type,
					"coordinates": geom,
				},
			},
		}
	}

	if search.GeometryCollection.Type == "GeometryCollection" {
		for _, geometryJSON := range search.GeometryCollection.Geometries {
			generic := models.GeoJSONGenericGeometry{}
			json.Unmarshal(geometryJSON, &generic)
			switch generic.Type {
			case "MultiPolygon":
				geom := models.GeoJSONMultiPolygon{}
				json.Unmarshal(geometryJSON, &geom)
				filter["geometry"] = bson.M{
					"$geoIntersects": bson.M{
						"$geometry": bson.M{
							"type":        geom.Type,
							"coordinates": geom.Coordinates,
						},
					},
				}

			case "Polygon", "MultiLineString":
				geom := models.GeoJSONPolygon{}
				json.Unmarshal(geometryJSON, &geom)
				filter["geometry"] = bson.M{
					"$geoIntersects": bson.M{
						"$geometry": bson.M{
							"type":        geom.Type,
							"coordinates": geom.Coordinates,
						},
					},
				}

			case "LineString", "MultiPoint":
				geom := models.GeoJSONLine{}
				json.Unmarshal(geometryJSON, &geom)
				filter["geometry"] = bson.M{
					"$geoIntersects": bson.M{
						"$geometry": bson.M{
							"type":        geom.Type,
							"coordinates": geom.Coordinates,
						},
					},
				}

			case "Point":
				geom := models.GeoJSONPoint{}
				json.Unmarshal(geometryJSON, &geom)
				filter["geometry"] = bson.M{
					"$geoIntersects": bson.M{
						"$geometry": bson.M{
							"type":        geom.Type,
							"coordinates": geom.Coordinates,
						},
					},
				}
			}
		}
	}
	if len(search.Collections) > 0 {
		filter["collection"] = bson.M{"$in": search.Collections}
	}
	if len(search.Ids) > 0 {
		filter["id"] = bson.M{"$in": search.Ids}
	}
	fmt.Println("Filter: ", filter)

	limit := 0
	if search.Limit > 0 {
		limit = search.Limit
	} else {
		limit = 100
	}

	opts := options.Find().SetLimit(int64(limit))
	if len(search.Sort) > 0 {
		field := "properties." + search.Sort[0].Field
		value := 1
		if search.Sort[0].Direction == "desc" {
			value = -1
		}
		opts = options.Find().SetLimit(int64(limit)).SetSort(bson.D{{Key: field, Value: value}})
	}

	results, err := stacItem.Find(ctx, filter, opts)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}
	defer results.Close(ctx)
	count := 0
	for results.Next(ctx) {
		var singleItem models.Item
		if err = results.Decode(&singleItem); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		}
		items = append(items, singleItem)
		count = count + 1
	}

	context := models.Context{
		Returned: count,
		Limit:    limit,
	}

	itemCollection := models.ItemCollection{
		Type:     "FeatureCollection",
		Context:  context,
		Features: items,
	}

	return c.Status(http.StatusOK).JSON(
		responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: itemCollection},
	)
}
