package urlnormalizer_test

import (
	"testing"

	"github.com/amaterasutears/url-shortener/pkg/urlnormalizer"
	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	type testCase struct {
		src      string
		expected string
	}

	testCases := []testCase{
		{
			src:      "https://www.google.com/",
			expected: "https://google.com",
		},
		{
			src:      "https://google.com/",
			expected: "https://google.com",
		},
		{
			src:      "https://www.google.com/somedata/",
			expected: "https://google.com/somedata",
		},
		{
			src:      "https://google.com/somedata/",
			expected: "https://google.com/somedata",
		},
		{
			src:      "https://www.google.com/somedata?somedata=somedata/",
			expected: "https://google.com/somedata?somedata=somedata",
		},
	}

	for _, tc := range testCases {
		actual, err := urlnormalizer.Normalize(tc.src)
		assert.Nil(t, err)
		assert.Equal(t, tc.expected, actual)
	}
}
