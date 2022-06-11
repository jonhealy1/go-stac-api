package models

type Item struct {
	Id             string        `json:"id,omitempty"`
	Type           string        `json:"type,omitempty"`
	Collection     string        `json:"collection,omitempty"`
	StacVersion    string        `json:"stac_version,omitempty"`
	StacExtensions []string      `json:"stac_extensions,omitempty"`
	Bbox           []float64     `json:"bbox,omitempty"`
	Geometry       interface{}   `json:"geometry,omitempty"`
	Properties     interface{}   `json:"properties,omitempty"`
	Assets         interface{}   `json:"assets,omitempty"`
	Links          []interface{} `json:"links,omitempty"`
}
