package services

import (
	"errors"
	"os"

	"github.com/herEmanuel/go-api/models"
	"github.com/herEmanuel/go-api/repositories"
	"github.com/herEmanuel/go-api/utils"
)

//Register service
func Register(bodyFields map[string]string) (map[string]interface{}, error) {

	var newUser models.User

	err := repositories.Register(bodyFields, &newUser)
	if err != nil {
		return nil, err
	}

	token, err := utils.CreateToken(newUser.ID, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	emailToken, err := utils.CreateToken(newUser.ID, os.Getenv("JWT_EMAIL_KEY"))
	if err != nil {
		return nil, err
	}

	htmlBody := "Access this link in order to verify your email: <a href=\"localhost:5000/verify/" + emailToken + "\">Verify</a>"
	err = utils.SendEmail([]string{newUser.Email}, "Verify your email\n", htmlBody)
	if err != nil {
		return nil, err
	}

	newUserAndToken := make(map[string]interface{})
	newUserAndToken["newUser"] = newUser
	newUserAndToken["token"] = token

	return newUserAndToken, nil
}

//Login service
func Login(bodyFields map[string]string) (map[string]interface{}, error) {

	var user models.User

	err := repositories.Login(bodyFields, &user)
	if err != nil {
		return nil, err
	}

	token, err := utils.CreateToken(user.ID, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	userAndToken := make(map[string]interface{})
	userAndToken["user"] = user
	userAndToken["token"] = token

	return userAndToken, nil
}

func ChangePassword(bodyFields map[string]string) error {

	if bodyFields["newPassword"] == bodyFields["confirmationPassword"] {
		return errors.New("The password don't match")
	}

	err := repositories.ChangePassword(bodyFields)
	if err != nil {
		return err
	}

	return nil
}
