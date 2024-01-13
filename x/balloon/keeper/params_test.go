package keeper_test

import (
	"testing"

	testkeeper "github.com/ComputerKeeda/balloon/testutil/keeper"
	"github.com/ComputerKeeda/balloon/x/balloon/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BalloonKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
