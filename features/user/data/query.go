package data

import (
	"BE-Inbuscap/features/user"
	"BE-Inbuscap/helper"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) user.Model {
	return &model{
		connection: db,
	}
}

func (m *model) Register(newData user.User) error {
	newData.CreatedAt = time.Now().UTC()
	newData.UpdatedAt = time.Now().UTC()
	newData.IsVerified = false
	newData.IsAdmin = false
	return m.connection.Create(&newData).Error
}

func (m *model) Login(input string) (user.User, error) {
	var result user.User
	err := m.connection.Where("email = ? OR handphone = ?", input, input).First(&result).Error
	return result, err
}

func (m *model) Profile(id string) (user.User, error) {
	var result user.User
	err := m.connection.Where("id = ?", id).Find(&result).Error
	return result, err
}

func (m *model) Update(data user.User) error {
	var selectUpdate []string
	if data.Fullname != "" {
		selectUpdate = append(selectUpdate, "fullname")
	}
	if data.Email != "" {
		selectUpdate = append(selectUpdate, "email")
	}
	if data.Handphone != "" {
		selectUpdate = append(selectUpdate, "handphone")
	}
	if data.KTP != "" {
		selectUpdate = append(selectUpdate, "ktp")
	}
	if data.NPWP != "" {
		selectUpdate = append(selectUpdate, "npwp")
	}
	if data.Password != "" {
		selectUpdate = append(selectUpdate, "password")
	}
	if len(selectUpdate) == 0 {
		return errors.New(helper.ErrorNoRowsAffected)
	}

	data.UpdatedAt = time.Now().UTC()

	if query := m.connection.Model(&data).Select(selectUpdate).Updates(&data); query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorNoRowsAffected)
	}
	return nil
}

func (m *model) Delete(id string) error {
	if query := m.connection.Table("users").Where("id = ?", id).Select("email", "handphone", "ktp", "npwp").Updates(user.User{Email: "", Handphone: "", KTP: "", NPWP: ""}); query.Error != nil {
		fmt.Println(query.Error)
		return errors.New(helper.ErrorDatabaseNotFound)
	} else {
		fmt.Println(query.RowsAffected)
	}
	if query := m.connection.Where("id = ?", id).Delete(&user.User{}); query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorDatabaseNotFound)
	}
	return nil
}
