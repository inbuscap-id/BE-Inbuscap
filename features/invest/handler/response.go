package handler

import (
	"time"
)

type InvestmentResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Capital     int       `json:"capital"`
	Share       int       `json:"share"`
	Status      int       `json:"status"`
	Amount      int       `json:"amount"`
}

type InvestmentDetileResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Amount    int       `json:"amount"`
	Status    int       `json:"status"`
}
