package models

type Project struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	ProjectID   string `json:"project_id" gorm:"unique" validate:"required,min=1,max=50"`
	TeamID      string `json:"team_id" validate:"required"`
	Team        Team   `json:"team" gorm:"foreignKey:TeamID"`
	ProjectName string `json:"project_name" gorm:"unique" validate:"required,min=3,max=50"`
	Private     bool   `json:"private"`
	Description string `json:"description" validate:"max=200"`
	URL         string `json:"url"`
}
