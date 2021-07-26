package http

import (
	"parse-config/internal/router"

	"github.com/gin-gonic/gin"
)

func InitServer(port string) {
	r := gin.Default()
	router.Route(r)
	r.Run(":" + port)
}
