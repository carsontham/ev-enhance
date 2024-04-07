package model

import (
	"time"
)

type EVInformation struct {
}

type ExternalPartnerObject struct {
}

type EVTransaction struct {
	Id            int       `json:"id"`
	User_id       int       `json:"user_id"`
	Operator_id   int       `json:"operator_id"`
	Point_id      int       `json:"point_id"`
	Charging_time time.Time `json:"charging_time"`
	Charging_type string    `json:"charging_type"`
	Payment_type  string    `json:"payment_type"`
	Status        string    `json:"status"`
}

type Operators struct {
	Charger_id  string  `json:"charger_id"`
	Operator_id int     `json:"operator_id"`
	Location    string  `json:"location"`
	Rate_kwh    float64 `json:"rate_kwh"`
	Name        string  `json:"name"`
}

type Operator struct {
	Charger_id string `json:"charger_id"`
}
