package models

type RiskLevel int

const (
	Low RiskLevel = 1 + iota
	Medium
	High
)

type Task struct {
	TaskID      string     `json:"task_id" bson:"task_id" validate:"required"`
	Technology  Technology `json:"technology" bson:"technology" validate:"required"`
	Name        string     `json:"name" bson:"name" validate:"required,min=3,max=100"`
	Description string     `json:"description" bson:"description" validate:"max=500"`
	RiskLevel   RiskLevel  `json:"risk_level" bson:"risk_level" validate:"required,min=1,max=3"`
	Completed   bool       `json:"completed" bson:"completed"`
	Ignored     bool       `json:"ignored" bson:"ignored"`
}

type PatchTaskRequest struct {
	TaskID      string      `json:"task_id" bson:"task_id"`
	Technology  *Technology `json:"technology,omitempty" bson:"technology,omitempty"`
	Name        *string     `json:"name,omitempty" bson:"name,omitempty"`
	Description *string     `json:"description,omitempty" bson:"description,omitempty"`
	RiskLevel   *RiskLevel  `json:"risk_level,omitempty" bson:"risk_level,omitempty"`
	Completed   *bool       `json:"completed,omitempty" bson:"completed,omitempty"`
	Ignored     *bool       `json:"ignored,omitempty" bson:"ignored,omitempty"`
}
