package balloon_test

import (
	"testing"

	keepertest "github.com/ComputerKeeda/balloon/testutil/keeper"
	"github.com/ComputerKeeda/balloon/testutil/nullify"
	"github.com/ComputerKeeda/balloon/x/balloon"
	"github.com/ComputerKeeda/balloon/x/balloon/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BalloonKeeper(t)
	balloon.InitGenesis(ctx, *k, genesisState)
	got := balloon.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
