package urlnormalizer

import (
	"net/url"
	"strings"
)

func Normalize(src string) (string, error) {
	parsed, err := url.Parse(src)
	if err != nil {
		return "", err
	}

	ok := strings.HasPrefix(parsed.Host, "www.")
	if ok {
		parsed.Host = strings.TrimPrefix(parsed.Host, "www.")
	}

	ok = strings.HasSuffix(parsed.Path, "/")
	if ok {
		parsed.Path = strings.TrimSuffix(parsed.Path, "/")
	}

	return parsed.String(), nil
}
