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
        "/collections": {
            "get": {
                "description": "Get all Collections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Get all Collections",
                "operationId": "get-all-collections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Collection"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a collection with a unique ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Create a STAC collection",
                "operationId": "post-collection",
                "parameters": [
                    {
                        "description": "STAC Collection json",
                        "name": "collection",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Collection"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/collections/{collectionId}": {
            "get": {
                "description": "Get a collection by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Get a Collection",
                "operationId": "get-collection-by-id",
                "parameters": [
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
                            "$ref": "#/definitions/models.Collection"
                        }
                    }
                }
            },
            "put": {
                "description": "Edit a collection by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Edit a Collection",
                "operationId": "get-collection-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection ID",
                        "name": "collectionId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "STAC Collection json",
                        "name": "collection",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Collection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Collection"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a collection by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Delete a Collection",
                "operationId": "delete-collection-by-id",
                "parameters": [
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
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/collections/{collectionId}/items": {
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
        "models.Collection": {
            "type": "object",
            "properties": {
                "crs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "extent": {},
                "id": {
                    "type": "string"
                },
                "itemType": {
                    "type": "string"
                },
                "keywords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "license": {
                    "type": "string"
                },
                "links": {
                    "type": "array",
                    "items": {}
                },
                "providers": {
                    "type": "array",
                    "items": {}
                },
                "stac_version": {
                    "type": "string"
                },
                "summary": {},
                "title": {
                    "type": "string"
                }
            }
        },
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
