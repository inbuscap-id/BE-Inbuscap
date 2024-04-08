package handler

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
	Fullname  string `json:"fullname" form:"fullname"`
	Email     string `json:"email" form:"email"`
	Handphone string `json:"handphone" form:"handphone"`
	KTP       string `json:"ktp" form:"ktp"`
	NPWP      string `json:"npwp" form:"npwp"`
	Saldo     int    `json:"saldo" form:"saldo"`
	Avatar    string `json:"avatar" form:"avatar"`
	PhotoKTP  string `json:"photo_ktp" form:"photo_ktp"`
	PhotoNPWP string `json:"photo_npwp" form:"photo_npwp"`
	PhotoSelf string `json:"photo_selfie" form:"photo_selfie"`
}

type GetVerificationsResponse struct {
	ID        uint   `json:"id" form:"id"`
	Fullname  string `json:"fullname" form:"fullname"`
	Handphone string `json:"handphone" form:"handphone"`
	KTP       string `json:"ktp" form:"ktp"`
	NPWP      string `json:"npwp" form:"npwp"`
	IsActive  int    `json:"is_active" form:"is_active"`
	PhotoKTP  string `json:"photo_ktp" form:"photo_ktp"`
	PhotoNPWP string `json:"photo_npwp" form:"photo_npwp"`
	PhotoSelf string `json:"photo_selfie" form:"photo_selfie"`
}
