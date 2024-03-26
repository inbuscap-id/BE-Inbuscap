package data

import (
	invest "BE-Inbuscap/features/invest/data"
	post "BE-Inbuscap/features/post/data"

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
	Verification Verification
	Posts        []post.Post
	Invest       []*invest.Invest `gorm:"many2many:user_invest;"`
	IsVerified   bool
}

type Admin struct {
	Fullname string
	Email    string `gorm:"unique"`
	Password string
}

type Verification struct {
	gorm.Model
	PhotoSelf string
	PhotoKTP  string
	PhotoNPWP string
	UserID    uint
	Submit    bool
}

type Proposal struct {
	gorm.Model
	//tanda pengenal item di publik, pengganti ID
	//format: [nomerid]/[nama usaha (disingkat)]/[bulan dalam bentuk romawi]/[tahun]
	//ex: 1923/AJY/X/2024
	SerialNumber  string
	Title         string
	Category      string
	Image         string
	Description   string
	Nominal       uint64
	Document      string //dokumen tentang detail usaha lebih lanjut
	Collected     uint64 //dana yang sudah terkumpul
	InvestorShare uint   //pembagian keuntungan buat investor
	OwnerShare    uint
}
