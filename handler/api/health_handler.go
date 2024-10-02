package api

import (
	"github.com/gin-gonic/gin"
)

//  HealthHandler service
//  Health example
//	@Summary		health check
//	@Description	health check
//	@ID				health
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{string} string	"ok"
//	@Router			/health [get]
func Health(c *gin.Context) {
	c.Writer.Write([]byte("OK"))
}
