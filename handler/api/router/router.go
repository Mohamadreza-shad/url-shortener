package router

import (
	"errors"
	"net/http"

	"github.com/Mohamadreza-shad/url-shortener/handler/api"
	"github.com/Mohamadreza-shad/url-shortener/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	Handler *gin.Engine
}

func New(
	urlHandler *api.UrlHandler,
	logger *logger.Logger,
) *Router {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(globalRecover(logger))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			api.ResponseFailure{
				Success: false,
				Error: api.ErrorCode{
					Code:    http.StatusNotFound,
					Message: "URL not found",
				},
			})
	})

	r.POST("/api/url/shorten", urlHandler.ShortenUrl)
	r.GET("/api/url/long", urlHandler.LongUrl)

	return &Router{
		Handler: r,
	}
}

func globalRecover(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {
			if rec := recover(); rec != nil {
				err := errors.New("error 500")
				if err != nil {
					logger.Error("error 500", zap.Error(err))
				}
				api.MakeErrorResponseWithCode(c.Writer, http.StatusInternalServerError, "error 500")
			}
		}(c)
		c.Next()
	}
}
