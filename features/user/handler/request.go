package handler

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Fullname  string `json:"fullname" form:"fullname"`
	Email     string `json:"email" form:"email"`
	Handphone string `json:"handphone" form:"handphone"`
	KTP       string `json:"ktp" form:"ktp"`
	NPWP      string `json:"npwp" form:"npwp"`
	Password  string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Fullname  string `json:"fullname" form:"fullname"`
	Email     string `json:"email" form:"email"`
	Handphone string `json:"handphone" form:"handphone"`
	KTP       string `json:"ktp" form:"ktp"`
	NPWP      string `json:"npwp" form:"npwp"`
	Password  string `json:"password" form:"password"`
}
