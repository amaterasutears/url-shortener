package urlnormalizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	type testCase struct {
		name        string
		src         string
		expected    string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "returns an address without www. and / at the end of the path",
			src:         "https://www.google.com/",
			expected:    "https://google.com",
			expectedErr: nil,
		},
		{
			name:        "returns an address without www. but leaves a slash at the end of the request parameter path",
			src:         "https://www.google.com/somaparam?=somedata/",
			expected:    "https://google.com/somaparam?=somedata/",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Normalize(tc.src)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
