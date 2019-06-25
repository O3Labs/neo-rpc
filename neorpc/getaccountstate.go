package neorpc

type GetAccountStateResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result GetAccountStateResult `json:"result"`
}

type GetAccountStateResult struct {
	Version    int                      `json:"version"`
	ScriptHash string                   `json:"script_hash"`
	Frozen     bool                     `json:"frozen"`
	Votes      interface{}              `json:"votes"`
	Balances   []GetAccountStateBalance `json:"balances"`
}

type GetAccountStateBalance struct {
	Asset string `json:"asset"`
	Value string `json:"value"`
}
