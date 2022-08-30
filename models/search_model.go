package models

type Search struct {
	Ids         []string `json:"ids,omitempty"`
	Collections []string `json:"collections,omitempty"`
	Limit       int      `json:"limit,omitempty"`
}
