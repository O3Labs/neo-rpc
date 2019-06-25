package neorpc

type GetAllTransactionsResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result []struct {
		Txid       string `json:"txid"`
		Size       int    `json:"size"`
		O3Type     int    `json:"o3type"`
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
		SysFee  string `json:"sys_fee"`
		NetFee  string `json:"net_fee"`
		Scripts []struct {
			Invocation   string `json:"invocation"`
			Verification string `json:"verification"`
		} `json:"scripts"`
		Script     string `json:"script,omitempty"`
		Gas        string `json:"gas,omitempty"`
		Blockindex int    `json:"blockindex"`
		Blockhash  string `json:"blockhash"`
		Blocktime  int    `json:"blocktime"`
		BlockIndex int    `json:"blockindex"`
		Sent       []struct {
			Timestamp           int    `json:"timestamp"`
			AssetHash           string `json:"asset_hash"`
			AssetName           string `json:"asset_name"`
			AssetSymbol         string `json:"asset_symbol"`
			AssetDecimals       int    `json:"asset_decimals"`
			TransferToAddress   string `json:"transfer_to_address"`
			TransferFromAddress string `json:"transfer_from_address"`
			Amount              string `json:"amount"`
			BlockIndex          int    `json:"block_index"`
			TransferNotifyIndex int    `json:"transfer_notify_index"`
			TxHash              string `json:"tx_hash"`
		} `json:"sent"`
		Received []struct {
			Timestamp           int    `json:"timestamp"`
			AssetHash           string `json:"asset_hash"`
			AssetName           string `json:"asset_name"`
			AssetSymbol         string `json:"asset_symbol"`
			AssetDecimals       int    `json:"asset_decimals"`
			TransferToAddress   string `json:"transfer_to_address"`
			TransferFromAddress string `json:"transfer_from_address"`
			Amount              string `json:"amount"`
			BlockIndex          int    `json:"block_index"`
			TransferNotifyIndex int    `json:"transfer_notify_index"`
			TxHash              string `json:"tx_hash"`
		} `json:"received"`
		Claims []struct {
			Txid string `json:"txid"`
			Vout int    `json:"vout"`
		} `json:"claims,omitempty"`
	} `json:"result"`
}
