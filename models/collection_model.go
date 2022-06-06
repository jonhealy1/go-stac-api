package models

type Collection struct {
	StacVersion string        `json:"stac_version,omitempty"`
	Id          string        `json:"id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Keywords    []string      `json:"keywords,omitempty"`
	License     string        `json:"license,omitempty"`
	Providers   []interface{} `json:"providers,omitempty"`
	Extent      interface{}   `json:"extent,omitempty"`
	Summaries   interface{}   `json:"summary,omitempty"`
	Links       []interface{} `json:"links,omitempty"`
	ItemType    string        `json:"itemType,omitempty"`
	Crs         []string      `json:"crs,omitempty"`
}
