package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	dextypes "github.com/ChihuahuaChain/chihuahua/x/dex/types"
)

func TestValidatePoolDenom(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		denom    string
		expected uint64
		valid    bool
	}{
		{
			desc:     "valid denom",
			denom:    "chihuahua/pool/0",
			expected: 0,
			valid:    true,
		},
		{
			desc:     "valid denom long",
			denom:    "chihuahua/pool/999999999",
			expected: 999999999,
			valid:    true,
		},
		{
			desc:  "invalid denom long",
			denom: "chihuahua/pool/99999999999999999999",
			valid: false,
		},
		{
			desc:  "invalid format 1",
			denom: "/chihuahua/pool/0",
			valid: false,
		},
		{
			desc:  "invalid format 2",
			denom: "chihuahua/pool/0/",
			valid: false,
		},
		{
			desc:  "invalid denom prefix",
			denom: "INVALID/pool/0",
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			validateError := dextypes.ValidatePoolDenom(tc.denom)
			id, err := dextypes.ParsePoolIDFromDenom(tc.denom)
			if tc.valid {
				require.NoError(t, err)
				require.NoError(t, validateError)
				require.Equal(t, tc.expected, id)
			} else {
				require.Error(t, err)
				require.Equal(t, err, validateError)
			}
		})
	}
}
