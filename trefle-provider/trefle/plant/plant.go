package plant

type Plant struct {
	ID             int64  `json:"id"`
	CommonName     string `json:"common_name"`
	ScientificName string `json:"scientific_name"`
	ImageUrl       string `json:"image_url"`
}

type RawSearchPayload struct {
	Data []Plant `json:"data"`
	Meta struct {
		Total int64 `json:"total"`
	} `json:"meta"`
}
