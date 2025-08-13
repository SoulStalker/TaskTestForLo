package domain

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Status string

const (
	StatusTodo  Status = "todo"
	StatusDoing Status = "doing"
	StatusDone  Status = "done"
)

type Task struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"  binding:"required"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepo interface {
	All(ctx context.Context, status *Status) ([]Task, error)
	GetById(ctx context.Context, id uint) (Task, error)
	Create(ctx context.Context, task Task) (Task, error)
}

func (t *Task) Create(now time.Time, id uint) {
	t.ID = id
	t.CreatedAt = now
	t.UpdatedAt = now
	if t.Status == "" {
		t.Status = StatusTodo
	}
}

func ParseStatus(s string) (Status, error) {
	switch strings.ToLower(s) {
	case string(StatusTodo):
		return StatusTodo, nil
	case string(StatusDoing):
		return StatusDoing, nil
	case string(StatusDone):
		return StatusDone, nil
	default:
		return "", fmt.Errorf("invalid status: %q", s)
	}
}
