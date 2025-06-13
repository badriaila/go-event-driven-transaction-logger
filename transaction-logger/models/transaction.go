package models

type Transaction struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
}
