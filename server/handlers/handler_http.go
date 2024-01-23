package handlers

import "net/http"

type HostRouteHandler struct {
	http.Handler

	hosts map[string]http.Handler
}

func (h *HostRouteHandler) RegisterHost(Name string, handler http.Handler) {
	if h.hosts == nil {
		h.hosts = make(map[string]http.Handler)
	}
	h.hosts[Name] = handler
}

func (h *HostRouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/healthCheck" {
		w.Write(make([]byte, 0))
		return
	}

	if host, ok := h.hosts[r.Host]; ok {
		host.ServeHTTP(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
