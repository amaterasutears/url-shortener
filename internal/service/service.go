package service

import (
	"net/url"
	"strings"

	"github.com/amaterasutears/url-shortener/internal/entity"
	"github.com/amaterasutears/url-shortener/internal/shortener"
)

type LinkRepository interface {
	Put(code, original string) error
	FindOne(code string) (*entity.Link, error)
}

type Service struct {
	linkRepository  LinkRepository
	cacheRepository LinkRepository
}

func New(linkRepository LinkRepository, cacheRepository LinkRepository) *Service {
	return &Service{
		linkRepository:  linkRepository,
		cacheRepository: cacheRepository,
	}
}

func (s *Service) Shorten(original string) (string, error) {
	original, err := s.normalizeURL(original)
	if err != nil {
		return "", err
	}

	code := shortener.Code(original)

	_, err = s.linkRepository.FindOne(code)
	if err != nil {
		s.linkRepository.Put(code, original)
		s.cacheRepository.Put(code, original)
		return code, nil
	}

	return code, nil
}

func (s *Service) Redirect(code string) (string, error) {
	link, err := s.cacheRepository.FindOne(code)
	if err != nil {
		link, err = s.linkRepository.FindOne(code)
		if err != nil {
			return "", err
		}
	}

	return link.Original, nil
}

func (s *Service) normalizeURL(original string) (string, error) {
	parsedURL, err := url.Parse(original)
	if err != nil {
		return "", err
	}

	ok := strings.HasPrefix(parsedURL.Host, "www.")
	if ok {
		parsedURL.Host = strings.TrimPrefix(parsedURL.Host, "www.")
	}

	ok = strings.HasSuffix(parsedURL.Path, "/")
	if ok {
		parsedURL.Path = strings.TrimSuffix(parsedURL.Path, "/")
	}

	return parsedURL.String(), nil
}
