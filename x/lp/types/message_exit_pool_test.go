package types

/*
func TestMsgWithdrawLPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgWithdrawLPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgWithdrawLPool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgWithdrawLPool{
				Creator: sample.AccAddress(),
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
*/
