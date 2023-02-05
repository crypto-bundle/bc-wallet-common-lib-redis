package healthcheck

import (
	"net/http"
)

type httpHandler struct {
	probeSrv             probeService
	directiveExecutorSrv appDirectiveExecutionService
}

func newHttpHandler(probeSrv probeService) *httpHandler {
	return &httpHandler{
		probeSrv: probeSrv,
	}
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthStatus := h.probeSrv.Do(r.Context())
	if healthStatus.Error != nil || !healthStatus.IsHealed {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if healthStatus.IsHealed {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Add("Content-Type", "text/plain")
	_, writeErr := w.Write([]byte(healthStatus.Message))
	if writeErr != nil {
		return
	}

	return
}
