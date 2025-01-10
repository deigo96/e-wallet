package models

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) TableName() string {
	return "users"
}

func NewUser(name, email string) *User {
	return &User{Name: name, Email: email}
}

func NewUserWithID(id int, name, email string) *User {
	return &User{ID: id, Name: name, Email: email}
}

func (u *User) GetID() int {
	return u.ID
}
