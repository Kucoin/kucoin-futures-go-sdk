package kumex

import "net/http"

// A ContractsModel is the struct.
type ContractsModel struct {
	BaseCurrency            string  `json:"baseCurrency"`
	QuoteCurrency           string  `json:"quoteCurrency"`
	SettleCurrency          string  `json:"settleCurrency"`
	FairMethod              string  `json:"fairMethod"`
	FundingBaseSymbol       string  `json:"fundingBaseSymbol"`
	FundingQuoteSymbol      string  `json:"fundingQuoteSymbol"`
	FundingRateSymbol       string  `json:"fundingRateSymbol"`
	IndexSymbol             string  `json:"indexSymbol"`
	InitialMargin           float32 `json:"initialMargin"`
	IsDeleverage            bool    `json:"isDeleverage"`
	IsInverse               bool    `json:"isInverse"`
	IsQuanto                bool    `json:"isQuanto"`
	LotSize                 float32 `json:"lotSize"`
	MaintainMargin          float32 `json:"maintainMargin"`
	MakerFeeRate            float32 `json:"makerFeeRate"`
	MakerFixFee             float32 `json:"makerFixFee"`
	MarkMethod              string  `json:"markMethod"`
	MaxOrderQty             float32 `json:"maxOrderQty"`
	MaxPrice                float32 `json:"maxPrice"`
	MaxRiskLimit            float32 `json:"maxRiskLimit"`
	MinRiskLimit            float32 `json:"minRiskLimit"`
	Multiplier              float32 `json:"multiplier"`
	RiskStep                int     `json:"riskStep"`
	RootSymbol              string  `json:"rootSymbol"`
	Status                  string  `json:"status"`
	Symbol                  string  `json:"symbol"`
	TakerFeeRate            float32 `json:"takerFeeRate"`
	TakerFixFee             float32 `json:"takerFixFee"`
	TickSize                float32 `json:"tickSize"`
	Type                    string  `json:"type"`
	MaxLeverage             float32 `json:"maxLeverage"`
	VolumeOf24h             float64 `json:"volumeOf24h"`
	TurnoverOf24h           float64 `json:"turnoverOf24h"`
	FundingFeeRate          float64 `json:"fundingFeeRate"`
	PredictedFundingFeeRate float64 `json:"predictedFundingFeeRate"`
	FundingRateGranularity  int64   `json:"fundingRateGranularity"`
	OpenInterest            string  `json:"openInterest"`
	LastTradePrice          float64 `json:"lastTradePrice"`
	NextFundingRateTime     int64   `json:"nextFundingRateTime"`
	LowPrice                float64 `json:"lowPrice"`
	HighPrice               float64 `json:"highPrice"`
	PriceChgPct             float64 `json:"priceChgPct"`
	PriceChg                float64 `json:"priceChg"`
}

type ContractsModels []*ContractsModel

// ActiveContracts Get Open Contract List.
func (as *ApiService) ActiveContracts() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/contracts/active", nil)
	return as.Call(req)
}

// Contracts Get Order Info. of the Contract.
func (as *ApiService) Contracts(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/contracts/"+symbol, nil)
	return as.Call(req)
}
