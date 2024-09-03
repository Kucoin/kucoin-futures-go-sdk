package kumex

import "testing"

func TestApiService_FundingHistory(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.FundingHistory(map[string]string{"symbol": "XBTUSDTM"})
	if err != nil {
		t.Fatal(err)
	}

	os := FundingListModel{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os.DataList {
		t.Log(ToJsonString(o))
		switch {
		case o.Id <= 0:
			t.Error("Empty key 'id'")
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		case o.MarkPrice == 0:
			t.Error("Empty key 'markPrice'")
		case o.FundingRate == 0:
			t.Error("Empty key 'fundingRate'")
		case o.Funding == 0:
			t.Error("Empty key 'funding'")
		case o.SettleCurrency == "":
			t.Error("Empty key 'settleCurrency'")
		}
	}
}

func TestApiService_FundingRatesTimeRange(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.FundingRatesTimeRange("XBTUSDTM", "1700310700000", "1702310700000")
	if err != nil {
		t.Fatal(err)
	}

	os := FundingTimeRangeRatesModel{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
	}
}

func TestApiService_TradeFeesV1(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.TradeFeesV1("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}

	os := TradeFeesV1Resp{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}
