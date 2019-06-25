package neorpc

import (
	"strconv"
)

type GetBlockResponse struct {
	JSONRPCResponse
	*ErrorResponse                //optional
	Result         GetBlockResult `json:"result"`
}

type Tx struct {
	Txid       string `json:"txid"`
	Size       int    `json:"size"`
	Type       string `json:"type"`
	Version    int    `json:"version"`
	Nonce      int    `json:"nonce"`
	Attributes []struct {
		Usage string `json:"usage"`
		Data  string `json:"data"`
	} `json:"attributes"`
	Vin []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	} `json:"vin"`
	Vout []struct {
		N       int    `json:"n"`
		Asset   string `json:"asset"`
		Value   string `json:"value"`
		Address string `json:"address"`
	} `json:"vout"`
	Claims []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	} `json:"claims"`
	SysFee  string `json:"sys_fee"`
	NetFee  string `json:"net_fee"`
	Scripts []struct {
		Invocation   string `json:"invocation"`
		Verification string `json:"verification"`
	} `json:"scripts"`
	Script string `json:"script"`
	Gas    string `json:"gas"`
}

type GetBlockResult struct {
	Hash              string `json:"hash"`
	Size              int    `json:"size"`
	Version           int    `json:"version"`
	Previousblockhash string `json:"previousblockhash"`
	Merkleroot        string `json:"merkleroot"`
	Time              int    `json:"time"`
	Index             int    `json:"index"`
	Nonce             string `json:"nonce"`
	Nextconsensus     string `json:"nextconsensus"`
	Script            struct {
		Invocation   string `json:"invocation"`
		Verification string `json:"verification"`
	} `json:"script"`
	Tx            []Tx   `json:"tx"`
	Confirmations int    `json:"confirmations"`
	Nextblockhash string `json:"nextblockhash"`
}

type TxFee struct {
	NetFee string
	SysFee string
}

func (g GetBlockResult) TotalFees() TxFee {
	total := TxFee{}
	totalNetFee := float64(0)
	totalSysFee := float64(0)
	for _, v := range g.Tx {
		netfeeFloat, err := strconv.ParseFloat(v.NetFee, 64)
		if err == nil {
			totalNetFee = totalNetFee + netfeeFloat
		}
		sysfeeFloat, err := strconv.ParseFloat(v.SysFee, 64)
		if err == nil {
			totalSysFee = totalSysFee + sysfeeFloat
		}
	}

	total.NetFee = strconv.FormatFloat(RoundFixed(totalNetFee, 8), 'f', -1, 64)
	total.SysFee = strconv.FormatFloat(RoundFixed(totalSysFee, 8), 'f', -1, 64)
	return total
}
