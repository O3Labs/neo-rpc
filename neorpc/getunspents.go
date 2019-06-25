package neorpc

type GetUnspentsResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result struct {
		Balance []struct {
			Unspent []struct {
				Txid  string  `json:"txid"`
				N     int     `json:"n"`
				Value float64 `json:"value"`
			} `json:"unspent"`
			AssetHash   string  `json:"asset_hash"`
			Asset       string  `json:"asset"`
			AssetSymbol string  `json:"asset_symbol"`
			Amount      float64 `json:"amount"`
		} `json:"balance"`
		Address string `json:"address"`
	} `json:"result"`
}
