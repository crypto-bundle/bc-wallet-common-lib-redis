package healthcheck

import (
	"net/http"
	"io"
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

func (h *httpHandler) livenessHandler(w http.ResponseWriter, r *http.Request) {
	err := h.livenessProbe.Do(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		io.WriteString(w, "ok")
	}
}

func (h *httpHandler) getReadinessHandler(w http.ResponseWriter, r *http.Request) {
	err := h.readinessProbe.Do(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		io.WriteString(w, "ok")
	}
}

func (h *httpHandler) getStartupHandler(w http.ResponseWriter, r *http.Request) {
	err := h.startupProbe.Do(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		io.WriteString(w, "ok")
	}
}
