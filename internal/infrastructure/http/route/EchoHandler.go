package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EchoHandler struct {
	log *log.Logger
}

func NewEchoHandler(log *log.Logger) *EchoHandler {
	return &EchoHandler{log}
}

func (*EchoHandler) Method() string {
	return http.MethodGet
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func (eh *EchoHandler) Handler(ctx *gin.Context) {
	ctx.JSON(200, "echo")
}
