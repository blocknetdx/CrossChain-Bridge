package tokens

import (
	"errors"
	"math/big"
)

type TokenConfig struct {
	BlockChain      string
	NetID           string
	ID              string `json:",omitempty"`
	Name            string
	Symbol          string
	Decimals        *uint8
	Description     string `json:",omitempty"`
	DcrmAddress     string
	ContractAddress string `json:",omitempty"`
	Confirmations   *uint64
	MaximumSwap     *float64 // whole unit (eg. BTC, ETH, FSN), not Satoshi
	MinimumSwap     *float64 // whole unit
	SwapFeeRate     *float64
}

func (c *TokenConfig) CheckConfig(isSrc bool) error {
	if c.BlockChain == "" {
		return errors.New("token must config 'BlockChain'")
	}
	if c.NetID == "" {
		return errors.New("token must config 'NetID'")
	}
	if c.Decimals == nil {
		return errors.New("token must config 'Decimals'")
	}
	if c.Confirmations == nil {
		return errors.New("token must config 'Confirmations'")
	}
	if c.MaximumSwap == nil {
		return errors.New("token must config 'MaximumSwap'")
	}
	if c.MinimumSwap == nil {
		return errors.New("token must config 'MinimumSwap'")
	}
	if c.SwapFeeRate == nil {
		return errors.New("token must config 'SwapFeeRate'")
	}
	if c.DcrmAddress == "" {
		return errors.New("token must config 'DcrmAddress'")
	}
	if !isSrc && c.ContractAddress == "" {
		return errors.New("token must config 'ContractAddress' for destination chain")
	}
	return nil
}

type GatewayConfig struct {
	ApiAddress string
}

type SwapType uint32

const (
	Swap_Unknown SwapType = iota
	Swap_Swapin
	Swap_Swapout
	Swap_Recall
)

type TxSwapInfo struct {
	Hash      string   `json:"hash"`
	Height    uint64   `json:"height"`
	Timestamp uint64   `json:"timestamp"`
	From      string   `json:"from"`
	To        string   `json:"to"`
	Bind      string   `json:"bind"`
	Value     *big.Int `json:"value"`
}

type TxStatus struct {
	Receipt       interface{} `json:"receipt,omitempty"`
	Confirmations uint64      `json:"confirmations"`
	Block_height  uint64      `json:"block_height"`
	Block_hash    string      `json:"block_hash"`
	Block_time    uint64      `json:"block_time"`
}

type SwapInfo struct {
	SwapID   string   `json:"swapid,omitempty"`
	SwapType SwapType `json:"swaptype,omitempty"`
}

type BuildTxArgs struct {
	*SwapInfo `json:"swapInfo`
	From      string      `json:"from"`
	To        string      `json:"to"`
	Value     *big.Int    `json:"value"`
	Memo      string      `json:"memo"`
	Input     *[]byte     `json:"input"`
	Extra     interface{} `json:"extra"`
}

func (args *BuildTxArgs) GetExtraArgs() *BuildTxArgs {
	return &BuildTxArgs{
		SwapInfo: args.SwapInfo,
		Extra:    args.Extra,
	}
}

type EthExtraArgs struct {
	Gas      *uint64  `json:"gas,omitempty"`
	GasPrice *big.Int `json:"gasPrice,omitempty"`
	Nonce    *uint64  `json:"nonce,omitempty"`
}

type FsnExtraArgs EthExtraArgs

type BtcExtraArgs struct {
	RelayFeePerKb *int64  `json:"relayFeePerKb,omitempty"`
	ChangeAddress *string `json:"changeAddress,omitempty"`
	FromPublicKey *string `json:"fromPublickey,omitempty"`
}
