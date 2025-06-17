package handlers

import (
	"context"
	"errors"
	"github.com/Killazius/workmate/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Service interface {
	CreateTask(ctx context.Context) (*domain.Task, error)
	GetTask(ctx context.Context, id int64) (*domain.Task, error)
	DeleteTask(ctx context.Context, id int64) error
}
type TaskHandler struct {
	service Service
	log     *zap.SugaredLogger
}

func New(s Service, log *zap.SugaredLogger) *TaskHandler {
	return &TaskHandler{
		service: s,
		log:     log,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	task, err := h.service.CreateTask(ctx)
	if err != nil {
		renderError(w, r, "failed to create task", http.StatusInternalServerError)
		h.log.Error(err)
		return
	}
	h.log.Infow("created task", "id", task.ID)
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, newTaskResponse(task))
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID, err := getTaskID(r)
	if err != nil {
		h.log.Error(err)
		renderError(w, r, "failed to get task ID", http.StatusInternalServerError)
		return
	}
	h.log.Infow("get task", "id", taskID)
	task, err := h.service.GetTask(ctx, taskID)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			renderError(w, r, "task not found", http.StatusNotFound)
		} else {
			renderError(w, r, "failed to get task", http.StatusInternalServerError)
		}
		h.log.Errorw("failed to get task", "error", err)
		return
	}

	render.JSON(w, r, newTaskResponse(task))
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID, err := getTaskID(r)
	if err != nil {
		h.log.Error(err)
		renderError(w, r, "failed to get task ID", http.StatusInternalServerError)
		return
	}
	h.log.Infow("delete task", "id", taskID)
	if err := h.service.DeleteTask(ctx, taskID); err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			renderError(w, r, "task not found", http.StatusNotFound)
		} else {
			renderError(w, r, "failed to delete task", http.StatusInternalServerError)
		}
		h.log.Errorw("failed to delete task", "error", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getTaskID(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
}

type TaskResponse struct {
	ID          int64      `json:"id"`
	Status      string     `json:"status"`
	Result      string     `json:"result,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	DurationSec float64    `json:"duration_sec,omitempty"`
}

func newTaskResponse(t *domain.Task) *TaskResponse {

	resp := &TaskResponse{
		ID:          t.ID,
		Status:      string(t.Status),
		CreatedAt:   t.CreatedAt,
		StartedAt:   t.StartedAt,
		CompletedAt: t.CompletedAt,
		Result:      t.Result,
	}
	if t.StartedAt != nil {
		endTime := time.Now()
		if t.CompletedAt != nil {
			endTime = *t.CompletedAt
		}
		resp.DurationSec = endTime.Sub(*t.StartedAt).Seconds()
	}

	return resp
}

func renderError(w http.ResponseWriter, r *http.Request, msg string, status int) {
	type errorResponse struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}

	render.Status(r, status)
	render.JSON(w, r, &errorResponse{
		Message:    msg,
		StatusCode: status,
	})
}
