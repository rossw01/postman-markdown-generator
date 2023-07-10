package models

type DocItem struct {
	Name     string   `json:"name"`
	Request  Request  `json:"request"`
	Response []string `json:"response"` // I don't think this is actually a string array, but from what I can tell, postman never uses this when creating documentation, so I guessed lol
}
