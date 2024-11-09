package test

import (
	// "context"
	"testing"

	// "github.com/Mohamadreza-shad/url-shortener/handler/api"
	// "github.com/Mohamadreza-shad/url-shortener/repository"
	// "github.com/Mohamadreza-shad/url-shortener/service/urls"
	// "github.com/go-playground/validator/v10"
	// "github.com/stretchr/testify/assert"
)

func Test_ShortenUrl_UrlIsAlreadyExist_InRedis(t *testing.T) {
	// ctx := context.Background()
	// assert := assert.New(t)
	// logger := getLogger()
	// validator := validator.New()
	// redisClient := getRedis()
	// err := redisClient.FlushAll(ctx).Err()
	// assert.Nil(err)
	// db := getDB()
	// err = truncateDB()
	// assert.Nil(err)
	// repo := repository.New()
	// urlService := urls.NewUrlService(db,repo,logger)
	// urlHandler := api.NewUrlHandler(urlService,validator)


}
func Test_ShortenUrl_UrlIsAlreadyExist_NotInRedis_ButInDb(t *testing.T) {}
func Test_ShortenUrl_Successfully(t *testing.T)                         {}

func Test_GetLongUrl_DataIsInRedis(t *testing.T)                            {}
func Test_GetLongUrl_DataIsNotInRedis_SearchingDb_UrlNotFound(t *testing.T) {}
func Test_GetLongUrl_DataIsNotInRedis_SearchingDb_UrlExpired(t *testing.T)  {}
func Test_GetLongUrl_DataIsNotInRedis_SearchingDb_Successful(t *testing.T)  {}
