package invest

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller interface {
	SendCapital() echo.HandlerFunc
	CancelSendCapital() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	// Edit() echo.HandlerFunc
	// Archive() echo.HandlerFunc
}

type Model interface {
	SendCapital(data Investment) error
	CancelSendCapital(data Investment) error
	GetAll(user_id uint, page int) (interface{}, int, error)
	GetDetail(user_id uint, proposal_id int) (interface{}, error)
	// Edit(user_id string, proposal_id string, data Investment) error
	// Archive() error
}

type Services interface {
	SendCapital(token *jwt.Token, proposal_id uint, amount int) error
	CancelSendCapital(token *jwt.Token, proposal_id string) error
	GetAll(token *jwt.Token, page string) (interface{}, int, error)
	GetDetail(token *jwt.Token, proposal_id string) (interface{}, error)
	// Edit() error
	// Archive() error
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
	Status      int
	Collected   int
	Investments []Investment
	Reports     []Report
}

type Investment struct {
	gorm.Model
	Proposal_id uint `gorm:"primarykey"`
	Proposal    Proposal
	User_id     uint `gorm:"primarykey"`
	Amount      int
	Status      int
}

type Report struct {
	gorm.Model
	Proposal_id uint `gorm:"primarykey"`
	Document    string
}
