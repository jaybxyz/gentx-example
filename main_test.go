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
	var resp = `{"height":"2274483","txhash":"CF5C0328C1243E50A4AAA536288FA2BD2DC9FD08837FFAF8E0AB3070B09D12A1","raw_log":"[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]","logs":[{"msg_index":"0","success":true,"log":""}],"gas_wanted":"200000","gas_used":"103821","tags":[{"key":"action","value":"create_validator"},{"key":"destination-validator","value":"cosmosvaloper1khnuwdpnyv6utvqx3plt3x6makh3rc9tv0wvv9"},{"key":"moniker","value":"cosmos-sta"},{"key":"identity"}],"tx":{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"cosmos-sta","identity":"","website":"","details":""},"commission":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"min_self_delegation":"1","delegator_address":"cosmos1khnuwdpnyv6utvqx3plt3x6makh3rc9tfm6eqk","validator_address":"cosmosvaloper1khnuwdpnyv6utvqx3plt3x6makh3rc9tv0wvv9","pubkey":"cosmosvalconspub1zcjduepqgygxh0pntlyhqhrdvzwchezzs0wut0ug7aa4dfym8xcyqce7qw8spflzc3","value":{"denom":"uatom","amount":"1000000"}}}],"fee":{"amount":[{"denom":"uatom","amount":"5000"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A5sGJWDI/DfZS5MNoaxGEgWEb0CARrWGPPj4KR/fNhTv"},"signature":"CPbUTKQ3LqWjx4Smmjy0IzTGYnuP5rgfBXfWdfezeVE79ZGcJSmDEGhotcrXN1Rgvo5mP+8y1nV+bRpJXMuMCA=="}],"memo":""}},"timestamp":"2019-10-22T03:30:17Z"}`

	var tx sdk.TxResponse

	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterInterface((*auth.Account)(nil), nil)
	cdc.RegisterInterface((*sdk.Tx)(nil), nil)
	cdc.RegisterConcrete(&auth.StdTx{}, "auth/StdTx", nil)
	cdc.RegisterConcrete(&staking.MsgCreateValidator{}, "cosmos-sdk/MsgCreateValidator", nil)

	cdc.MustUnmarshalJSON([]byte(resp), &tx)

	// fmt.Println(tx)
	// fmt.Println("")

	stdTx := tx.Tx.(*auth.StdTx)

	// stdSigs contains the sequence number, account number, and signatures.
	stdSigs := stdTx.GetSignatures()
	for i := 0; i < len(stdSigs); i++ {
		// TODO: Get signer account from somewhere you trust
		// ----------- e.g. ------------
		signer, _ := sdk.AccAddressFromBech32("cosmos1khnuwdpnyv6utvqx3plt3x6makh4ee1tfm6eqe")
		acc := auth.BaseAccount{
			Address:       signer,
			AccountNumber: 21634, // from lcd
			Sequence:      1,     //from lcd
			PubKey:        stdSigs[i].PubKey,
		}
		// ----------- e.g. ------------

		signBytes := auth.GetSignBytes("cosmoshub-2", *stdTx, &acc, false)

		fmt.Println(string(signBytes))

		pubKey := stdSigs[i].PubKey
		fmt.Printf("Verified: %t", pubKey.VerifyBytes(signBytes, stdSigs[i].Signature))
	}
}

func TestVerify2(t *testing.T) {
	// -------- failed gentx ---------
	// var resp = `{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"Cosmostation","identity":"AE4C403A6E7AA1AC","website":"https://www.cosmostation.io","details":"CØSMOSTATION Validator. Delegate your atoms and Start Earning Staking Rewards"},"commission":{"rate":"0.120000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.100000000000000000"},"min_self_delegation":"10","delegator_address":"cosmos1clpqr4nrk4khgkxj78fcwwh6dl3uw4ep4tgu9q","validator_address":"cosmosvaloper1clpqr4nrk4khgkxj78fcwwh6dl3uw4epsluffn","pubkey":"cosmosvalconspub1zcjduepq0dc9apn3pz2x2qyujcnl2heqq4aceput2uaucuvhrjts75q0rv5smjjn7v","value":{"denom":"uatom","amount":"30000000000"}}}],"fee":{"amount":null,"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"An0yrOygz23oiJQZg63gJSbg4nkrWmHBC02/6Am2oDrS"},"signature":"7tGEPPbDpcw16rYHBZlulA552vn/J6ZBc4XMFzdyYzk5kD5BNWlm+V+5Bh/rzLcrTOCealX1/69g4hRKxM66vA=="}],"memo":""}}`

	// -------- verified gentx ---------
	var resp = `{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"Cosmostation","identity":"AE4C403A6E7AA1AC","website":"https://www.cosmostation.io","details":"CØSMOSTATION Validator. Delegate your atoms and Start Earning Staking Rewards"},"commission":{"rate":"0.120000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.100000000000000000"},"min_self_delegation":"10","delegator_address":"cosmos1clpqr4nrk4khgkxj78fcwwh6dl3uw4ep4tgu9q","validator_address":"cosmosvaloper1clpqr4nrk4khgkxj78fcwwh6dl3uw4epsluffn","pubkey":"cosmosvalconspub1zcjduepq0dc9apn3pz2x2qyujcnl2heqq4aceput2uaucuvhrjts75q0rv5smjjn7v","value":{"denom":"uatom","amount":"30000000000"}}}],"fee":{"amount":null,"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"An0yrOygz23oiJQZg63gJSbg4nkrWmHBC02/6Am2oDrS"},"signature":"6l0B/lTVJxiKoNmo3F4qH5bRPmA9uh2dL43b9w5beUdODGYEXNSe2DG+rSfjg1JBrBvhFKYTt7Rl/5PpdlgKgg=="}],"memo":""}}`

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

	stdTx := tx.Tx.(*auth.StdTx)

	// stdSigs contains the sequence number, account number, and signatures.
	stdSigs := stdTx.GetSignatures()
	for i := 0; i < len(stdSigs); i++ {
		// TODO: Get signer account from somewhere you trust
		// ----------- e.g. ------------
		signer, _ := sdk.AccAddressFromBech32("cosmos1khnuwdpnyv6utvqx3plt3x6makh4ee1tfm6eqe")
		acc := auth.BaseAccount{
			Address:       signer,
			AccountNumber: 21634, // from lcd
			Sequence:      1,     //from lcd
			PubKey:        stdSigs[i].PubKey,
		}
		// ----------- e.g. ------------

		signBytes := auth.GetSignBytes("cosmoshub-2", *stdTx, &acc, false)

		fmt.Println(string(signBytes))

		pubKey := stdSigs[i].PubKey
		fmt.Printf("Verified: %t", pubKey.VerifyBytes(signBytes, stdSigs[i].Signature))
	}
}
