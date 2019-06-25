package neorpc

type GetContractStateResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result struct {
		Version     int      `json:"version"`
		Hash        string   `json:"hash"`
		Script      string   `json:"script"`
		Parameters  []string `json:"parameters"`
		Returntype  string   `json:"returntype"`
		Name        string   `json:"name"`
		CodeVersion string   `json:"code_version"`
		Author      string   `json:"author"`
		Email       string   `json:"email"`
		Description string   `json:"description"`
		Properties  struct {
			Storage       bool `json:"storage"`
			DynamicInvoke bool `json:"dynamic_invoke"`
			// DynamicInvoke bool `json:"payable"` // this property is not even one the code yet
			//https://github.com/neo-project/neo/blob/master/neo/Ledger/ContractState.cs
		} `json:"properties"`
	} `json:"result"`
}
