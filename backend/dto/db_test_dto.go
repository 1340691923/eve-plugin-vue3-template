package dto

type DbTestSearchReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type DbTestInsertReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
