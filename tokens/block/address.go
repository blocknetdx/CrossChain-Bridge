package block

import (
	"encoding/json"
	"strings"

	"github.com/blocknetdx/btcd/chaincfg"
	btcsuitechaincfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// IsValidAddress check address
func (b *Bridge) IsValidAddress(addr string) bool {
	chainConfig := b.GetChainConfig()
	address, err := btcutil.DecodeAddress(addr, chainConfig)
	if err != nil {
		return false
	}
	return address.IsForNet(chainConfig)
}

// IsP2pkhAddress check p2pkh addrss
func (b *Bridge) IsP2pkhAddress(addr string) bool {
	chainConfig := b.GetChainConfig()
	address, err := btcutil.DecodeAddress(addr, chainConfig)
	if err != nil {
		return false
	}
	if !address.IsForNet(chainConfig) {
		return false
	}
	_, ok := address.(*btcutil.AddressPubKeyHash)
	return ok
}

// IsP2shAddress check p2sh addrss
func (b *Bridge) IsP2shAddress(addr string) bool {
	chainConfig := b.GetChainConfig()
	address, err := btcutil.DecodeAddress(addr, chainConfig)
	if err != nil {
		return false
	}
	if !address.IsForNet(chainConfig) {
		return false
	}
	_, ok := address.(*btcutil.AddressScriptHash)
	return ok
}

func convertToBTCSuiteChainCfg(cfg interface{}) *btcsuitechaincfg.Params {
	bcfg := btcsuitechaincfg.Params{}
	bz, err := json.Marshal(cfg)
	if err != nil {
		panic("invalid config")
	}
	err = json.Unmarshal(bz, &bcfg)
	if err != nil {
		panic("error unmarshaling config to btcsuite")
	}
	return &bcfg
}

// GetChainConfig get chain config (net params)
func (b *Bridge) GetChainConfig() *btcsuitechaincfg.Params {
	token := b.TokenConfig
	networkID := strings.ToLower(token.NetID)
	switch networkID {
	case netMainnet:
		return convertToBTCSuiteChainCfg(chaincfg.MainNetParams)
	case netTestnet3:
		return convertToBTCSuiteChainCfg(chaincfg.TestNet3Params)
	}
	return convertToBTCSuiteChainCfg(chaincfg.TestNet3Params)
}
