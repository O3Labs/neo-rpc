package neorpc

type GetBlockCountResponse struct {
	JSONRPCResponse
	*ErrorResponse     //optional
	Result         int `json:"result"`
}
