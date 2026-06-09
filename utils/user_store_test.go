package utils

import (
	"os"
	"path/filepath"
	"testing"

	"TravelSphere/models"
)
func redirectUsers(t *testing.T) {
	t.Helper()
	original := userFile
	userFile = filepath.Join(t.TempDir(), "users.json")
	t.Cleanup(func() { userFile = original })
}

func TestReadUsers_MissingFile(t *testing.T) {
	redirectUsers(t)

	users, err := ReadUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("expected empty list, got %d", len(users))
	}
}

func TestReadUsers_EmptyFile(t *testing.T) {
	redirectUsers(t)

	if err := os.WriteFile(userFile, []byte{}, 0o644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	users, err := ReadUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("expected empty list, got %d", len(users))
	}
}

func TestReadUsers_CorruptJSON(t *testing.T) {
	redirectUsers(t)

	if err := os.WriteFile(userFile, []byte("{not valid"), 0o644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	if _, err := ReadUsers(); err == nil {
		t.Error("expected a decode error, got nil")
	}
}

func TestReadUsers_ReadError(t *testing.T) {
	redirectUsers(t)
	dir := filepath.Join(t.TempDir(), "adir")
	if err := os.Mkdir(dir, 0o755); err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	userFile = dir

	if _, err := ReadUsers(); err == nil {
		t.Error("expected a read error, got nil")
	}
}

func TestCreateUser_AndFind(t *testing.T) {
	redirectUsers(t)

	u := models.User{Username: "beta", PasswordHash: "hash", CreatedAt: "2026-01-01T00:00:00Z"}
	if err := CreateUser(u); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	found, ok, err := FindUser("BETA")
	if err != nil {
		t.Fatalf("find error: %v", err)
	}
	if !ok {
		t.Fatal("expected to find the user, got ok=false")
	}
	if found.Username != "beta" {
		t.Errorf("got username %q, want beta", found.Username)
	}
}

func TestFindUser_NotFound(t *testing.T) {
	redirectUsers(t)

	_, ok, err := FindUser("ghost")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ok {
		t.Error("expected ok=false for a missing user")
	}
}

func TestCreateUser_Duplicate(t *testing.T) {
	redirectUsers(t)

	u := models.User{Username: "beta", PasswordHash: "hash"}
	if err := CreateUser(u); err != nil {
		t.Fatalf("first create failed: %v", err)
	}
	dup := models.User{Username: "Beta", PasswordHash: "other"}
	if err := CreateUser(dup); err == nil {
		t.Error("expected duplicate username to be rejected, got nil")
	}
}

func TestFindUser_ReadError(t *testing.T) {
	redirectUsers(t)
	if err := os.WriteFile(userFile, []byte("{bad"), 0o644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	if _, _, err := FindUser("anyone"); err == nil {
		t.Error("expected FindUser to propagate the read error, got nil")
	}
}

func TestCreateUser_ReadError(t *testing.T) {
	redirectUsers(t)

	if err := os.WriteFile(userFile, []byte("{bad"), 0o644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	if err := CreateUser(models.User{Username: "x"}); err == nil {
		t.Error("expected CreateUser to propagate the read error, got nil")
	}
}
