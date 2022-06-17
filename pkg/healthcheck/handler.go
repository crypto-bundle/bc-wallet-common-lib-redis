package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	livenessProbe  Probe
	readinessProbe Probe
	startupProbe   Probe
}

func NewHttpHandler(livenessProbe, readinessProbe, startupProbe Probe) *httpHandler {
	return &httpHandler{
		livenessProbe:  livenessProbe,
		readinessProbe: readinessProbe,
		startupProbe:   startupProbe,
	}
}

func (h *httpHandler) Liveness(ctx *gin.Context) {
	err := h.livenessProbe(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.String(http.StatusOK, "ok")
}

func (h *httpHandler) Readiness(ctx *gin.Context) {
	err := h.readinessProbe(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.String(http.StatusOK, "ok")
}

func (h *httpHandler) Startup(ctx *gin.Context) {
	err := h.startupProbe(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.String(http.StatusOK, "ok")
}
