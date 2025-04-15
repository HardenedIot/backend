package models

type RiskLevel int

const (
	Low RiskLevel = 1 + iota
	Medium
	High
)

type Task struct {
	TaskID      string     `json:"task_id" validate:"required"`
	Technology  Technology `json:"technology" validate:"required"`
	Name        string     `json:"name" validate:"required,min=3,max=100"`
	Description string     `json:"description" validate:"max=500"`
	RiskLevel   RiskLevel  `json:"risk_level" validate:"required,min=1,max=3"`
	Completed   bool       `json:"completed"`
	Ignored     bool       `json:"ignored"`
}

type PatchTaskRequest struct {
	TaskID      *string     `json:"task_id,omitempty"`
	Technology  *Technology `json:"technology,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	RiskLevel   *RiskLevel  `json:"risk_level,omitempty"`
	Completed   *bool       `json:"completed,omitempty"`
	Ignored     bool        `json:"ignored,omitempty"`
}
