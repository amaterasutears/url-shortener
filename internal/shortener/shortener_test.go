package shortener_test

import (
	"testing"

	"github.com/amaterasutears/url-shortener/internal/shortener"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	t.Run("returns code", func(t *testing.T) {
		type testCase struct {
			src      string
			expected string
		}

		testCases := []testCase{
			{
				src:      "https://google.com",
				expected: "05046f26",
			},
			{
				src:      "https://google.com/somesymbols",
				expected: "dd4a7086",
			},
			{
				src:      "https://google.com/somesymbols?someparam=somesymbols",
				expected: "174fa2bb",
			},
		}
		t.Run("alpha, numeric or alphanumeric", func(t *testing.T) {
			for _, tc := range testCases {
				actual := shortener.Code(tc.src)
				assert.Equal(t, tc.expected, actual)
			}
		})

		t.Run("length is 8", func(t *testing.T) {
			for _, tc := range testCases {
				actual := shortener.Code(tc.src)
				assert.Equal(t, 8, len(actual))
			}
		})
	})

	t.Run("idempotent", func(t *testing.T) {
		testCase := struct {
			src      string
			expected string
		}{
			src:      "https://google.com",
			expected: "05046f26",
		}

		for i := 0; i < 100; i++ {
			actual := shortener.Code(testCase.src)
			assert.Equal(t, testCase.expected, actual)
		}
	})
}
