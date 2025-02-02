package types

type JSONAPIResource struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	Atributes     map[string]interface{} `json:"atributes"`
	Relationships map[string]interface{} `json:"relationships"`
}

type JSONAPIResponse struct {
	Data JSONAPIResource `json:"data"`
}
