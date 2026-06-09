package services

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"TravelSphere/models"
	"TravelSphere/utils"
)


const minPasswordLen = 6

func RegisterUser(username, password string) (models.User, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return models.User{}, fmt.Errorf("username is required")
	}
	if len(password) < minPasswordLen {
		return models.User{}, fmt.Errorf("password must be at least %d characters", minPasswordLen)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username:     username,
		PasswordHash: string(hash),
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}
	
	if err := utils.CreateUser(user); err != nil {
		return models.User{}, err
	}
	return user, nil
}


func AuthenticateUser(username, password string) (string, error) {
	user, found, err := utils.FindUser(strings.TrimSpace(username))
	if err != nil {
		return "", err
	}
	if !found {
		return "", fmt.Errorf("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	return user.Username, nil
}
