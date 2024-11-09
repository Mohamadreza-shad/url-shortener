package api

import (
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.Writer.Write([]byte("OK"))
}
