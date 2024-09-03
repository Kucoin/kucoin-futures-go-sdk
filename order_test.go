package kumex

import (
	"testing"
	"time"
)

func TestApiService_CreateOrder(t *testing.T) {

	s := NewApiServiceFromEnv()
	clientId := IntToString(time.Now().UnixNano())
	t.Log(clientId)
	p := map[string]string{
		"clientOid": clientId,
		"side":      "buy",
		"symbol":    "XBTUSDTM",
		"price":     "0.2",
		"size":      "1",
		"type":      "limit",
		"leverage":  "1",
	}
	rsp, err := s.CreateOrder(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &CreateOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.OrderId == "":
		t.Error("Empty key 'OrderId'")
	}
}

func TestApiService_CancelOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrder("order id")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case len(o.CancelledOrderIds) == 0:
		t.Error("Empty key 'cancelledOrderIds'")
	}
}

func TestApiService_CancelOrders(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrders("BTC")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case len(o.CancelledOrderIds) == 0:
		t.Error("Empty key 'cancelledOrderIds'")
	}
}

func TestApiService_Orders(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.Orders(map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}

	os := OrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		case o.Type == "":
			t.Error("Empty key 'type'")
		case o.Side == "":
			t.Error("Empty key 'side'")
		case o.SettleCurrency == "":
			t.Error("Empty key 'settleCurrency'")
		case o.Status == "":
			t.Error("Empty key 'status'")
		case o.UpdatedAt == 0:
			t.Error("Empty key 'UpdatedAt'")
		}
	}
}

func TestApiService_Order(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.Orders(map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}

	os := OrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	if len(os) == 0 {
		t.SkipNow()
	}

	rsp, err = s.Order(os[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	o := &OrderModel{}
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.Type == "":
		t.Error("Empty key 'type'")
	case o.Side == "":
		t.Error("Empty key 'side'")
	case o.SettleCurrency == "":
		t.Error("Empty key 'settleCurrency'")
	case o.Status == "":
		t.Error("Empty key 'status'")
	case o.UpdatedAt == 0:
		t.Error("Empty key 'UpdatedAt'")
	}
}

func TestApiService_OrderByClientOid(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.OrderByClientOid("eresc138b21023a909e5ad59")
	if err != nil {
		t.Fatal(err)
	}
	o := &OrderModel{}
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.Type == "":
		t.Error("Empty key 'type'")
	case o.Side == "":
		t.Error("Empty key 'side'")
	case o.SettleCurrency == "":
		t.Error("Empty key 'settleCurrency'")
	case o.Status == "":
		t.Error("Empty key 'status'")
	case o.UpdatedAt == 0:
		t.Error("Empty key 'UpdatedAt'")
	}
}

func TestApiService_RecentOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.RecentDoneOrders("MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}

	os := OrdersModel{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		case o.Type == "":
			t.Error("Empty key 'type'")
		case o.Side == "":
			t.Error("Empty key 'side'")
		case o.SettleCurrency == "":
			t.Error("Empty key 'settleCurrency'")
		case o.Status == "":
			t.Error("Empty key 'status'")
		case o.UpdatedAt == 0:
			t.Error("Empty key 'UpdatedAt'")
		}
	}
}

func TestApiService_StopOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.ObtainStopOrders("EDU-USDT", p)
	if err != nil {
		t.Fatal(err)
	}

	os := OrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		case o.Type == "":
			t.Error("Empty key 'type'")
		case o.Side == "":
			t.Error("Empty key 'side'")
		case o.SettleCurrency == "":
			t.Error("Empty key 'settleCurrency'")
		case o.Status == "":
			t.Error("Empty key 'status'")
		case o.UpdatedAt == 0:
			t.Error("Empty key 'UpdatedAt'")
		}
	}
}

func TestApiService_CancelOrderClientId(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrderClientId("1709979337447958000", "XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderClientIdResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_CreateMultiOrders(t *testing.T) {

	s := NewApiServiceFromEnv()
	p := make([]*CreateOrderReq, 0)
	p = append(p, &CreateOrderReq{
		ClientOid: IntToString(time.Now().UnixNano()),
		Side:      "buy",
		Symbol:    "XBTUSDTM",
		Leverage:  "1",
		Type:      "limit",
		Size:      "1",
		Price:     "0.3",
	})

	p = append(p, &CreateOrderReq{
		ClientOid: IntToString(time.Now().UnixNano()),
		Side:      "buy",
		Symbol:    "XBTUSDTM",
		Leverage:  "1",
		Type:      "limit",
		Size:      "1",
		Price:     "0.2",
	})

	rsp, err := s.CreateMultiOrders(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &CreateMultiOrdersRes{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_CreateSTOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	req := STOrderReq{
		ClientOid:            IntToString(time.Now().UnixNano()),
		Side:                 "buy",
		Symbol:               "XBTUSDM",
		Leverage:             "20",
		Type:                 "limit",
		Price:                "8000",
		Size:                 1,
		StopPriceType:        "TP",
		MarginMode:           "ISOLATED",
		TriggerStopUpPrice:   "9000",
		TriggerStopDownPrice: "8000",
		TimeInForce:          "GTC",
	}
	rsp, err := s.CreateSTOrder(&req)
	if err != nil {
		t.Fatal(err)
	}
	o := &STOrderRes{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}
