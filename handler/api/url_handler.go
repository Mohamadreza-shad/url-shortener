package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Mohamadreza-shad/url-shortener/helper"
	"github.com/Mohamadreza-shad/url-shortener/service/urls"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UrlHandler struct {
	service   *urls.UrlService
	validator *validator.Validate
}

func (h *UrlHandler) ShortenUrl(c *gin.Context) {
	var params urls.ShortenUrlParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		MakeErrorResponseWithCode(c.Writer, http.StatusBadRequest, "invalid request")
		return
	}
	if strings.EqualFold(params.LongUrl, "") {
		MakeErrorResponseWithCode(c.Writer, http.StatusBadRequest, "long url is required")
		return
	}
	if !helper.IsValidURL(params.LongUrl) {
		MakeErrorResponseWithCode(c.Writer, http.StatusBadRequest, "invalid url")
		return
	}
	err = h.service.ShortenUrl(c, params)
	if errors.Is(err, urls.ErrUrlIsAlreadyExist) {
		MakeErrorResponseWithCode(c.Writer, http.StatusConflict, urls.ErrUrlIsAlreadyExist.Error())
		return
	}
	if err != nil {
		MakeErrorResponseWithoutCode(c.Writer, urls.ErrSomethingWentWrong)
		return
	}
	MakeSuccessResponse(c.Writer, nil, "url shortened successfully")
}

func (h *UrlHandler) LongUrl(c *gin.Context) {
	var params urls.LongUrlParams
	err := c.ShouldBindQuery(&params)
	if err != nil {
		MakeErrorResponseWithCode(c.Writer, http.StatusBadRequest, "invalid request")
		return
	}
	if strings.EqualFold(params.ShortUrl, "") {
		MakeErrorResponseWithCode(c.Writer, http.StatusBadRequest, "short-url is required")
		return
	}
	if !helper.IsValidURL(params.ShortUrl) {
		MakeErrorResponseWithCode(c.Writer, http.StatusBadRequest, "invalid url")
		return
	}
	longUrl, err := h.service.LongUrl(c, params)
	if errors.Is(err, urls.ErrUrlNotFound) {
		MakeErrorResponseWithCode(c.Writer, http.StatusNotFound, urls.ErrUrlNotFound.Error())
		return
	}
	if errors.Is(err, urls.ErrExpiredUrl) {
		MakeErrorResponseWithCode(c.Writer, http.StatusGone, urls.ErrExpiredUrl.Error())
		return
	}
	if err != nil {
		MakeErrorResponseWithoutCode(c.Writer, urls.ErrSomethingWentWrong)
		return
	}
	MakeSuccessResponse(c.Writer, longUrl.Url, "long url retrieved successfully")
}

func NewUrlHandler(
	service *urls.UrlService,
	validator *validator.Validate,
) *UrlHandler {
	return &UrlHandler{
		service:   service,
		validator: validator,
	}
}
