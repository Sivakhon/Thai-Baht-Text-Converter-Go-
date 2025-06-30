package dto

type Money struct {
	Money float64 `json:"money" example: 21212121.21`
}

type MoneyResponse struct {
	Text string `json:"text"`
}
