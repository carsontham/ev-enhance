package service

type EVRequestForm struct {
	operator int `json:operator_id`
}

type EVResponse struct {
	operator_id int `json:operator_id`
	operator_name string `json:operator_name`
	rate float64 `json:rate_kwh`
}
