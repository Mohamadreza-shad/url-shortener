package urls

import (
	"context"
	"crypto/md5"
	"encoding/binary"

	"github.com/Mohamadreza-shad/url-shortener/client"
	"github.com/Mohamadreza-shad/url-shortener/config"
	"github.com/Mohamadreza-shad/url-shortener/logger"
	"github.com/Mohamadreza-shad/url-shortener/repository"
	"github.com/speps/go-hashids"
)

type UrlService struct {
	db     client.PgxInterface
	repo   *repository.Queries
	logger *logger.Logger
}

type ShortenUrlParams struct {
	LongUrl string
}

func (s *UrlService) ShortenUrl(ctx context.Context, params ShortenUrlParams) error {
	return nil
}

func generateShortUrl(longUrl string) (string, error) {
	hash := md5.Sum([]byte(longUrl))
	num := int64(binary.BigEndian.Uint64(hash[:8]))

	hd := hashids.NewData()
	hd.Salt = config.SaltKey()
	hd.MinLength = 12
	h := hashids.NewWithData(hd)
	e, err := h.EncodeInt64([]int64{num})
	if err != nil {
		return "", err
	}
	return e, nil
}

func NewUrlService(
	db client.PgxInterface,
	repo *repository.Queries,
	logger *logger.Logger,
) *UrlService {
	return &UrlService{
		db:     db,
		repo:   repo,
		logger: logger,
	}
}
