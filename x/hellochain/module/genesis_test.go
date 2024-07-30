package hellochain_test

import (
	"testing"

	keepertest "hellochain/testutil/keeper"
	"hellochain/testutil/nullify"
	hellochain "hellochain/x/hellochain/module"
	"hellochain/x/hellochain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HellochainKeeper(t)
	hellochain.InitGenesis(ctx, k, genesisState)
	got := hellochain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
