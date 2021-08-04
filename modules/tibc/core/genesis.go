package ibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	client "github.com/bianjieai/tibc-go/modules/tibc/core/02-client"
	packet "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet"
	"github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
	"github.com/bianjieai/tibc-go/modules/tibc/core/types"
)

// InitGenesis initializes the ibc state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, createLocalhost bool, gs *types.GenesisState) {
	client.InitGenesis(ctx, k.ClientKeeper, gs.ClientGenesis)
	packet.InitGenesis(ctx, k.Packetkeeper, gs.PacketGenesis)
}

// ExportGenesis returns the ibc exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		ClientGenesis: client.ExportGenesis(ctx, k.ClientKeeper),
		PacketGenesis: packet.ExportGenesis(ctx, k.Packetkeeper),
	}
}