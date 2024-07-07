package service

import (
	"log"
	"net/url"
	"strings"

	"github.com/amaterasutears/url-shortener/internal/shortener"
)

type LinksRepository interface {
	Put(code, original string) error
	FindOne(code string) (string, error)
}

type Service struct {
	linksRepository      LinksRepository
	cacheLinksRepository LinksRepository
}

func New(linkRepository LinksRepository, cacheLinksRepository LinksRepository) *Service {
	return &Service{
		linksRepository:      linkRepository,
		cacheLinksRepository: cacheLinksRepository,
	}
}

func (s *Service) Shorten(original string) (string, error) {
	original, err := s.normalizeURL(original)
	if err != nil {
		return "", err
	}

	code := shortener.Code(original)

	_, err = s.linksRepository.FindOne(code)
	if err != nil {
		s.linksRepository.Put(code, original)
		s.cacheLinksRepository.Put(code, original)
		return code, nil
	}

	return code, nil
}

func (s *Service) Redirect(code string) (string, error) {
	original, err := s.cacheLinksRepository.FindOne(code)
	log.Println(original)
	if err != nil {
		original, err = s.linksRepository.FindOne(code)
		if err != nil {
			return "", err
		}
	}

	return original, nil
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
