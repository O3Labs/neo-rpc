package neorpc

type GetBlocksResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result []struct {
		Hash         string `json:"hash"`
		Size         int    `json:"size"`
		Time         int    `json:"time"`
		Version      int    `json:"version"`
		Index        int    `json:"index"`
		Validator    string `json:"validator"`
		Transactions int    `json:"transactions"`
	} `json:"result"`
}
