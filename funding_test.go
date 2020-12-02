package kumex

import "testing"

func TestApiService_FundingHistory(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.FundingHistory(map[string]string{"symbol": "XBTUSDM"})
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
