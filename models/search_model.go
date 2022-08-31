package models

type Search struct {
	Ids         []string  `json:"ids,omitempty"`
	Collections []string  `json:"collections,omitempty"`
	Limit       int       `json:"limit,omitempty"`
	Bbox        []float64 `json:"bbox,omitempty"`
	Geometry    GeoJson   `json:"geometry,omitempty"`
}

type GeoJson struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
