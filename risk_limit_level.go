package kumex

import (
	"net/http"
)

// RiskLimitLevelModel represents a Contract risk limit level info.
type RiskLimitLevelModel struct {
	Symbol         string  `json:"symbol"`
	Level          int64   `json:"level"`
	MaxRiskLimit   int64   `json:"maxRiskLimit"`
	MinRiskLimit   int64   `json:"minRiskLimit"`
	MaxLeverage    int64   `json:"maxLeverage"`
	InitialMargin  float64 `json:"initialMargin"`
	MaintainMargin float64 `json:"maintainMargin"`
}

// ContractRiskLimitLevelModel represents a Contract risk limit level info.
type ContractsRiskLimitLevelModel []*RiskLimitLevelModel

// ContractsRiskLimitLevel obtain information about risk limit level of a specific contract
func (as *ApiService) ContractsRiskLimitLevel(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/contracts/risk-limit/"+symbol, nil)
	return as.Call(req)
}

// ContractsRiskLimit adjust contract risk limit level
func (as *ApiService) ChangeRiskLimitLevel(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/position/risk-limit-level/change", params)
	return as.Call(req)
}
