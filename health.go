package health

import (
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IHealth interface {
	Healthy() bool
	SetSystemState(name string, up bool)
}

type Health struct {
	systems sync.Map
	logger  *zap.Logger
}

func New(logger *zap.Logger) *Health {
	logger = logger.With(zap.String("module", "health"))

	h := &Health{
		systems: sync.Map{},
		logger:  logger,
	}

	return h
}

func (h *Health) RegisterEndpoint(r gin.IRouter) {
	r.GET("/health", h.Handle)
}

func (h *Health) Healthy() bool {
	healthy := true
	h.systems.Range(func(key, value interface{}) bool {
		if !(value.(bool)) {
			healthy = false
			return false
		}
		return true
	})
	return healthy
}

func (h *Health) SetSystemState(name string, value bool) {
	h.systems.Store(name, value)
}

func (h *Health) Handle(ctx *gin.Context) {
	if !h.Healthy() {
		ctx.String(500, "NOT OK")
		return
	}

	ctx.String(200, "OK")
}
