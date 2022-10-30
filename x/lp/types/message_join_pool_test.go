package types

/*
func TestMsgDepositLPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDepositLPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDepositLPool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDepositLPool{
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
