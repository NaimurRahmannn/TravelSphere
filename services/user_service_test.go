package services

import (
	"path/filepath"
	"testing"

	"TravelSphere/utils"
)
func redirectUsers(t *testing.T) {
	t.Helper()
	restore := utils.SetUserFile(filepath.Join(t.TempDir(), "users.json"))
	t.Cleanup(restore)
}

func TestRegisterUser_Success(t *testing.T) {
	redirectUsers(t)

	user, err := RegisterUser("alice", "secret123")
	if err != nil {
		t.Fatalf("register failed: %v", err)
	}
	if user.Username != "alice" {
		t.Errorf("got username %q, want alice", user.Username)
	}
	if user.PasswordHash == "secret123" || user.PasswordHash == "" {
		t.Error("password was not hashed")
	}
}

func TestRegisterUser_EmptyUsername(t *testing.T) {
	redirectUsers(t)

	if _, err := RegisterUser("   ", "secret123"); err == nil {
		t.Error("expected error for empty username, got nil")
	}
}

func TestRegisterUser_ShortPassword(t *testing.T) {
	redirectUsers(t)

	if _, err := RegisterUser("alice", "123"); err == nil {
		t.Error("expected error for short password, got nil")
	}
}

func TestRegisterUser_Duplicate(t *testing.T) {
	redirectUsers(t)

	if _, err := RegisterUser("alice", "secret123"); err != nil {
		t.Fatalf("first register failed: %v", err)
	}
	if _, err := RegisterUser("alice", "another123"); err == nil {
		t.Error("expected duplicate username error, got nil")
	}
}

func TestAuthenticateUser_Success(t *testing.T) {
	redirectUsers(t)

	RegisterUser("alice", "secret123")

	username, err := AuthenticateUser("alice", "secret123")
	if err != nil {
		t.Fatalf("auth failed: %v", err)
	}
	if username != "alice" {
		t.Errorf("got %q, want alice", username)
	}
}

func TestAuthenticateUser_WrongPassword(t *testing.T) {
	redirectUsers(t)

	RegisterUser("alice", "secret123")

	if _, err := AuthenticateUser("alice", "wrongpass"); err == nil {
		t.Error("expected error for wrong password, got nil")
	}
}

func TestAuthenticateUser_UnknownUser(t *testing.T) {
	redirectUsers(t)

	if _, err := AuthenticateUser("ghost", "whatever"); err == nil {
		t.Error("expected error for unknown user, got nil")
	}
}