package client

import (
	"github.com/cosmos/cosmos-sdk/codec"

	customcdc "github.com/kogisin/gentx-example/codec"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

// Client implements a wrapper around both a Tendermint RPC client and a
// Cosmos SDK REST client that allows for essential data queries.
type Client struct {
	RPCNode     rpcclient.Client // Tendermint RPC node
	LCDEndpoint string           // Full node
	Codec       *codec.Codec
}

func New(rpcNode, clientNode string) (Client, error) {
	rpcClient := rpcclient.NewHTTP(rpcNode, "/websocket")

	if err := rpcClient.Start(); err != nil {
		return Client{}, err
	}

	return Client{rpcClient, clientNode, customcdc.Codec}, nil
}
