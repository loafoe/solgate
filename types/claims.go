package types

type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Verified bool   `json:"email_verified"`
}
