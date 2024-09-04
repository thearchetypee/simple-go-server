package models

type InMemoryResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InMemorySuccessReponse struct {
	Msg string `json:"msg"`
}
