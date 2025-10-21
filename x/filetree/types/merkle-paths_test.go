package types_test

import (
	"testing"

	"github.com/jackalLabs/canine-chain/v5/x/filetree/types"
	"github.com/stretchr/testify/require"
)

func TestMerklePath(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected string
	}{
		{
			input:    "hello/world/path",
			expected: "3867baa2724c672442e4ba21b6fa532a6380d06a2f8779f11d626bd840d1cdee",
		},
		{
			input:    "hello/world/path/",
			expected: "3867baa2724c672442e4ba21b6fa532a6380d06a2f8779f11d626bd840d1cdee",
		},
		{
			input:    "jackal/maxi",
			expected: "442db08812d5fc05653ba5621cc000f40fab3a326532a3584df515b046bd86c0",
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			result := types.MerklePath(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestAddToMerkle(t *testing.T) {
	for _, tc := range []struct {
		input1   string
		input2   string
		expected string
	}{
		{
			input1:   "path",
			input2:   "append",
			expected: "82568d2e91733544857522d73f31cc10387000a369f9f93cf94724fec4594fed",
		},
		{
			input1:   "hello",
			input2:   "world",
			expected: "936a185caaa266bb9cbe981e9e05cb78cd732b0b3280eb944412bb6f8f8f07af",
		},
	} {
		t.Run("", func(t *testing.T) {
			result := types.AddToMerkle(tc.input1, tc.input2)
			require.Equal(t, tc.expected, result)
		})
	}
}
