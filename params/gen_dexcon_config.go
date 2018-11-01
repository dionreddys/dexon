// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package params

import (
	"encoding/json"
	"math/big"

	"github.com/dexon-foundation/dexon/common"
	"github.com/dexon-foundation/dexon/common/math"
)

var _ = (*dexconConfigSpecMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (d DexconConfig) MarshalJSON() ([]byte, error) {
	type DexconConfig struct {
		GenesisCRSText   string                `json:"genesisCRSText"`
		Owner            common.Address        `json:"owner"`
		BlockReward      *math.HexOrDecimal256 `json:"blockReward"`
		BlockGasLimit    uint64                `json:"blockGasLimit"`
		NumChains        uint32                `json:"numChains"`
		LambdaBA         uint64                `json:"lambdaBA"`
		LambdaDKG        uint64                `json:"lambdaDKG"`
		K                int                   `json:"k"`
		PhiRatio         float32               `json:"phiRatio"`
		NotarySetSize    uint32                `json:"notarySetSize"`
		DKGSetSize       uint32                `json:"dkgSetSize"`
		RoundInterval    uint64                `json:"roundInterval"`
		MinBlockInterval uint64                `json:"minBlockInterval"`
		MaxBlockInterval uint64                `json:"maxBlockInterval"`
	}
	var enc DexconConfig
	enc.GenesisCRSText = d.GenesisCRSText
	enc.Owner = d.Owner
	enc.BlockReward = (*math.HexOrDecimal256)(d.BlockReward)
	enc.BlockGasLimit = d.BlockGasLimit
	enc.NumChains = d.NumChains
	enc.LambdaBA = d.LambdaBA
	enc.LambdaDKG = d.LambdaDKG
	enc.K = d.K
	enc.PhiRatio = d.PhiRatio
	enc.NotarySetSize = d.NotarySetSize
	enc.DKGSetSize = d.DKGSetSize
	enc.RoundInterval = d.RoundInterval
	enc.MinBlockInterval = d.MinBlockInterval
	enc.MaxBlockInterval = d.MaxBlockInterval
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (d *DexconConfig) UnmarshalJSON(input []byte) error {
	type DexconConfig struct {
		GenesisCRSText   *string               `json:"genesisCRSText"`
		Owner            *common.Address       `json:"owner"`
		BlockReward      *math.HexOrDecimal256 `json:"blockReward"`
		BlockGasLimit    *uint64               `json:"blockGasLimit"`
		NumChains        *uint32               `json:"numChains"`
		LambdaBA         *uint64               `json:"lambdaBA"`
		LambdaDKG        *uint64               `json:"lambdaDKG"`
		K                *int                  `json:"k"`
		PhiRatio         *float32              `json:"phiRatio"`
		NotarySetSize    *uint32               `json:"notarySetSize"`
		DKGSetSize       *uint32               `json:"dkgSetSize"`
		RoundInterval    *uint64               `json:"roundInterval"`
		MinBlockInterval *uint64               `json:"minBlockInterval"`
		MaxBlockInterval *uint64               `json:"maxBlockInterval"`
	}
	var dec DexconConfig
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.GenesisCRSText != nil {
		d.GenesisCRSText = *dec.GenesisCRSText
	}
	if dec.Owner != nil {
		d.Owner = *dec.Owner
	}
	if dec.BlockReward != nil {
		d.BlockReward = (*big.Int)(dec.BlockReward)
	}
	if dec.BlockGasLimit != nil {
		d.BlockGasLimit = *dec.BlockGasLimit
	}
	if dec.NumChains != nil {
		d.NumChains = *dec.NumChains
	}
	if dec.LambdaBA != nil {
		d.LambdaBA = *dec.LambdaBA
	}
	if dec.LambdaDKG != nil {
		d.LambdaDKG = *dec.LambdaDKG
	}
	if dec.K != nil {
		d.K = *dec.K
	}
	if dec.PhiRatio != nil {
		d.PhiRatio = *dec.PhiRatio
	}
	if dec.NotarySetSize != nil {
		d.NotarySetSize = *dec.NotarySetSize
	}
	if dec.DKGSetSize != nil {
		d.DKGSetSize = *dec.DKGSetSize
	}
	if dec.RoundInterval != nil {
		d.RoundInterval = *dec.RoundInterval
	}
	if dec.MinBlockInterval != nil {
		d.MinBlockInterval = *dec.MinBlockInterval
	}
	if dec.MaxBlockInterval != nil {
		d.MaxBlockInterval = *dec.MaxBlockInterval
	}
	return nil
}