package route

import (
	"io"
	"log"
	"net/http"
)

type EchoHandler struct {
	log *log.Logger
}

func NewEchoHandler(log *log.Logger) *EchoHandler {
	return &EchoHandler{log}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func (eh *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		eh.log.Print("Failed to handle request:", err)
	}
}
