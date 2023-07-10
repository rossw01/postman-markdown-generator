package models

type Documentation struct {
	Info Info         `json:"info"`
	Item []Collection `json:"item"`
}
