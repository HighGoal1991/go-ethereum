// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package catalyst

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*executableDataMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (e ExecutableData) MarshalJSON() ([]byte, error) {
	type ExecutableData struct {
		BlockHash     common.Hash     `json:"blockHash"     gencodec:"required"`
		ParentHash    common.Hash     `json:"parentHash"    gencodec:"required"`
		Coinbase      common.Address  `json:"coinbase"      gencodec:"required"`
		StateRoot     common.Hash     `json:"stateRoot"     gencodec:"required"`
		ReceiptRoot   common.Hash     `json:"receiptRoot"   gencodec:"required"`
		LogsBloom     hexutil.Bytes   `json:"logsBloom"     gencodec:"required"`
		Random        common.Hash     `json:"random"        gencodec:"required"`
		Number        hexutil.Uint64  `json:"blockNumber"   gencodec:"required"`
		GasLimit      hexutil.Uint64  `json:"gasLimit"      gencodec:"required"`
		GasUsed       hexutil.Uint64  `json:"gasUsed"       gencodec:"required"`
		Timestamp     hexutil.Uint64  `json:"timestamp"     gencodec:"required"`
		ExtraData     hexutil.Bytes   `json:"extraData"     gencodec:"required"`
		BaseFeePerGas *hexutil.Big    `json:"baseFeePerGas" gencodec:"required"`
		Transactions  []hexutil.Bytes `json:"transactions"  gencodec:"required"`
	}
	var enc ExecutableData
	enc.BlockHash = e.BlockHash
	enc.ParentHash = e.ParentHash
	enc.Coinbase = e.Coinbase
	enc.StateRoot = e.StateRoot
	enc.ReceiptRoot = e.ReceiptRoot
	enc.LogsBloom = e.LogsBloom
	enc.Random = e.Random
	enc.Number = hexutil.Uint64(e.Number)
	enc.GasLimit = hexutil.Uint64(e.GasLimit)
	enc.GasUsed = hexutil.Uint64(e.GasUsed)
	enc.Timestamp = hexutil.Uint64(e.Timestamp)
	enc.ExtraData = e.ExtraData
	enc.BaseFeePerGas = (*hexutil.Big)(e.BaseFeePerGas)
	if e.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(e.Transactions))
		for k, v := range e.Transactions {
			enc.Transactions[k] = v
		}
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *ExecutableData) UnmarshalJSON(input []byte) error {
	type ExecutableData struct {
		BlockHash     *common.Hash    `json:"blockHash"     gencodec:"required"`
		ParentHash    *common.Hash    `json:"parentHash"    gencodec:"required"`
		Coinbase      *common.Address `json:"coinbase"      gencodec:"required"`
		StateRoot     *common.Hash    `json:"stateRoot"     gencodec:"required"`
		ReceiptRoot   *common.Hash    `json:"receiptRoot"   gencodec:"required"`
		LogsBloom     *hexutil.Bytes  `json:"logsBloom"     gencodec:"required"`
		Random        *common.Hash    `json:"random"        gencodec:"required"`
		Number        *hexutil.Uint64 `json:"blockNumber"   gencodec:"required"`
		GasLimit      *hexutil.Uint64 `json:"gasLimit"      gencodec:"required"`
		GasUsed       *hexutil.Uint64 `json:"gasUsed"       gencodec:"required"`
		Timestamp     *hexutil.Uint64 `json:"timestamp"     gencodec:"required"`
		ExtraData     *hexutil.Bytes  `json:"extraData"     gencodec:"required"`
		BaseFeePerGas *hexutil.Big    `json:"baseFeePerGas" gencodec:"required"`
		Transactions  []hexutil.Bytes `json:"transactions"  gencodec:"required"`
	}
	var dec ExecutableData
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BlockHash == nil {
		return errors.New("missing required field 'blockHash' for ExecutableData")
	}
	e.BlockHash = *dec.BlockHash
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for ExecutableData")
	}
	e.ParentHash = *dec.ParentHash
	if dec.Coinbase == nil {
		return errors.New("missing required field 'coinbase' for ExecutableData")
	}
	e.Coinbase = *dec.Coinbase
	if dec.StateRoot == nil {
		return errors.New("missing required field 'stateRoot' for ExecutableData")
	}
	e.StateRoot = *dec.StateRoot
	if dec.ReceiptRoot == nil {
		return errors.New("missing required field 'receiptRoot' for ExecutableData")
	}
	e.ReceiptRoot = *dec.ReceiptRoot
	if dec.LogsBloom == nil {
		return errors.New("missing required field 'logsBloom' for ExecutableData")
	}
	e.LogsBloom = *dec.LogsBloom
	if dec.Random == nil {
		return errors.New("missing required field 'random' for ExecutableData")
	}
	e.Random = *dec.Random
	if dec.Number == nil {
		return errors.New("missing required field 'blockNumber' for ExecutableData")
	}
	e.Number = uint64(*dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for ExecutableData")
	}
	e.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for ExecutableData")
	}
	e.GasUsed = uint64(*dec.GasUsed)
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for ExecutableData")
	}
	e.Timestamp = uint64(*dec.Timestamp)
	if dec.ExtraData == nil {
		return errors.New("missing required field 'extraData' for ExecutableData")
	}
	e.ExtraData = *dec.ExtraData
	if dec.BaseFeePerGas == nil {
		return errors.New("missing required field 'baseFeePerGas' for ExecutableData")
	}
	e.BaseFeePerGas = (*big.Int)(dec.BaseFeePerGas)
	if dec.Transactions == nil {
		return errors.New("missing required field 'transactions' for ExecutableData")
	}
	e.Transactions = make([][]byte, len(dec.Transactions))
	for k, v := range dec.Transactions {
		e.Transactions[k] = v
	}
	return nil
}
