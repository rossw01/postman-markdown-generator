package models

type Body struct {
	Mode    string      `json:"mode"`
	Raw     string      `json:""`
	Options BodyOptions `json:"options"`
}
