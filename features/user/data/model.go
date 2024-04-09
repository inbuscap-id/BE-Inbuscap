package data

import (
	"BE-Inbuscap/features/invest"
	"BE-Inbuscap/features/proposal"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname     string
	Email        string `gorm:"unique"`
	Handphone    string `gorm:"unique"`
	KTP          string `gorm:"unique"`
	NPWP         string
	Password     string
	PhotoKTP     string
	PhotoNPWP    string
	PhotoSelf    string
	IsActive     int
	IsAdmin      bool
	Saldo        int
	Proposals    []Proposal
	Investments  []Investment
	Transactions []Transaction
	Avatar       string
}

type Proposal proposal.Proposal

type Investment invest.Investment

type Report struct {
	gorm.Model
	Proposal_id uint `gorm:"primarykey"`
	Document    string
}

type Transaction struct {
	gorm.Model
	OrderID string `gorm:"unique"`
	UserID  uint
	Amount  int
	Status  string
	Token   string
	Url     string
	User    User `gorm:"foreignKey:UserID"`
}
