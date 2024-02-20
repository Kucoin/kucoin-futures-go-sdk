package kumex

import (
	"net/http"
)

// A FundingModel represents a funding record.
type FundingModel struct {
	Id             int64   `json:"id"`
	Symbol         string  `json:"symbol"`
	TimePoint      int64   `json:"timePoint"`
	FundingRate    float64 `json:"fundingRate"`
	MarkPrice      float64 `json:"markPrice"`
	PositionQty    float32 `json:"positionQty"`
	PositionCost   float64 `json:"positionCost"`
	Funding        float64 `json:"funding"`
	SettleCurrency string  `json:"settleCurrency"`
}

// A FundingListModel is the set of *FundingModel.
type FundingListModel struct {
	HasMore  bool            `json:"hasMore"`
	DataList []*FundingModel `json:"dataList"` // delay parsing
}

// FundingHistory Get Funding History.
func (as *ApiService) FundingHistory(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/funding-history", params)
	return as.Call(req)
}

type FundingTimeRangeRatesModel []*FundingTimeRangeRateModel

type FundingTimeRangeRateModel struct {
	Symbol      string  `json:"symbol"`
	TimePoint   float64 `json:"timePoint"`
	FundingRate float64 `json:"fundingRate"`
}

// FundingRatesTimeRange Get Funding rates info .
func (as *ApiService) FundingRatesTimeRange(symbol, from, to string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/contract/funding-rates", map[string]string{
		"symbol": symbol,
		"from":   from,
		"to":     to,
	})
	return as.Call(req)
}
