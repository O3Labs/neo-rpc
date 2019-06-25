package neorpc

import (
	"math"
	"strconv"
)

type GetNEP5BalancesResponse struct {
	JSONRPCResponse
	*ErrorResponse
	Result struct {
		Balance []NEP5Balance `json:"balance"`
		Address string        `json:"address"`
	} `json:"result"`
}

type NEP5Balance struct {
	AssetHash        string `json:"asset_hash"`
	AssetName        string `json:"asset_name"`
	AssetSymbol      string `json:"asset_symbol"`
	AssetDecimals    int    `json:"asset_decimals"`
	Amount           string `json:"amount"`
	LastUpdatedBlock int    `json:"last_updated_block"`
}

func (n NEP5Balance) FormattedAmount() string {
	amountUnit64, err := strconv.ParseUint(n.Amount, 10, 64)
	if err != nil {
		return "0.0"
	}
	return strconv.FormatFloat(float64(amountUnit64)/math.Pow10(n.AssetDecimals), 'f', -1, 64)
}
