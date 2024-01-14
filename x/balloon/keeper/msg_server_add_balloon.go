package keeper

import (
	"context"
	"fmt"

	"github.com/ComputerKeeda/balloon/x/balloon/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AddBalloon(goCtx context.Context, msg *types.MsgAddBalloon) (*types.MsgAddBalloonResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	balloonName := msg.BalloonName
	balloonHeight := msg.BalloonHeight
	balloonNameByte := []byte(balloonName)
	balloonHeightByte := []byte(balloonHeight)
	balloonStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BalloonStore))

	getBalloon := balloonStore.Get(balloonNameByte)
	if getBalloon != nil {
		return &types.MsgAddBalloonResponse{
			Message: fmt.Sprintf("balloon with name `%s` already exist", msg.BalloonName),
		}, status.Error(codes.InvalidArgument, fmt.Sprintf("balloon with name `%s` already exist", msg.BalloonName))
	}

	balloonStore.Set(balloonNameByte, balloonHeightByte)

	return &types.MsgAddBalloonResponse{
		Message: "Balloon added successfully",
	}, nil
}
