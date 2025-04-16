package models

type Project struct {
	ID           uint64       `json:"id" gorm:"primaryKey"`
	ProjectID    string       `json:"project_id" gorm:"unique" validate:"required,min=1,max=50"`
	TeamID       string       `json:"team_id" validate:"required"`
	Team         *Team        `json:"team" gorm:"foreignKey:TeamID;references:TeamID" validate:"-"`
	ProjectName  string       `json:"project_name" gorm:"unique" validate:"required,min=3,max=50"`
	Private      *bool        `json:"private"`
	Description  string       `json:"description" validate:"max=500"`
	URL          string       `json:"url"`
	Technologies []Technology `json:"technologies" validate:"required"`
}

type PatchProjectRequest struct {
	ProjectName  *string       `json:"project_name,omitempty"`
	TeamID       *string       `json:"team_id,omitempty"`
	Private      *bool         `json:"private,omitempty"`
	Description  *string       `json:"description,omitempty"`
	URL          *string       `json:"url,omitempty"`
	Technologies *[]Technology `json:"technologies,omitempty"`
}
