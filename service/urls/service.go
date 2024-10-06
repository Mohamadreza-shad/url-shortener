package urls

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/Mohamadreza-shad/url-shortener/client"
	"github.com/Mohamadreza-shad/url-shortener/config"
	"github.com/Mohamadreza-shad/url-shortener/logger"
	"github.com/Mohamadreza-shad/url-shortener/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"github.com/speps/go-hashids"
)

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrUrlIsAlreadyExist  = errors.New("url is already exist")
	ErrUrlNotFound        = errors.New("url not found")
	ErrExpiredUrl         = errors.New("url is expired")

	ExpireDuration = time.Hour * 24 * 7
	ExpireTime     = time.Now().Add(ExpireDuration)
)

type UrlService struct {
	db          client.PgxInterface
	repo        *repository.Queries
	redisClient redis.UniversalClient
	logger      *logger.Logger
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
	shortUrl, err := generateShortUrl(params.LongUrl)
	if err != nil {
		return ErrSomethingWentWrong
	}
	err = s.redisClient.Get(ctx, shortUrl).Err()
	if err == nil {
		return ErrUrlIsAlreadyExist
	}
	_, err = s.repo.UrlByLongUrl(ctx, s.db, params.LongUrl)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		s.logger.Error(fmt.Sprintf("error while checking long url: %v", err))
		return ErrSomethingWentWrong
	}
	if err == nil {
		return ErrUrlIsAlreadyExist
	}
	_, err = s.repo.ShortenUrl(ctx, s.db, repository.ShortenUrlParams{
		LongUrl:   params.LongUrl,
		ShortUrl:  shortUrl,
		ExpiredAt: pgtype.Timestamp{Time: ExpireTime, Valid: true},
	})
	if err != nil {
		s.logger.Error(fmt.Sprintf("error while inserting shortened url to db: %v", err))
		return ErrSomethingWentWrong
	}
	err = s.redisClient.Set(ctx, shortUrl, params.LongUrl, ExpireDuration).Err()
	if err != nil {
		s.logger.Error(fmt.Sprintf("error while setting url to redis: %v", err))
	}
	return nil
}

func (s *UrlService) LongUrl(ctx context.Context, params LongUrlParams) (LongUrl, error) {
	longUrl, err := s.redisClient.Get(ctx, params.ShortUrl).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		s.logger.Error(fmt.Sprintf("error while getting url from redis: %v", err))
		return LongUrl{}, ErrSomethingWentWrong
	}
	if errors.Is(err, redis.Nil) {
		wholeUrl, err := s.repo.UrlByShortUrl(ctx, s.db, params.ShortUrl)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return LongUrl{}, ErrSomethingWentWrong
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return LongUrl{}, ErrUrlNotFound
		}
		if wholeUrl.ExpiredAt.Time.Before(time.Now()) {
			return LongUrl{}, ErrExpiredUrl
		}
		longUrl = wholeUrl.LongUrl
	}
	return LongUrl{Url: longUrl}, nil
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
