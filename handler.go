package godoauth

import "net/http"

type Handler struct {
	*http.ServeMux
}

// NewHandler returns a new instance of Handler built from a config.
func NewHandler(authHandler *TokenAuthHandler) *Handler {
	s := &Handler{
		ServeMux: http.NewServeMux(),
	}

	s.Handle("/auth", authHandler)
	s.HandleFunc("/server-ping", s.ping)
	return s
}

// ping is an health check handler, so we can use ELB/HA proxy health check
func (Handler) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Save the Whales !\"}\n\r"))
}
