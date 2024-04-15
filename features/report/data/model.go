package data

import (
	"gorm.io/gorm"
)

type report struct {
	gorm.Model
	Title   string
	Image   string
	Caption string
	Nominal uint64
}
