package handler

import "time"

// type LoginResponse struct {
// 	CreatedAt time.Time `json:"created_at" form:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
// 	Fullname  string    `json:"fullname" form:"fullname"`
// 	Username  string    `json:"username" form:"username"`
// 	Handphone string    `json:"handphone" form:"handphone"`
// 	Email     string    `json:"email" form:"email"`
// 	Biodata   string    `json:"biodata" form:"biodata"`
// 	Token     string    `json:"token" form:"token"`
// }

type ProfileResponse struct {
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
	Fullname   string    `json:"fullname" form:"fullname"`
	Email      string    `json:"email" form:"email"`
	Handphone  string    `json:"handphone" form:"handphone"`
	KTP        string    `json:"ktp" form:"ktp"`
	NPWP       string    `json:"npwp" form:"npwp"`
	IsVerified bool      `json:"is_active" form:"is_active"`
}
