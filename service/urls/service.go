package urls

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"time"

	"github.com/Mohamadreza-shad/url-shortener/client"
	"github.com/Mohamadreza-shad/url-shortener/config"
	"github.com/Mohamadreza-shad/url-shortener/logger"
	"github.com/Mohamadreza-shad/url-shortener/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/speps/go-hashids"
)

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrUrlIsAlreadyExist  = errors.New("url is already exist")
	ErrUrlNotFound        = errors.New("url not found")
	ErrExpiredUrl         = errors.New("url is expired")

	ExpireTime = time.Now().Add(time.Hour * 24 * 7)
)

type UrlService struct {
	db     client.PgxInterface
	repo   *repository.Queries
	logger *logger.Logger
}

type ShortenUrlParams struct {
	LongUrl string `json:"longUrl"`
}

type LongUrl struct {
	Url string `json:"url"`
}
type LongUrlParams struct {
	ShortUrl string `json:"shortUrl" validate:"required"`
}

func (s *UrlService) ShortenUrl(ctx context.Context, params ShortenUrlParams) error {
	_, err := s.repo.UrlByLongUrl(ctx, s.db, params.LongUrl)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return ErrSomethingWentWrong
	}
	if err == nil {
		return ErrUrlIsAlreadyExist
	}
	shortUrl, err := generateShortUrl(params.LongUrl)
	if err != nil {
		return ErrSomethingWentWrong
	}
	_, err = s.repo.ShortenUrl(ctx, s.db, repository.ShortenUrlParams{
		LongUrl:   params.LongUrl,
		ShortUrl:  shortUrl,
		ExpiredAt: pgtype.Timestamp{Time: ExpireTime, Valid: true},
	})
	if err != nil {
		return ErrSomethingWentWrong
	}
	return nil
}

func (s *UrlService) LongUrl(ctx context.Context, params LongUrlParams) (LongUrl, error) {
	url, err := s.repo.UrlByShortUrl(ctx, s.db, params.ShortUrl)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return LongUrl{}, ErrSomethingWentWrong
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return LongUrl{}, ErrUrlNotFound
	}
	if url.ExpiredAt.Time.Before(time.Now()) {
		return LongUrl{}, ErrExpiredUrl
	}
	return LongUrl{Url: url.LongUrl}, nil
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
