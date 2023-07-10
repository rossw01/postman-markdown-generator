package models

type Info struct {
	PostmanId   string `json:"_postman_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
	ExporterId  string `json:"_exporter_id"`
}
