package models

// User is the type interface of a user.
type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// TableName returns string of "user".
func (b *User) TableName() string {
	return "user"
}
