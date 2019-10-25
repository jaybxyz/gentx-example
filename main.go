package main

import (
	"encoding/json"
	"fmt"

	"github.com/kogisin/gentx-example/client"
	"github.com/kogisin/gentx-example/config"
	"github.com/kogisin/gentx-example/types"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/x/auth"

	resty "gopkg.in/resty.v1"
)

const (
	cfgFile = "./config.toml"
)

var (
	chainID = "cosmoshub-2"
	txHash  = "CF5C0328C1243E50A4AAA536288FA2BD2DC9FD08837FFAF8E0AB3070B09D12A1"
)

func main() {
	// Configuration in config.toml
	cfg := config.ParseConfig(cfgFile)

	// Connect to Tendermint RPC client
	cp, err := client.New(cfg.RPCNode, cfg.LCDEndpoint)
	if err != nil {
		fmt.Println("failed to start RPC client: ", err)
	}

	// Fetch tx data
	resp, err := resty.R().Get(cp.LCDEndpoint + "/txs/" + txHash)
	if err != nil {
		fmt.Printf("failed to fetch data %s: ", err)
	}

	var tx types.MsgCreateValidatorTx
	err = json.Unmarshal(resp.Body(), &tx)
	if err != nil {
		fmt.Printf("failed to unmarshal tx %s", err)
	}

	fmt.Println(tx.Tx.Value.Msg)
	fmt.Println(tx.Tx.Value.Signatures)

	// check signature, return account with incremented nonce
	// signBytes := auth.GetSignBytes(chainID, stdTx, signerAccs[i], isGenesis)

	// Verify
	// result := pubKeyTest.VerifyBytes(newStdSignBytes, signStdSignMsgBytes)

}
