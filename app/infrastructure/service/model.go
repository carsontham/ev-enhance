package service

type EVRequestForm struct {
	operator int `json:operator_id`
}

type EVResponse struct {
	Operator_id   int     `json:"operator_id"`
	Operator_name string  `json:"operator_name"`
	Rate          float64 `json:"rate_kwh"`
}
