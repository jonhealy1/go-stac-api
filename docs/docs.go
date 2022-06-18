// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jonathan Healy",
            "email": "jonathan.d.healy@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/collections/{collectionId}/items/": {
            "post": {
                "description": "Create an item with an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Create a STAC item",
                "operationId": "post-item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection ID",
                        "name": "collectionId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "STAC Item json",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Item"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/collections/{collectionId}/items/{itemId}": {
            "get": {
                "description": "Get an item by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Get an item",
                "operationId": "get-item-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "itemId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Collection ID",
                        "name": "collectionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Item"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Item": {
            "type": "object",
            "properties": {
                "assets": {},
                "bbox": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "collection": {
                    "type": "string"
                },
                "geometry": {},
                "id": {
                    "type": "string"
                },
                "links": {
                    "type": "array",
                    "items": {}
                },
                "properties": {},
                "stac_extensions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stac_version": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0",
	Host:             "localhost:6001",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "go-stac-api",
	Description:      "STAC api written in go with fiber and mongodb",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
