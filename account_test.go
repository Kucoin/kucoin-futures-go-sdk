package kumex

import (
	"testing"
)

func TestApiService_AccountOverview(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := map[string]string{}
	p["currency"] = "XBT"

	rsp, err := s.AccountOverview(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &AccountModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.AccountEquity == 0:
		t.Error("Empty key 'accountEquity'")
	case o.MarginBalance < 0:
		t.Error("Empty key 'marginBalance'")
	case o.PositionMargin < 0:
		t.Error("Empty key 'positionMargin'")
	case o.OrderMargin < 0:
		t.Error("Empty key 'orderMargin'")
	case o.FrozenFunds < 0:
		t.Error("Empty key 'frozenFunds'")
	case o.AvailableBalance <= 0:
		t.Error("Empty key 'availableBalance'")
	case o.Currency == "":
		t.Error("Empty key 'currency'")
	}
}

func TestApiService_TransactionHistory(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	pp := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.TransactionHistory(p, pp)
	if err != nil {
		t.Fatal(err)
	}
	ds := TransactionHistoryListModel{}
	if _, err := rsp.ReadPaginationData(&ds); err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		t.Log(ToJsonString(d))
		switch {
		case d.Time == "":
			t.Error("Empty key 'time'")
		case d.Type == "":
			t.Error("Empty key 'type'")
		case d.Amount == "":
			t.Error("Empty key 'amount'")
		case d.Status == "":
			t.Error("Empty key 'status'")
		case d.Fee == "":
			t.Error("Empty key 'fee'")
		case d.AccountEquity == "":
			t.Error("Empty key 'accountEquity'")
		case d.Remarks == "":
			t.Error("Empty key 'remark'")
		case d.Offset == "":
			t.Error("Empty key 'Offset'")
		}
	}
}
