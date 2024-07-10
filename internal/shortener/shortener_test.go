package shortener_test

import (
	"testing"

	"github.com/amaterasutears/url-shortener/internal/shortener"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	t.Run("returns the code", func(t *testing.T) {
		type testCase struct {
			original string
			expected string
		}

		testCases := []testCase{
			{
				original: "https://google.com",
				expected: "05046f26",
			},
			{
				original: "https://vk.com/profile/135134",
				expected: "48039f47",
			},
			{
				original: "https://github.com/altkraft/for-applicants/blob/master/backend/shortener/task.md",
				expected: "5a63e03d",
			},
		}

		t.Run("consistings of alpha/numeric/alphanumeric symbols", func(t *testing.T) {
			for _, tc := range testCases {
				code := shortener.Code(tc.original)
				assert.Equal(t, tc.expected, code)
			}
		})

		t.Run("len(code) = 8", func(t *testing.T) {
			for _, tc := range testCases {
				code := shortener.Code(tc.original)
				assert.Equal(t, 8, len(code))
			}
		})
	})

	t.Run("idempotent", func(t *testing.T) {
		testCase := struct {
			original string
			expected string
		}{
			original: "https://google.com",
			expected: "05046f26",
		}

		for i := 0; i < 100; i++ {
			code := shortener.Code(testCase.original)
			assert.Equal(t, testCase.expected, code)
		}
	})
}
