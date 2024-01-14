package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddBalloon = "add_balloon"

var _ sdk.Msg = &MsgAddBalloon{}

func NewMsgAddBalloon(creator string, balloonName string, balloonHeight string) *MsgAddBalloon {
	return &MsgAddBalloon{
		Creator:       creator,
		BalloonName:   balloonName,
		BalloonHeight: balloonHeight,
	}
}

func (msg *MsgAddBalloon) Route() string {
	return RouterKey
}

func (msg *MsgAddBalloon) Type() string {
	return TypeMsgAddBalloon
}

func (msg *MsgAddBalloon) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddBalloon) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddBalloon) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
