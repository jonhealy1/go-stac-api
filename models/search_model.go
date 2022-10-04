package models

import "encoding/json"

type Search struct {
	Ids                []string                  `json:"ids,omitempty"`
	Collections        []string                  `json:"collections,omitempty"`
	Limit              int                       `json:"limit,omitempty"`
	Sort               []Sort                    `json:"sort,omitempty"`
	Datetime           string                    `json:"datetime,omitempty"`
	Bbox               []float64                 `json:"bbox,omitempty"`
	Geometry           GeoJSONGenericGeometry    `json:"geometry,omitempty"`
	GeometryCollection GeoJSONGeometryCollection `json:"geometrycollection,omitempty"`
}

type Sort struct {
	Field     string `json:"field,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type GeoJSONGeometryCollection struct {
	Type       string            `json:"type"` // will always be "GeometryCollection"
	Geometries []json.RawMessage `json:"geometries"`
}

type GeoJSONGenericGeometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}

type GeoJSONPoint struct {
	Type        string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

// line or multipoint
type GeoJSONLine struct {
	Type        string       `json:"type"`
	Coordinates [][2]float64 `json:"coordinates"`
}

// polygon or multiline
type GeoJSONPolygon struct {
	Type        string         `json:"type"`
	Coordinates [][][2]float64 `json:"coordinates"`
}

type GeoJSONMultiPolygon struct {
	Type        string           `json:"type"`
	Coordinates [][][][2]float64 `json:"coordinates"`
}
