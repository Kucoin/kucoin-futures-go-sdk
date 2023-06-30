package kumex

import (
	"testing"
)

func TestApiService_TransferOut(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.TransferOut("123", "0.1")
	if err != nil {
		t.Fatal(err)
	}
	w := &TransferOutModel{}
	if err := rsp.ReadData(w); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(w))
	switch {
	case w.ApplyId == "":
		t.Error("Empty key 'applyId'")
	}
}

func TestApiService_TransferList(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	pp := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.TransferList(p, pp)
	if err != nil {
		t.Fatal(err)
	}
	ds := TransfersModel{}
	if _, err := rsp.ReadPaginationData(&ds); err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		t.Log(ToJsonString(d))
		switch {
		case d.ApplyId == "":
			t.Error("Empty key 'applyId'")
		case d.Amount == "":
			t.Error("Empty key 'amount'")
		case d.Currency == "":
			t.Error("Empty key 'currency'")
		case d.Offset == 0:
			t.Error("Empty key 'offset'")
		case d.Status == "":
			t.Error("Empty key 'status'")
		case d.CreatedAt == 0:
			t.Error("Empty key 'createdAt'")
		}
	}
}

func TestApiService_CancelTransfer(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelTransfer("xxx")
	if err != nil {
		t.Fatal(err)
	}
	w := &CancelTransferModel{}
	if err := rsp.ReadData(w); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(w))
	switch {
	case len(w.ApplyId) == 0:
		t.Error("Empty key 'applyId'")
	}
}

func TestApiService_TransferIn(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.TransferIn("USDT", "TRADE", "2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(rsp.Message))
	t.Log(ToJsonString(rsp.Code))
}

func TestApiService_TransferOutV3(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.TransferOutV3("USDT", "TRADE", "0.5")
	if err != nil {
		t.Fatal(err)
	}
	w := &TransferOutV3Res{}
	if err := rsp.ReadData(w); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(w))
}
