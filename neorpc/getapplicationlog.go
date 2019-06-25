package neorpc

type GetApplicationLogResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result struct {
		Txid       string `json:"txid"`
		Executions []struct {
			Trigger     string `json:"trigger"`
			Contract    string `json:"contract"`
			Vmstate     string `json:"vmstate"`
			GasConsumed string `json:"gas_consumed"`
			Stack       []struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"stack"`
			Notifications []struct {
				Contract string `json:"contract"`
				State    struct {
					Type  string `json:"type"`
					Value []struct {
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"value"`
				} `json:"state"`
			} `json:"notifications"`
		} `json:"executions"`
	} `json:"result"`
}
