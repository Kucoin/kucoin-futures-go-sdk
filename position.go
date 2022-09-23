package kumex

import (
	"net/http"
)

// A PositionModel represents a position info.
type PositionModel struct {
	Id                string  `json:"id"`
	Symbol            string  `json:"symbol"`
	AutoDeposit       bool    `json:"autoDeposit"`
	MaintMarginReq    float64 `json:"maintMarginReq"`
	RiskLimit         int     `json:"riskLimit"`
	RealLeverage      float64 `json:"realLeverage"`
	CrossMode         bool    `json:"crossMode"`
	DelevPercentage   float64 `json:"delevPercentage"`
	OpeningTimestamp  int64   `json:"openingTimestamp"`
	CurrentTimestamp  int64   `json:"currentTimestamp"`
	CurrentQty        int     `json:"currentQty"`
	CurrentCost       float64 `json:"currentCost"`
	CurrentComm       float64 `json:"currentComm"`
	UnrealisedCost    float64 `json:"unrealisedCost"`
	RealisedGrossCost float64 `json:"realisedGrossCost"`
	RealisedCost      float64 `json:"realisedCost"`
	IsOpen            bool    `json:"isOpen"`
	MarkPrice         float64 `json:"markPrice"`
	MarkValue         float64 `json:"markValue"`
	PosCost           float64 `json:"posCost"`
	PosCross          float64 `json:"posCross"`
	PosInit           float64 `json:"posInit"`
	PosComm           float64 `json:"posComm"`
	PosLoss           float64 `json:"posLoss"`
	PosMargin         float64 `json:"posMargin"`
	PosMaint          float64 `json:"posMaint"`
	MaintMargin       float64 `json:"maintMargin"`
	RealisedGrossPnl  float64 `json:"realisedGrossPnl"`
	RealisedPnl       float64 `json:"realisedPnl"`
	UnrealisedPnl     float64 `json:"unrealisedPnl"`
	UnrealisedPnlPcnt float64 `json:"unrealisedPnlPcnt"`
	UnrealisedRoePcnt float64 `json:"unrealisedRoePcnt"`
	AvgEntryPrice     float64 `json:"avgEntryPrice"`
	LiquidationPrice  float64 `json:"liquidationPrice"`
	BankruptPrice     float64 `json:"bankruptPrice"`
	SettleCurrency    string  `json:"settleCurrency"`
	IsInverse         bool    `json:"isInverse"`
	MaintainMargin    float64 `json:"maintainMargin"`
}

// Position Get Position Details.
func (as *ApiService) Position(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodGet, "/api/v1/position", p)
	return as.Call(req)
}

// Positions Get Position List.
func (as *ApiService) Positions() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/positions", nil)
	return as.Call(req)
}

// AutoDepositStatus Enable/Disable of Auto-Deposit Margin.
func (as *ApiService) AutoDepositStatus(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/position/margin/auto-deposit-status", params)
	return as.Call(req)
}

// DepositMargin Add Margin Manually.
func (as *ApiService) DepositMargin(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/position/margin/deposit-margin", params)
	return as.Call(req)
}
