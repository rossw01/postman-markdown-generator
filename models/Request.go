package models

type Request struct {
	Method      string   `json:"method"`
	Header      []string `json:"header"`
	Body        Body     `json:"body"`
	Url         Url      `json:"url"`
	Description string   `json:"description"`
}
