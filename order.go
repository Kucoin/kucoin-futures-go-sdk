package kumex

import "net/http"

// A CreateOrderResultModel represents the result of CreateOrder().
type CreateOrderResultModel struct {
	OrderId string `json:"orderId"`
}

// CreateOrder places a new order.
func (as *ApiService) CreateOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/orders", params)
	return as.Call(req)
}

// A CancelOrderResultModel represents the result of CancelOrder().
type CancelOrderResultModel struct {
	CancelledOrderIds []string `json:"cancelledOrderIds"`
}

// CancelOrder cancels a previously placed order.
func (as *ApiService) CancelOrder(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/orders/"+orderId, nil)
	return as.Call(req)
}

// CancelOrders cancels all orders of the symbol.
// With best effort, cancel all open orders. The response is a list of ids of the canceled orders.
func (as *ApiService) CancelOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodDelete, "/api/v1/orders", p)
	return as.Call(req)
}

// StopOrders represents an order.
func (as *ApiService) StopOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodDelete, "/api/v1/stopOrders", p)
	return as.Call(req)
}

// ObtainStopOrders represents an order.
func (as *ApiService) ObtainStopOrders(symbol string, page *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	page.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v1/stopOrders", p)
	return as.Call(req)
}

// An OrderModel represents an order.
type OrderModel struct {
	Id             string `json:"id"`
	Symbol         string `json:"symbol"`
	Type           string `json:"type"`
	Side           string `json:"side"`
	Price          string `json:"price"`
	Size           int64  `json:"size"`
	Value          string `json:"value"`
	DealValue      string `json:"dealValue"`
	DealSize       int64  `json:"dealSize"`
	Stp            string `json:"stp"`
	Stop           string `json:"stop"`
	StopPriceType  string `json:"stopPriceType"`
	StopTriggered  bool   `json:"stopTriggered"`
	StopPrice      string `json:"stopPrice"`
	TimeInForce    string `json:"timeInForce"`
	PostOnly       bool   `json:"postOnly"`
	Hidden         bool   `json:"hidden"`
	IceBerg        bool   `json:"iceberg"`
	VisibleSize    string `json:"visibleSize"`
	Leverage       string `json:"leverage"`
	ForceHold      bool   `json:"forceHold"`
	CloseOrder     bool   `json:"closeOrder"`
	CloseOnly      bool   `json:"closeOnly"`
	ClientOid      string `json:"clientOid"`
	Remark         string `json:"remark"`
	IsActive       bool   `json:"isActive"`
	CancelExist    bool   `json:"cancelExist"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	SettleCurrency string `json:"settleCurrency"`
	Status         string `json:"status"`
}

// A OrdersModel is the set of *OrderModel.
type OrdersModel []*OrderModel

// Orders returns a list your current orders.
func (as *ApiService) Orders(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/orders", params)
	return as.Call(req)
}

// Order returns a single order by order id.
func (as *ApiService) Order(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/orders/"+orderId, nil)
	return as.Call(req)
}

// Order returns a single order by client Oid.
func (as *ApiService) OrderByClientOid(clientOid string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/orders/byClientOid?clientOid="+clientOid, nil)
	return as.Call(req)
}

// RecentDoneOrders returns the recent orders of the latest transactions within 24 hours.
func (as *ApiService) RecentDoneOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/recentDoneOrders", p)
	return as.Call(req)
}
