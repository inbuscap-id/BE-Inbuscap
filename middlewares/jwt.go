package middlewares

import (
	"BE-Inbuscap/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id string, status int, role bool) (string, error) {
	var data = jwt.MapClaims{}
	data["id"] = id
	data["is_active"] = status
	data["is_admin"] = role
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var proccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := proccessToken.SignedString([]byte(config.JWTSECRET))
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

func DecodeToken(i interface{}) string {
	var claim = i.(*jwt.Token).Claims.(jwt.MapClaims)
	var result string

	if val, found := claim["id"]; found {
		result = val.(string)
	}

	return result
}

func DecodeRole(i interface{}) (bool, error) {
	var claim = i.(*jwt.Token).Claims.(jwt.MapClaims)
	var result bool

	if val, found := claim["is_admin"]; found {
		result = val.(bool)
		return result, nil
	}

	return false, nil
}
func DecodeStatus(i interface{}) (int, error) {
	var claim = i.(*jwt.Token).Claims.(jwt.MapClaims)
	var result int

	if val, found := claim["is_active"]; found {
		result = val.(int)
		return result, nil
	}

	return 0, nil
}
