package models

// Invoice represents an invoice entity
type Invoice struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}
