package user

import (
	"BE-Inbuscap/helper"
	"mime/multipart"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	AddVerification() echo.HandlerFunc
	GetVerifications() echo.HandlerFunc
	// Verify() echo.HandlerFunc
}

type Service interface {
	Register(register_data Register) error
	Login(login_data User) (string, error)
	Profile(token *jwt.Token) (User, error)
	Update(token *jwt.Token, update_data User, avatar *multipart.FileHeader) error
	Delete(token *jwt.Token) error
	AddVerification(token *jwt.Token, uploads []*multipart.FileHeader) error
	GetVerifications(paginasi helper.Pagination, status int) ([]User, int, error)
}

type Model interface {
	Register(register_data User) error
	Login(input string) (User, error)
	Profile(id string) (User, error)
	Update(data User) error
	Delete(id string) error
	GetVerifications(paginasi helper.Pagination, status int) ([]User, int, error)
}

// Structur Data
type User struct {
	gorm.Model
	Fullname    string
	Email       string `gorm:"unique"`
	Handphone   string `gorm:"unique"`
	KTP         string `gorm:"unique"`
	NPWP        string
	Password    string
	PhotoKTP    string
	PhotoNPWP   string
	PhotoSelf   string
	IsActive    int
	IsAdmin     bool
	Avatar      string
	Saldo       int
	Proposals   []Proposal
	Investments []Investment
}

type Proposal struct {
	gorm.Model
	User_id     uint
	User        *User
	Title       string
	Image       string
	Document    string
	Description string
	Capital     int
	Share       int
	Status      int
	Investments []Investment
	Reports     []Report
}

type Investment struct {
	gorm.Model
	Proposal_id uint `gorm:"primarykey"`
	User_id     uint `gorm:"primarykey"`
	Amount      int
	Status      int
}

type Report struct {
	gorm.Model
	Proposal_id uint `gorm:"primarykey"`
	Document    string
}

// Validate
type Register struct {
	Fullname  string `validate:"required,min=5"`
	Email     string `validate:"required,email"`
	Handphone string `validate:"required,number,min=11,max=14"`
	KTP       string `validate:"required,number,min=16,max=16"`
	NPWP      string `validate:"required,number,min=15,max=16"`
	Password  string `validate:"required,min=8"`
}

type Update struct {
	Fullname  string `validate:"required,min=5"`
	Email     string `validate:"required,email"`
	Handphone string `validate:"required,number,min=11,max=14"`
	KTP       string `validate:"required,number,min=16,max=16"`
	NPWP      string `validate:"required,number,min=15,max=16"`
	Password  string `validate:"required,min=8"`
}

// Response
// type LoginResponse struct {
// 	CreatedAt time.Time `json:"created_at" form:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
// 	Fullname  string    `json:"fullname" form:"fullname"`
// 	Handphone string    `json:"handphone" form:"handphone"`
// 	Email     string    `json:"email" form:"email"`
// 	Biodata   string    `json:"biodata" form:"biodata"`
// 	Token     string    `json:"token" form:"token"`
// }
