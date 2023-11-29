package model

type DeleteResponse struct {
	Status  int    `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Success bool   `json:"success" form:"success"`
}

type PutPostResponse struct {
	Data    interface{} `json:"data" form:"data"`
	Message string      `json:"message" form:"message"`
	Success bool        `json:"success" form:"success"`
}

type DataResponse struct {
	Data []Orders `json:"data"`
}
