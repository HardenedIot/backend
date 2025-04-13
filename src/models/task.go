package models

import "github.com/google/uuid"

type RiskLevel int

const (
	Low RiskLevel = 1 + iota
	Medium
	High
)

type Task struct {
	TaskID      uuid.UUID  `json:"task_id"`
	Technology  Technology `json:"technology"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	RiskLevel   RiskLevel  `json:"risk_level"`
	Completed   bool       `json:"completed"`
}
