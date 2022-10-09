package models

type Item struct {
	Id             string                 `json:"id,omitempty"`
	Type           string                 `json:"type,omitempty"`
	Collection     string                 `json:"collection,omitempty"`
	StacVersion    string                 `json:"stac_version,omitempty"`
	StacExtensions []string               `json:"stac_extensions,omitempty"`
	Bbox           []float64              `json:"bbox,omitempty"`
	Geometry       map[string]interface{} `json:"geometry,omitempty"`
	Properties     map[string]interface{} `json:"properties,omitempty"`
	Assets         map[string]interface{} `json:"assets,omitempty"`
	Links          []Link                 `json:"links,omitempty"`
}

type Context struct {
	Returned int `json:"returned,omitempty"`
	Limit    int `json:"limit,omitempty"`
}

type ItemCollection struct {
	Type     string  `json:"type,omitempty"`
	Context  Context `json:"context,omitempty"`
	Features []Item  `json:"features,omitempty"`
}

type ItemCollection2 struct {
	Type     string                   `json:"type,omitempty"`
	Context  Context                  `json:"context,omitempty"`
	Features []map[string]interface{} `json:"features,omitempty"`
}
