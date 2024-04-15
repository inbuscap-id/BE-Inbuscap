package data

import (
	"BE-Inbuscap/features/user"
	"BE-Inbuscap/helper"
	"errors"
	"log"
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
	newData.IsActive = 0
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
	err := m.connection.Where("id = ?", id).First(&result).Error

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
	if data.PhotoKTP != "" {
		selectUpdate = append(selectUpdate, "photo_ktp")
	}
	if data.PhotoNPWP != "" {
		selectUpdate = append(selectUpdate, "photo_npwp")
	}
	if data.PhotoSelf != "" {
		selectUpdate = append(selectUpdate, "photo_self")
	}
	if data.Avatar != "" {
		selectUpdate = append(selectUpdate, "avatar")
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
	if query := m.connection.Where("id = ?", id).Delete(&user.User{}); query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorDatabaseNotFound)
	}
	return nil
}

func (m *model) GetVerifications(paginasi helper.Pagination, status int) ([]user.User, int, error) {
	var proses = new([]User)
	var count int64
	offset := (paginasi.Page - 1) * paginasi.PageSize
	if err := m.connection.Find(&proses).Where(" is_active = ? AND is_admin =? ", status, false).Count(&count).Error; err != nil {
		log.Println("repo error: ", err.Error())
		return nil, 0, err
	}

	var selected = new([]User)
	if err := m.connection.Order("updated_at desc").Where(" is_active = ? AND is_admin =? ", status, false).Offset(offset).
		Limit(paginasi.PageSize).Find(&selected).Error; err != nil {
		log.Println("repo error: ", err.Error())
		return nil, 0, err
	}
	var results []user.User
	for _, val := range *selected {
		var result = user.User{
			Fullname:  val.Fullname,
			Handphone: val.Handphone,
			KTP:       val.KTP,
			NPWP:      val.NPWP,
			PhotoKTP:  val.PhotoKTP,
			PhotoNPWP: val.PhotoNPWP,
			PhotoSelf: val.PhotoSelf,
			IsActive:  val.IsActive,
		}
		result.ID = val.ID
		results = append(results, result)

	}
	return results, int(count), nil
}

func (m *model) ChangeStatus(userID uint, status int) error {
	var result user.User
	err := m.connection.Where("id = ?", userID).First(&result).Error
	if err != nil {
		log.Println("error mengambil data", err.Error())
		return err
	}
	result.IsActive = status
	err = m.connection.Save(&result).Error

	return err
}
