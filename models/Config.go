package models

type Config struct {
	EndpointName        string `json:"endpoint-name"`
	EndpointDescription string `json:"endpoint-description"`
	EndpointMethod      string `json:"endpoint-method"`
	EndpointUrl         string `json:"endpoint-url"`
	EndpointBodyFormat  string `json:"endpoint-body-format"`
	EndpointBody        string `json:"endpoint-body"`
	UseMethodImage      bool   `json:"use-method-image"`
	GetImageUrl         string `json:"get-image-url"`
	PutImageUrl         string `json:"put-image-url"`
	DeleteImageUrl      string `json:"delete-image-url"`
	PostImageUrl        string `json:"post-image-url"`
	OptionsImageUrl     string `json:"options-image-url"`
	PatchImageUrl       string `json:"patch-image-url"`
	HeadImageUrl        string `json:"head-image-url"`
}
