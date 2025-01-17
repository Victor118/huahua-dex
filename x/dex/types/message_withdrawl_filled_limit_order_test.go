package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ChihuahuaChain/chihuahua/testutil/common/sample"
	dextypes "github.com/ChihuahuaChain/chihuahua/x/dex/types"
)

func TestMsgWithdrawFilledLimitOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  dextypes.MsgWithdrawFilledLimitOrder
		err  error
	}{
		{
			name: "invalid creator",
			msg: dextypes.MsgWithdrawFilledLimitOrder{
				Creator:    "invalid_address",
				TrancheKey: "ORDER123",
			},
			err: dextypes.ErrInvalidAddress,
		}, {
			name: "valid msg",
			msg: dextypes.MsgWithdrawFilledLimitOrder{
				Creator:    sample.AccAddress(),
				TrancheKey: "ORDER123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
