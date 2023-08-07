package model

type AddRequest struct {
	Param1 int32 `json:"param_1"`
	Param2 int32 `json:"param_2"`
}

type AddResponse struct {
	Response int32 `json:"response"`
}
