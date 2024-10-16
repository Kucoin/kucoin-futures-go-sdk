package kumex

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type PositionsModel []*PositionModel

// A PositionModel represents a position info.
type PositionModel struct {
	UserId            string `json:"userId"`
	Id                string `json:"id"`
	Symbol            string `json:"symbol"`
	AutoDeposit       bool   `json:"autoDeposit"`
	MaintMarginReq    string `json:"maintMarginReq"`
	RiskLimit         string `json:"riskLimit"`
	RealLeverage      string `json:"realLeverage"`
	CrossMode         bool   `json:"crossMode"`
	DelevPercentage   string `json:"delevPercentage"`
	OpeningTimestamp  string `json:"openingTimestamp"`
	CurrentTimestamp  string `json:"currentTimestamp"`
	CurrentQty        string `json:"currentQty"`
	CurrentCost       string `json:"currentCost"`
	CurrentComm       string `json:"currentComm"`
	UnrealisedCost    string `json:"unrealisedCost"`
	RealisedGrossCost string `json:"realisedGrossCost"`
	RealisedCost      string `json:"realisedCost"`
	IsOpen            bool   `json:"isOpen"`
	MarkPrice         string `json:"markPrice"`
	MarkValue         string `json:"markValue"`
	PosCost           string `json:"posCost"`
	PosCross          string `json:"posCross"`
	PosInit           string `json:"posInit"`
	PosComm           string `json:"posComm"`
	PosLoss           string `json:"posLoss"`
	PosMargin         string `json:"posMargin"`
	PosMaint          string `json:"posMaint"`
	MaintMargin       string `json:"maintMargin"`
	RealisedGrossPnl  string `json:"realisedGrossPnl"`
	RealisedPnl       string `json:"realisedPnl"`
	UnrealisedPnl     string `json:"unrealisedPnl"`
	UnrealisedPnlPcnt string `json:"unrealisedPnlPcnt"`
	UnrealisedRoePcnt string `json:"unrealisedRoePcnt"`
	AvgEntryPrice     string `json:"avgEntryPrice"`
	LiquidationPrice  string `json:"liquidationPrice"`
	BankruptPrice     string `json:"bankruptPrice"`
	SettleCurrency    string `json:"settleCurrency"`
	MaintainMargin    string `json:"maintainMargin"`
	RiskLimitLevel    string `json:"riskLimitLevel"`
	MarginMode        string `json:"marginMode"`
	PositionSide      string `json:"positionSide"`
	Leverage          string `json:"leverage"`
	PosFunding        string `json:"posFunding"`
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
func (as *ApiService) Positions(currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/positions", p)
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

type MarginModeModel struct {
	Symbol     string `json:"symbol"`     //symbol of the contract
	MarginMode string `json:"marginMode"` //Margin mode: ISOLATED (isolated), CROSS (cross margin).
}

// GetMarginMode can query the margin mode of the current symbol.
func (as *ApiService) GetMarginMode(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v2/position/getMarginMode", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// ChangeMarginMode modify the margin mode of the current symbol
func (as *ApiService) ChangeMarginMode(symbol string, marginMode string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v2/position/changeMarginMode", map[string]string{"symbol": symbol, "marginMode": marginMode})
	return as.Call(req)
}

type MarginUserLeverage struct {
	Symbol   string `json:"symbol"`   //symbol of the contract
	Leverage string `json:"leverage"` //Leverage multiple
}

// GetCrossUserLeverage query the current symbol’s cross-margin leverage multiple
func (as *ApiService) GetCrossUserLeverage(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v2/getCrossUserLeverage", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// ChangeCrossUserLeverage  modify the current symbol’s cross-margin leverage multiple
func (as *ApiService) ChangeCrossUserLeverage(symbol string, leverage string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v2/changeCrossUserLeverage", map[string]string{"symbol": symbol, "leverage": leverage})
	return as.Call(req)
}

// MaxWithdrawMarginV1 This interface can query the maximum amount of margin that the current position supports withdrawal.
func (as *ApiService) MaxWithdrawMarginV1(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/maxWithdrawMargin", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type WithdrawMarginV1Req struct {
	Symbol         string `json:"symbol"`
	WithdrawAmount string `json:"withdrawAmount"`
}

// WithdrawMarginV1 Remove Margin Manually
func (as *ApiService) WithdrawMarginV1(r *WithdrawMarginV1Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/withdrawMargin", r)
	return as.Call(req)
}

type GetPositionsHistoryV1Req struct {
	Symbol string `url:"symbol,omitempty"`
	From   int64  `url:"from,omitempty"`
	To     int64  `url:"to,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	PageID int    `url:"pageId,omitempty"`
}

type PositionsHistoryItem struct {
	CloseID           string   `json:"closeId"`
	PositionID        string   `json:"positionId"`
	UID               int64    `json:"uid"`
	UserID            string   `json:"userId"`
	Symbol            string   `json:"symbol"`
	SettleCurrency    string   `json:"settleCurrency"`
	Leverage          string   `json:"leverage"`
	Type              string   `json:"type"`
	Side              *string  `json:"side"`
	CloseSize         *float64 `json:"closeSize"`
	PNL               string   `json:"pnl"`
	RealisedGrossCost string   `json:"realisedGrossCost"`
	WithdrawPNL       string   `json:"withdrawPnl"`
	ROE               *float64 `json:"roe"`
	TradeFee          string   `json:"tradeFee"`
	FundingFee        string   `json:"fundingFee"`
	OpenTime          int64    `json:"openTime"`
	CloseTime         int64    `json:"closeTime"`
	OpenPrice         *float64 `json:"openPrice"`
	ClosePrice        *float64 `json:"closePrice"`
}

type GetPositionsHistoryV1Resp struct {
	CurrentPage int                    `json:"currentPage"`
	PageSize    int                    `json:"pageSize"`
	TotalNum    int                    `json:"totalNum"`
	TotalPage   int                    `json:"totalPage"`
	Items       []PositionsHistoryItem `json:"items"`
}

// GetPositionsHistoryV1 This interface can query position history information records
func (as *ApiService) GetPositionsHistoryV1(r *GetPositionsHistoryV1Req) (*ApiResponse, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v1/history-positions", v)
	return as.Call(req)
}

type GetMaxOpenSizeReq struct {
	Symbol   string `url:"symbol"`
	Price    string `url:"price"`
	Leverage string `url:"leverage"`
}

type GetMaxOpenSizeResp struct {
	Symbol          string `json:"symbol"`
	MaxBuyOpenSize  int    `json:"maxBuyOpenSize"`
	MaxSellOpenSize int    `json:"leverage"`
}

// GetMaxOpenSize Get Maximum Open Position Size
func (as *ApiService) GetMaxOpenSize(r *GetMaxOpenSizeReq) (*ApiResponse, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v2/getMaxOpenSize", v)
	return as.Call(req)
}
