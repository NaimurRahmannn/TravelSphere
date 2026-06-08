package models

// User is a registered account. Only the bcrypt hash of the password is stored;
// the plaintext password never touches disk.
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    string `json:"created_at"`
}
