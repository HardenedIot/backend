package models

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique" validate:"required,min=3,max=20"`
	Name     string `json:"name" validate:"required,min=1,max=50"`
	Surname  string `json:"surname" validate:"required,min=1,max=50"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Private  bool   `json:"private"`
	Teams    []Team `json:"teams" gorm:"many2many:user_teams;"`
}
