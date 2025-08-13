package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soulstalker/task-api/internal/domain"
	"github.com/soulstalker/task-api/internal/logger"
	"github.com/soulstalker/task-api/internal/usecase"
)

type Handler struct {
	UC     *usecase.TaskUC
	Logger *logger.AsyncLogger
}

func NewHandler(uc *usecase.TaskUC, lg *logger.AsyncLogger) *Handler {
	return &Handler{UC: uc, Logger: lg}
}

// GET /tasks
func (h *Handler) All(c *gin.Context) {
	ctx := c.Request.Context()
	statusParam := c.Query("status")
	var status *domain.Status
	if statusParam != "" {
		if s, err := domain.ParseStatus(statusParam); err == nil {
			status = &s
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
			return
		}
	}
	tasks, err := h.UC.All(ctx, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
	h.Logger.Log("tasks_list", gin.H{"status": c.Query("status"), "count": len(tasks)})

}

// GET /tasks/:id
func (h *Handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.UC.GetById(ctx, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": usecase.ErrNotFound})
		return
	}
	c.JSON(http.StatusOK, task)
	h.Logger.Log("task_get", gin.H{"id": id})

}

// Create /tasks
func (h *Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var task domain.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask, err := h.UC.Create(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
	h.Logger.Log("task_created", gin.H{"id": newTask.ID, "status": newTask.Status})

}
