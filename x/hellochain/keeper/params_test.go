package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "hellochain/testutil/keeper"
	"hellochain/x/hellochain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HellochainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
