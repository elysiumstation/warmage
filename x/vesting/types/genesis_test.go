package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/petri-labs/warmage/app"
	"github.com/petri-labs/warmage/x/vesting/types"
	"github.com/stretchr/testify/require"
)

func TestAllocationAddresses_GetStrategicReserveCustodianAddr(t *testing.T) {
	app.Setup(false)

	addr := types.AllocationAddresses{}
	require.Equal(t, sdk.AccAddress{}, addr.GetStrategicReserveCustodianAddr())

	addrStr := "war1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5"
	addr = types.AllocationAddresses{
		TeamVestingAddr:               addrStr,
		StrategicReserveCustodianAddr: addrStr,
	}
	require.Equal(t, addrStr, addr.GetStrategicReserveCustodianAddr().String())
}

func TestAllocationAddresses_GetTeamVestingAddr(t *testing.T) {
	app.Setup(false)

	addr := types.AllocationAddresses{}
	require.Equal(t, sdk.AccAddress{}, addr.GetTeamVestingAddr())

	addrStr := "war1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5"
	addr = types.AllocationAddresses{
		TeamVestingAddr:               addrStr,
		StrategicReserveCustodianAddr: addrStr,
	}
	require.Equal(t, addrStr, addr.GetTeamVestingAddr().String())
}

func TestAirdrop_Empty(t *testing.T) {
	app.Setup(false)

	airdrop := types.Airdrop{}
	require.Equal(t, true, airdrop.Empty())
}

func TestAirdrop_GetTargetAddr(t *testing.T) {
	app.Setup(false)

	airdrop := types.Airdrop{}
	require.Equal(t, sdk.AccAddress{}, airdrop.GetTargetAddr())

	addrStr := "war1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5"
	airdrop = types.Airdrop{
		TargetAddr: addrStr,
	}
	require.Equal(t, addrStr, airdrop.GetTargetAddr().String())
}