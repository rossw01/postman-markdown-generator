package models

type Collection struct {
	Name string    `json:"name"`
	Item []DocItem `json:"item"`
}
