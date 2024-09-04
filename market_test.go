package kumex

import (
	"testing"
)

func TestApiService_TickerLevel1(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Ticker("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	tk := &TickerLevel1Model{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))
	switch {
	case tk.Sequence <= 0:
		t.Error("Empty key 'sequence'")
	case tk.Price == "":
		t.Error("Empty key 'price'")
	case tk.Symbol == "":
		t.Error("Empty key 'Symbol'")
	case tk.BestBidPrice == "":
		t.Error("Empty key 'bestBidPrice'")
	case tk.BestBidSize <= 0:
		t.Error("Empty key 'bestBidSize'")
	case tk.BestAskPrice == "":
		t.Error("Empty key 'bestAskPrice'")
	case tk.BestAskSize <= 0:
		t.Error("Empty key 'bestAskSize'")
	}
}

func TestApiService_Level2Snapshot(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Level2Snapshot("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	tk := &Level2SnapshotModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))
	switch {
	case tk.Sequence <= 0:
		t.Error("Empty key 'sequence'")
	case len(tk.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(tk.Asks[0]) != 2:
		t.Error("Invalid ask length")
	case tk.Symbol == "":
		t.Error("Empty key 'Symbol'")
	case len(tk.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(tk.Bids[0]) != 2:
		t.Error("Invalid bid length")
	}
}

func TestApiService_Level2MessageQuery(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Level2MessageQuery("XBTUSDTM", 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	l := Level2MessageQueryListModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Sequence <= 0:
			t.Error("Empty key 'sequence'")
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.Change == "":
			t.Error("Empty key 'change'")
		}
	}
}

func TestApiService_Level3Snapshot(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Level3Snapshot("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	tk := &Level3SnapshotModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))

	if tk.Sequence <= 0 {
		t.Error("Empty key 'sequence'")
	}
	//switch {
	//case tk.Sequence <= 0:
	//	t.Error("Empty key 'sequence'")
	//case tk.Symbol == "":
	//	t.Error("Empty key 'Symbol'")
	//case len(tk.Asks) == 0:
	//	t.Error("Empty key 'asks'")
	//case len(tk.Asks[0]) != 3:
	//	t.Error("Invalid ask length")
	//case len(tk.Bids) == 0:
	//	t.Error("Empty key 'bids'")
	//case len(tk.Bids[0]) != 3:
	//	t.Error("Invalid bid length")
	//}
}

func TestApiService_Level3SnapshotV2(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Level3SnapshotV2("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	tk := &Level3SnapshotV2Model{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))

	if tk.Sequence <= 0 {
		t.Error("Empty key 'sequence'")
	}
}

func TestApiService_Level3MessageQuery(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Level3MessageQuery("XBTUSDTM", 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	l := Level3MessageQueryListModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Sequence <= 0:
			t.Error("Empty key 'sequence'")
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.OrderId == "":
			t.Error("Empty key 'orderId'")
		}
	}
}

func TestApiService_TradeHistory(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.TradeHistory("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	l := TradesHistoryModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Sequence <= 0:
			t.Error("Empty key 'sequence'")
		case c.TradeId == "":
			t.Error("Empty key 'tradeId'")
		case c.Price == "":
			t.Error("Empty key 'price'")
		case c.Size <= 0:
			t.Error("Empty key 'size'")
		case c.Side == "":
			t.Error("Empty key 'side'")
		}
	}
}

func TestApiService_InterestQuery(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	p["symbol"] = ".XBTINT"
	pp := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.InterestQuery(p, pp)
	if err != nil {
		t.Fatal(err)
	}
	l := InterestsModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}

	for _, c := range l.DataList {
		t.Log(ToJsonString(c))
		switch {
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.Granularity <= 0:
			t.Error("Empty key 'granularity'")
		case c.TimePoint == 0:
			t.Error("Empty key 'timePoint'")
		case c.Value <= 0:
			t.Error("Empty key 'value'")
		}
	}
}

func TestApiService_IndexQuery(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	p["symbol"] = "XBTUSDTM"
	pp := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.IndexQuery(p, pp)
	if err != nil {
		t.Fatal(err)
	}
	l := IndexQueryModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l.DataList {
		t.Log(ToJsonString(c))
		switch {
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.Granularity <= 0:
			t.Error("Empty key 'granularity'")
		case c.TimePoint <= 0:
			t.Error("Empty key 'timePoint'")
		case c.Value <= 0:
			t.Error("Empty key 'value'")
		}
	}
}

func TestApiService_MarkPrice(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarkPrice("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	tk := &MarkPriceModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))
	switch {
	case tk.Granularity <= 0:
		t.Error("Empty key 'granularity'")
	case tk.TimePoint <= 0:
		t.Error("Empty key 'timePoint'")
	case tk.Symbol == "":
		t.Error("Empty key 'symbol'")
	case tk.IndexPrice <= 0:
		t.Error("Empty key 'indexPrice'")
	}
}

func TestApiService_PremiumQuery(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	p["symbol"] = "XBTUSDTM"
	pp := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.PremiumQuery(p, pp)
	if err != nil {
		t.Fatal(err)
	}
	l := PremiumsModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l.DataList {
		t.Log(ToJsonString(c))
		switch {
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.Granularity == "":
			t.Error("Empty key 'granularity'")
		case c.TimePoint == "":
			t.Error("Empty key 'timePoint'")
		case c.Value == "":
			t.Error("Empty key 'value'")
		}
	}
}

func TestApiService_FundingRate(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.FundingRate(".XBTUSDTMFPI8H")
	if err != nil {
		t.Fatal(err)
	}
	tk := &FundingRateModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))
	switch {
	case tk.Granularity <= 0:
		t.Error("Empty key 'granularity'")
	case tk.TimePoint <= 0:
		t.Error("Empty key 'timePoint'")
	case tk.Symbol == "":
		t.Error("Empty key 'symbol'")
	case tk.PredictedValue == 0:
		t.Error("Empty key 'predictedValue'")
	}
}

func TestApiService_TradeStatistics(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.TradeStatistics()
	if err != nil {
		t.Fatal(err)
	}
	tk := &TradeStatisticsModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))
}

func TestApiService_AllTickers(t *testing.T) {
	s := NewApiServiceFromEnv()
	resp, err := s.AllTickers()
	if err != nil {
		t.Fatal(err)
	}
	m := &AllTickersModel{}
	if err := resp.ReadData(m); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(m))
}
