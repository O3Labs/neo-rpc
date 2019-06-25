package neorpc

type GetClaimableResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result struct {
		Claimable []struct {
			Txid        string  `json:"txid"`
			N           int     `json:"n"`
			Value       int     `json:"value"`
			StartHeight int     `json:"start_height"`
			EndHeight   int     `json:"end_height"`
			Generated   float64 `json:"generated"`
			SysFee      float64 `json:"sys_fee"`
			Unclaimed   float64 `json:"unclaimed"`
		} `json:"claimable"`
		Address   string  `json:"address"`
		Unclaimed float64 `json:"unclaimed"`
	} `json:"result"`
}
