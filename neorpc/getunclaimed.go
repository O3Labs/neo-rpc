package neorpc

type GetUnclaimedResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result struct {
		Available   float64 `json:"available"`
		Unavailable float64 `json:"unavailable"`
		Unclaimed   float64 `json:"unclaimed"`
	} `json:"result"`
}
