package main

import (
	"encoding/json"
	"fmt"

	"github.com/kogisin/gentx-example/client"
	"github.com/kogisin/gentx-example/config"
	"github.com/kogisin/gentx-example/types"

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

	var txResp types.MsgCreateValidatorTx
	err = json.Unmarshal(resp.Body(), &txResp)
	if err != nil {
		fmt.Printf("failed to unmarshal tx %s", err)
	}

	fmt.Println(txResp)

	// StdTx is a standard way to wrap a Msg with Fee and Signatures.
	// NOTE: the first signature is the fee payer (Signatures must not be nil).
	// type StdTx struct {
	// 	Msgs       []sdk.Msg      `json:"msg"`
	// 	Fee        StdFee         `json:"fee"`
	// 	Signatures []StdSignature `json:"signatures"`
	// 	Memo       string         `json:"memo"`
	// }

	// msgs := []sdk.Msg{}
	// fee := auth.NewStdFee(50000,
	// 	sdk.NewCoins(sdk.NewInt64Coin(tx.Tx.Value.Fee.Amount[0].Denom, tx.Tx.Value.Fee.Amount[0].Amount)),
	// )
	// sigs := auth.StdSignature{
	// 	PubKey:     []byte(tx.Tx.Value.Signatures[0].PubKey.Value),
	// 	Signatures: []byte(tx.Tx.Value.Signatures[0].Signature),
	// }

	// stdTx := NewStdTx(msgs, fee, sigs, "")

	// check signature, return account with incremented nonce
	// signBytes := auth.GetSignBytes(chainID, stdTx, signerAccs[i], isGenesis)

	// Verify
	// result := pubKeyTest.VerifyBytes(newStdSignBytes, signStdSignMsgBytes)

}
