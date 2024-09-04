package kumex

import (
	"testing"
)

func TestApiService_Fills(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.Fills(map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}

	fs := FillsModel{}
	if _, err := rsp.ReadPaginationData(&fs); err != nil {
		t.Fatal(err)
	}
	for _, f := range fs {
		t.Log(ToJsonString(f))
		switch {
		case f.Symbol == "":
			t.Error("Empty key 'symbol'")
		case f.TradeId == "":
			t.Error("Empty key 'tradeId'")
		case f.OrderId == "":
			t.Error("Empty key 'orderId'")
		case f.Side == "":
			t.Error("Empty key 'side'")
		case f.SettleCurrency == "":
			t.Error("Empty key 'settleCurrency'")
		case f.TradeTime == 0:
			t.Error("Empty key 'tradeTime'")
		}
	}
}

func TestApiService_RecentFills(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.RecentFills()
	if err != nil {
		t.Fatal(err)
	}

	fs := FillsModel{}
	if err := rsp.ReadData(&fs); err != nil {
		t.Fatal(err)
	}
	for _, f := range fs {
		t.Log(ToJsonString(f))
		switch {
		case f.Symbol == "":
			t.Error("Empty key 'symbol'")
		case f.TradeId == "":
			t.Error("Empty key 'tradeId'")
		case f.OrderId == "":
			t.Error("Empty key 'orderId'")
		case f.Side == "":
			t.Error("Empty key 'side'")
		case f.SettleCurrency == "":
			t.Error("Empty key 'settleCurrency'")
		case f.TradeTime == 0:
			t.Error("Empty key 'tradeTime'")
		}
	}
}

func TestApiService_OpenOrderStatistics(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.OpenOrderStatistics("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	o := &OpenOrderStatisticsModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.OpenOrderBuySize < 0:
		t.Error("Empty key 'OpenOrderBuySize'")
	case o.OpenOrderSellSize < 0:
		t.Error("Empty key 'OpenOrderSellSize'")
	case o.OpenOrderBuyCost == "":
		t.Error("Empty key 'OpenOrderBuyCost'")
	case o.OpenOrderSellCost == "":
		t.Error("Empty key 'OpenOrderSellCost'")
	case o.SettleCurrency == "":
		t.Error("Empty key 'SettleCurrency'")
	}
}
