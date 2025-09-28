package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// TODO: Not sure if TaskType needs to be its own type. Will consider changing later.
// TaskType represents the type of task
type TaskType string

const (
	TaskTypeEmail          TaskType = "email"
	TaskTypeVideoProcess   TaskType = "video_process"
	TaskTypeReportGenerate TaskType = "report_generate"
)

// TaskStatus represents the current state of the task
type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusProcessing TaskStatus = "processing"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
	TaskStatusDead       TaskStatus = "dead" // After max retry attempts
)

// TaskPriority represents the priority level of the task for queue ordering
type TaskPriority int

const (
	PriorityLow TaskPriority = iota
	PriorityNormal
	PriorityHigh
	PriorityUrgent
)

type Task struct {
	ID       uuid.UUID
	Type     TaskType
	Payload  json.RawMessage
	Status   TaskStatus
	Priority TaskPriority

	// TODO: Might get rid of TaskMetadata and replace with flat struct fields
	// Metadata
	Metadata TaskMetadata
}

type TaskMetadata struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Attempts    int
	MaxAttempts int
	ErrorMsg    string
}
