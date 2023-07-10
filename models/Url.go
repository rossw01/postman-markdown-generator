package models

type Url struct {
	Raw      string   `json:"raw"`
	Protocol string   `json:"protocol"`
	Host     []string `json:"host"`
	Port     string   `json:"port"` // This must be a string to avoid issues with trailing 0s
	Path     []string `json:"path"`
}
