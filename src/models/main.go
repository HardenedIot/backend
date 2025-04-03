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

type Team struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	TeamName    string `json:"team_name" gorm:"unique" validate:"required,min=3,max=50"`
	TeamID      string `json:"team_id" gorm:"unique" validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"max=200"`
	Users       []User `json:"users" gorm:"many2many:user_teams;"`
	Private     bool   `json:"private"`
}

type Project struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	ProjectName string `json:"project_name" gorm:"unique" validate:"required,min=3,max=50"`
	ProjectID   string `json:"project_id" gorm:"unique" validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"max=200"`
	TeamID      string `json:"team_id" validate:"required"`
	Team        Team   `json:"team" gorm:"foreignKey:TeamID"`
	Private     bool   `json:"private"`
}
