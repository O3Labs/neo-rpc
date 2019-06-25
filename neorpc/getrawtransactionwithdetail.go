package neorpc

import (
	"math"
	"strconv"
)

type NEP5Tx struct {
	Timestamp           int    `json:"timestamp"`
	AssetHash           string `json:"asset_hash"`
	AssetName           string `json:"asset_name"`
	AssetSymbol         string `json:"asset_symbol"`
	AssetDecimals       int    `json:"asset_decimals"`
	TransferFromAddress string `json:"transfer_from_address"`
	TransferToAddress   string `json:"transfer_to_address"`
	Amount              string `json:"amount"`
	BlockIndex          int    `json:"block_index"`
	TransferNotifyIndex int    `json:"transfer_notify_index"`
	TxHash              string `json:"tx_hash"`
}
type GetRawTransactionWithDetailResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result struct {
		Txid       string `json:"txid"`
		Size       int    `json:"size"`
		Type       string `json:"type"`
		Version    int    `json:"version"`
		Attributes []struct {
			Usage string `json:"usage"`
			Data  string `json:"data"`
		} `json:"attributes"`
		Vin []struct {
			Txid    string `json:"txid"`
			Vout    int    `json:"vout"`
			Asset   string `json:"asset"`
			Value   string `json:"value"`
			Address string `json:"address"`
		} `json:"vin"`
		Vout []struct {
			N       int    `json:"n"`
			Asset   string `json:"asset"`
			Value   string `json:"value"`
			Address string `json:"address"`
		} `json:"vout"`
		Claims []struct {
			Txid    string `json:"txid"`
			Vout    int    `json:"vout"`
			Asset   string `json:"asset"`
			Value   string `json:"value"`
			Address string `json:"address"`
		} `json:"claims"`
		SysFee  string `json:"sys_fee"`
		NetFee  string `json:"net_fee"`
		Gas     string `json:"gas"`
		Script  string `json:"script"`
		Scripts []struct {
			Invocation   string `json:"invocation"`
			Verification string `json:"verification"`
		} `json:"scripts"`
		Blockhash     string   `json:"blockhash"`
		Confirmations int      `json:"confirmations"`
		Blocktime     int      `json:"blocktime"`
		BlockIndex    int      `json:"blockindex"`
		Nep5          []NEP5Tx `json:"nep5"`
	} `json:"result"`
}

func (n NEP5Tx) FormattedAmount() string {
	amountUnit64, err := strconv.ParseUint(n.Amount, 10, 64)
	if err != nil {
		return "0.0"
	}
	return strconv.FormatFloat(float64(amountUnit64)/math.Pow10(n.AssetDecimals), 'f', -1, 64)
}

func (g GetRawTransactionWithDetailResponse) HasInOutOrNEP5() bool {
	if len(g.Result.Vin) > 0 || len(g.Result.Vout) > 0 || len(g.Result.Nep5) > 0 {
		return true
	}
	return false
}
