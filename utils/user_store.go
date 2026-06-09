package utils

import (
	"TravelSphere/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"path/filepath"
)


var userFile = "data/users.json"


var userMu sync.Mutex


func ReadUsers() ([]models.User, error) {
	userMu.Lock()
	defer userMu.Unlock()
	return readUsersUnlocked()
}

// Used when the caller already holds userMu.
func readUsersUnlocked() ([]models.User, error) {
	bytes, err := os.ReadFile(userFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.User{}, nil
		}
		return nil, err
	}
	if len(bytes) == 0 {
		return []models.User{}, nil
	}
	var users []models.User
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// writeUsersUnlocked persists without locking,for reuse under a held lock.
func writeUsersUnlocked(users []models.User) error {
	if err := os.MkdirAll(filepath.Dir(userFile), 0o755); err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(userFile, bytes, 0o644)
}


func FindUser(username string) (models.User, bool, error) {
	users, err := ReadUsers()
	if err != nil {
		return models.User{}, false, err
	}
	for _, u := range users {
		if strings.EqualFold(u.Username, username) {
			return u, true, nil
		}
	}
	return models.User{}, false, nil
}

func CreateUser(user models.User) error {
	userMu.Lock()
	defer userMu.Unlock()

	users, err := readUsersUnlocked()
	if err != nil {
		return err
	}
	for _, u := range users {
		if strings.EqualFold(u.Username, user.Username) {
			return fmt.Errorf("username %q is already taken", user.Username)
		}
	}
	users = append(users, user)
	return writeUsersUnlocked(users)
}
