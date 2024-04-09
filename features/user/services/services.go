package service

import (
	"BE-Inbuscap/features/user"
	"BE-Inbuscap/helper"
	"BE-Inbuscap/middlewares"
	utils "BE-Inbuscap/utils/cloudinary"
	"errors"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	model user.Model
	pm    helper.PasswordManager
	v     *validator.Validate
}

func NewService(m user.Model) user.Service {
	return &service{
		model: m,
		pm:    helper.NewPasswordManager(),
		v:     validator.New(),
	}
}

func (s *service) Register(register_data user.Register) error {
	// Check Validate
	var validate user.Register
	helper.ConvertStruct(&register_data, &validate)
	err := s.v.Struct(&validate)
	if err != nil {
		log.Println(err.Error())
		log.Println(register_data)
		return errors.New(helper.ErrorInvalidValidate)
	}

	// Hashing Password
	newPassword, err := s.pm.HashPassword(register_data.Password)
	if err != nil {
		return errors.New(helper.ErrorGeneralServer)
	}

	var new_user user.User
	helper.ConvertStruct(&register_data, &new_user)
	new_user.Password = newPassword

	// Do Register
	err = s.model.Register(new_user)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return errors.New(mysqlErr.Message)
		}
		return errors.New(helper.ErrorGeneralServer)
	}

	return nil
}

func (s *service) Login(login_data user.User) (string, error) {
	// Do Login & Get Password
	dbData, err := s.model.Login(login_data.Email)
	if err != nil {
		return "", errors.New(helper.ErrorDatabaseNotFound)
	}

	// Compare Password
	if err := s.pm.ComparePassword(login_data.Password, dbData.Password); err != nil {
		return "", errors.New(helper.ErrorUserCredential)
	}

	// Create Token
	token, err := middlewares.GenerateJWT(strconv.Itoa(int(dbData.ID)), dbData.IsActive, dbData.IsAdmin)
	if err != nil {
		return "", errors.New(helper.ErrorGeneralServer)
	}

	// Finished
	return token, nil
}

func (s *service) Profile(token *jwt.Token) (user.User, error) {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Get Profile
	result, err := s.model.Profile(decodeID)
	if err != nil {
		return user.User{}, err
	}

	// Finished
	// result.CreatedAt = result.CreatedAt.UTC()
	// result.UpdatedAt = result.UpdatedAt.UTC()
	return result, nil
}

func (s *service) Update(token *jwt.Token, update_data user.User, avatar *multipart.FileHeader) error {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Check Validate Password & Others
	var validate user.Update
	helper.ConvertStruct(&update_data, &validate)
	err := s.v.Struct(&validate)
	if err != nil {
		if strings.Contains(err.Error(), "Fullname") {
			update_data.Fullname = ""
		}
		if strings.Contains(err.Error(), "Email") {
			update_data.Email = ""
		}
		if strings.Contains(err.Error(), "Handphone") {
			update_data.Handphone = ""
		}
		if strings.Contains(err.Error(), "Password") {
			update_data.Password = ""
		}
		if strings.Contains(err.Error(), "ktp") {
			update_data.KTP = ""
		}
		if strings.Contains(err.Error(), "npwp") {
			update_data.NPWP = ""
		}
	}

	// Convert id to uint
	id_int, err := strconv.ParseUint(decodeID, 10, 32)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}
	update_data.ID = uint(id_int)

	// Hashing Password
	if update_data.Password != "" {
		newPassword, err := s.pm.HashPassword(update_data.Password)
		if err != nil {
			return errors.New(helper.ErrorGeneralServer)
		}
		update_data.Password = newPassword
	}
	// update avatar bila ada
	log.Println(avatar.Filename)
	if avatar != nil {
		link, err := utils.UploadImage(avatar)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update_data.Avatar = link
	}

	// Update Data
	if err := s.model.Update(update_data); err != nil {
		return err
	}

	// Finished
	return nil
}

func (s *service) Delete(token *jwt.Token) error {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Delete Date
	if err := s.model.Delete(decodeID); err != nil {
		return err
	}

	// Finished
	return nil
}

func (s *service) AddVerification(token *jwt.Token, uploads []*multipart.FileHeader) error {
	var links []string
	for _, val := range uploads {
		link, err := utils.UploadImage(val)
		if err != nil {
			log.Println("error di service:", err.Error())
			return err
		}
		links = append(links, link)
	}
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Get Profile
	result, err := s.model.Profile(decodeID)
	if err != nil {
		log.Println("error di service:", err.Error())

		return err
	}
	result.PhotoKTP = links[0]
	result.PhotoNPWP = links[1]
	result.PhotoSelf = links[2]

	if err := s.model.Update(result); err != nil {
		log.Println("error di service:", err.Error())

		return err
	}
	return nil
}

func (s *service) GetVerifications(paginasi helper.Pagination, status int) ([]user.User, int, error) {
	result, count, err := s.model.GetVerifications(paginasi, status)
	if err != nil {
		log.Println(err.Error(), "service")
		return nil, 0, err
	}
	return result, count, nil
}

func (s *service) ChangeStatus(userID uint, status int) error {
	err := s.model.ChangeStatus(userID, status)
	return err
}
