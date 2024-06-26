package proposal

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetAllMy() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Archive() echo.HandlerFunc
	GetVerifications() echo.HandlerFunc
	GetVerification() echo.HandlerFunc
	ChangeStatus() echo.HandlerFunc
	GetUpload() echo.HandlerFunc
}

type Model interface {
	Create(data Proposal) error
	Update(user_id string, proposal_id string, data Proposal) error
	GetAll(page int) ([]Proposal, int, error)
	GetAllMy(page int, user_id string) ([]Proposal, int, error)
	GetDetail(id_proposal string) (Proposal, error)
	Delete(id string, prososal_id string) error
	Archive() error
	GetVerifications(page int, status int) ([]Proposal, int, []string, error)
	GetVerification(proposalId uint) (Proposal, string, error)

	ChangeStatus(id uint, status int) error
}

type Services interface {
	Create(token *jwt.Token, image *multipart.FileHeader, document *multipart.FileHeader, data Proposal) error
	Update(token *jwt.Token, proposal_id string, image *multipart.FileHeader, document *multipart.FileHeader, data Proposal) error
	GetAll(page string) ([]Proposal, int, error)
	GetAllMy(token *jwt.Token, page string) ([]Proposal, int, error)
	GetDetail(id_proposal string) (Proposal, error)
	Delete(token *jwt.Token, prososal_id string) error
	Archive() error
	GetVerifications(page int, status int) ([]Proposal, int, []string, error)
	GetVerification(proposalId uint) (Proposal, string, error)

	ChangeStatus(id uint, status int) error
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
