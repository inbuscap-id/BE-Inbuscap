package data

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Title   string
	Image   string
	Caption string
	Nominal uint64
}
