package domain

import (
	"maps"
	"time"
)

type Priority int

const (
	Low Priority = iota
	High
	Medium
	Critical
)

func (p Priority) string() string {
	switch p {
	case Low:
		return "LOW"
	case Medium:
		return "MEDIUM"
	case Critical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

type TaskStatus string

const (
	TaskPending   TaskStatus = "PENDING"
	TaskScheduled TaskStatus = "SCHEDULED"
	TaskCompleted TaskStatus = "COMPLETED"
	TaskBlocked   TaskStatus = "BLOCKED"
	TaskCancelled TaskStatus = "CANCELLED"
)

type Task struct {
	ID                string
	Title             string
	Description       string
	Priority          Priority
	Status            TaskStatus
	EstimatedDuration time.Duration
	RemainingDuration time.Duration

	ExpectedRevenue Money
	ExpectedCost    Money
	GravityScore    int

	CanRunConcurrent bool
	Dependencies     []string
	ResourceType     string

	CreatedAt time.Time
	DueDate   *time.Time
}

func NewTask(
	id string,
	title string,
	priority Priority,
	duration time.Duration,
	expectedRevenue Money,
	expectedCost Money,
	gravityScore int,
	canRunConcurrent bool,
	dependencies []string,
	resourceType string,
	createdAt time.Time,
	dueDate *time.Time,
) Task {
	return Task{
		ID:                id,
		Title:             title,
		Priority:          priority,
		Status:            TaskPending,
		EstimatedDuration: duration,
		RemainingDuration: duration,
		ExpectedRevenue:   expectedRevenue,
		ExpectedCost:      expectedCost,
		GravityScore:      gravityScore,
		CanRunConcurrent:  canRunConcurrent,
		Dependencies:      dependencies,
		ResourceType:      resourceType,
		CreatedAt:         createdAt,
		DueDate:           dueDate,
	}

}

func (t Task) ExpectedProfit() Money {
	return t.ExpectedRevenue.Sub(t.ExpectedCost)
}

func (t Task) IsReady(completed map[string]bool) bool {
	for _, dep := range t.Dependencies {
		if !completed[dep] {
			return false
		}
	}
	return true
}
