package domain

import "time"

const (
	StatusTodo  = "todo"
	StatusDoing = "doing"
	StatusDone  = "done"
)

type Task struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepo interface {
	All() []Task
	GetById(id uint) (Task, error)
	Create(task Task) (Task, error)
}

// func (t *Task) Create(now time.Time, id uint) {
// 	t.ID = id
// 	t.CreatedAt = now
// 	t.UpdatedAt = now
// 	if t.Status == "" {
// 		t.Status = StatusTodo
// 	}
// }
