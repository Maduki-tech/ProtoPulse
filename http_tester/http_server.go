package http_tester

import "net/http"

type HTTPServer struct {
	server *http.Server
}

func New(port string) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr: port,
		},
	}
}

func (h *HTTPServer) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World"}`))
	})

	h.server.ListenAndServe()
}
