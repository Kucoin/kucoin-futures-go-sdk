package kumex

import "net/http"

// A FillModel represents the structure of fill.
type FillModel struct {
	Symbol         string  `json:"symbol"`
	TradeId        string  `json:"tradeId"`
	OrderId        string  `json:"orderId"`
	Side           string  `json:"side"`
	Liquidity      string  `json:"liquidity"`
	Price          string  `json:"price"`
	Size           float64 `json:"size"`
	Value          string  `json:"value"`
	FeeRate        string  `json:"feeRate"`
	FixFee         string  `json:"fixFee"`
	FeeCurrency    string  `json:"feeCurrency"`
	Stop           string  `json:"stop"`
	Fee            string  `json:"fee"`
	OrderType      string  `json:"orderType"`
	TradeType      string  `json:"tradeType"`
	CreatedAt      int64   `json:"createdAt"`
	SettleCurrency string  `json:"settleCurrency"`
	TradeTime      int64   `json:"tradeTime"`
	OpenFeePay     string  `json:"openFeePay"`
	CloseFeePay    string  `json:"closeFeePay"`
	MarginMode     string  `json:"marginMode"`
}

// A FillsModel is the set of *FillModel.
type FillsModel []*FillModel

// Fills returns a list of recent fills.
func (as *ApiService) Fills(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/fills", params)
	return as.Call(req)
}

// RecentFills returns the recent fills of the latest transactions within 24 hours.
func (as *ApiService) RecentFills() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/recentFills", nil)
	return as.Call(req)
}

// A OpenOrderStatisticsModel represents the struct of fill.
type OpenOrderStatisticsModel struct {
	OpenOrderBuySize  int32  `json:"openOrderBuySize"`
	OpenOrderSellSize int32  `json:"openOrderSellSize"`
	OpenOrderBuyCost  string `json:"openOrderBuyCost"`
	OpenOrderSellCost string `json:"openOrderSellCost"`
	SettleCurrency    string `json:"settleCurrency"`
}

// OpenOrderStatistics Active Order Value Calculation.
func (as *ApiService) OpenOrderStatistics(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodDelete, "/api/v1/openOrderStatistics", p)
	return as.Call(req)
}
