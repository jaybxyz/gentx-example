package main

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func TestVerify(t *testing.T) {
	// failed gentx
	var resp = `{"type":"cosmos-sdk/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"Cosmostation","identity":"","website":"","details":""},"commission":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"min_self_delegation":"1","delegator_address":"kava1qrlge6kqjz2763yp6ghws9ekv8u62dva9hs86p","validator_address":"kavavaloper1qrlge6kqjz2763yp6ghws9ekv8u62dvagp20zk","pubkey":"kavavalconspub1zcjduepqnlcxz09tphlp2skvmkyxt6rq0jhfhnwy6thq3v2j4t8cj4pvnn3sp76zg8","value":{"denom":"ukava","amount":"50000000"}}}],"fee":{"amount":[],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"AxhDEq4xW6Yfsl/VWN+ASRJV5mf2MGFyhj8sewIhE9Hj"},"signature":"I02IWE6EdHw377OD6YefRlY3/FTjF4ePQGHKZmYy6n5ToD9M8ZTso/ceq3y4oTYii66EvUC1Pgm74gMGNomhxw=="}],"memo":"cd2705b287f8a28b4bdb6f37dc23511cd76ce2fe@182.77.0.5:26656"}}`

	var tx sdk.TxResponse

	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterInterface((*auth.Account)(nil), nil)
	cdc.RegisterInterface((*sdk.Tx)(nil), nil)
	cdc.RegisterConcrete(&auth.StdTx{}, "auth/StdTx", nil)
	cdc.RegisterConcrete(&staking.MsgCreateValidator{}, "cosmos-sdk/MsgCreateValidator", nil)

	cdc.MustUnmarshalJSON([]byte(resp), &tx)

	fmt.Println(tx)
	fmt.Println("")

	// stdTx := tx.Tx.(*auth.StdTx)

	// // stdSigs contains the sequence number, account number, and signatures.
	// stdSigs := stdTx.GetSignatures()
	// for i := 0; i < len(stdSigs); i++ {
	// 	// TODO: Get signer account from somewhere you trust
	// 	// ----------- e.g. ------------
	// 	signer, _ := sdk.AccAddressFromBech32("cosmos1khnuwdpnyv6utvqx3plt3x6makh4ee1tfm6eqe")
	// 	acc := auth.BaseAccount{
	// 		Address:       signer,
	// 		AccountNumber: 21634, // from lcd
	// 		Sequence:      1,     //from lcd
	// 		PubKey:        stdSigs[i].PubKey,
	// 	}
	// 	// ----------- e.g. ------------

	// 	signBytes := auth.GetSignBytes("cosmoshub-2", *stdTx, &acc, false)

	// 	fmt.Println(string(signBytes))

	// 	pubKey := stdSigs[i].PubKey
	// 	fmt.Printf("Verified: %t", pubKey.VerifyBytes(signBytes, stdSigs[i].Signature))
	// }
}
