package models

import "time"

type Revenue struct {
	From  *time.Time `json:"from"`
	To    *time.Time `json:"to"`
	Total float64    `json:"total"`
}

type RevenueReport struct {
	Current          *Revenue `json:"current"`
	Previous         *Revenue `json:"previous"`
	PercentageChange float64  `json:"percentage_change"`
}
