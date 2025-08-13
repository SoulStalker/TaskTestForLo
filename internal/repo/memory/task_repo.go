package memory

import (
	"context"
	"sync"
	"time"

	"github.com/soulstalker/task-api/internal/domain"
	"github.com/soulstalker/task-api/internal/usecase"
)

type TaskRepoIM struct {
	mu     sync.RWMutex
	data   map[uint]domain.Task
	nextID uint
}

func NewTaskRepoIM() *TaskRepoIM {
	r := &TaskRepoIM{data: make(map[uint]domain.Task)}
	r.nextID = 1
	return r
}

func (r *TaskRepoIM) Create(ctx context.Context, t domain.Task) (domain.Task, error) {
	select {
	case <-ctx.Done():
		return domain.Task{}, ctx.Err()
	default:
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID - 1
	now := time.Now().UTC()
	t.Create(now, id)
	r.nextID++
	r.data[id] = t
	return t, nil
}

func (r *TaskRepoIM) GetById(ctx context.Context, id uint) (domain.Task, error) {
	select {
	case <-ctx.Done():
		return domain.Task{}, ctx.Err()
	default:
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, ok := r.data[id]
	if !ok {
		return domain.Task{}, usecase.ErrNotFound
	}
	return t, nil
}

func (r *TaskRepoIM) All(ctx context.Context) ([]domain.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]domain.Task, 0, len(r.data))
	for _, task := range r.data {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
