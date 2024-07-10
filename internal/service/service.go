package service

import (
	"github.com/amaterasutears/url-shortener/internal/handler"
	"github.com/amaterasutears/url-shortener/internal/shortener"
)

type LinksRepository interface {
	Put(code, original string) error
	FindOne(code string) (string, error)
}

type Service struct {
	mongoRepository LinksRepository
	redisRepository LinksRepository
}

var _ handler.ShortenerService = (*Service)(nil)

func New(mongoRepository LinksRepository, redisRepository LinksRepository) *Service {
	return &Service{
		mongoRepository: mongoRepository,
		redisRepository: redisRepository,
	}
}

func (s *Service) Shorten(original string) (string, error) {
	code := shortener.Code(original)

	_, err := s.mongoRepository.FindOne(code)
	if err != nil {
		s.mongoRepository.Put(code, original)
		s.redisRepository.Put(code, original)
		return code, nil
	}

	return code, nil
}

func (s *Service) Redirect(code string) (string, error) {
	original, err := s.redisRepository.FindOne(code)
	if err != nil {
		original, err = s.mongoRepository.FindOne(code)
		if err != nil {
			return "", err
		}
	}

	return original, nil
}
