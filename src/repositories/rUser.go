package repositories

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/herEmanuel/go-api/database"
	"github.com/herEmanuel/go-api/models"
)

//Register a user (in the database)
func Register(bodyFields map[string]string, varToStore *models.User) error {

	var user models.User
	database.Db.Where("email = ?", bodyFields["email"]).First(&user)
	if user.ID != 0 {
		return errors.New("This email already exists")
	}

	desiredPassword := []byte(bodyFields["password"])
	hashedPassword, err := bcrypt.GenerateFromPassword(desiredPassword, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Could not hash the new password")
	}

	newUser := models.User{
		Name:     bodyFields["name"],
		Email:    bodyFields["email"],
		Password: string(hashedPassword),
		Country:  bodyFields["country"],
	}

	result := database.Db.Create(&newUser)
	if result.Error != nil {
		return errors.New("Could not save the new user")
	}

	*varToStore = newUser

	return nil
}

//Login logs in a user
func Login(bodyFields map[string]string, varToStore *models.User) error {

	var user models.User

	database.Db.Where("email = ?", bodyFields["email"]).First(&user)
	if user.ID == 0 {
		return errors.New("Wrong email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bodyFields["password"]))
	if err != nil {
		return errors.New("Wrong email or password")
	}

	*varToStore = user

	return nil
}

//ChangePassword changes user's password
func ChangePassword(bodyFields map[string]string) error {

	var user models.User
	database.Db.First(&user, bodyFields["id"])

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bodyFields["oldPassword"]))
	if err != nil {
		return errors.New("Wrong password")
	}

	if bodyFields["newPassword"] == bodyFields["oldPassword"] {
		return errors.New("The new password can not be the same as the old one")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bodyFields["newPassword"]), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Could not hash the new password")
	}

	user.Password = string(hashedPassword)

	result := database.Db.Save(&user)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}
