package kumex

import (
	"net/http"
)

// A TickerLevel1Model represents ticker include only the inside (i.e. best) bid and ask data, last price and last trade size.
type TickerLevel1Model struct {
	Sequence     int    `json:"sequence"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	Size         int    `json:"size"`
	Price        string `json:"price"`
	BestBidSize  int    `json:"bestBidSize"`
	BestBidPrice string `json:"bestBidPrice"`
	BestAskSize  int    `json:"bestAskSize"`
	BestAskPrice string `json:"bestAskPrice"`
	TradeId      string `json:"tradeId"`
	Ts           int64  `json:"ts"`
}

// Ticker Get Real-Time Ticker.
func (as *ApiService) Ticker(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/ticker", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// Level2SnapshotModel represents level2 ticker.
type Level2SnapshotModel struct {
	Symbol   string      `json:"symbol"`
	Sequence int         `json:"sequence"`
	Asks     [][]float32 `json:"asks"`
	Bids     [][]float32 `json:"bids"`
}

// Level2Snapshot Get Full Order Book - Level 2.
func (as *ApiService) Level2Snapshot(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level2/snapshot", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// Level2MessageQueryModel represents level2 ticker message.
type Level2MessageQueryModel struct {
	Symbol   string `json:"symbol"`
	Sequence int    `json:"sequence"`
	Change   string `json:"change"`
}

// Level2MessageQueryListModel the set of *Level2MessageQueryModel.
type Level2MessageQueryListModel []*Level2MessageQueryModel

// Level2MessageQuery Level 2 Pulling Messages.
// Deprecated
func (as *ApiService) Level2MessageQuery(symbol string, start, end int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level2/message/query", map[string]string{
		"symbol": symbol,
		"start":  IntToString(start),
		"end":    IntToString(end),
	})
	return as.Call(req)
}

// Level3SnapshotModel represents level3 ticker message.
type Level3SnapshotModel struct {
	Symbol   string          `json:"symbol"`
	Sequence int             `json:"sequence"`
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
}

// Level3Snapshot Get Full Order Book - Level 3.
func (as *ApiService) Level3Snapshot(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level3/snapshot", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// Level3SnapshotV2Model represents level3 ticker message.
type Level3SnapshotV2Model struct {
	Symbol   string          `json:"symbol"`
	Sequence int             `json:"sequence"`
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
	Ts       int64           `json:"ts"`
}

// Level3SnapshotV2 Get Full Order Book - Level 3.
func (as *ApiService) Level3SnapshotV2(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v2/level3/snapshot", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// Level3MessageQueryModel represents level3 ticker message.
type Level3MessageQueryModel struct {
	Symbol    string `json:"symbol"`
	Sequence  int    `json:"sequence"`
	Side      string `json:"side"`
	OrderTime int64  `json:"orderTime"`
	Size      int    `json:"size"`
	OrderId   string `json:"orderId"`
	Price     string `json:"price"`
	Type      string `json:"type"`
	ClientOid string `json:"clientOid"`
	Ts        int64  `json:"ts"`
}

// Level3MessageQueryListModel is the set of *Level3MessageQueryModel
type Level3MessageQueryListModel []*Level3MessageQueryModel

// Level3MessageQuery Level 3 Pulling Messages.
func (as *ApiService) Level3MessageQuery(symbol string, start, end int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level3/message/query", map[string]string{
		"symbol": symbol,
		"start":  IntToString(start),
		"end":    IntToString(end),
	})
	return as.Call(req)
}

// TradeHistoryModel represents a the latest trades for a symbol.
type TradeHistoryModel struct {
	Sequence     int    `json:"sequence"`
	TradeId      string `json:"tradeId"`
	TakerOrderId string `json:"takerOrderId"`
	MakerOrderId string `json:"makerOrderId"`
	Price        string `json:"price"`
	Size         int    `json:"size"`
	Side         string `json:"side"`
	Time         int64  `json:"t"`
}

// TradesHistoryModel is the set of *TradeHistoryModel.
type TradesHistoryModel []*TradeHistoryModel

// TradeHistory returns a list the latest trades for a symbol.
func (as *ApiService) TradeHistory(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/trade/history", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// InterestModel is the struct.
type InterestModel struct {
	Symbol      string  `json:"symbol"`
	Granularity int     `json:"granularity"`
	TimePoint   int64   `json:"timePoint"`
	Value       float32 `json:"value"`
}

// InterestsModel is the set of *InterestModel.
type InterestsModel struct {
	HasMore  bool             `json:"hasMore"`
	DataList []*InterestModel `json:"dataList"` // delay parsing
}

// InterestQuery Get Interest Rate List .
func (as *ApiService) InterestQuery(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/interest/query", params)
	return as.Call(req)
}

// IndexModel is the struct.
type IndexModel struct {
	Symbol          string          `json:"symbol"`
	Granularity     int             `json:"granularity"`
	TimePoint       int64           `json:"timePoint"`
	Value           float32         `json:"value"`
	DecomposionList [][]interface{} `json:"decomposionList"`
}

// A IndexQueryModel is the set of *IndexModel.
type IndexQueryModel struct {
	HasMore  bool          `json:"hasMore"`
	DataList []*IndexModel `json:"dataList"` // delay parsing
}

// IndexQuery Get Index List.
func (as *ApiService) IndexQuery(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/interest/query", params)
	return as.Call(req)
}

// A MarkPriceModel is the struct.
type MarkPriceModel struct {
	Symbol      string  `json:"symbol"`
	Granularity float32 `json:"granularity"`
	TimePoint   int64   `json:"timePoint"`
	Value       float32 `json:"value"`
	IndexPrice  float32 `json:"indexPrice"`
}

// MarkPrice Get Current Mark Price
func (as *ApiService) MarkPrice(Symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/mark-price/"+Symbol+"/current", nil)
	return as.Call(req)
}

// A PremiumModel is the struct.
type PremiumModel struct {
	Symbol      string `json:"symbol"`
	Granularity string `json:"granularity"`
	TimePoint   string `json:"timePoint"`
	Value       string `json:"value"`
}

// A PremiumsModel is the set of *PremiumModel.
type PremiumsModel struct {
	HasMore  bool            `json:"hasMore"`
	DataList []*PremiumModel `json:"dataList"` // delay parsing
}

// PremiumQuery Get Premium Index.
func (as *ApiService) PremiumQuery(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/premium/query", params)
	return as.Call(req)
}

// A FundingRateModel is the struct.
type FundingRateModel struct {
	Symbol         string  `json:"symbol"`
	Granularity    int64   `json:"granularity"`
	TimePoint      int64   `json:"timePoint"`
	Value          float32 `json:"value"`
	PredictedValue float32 `json:"predictedValue"`
}

// FundingRate Get Current Funding Rate.
func (as *ApiService) FundingRate(Symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/funding-rate/"+Symbol+"/current", nil)
	return as.Call(req)
}
