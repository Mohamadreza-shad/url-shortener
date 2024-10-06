package test

import "testing"

func Test_ShortenUrl_UrlIsAlreadyExist_InRedis(t *testing.T)            {}
func Test_ShortenUrl_UrlIsAlreadyExist_NotInRedis_ButInDb(t *testing.T) {}
func Test_ShortenUrl_Successfully(t *testing.T)                         {}

func Test_GetLongUrl_DataIsInRedis(t *testing.T)                            {}
func Test_GetLongUrl_DataIsNotInRedis_SearchingDb_UrlNotFound(t *testing.T) {}
func Test_GetLongUrl_DataIsNotInRedis_SearchingDb_UrlExpired(t *testing.T)  {}
func Test_GetLongUrl_DataIsNotInRedis_SearchingDb_Successful(t *testing.T)  {}
