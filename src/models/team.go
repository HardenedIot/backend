package models

type Team struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	TeamName    string `json:"team_name" gorm:"unique" validate:"required,min=3,max=50"`
	TeamID      string `json:"team_id" gorm:"unique" validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"max=200"`
	Users       []User `json:"users" gorm:"many2many:user_teams;" validate:"-"`
	Private     bool   `json:"private"`
}

type PatchTeamRequest struct {
	TeamName    *string   `json:"team_name,omitempty"`
	TeamID      *string   `json:"team_id,omitempty"`
	Description *string   `json:"description,omitempty"`
	Users       *[]string `json:"users,omitempty"`
	Private     *bool     `json:"private,omitempty"`
}
