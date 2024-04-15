package report

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller interface {
	AddReport() echo.HandlerFunc
	Edit() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Archive() echo.HandlerFunc
}

type Model interface {
	AddReport(data Report) error
	Edit() error
	GetAll() error
	GetDetail() error
	Delete() error
	Archive() error
}

type Services interface {
	AddReport(token interface{}, proposal_id string, document *multipart.FileHeader) error
	Edit() error
	GetAll() error
	GetDetail() error
	Delete() error
	Archive() error
}

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
	IsVerified  bool
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
	Collected   int
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
