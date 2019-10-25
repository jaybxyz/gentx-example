package types

import (
	"time"
)

type MsgCreateValidatorTx struct {
	Height string `json:"height"`
	Txhash string `json:"txhash"`
	RawLog string `json:"raw_log"`
	Logs   []struct {
		MsgIndex string `json:"msg_index"`
		Success  bool   `json:"success"`
		Log      string `json:"log"`
	} `json:"logs"`
	GasWanted string `json:"gas_wanted"`
	GasUsed   string `json:"gas_used"`
	Tags      []struct {
		Key   string `json:"key"`
		Value string `json:"value,omitempty"`
	} `json:"tags"`
	Tx struct {
		Type  string `json:"type"`
		Value struct {
			Msg []struct {
				Type  string `json:"type"`
				Value struct {
					Description struct {
						Moniker  string `json:"moniker"`
						Identity string `json:"identity"`
						Website  string `json:"website"`
						Details  string `json:"details"`
					} `json:"description"`
					Commission struct {
						Rate          string `json:"rate"`
						MaxRate       string `json:"max_rate"`
						MaxChangeRate string `json:"max_change_rate"`
					} `json:"commission"`
					MinSelfDelegation string `json:"min_self_delegation"`
					DelegatorAddress  string `json:"delegator_address"`
					ValidatorAddress  string `json:"validator_address"`
					Pubkey            string `json:"pubkey"`
					Value             struct {
						Denom  string `json:"denom"`
						Amount string `json:"amount"`
					} `json:"value"`
				} `json:"value"`
			} `json:"msg"`
			Fee struct {
				Amount []struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"amount"`
				Gas string `json:"gas"`
			} `json:"fee"`
			Signatures []struct {
				PubKey struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"pub_key"`
				Signature string `json:"signature"`
			} `json:"signatures"`
			Memo string `json:"memo"`
		} `json:"value"`
	} `json:"tx"`
	Timestamp time.Time `json:"timestamp"`
}
