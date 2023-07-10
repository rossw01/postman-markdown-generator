package models

type Response struct {
	name string ``

	Response []struct {
		Name            string `json:"name"`
		OriginalRequest struct {
			Method string        `json:"method"`
			Header []interface{} `json:"header"`
			Body   struct {
				Mode    string `json:"mode"`
				Raw     string `json:"raw"`
				Options struct {
					Raw struct {
						Language string `json:"language"`
					} `json:"raw"`
				} `json:"options"`
			} `json:"body"`
			Url struct {
				Raw      string   `json:"raw"`
				Protocol string   `json:"protocol"`
				Host     []string `json:"host"`
				Port     string   `json:"port"`
				Path     []string `json:"path"`
			} `json:"url"`
		} `json:"originalRequest"`
		Status                 string `json:"status"`
		Code                   int    `json:"code"`
		PostmanPreviewlanguage string `json:"_postman_previewlanguage"`
		Header                 []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"header"`
		Cookie []interface{} `json:"cookie"`
		Body   string        `json:"body"`
	} `json:"response"`
}
