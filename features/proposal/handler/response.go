package handler

import "time"

type ProposalResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Capital     int       `json:"capital"`
	Share       int       `json:"share"`
	Status      int       `json:"status"`
	Collected   int       `json:"collected"`
}

type ProposalDetailResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      struct {
		Fullname  string `json:"fullname"`
		Email     string `json:"email"`
		Handphone string `json:"handphone"`
	} `json:"user"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Document    string `json:"document"`
	Description string `json:"description"`
	Capital     int    `json:"capital"`
	Share       int    `json:"share"`
	Status      int    `json:"status"`
	Collected   int    `json:"collected"`
}

type VerificationResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
	Capital     int    `json:"capital"`
	Share       int    `json:"share"`
	Status      int    `json:"status"`
	Document    string `json:"proposal"`
}
