package models

type Search struct {
	Ids        []string `json:"ids,omitempty"`
	Collection string   `json:"collection,omitempty"`
}
