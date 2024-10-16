package kumex

import (
	"testing"
	"time"
)

func TestApiService_Position(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.Position("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	o := &PositionModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.MarkPrice == "":
		t.Error("Empty key 'markPrice'")
	case o.MarkValue == "":
		t.Error("Empty key 'markValue'")
	case o.SettleCurrency == "":
		t.Error("Empty key 'SettleCurrency'")
	}
}

func TestApiService_Positions(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.Positions("USDT")
	if err != nil {
		t.Fatal(err)
	}
	os := PositionsModel{}
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
		case o.MarkPrice == "":
			t.Error("Empty key 'markPrice'")
		case o.MarkValue == "":
			t.Error("Empty key 'markValue'")
		case o.SettleCurrency == "":
			t.Error("Empty key 'SettleCurrency'")
		}
	}
}

func TestApiService_autoDepositStatus(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol": "XBTUSDTM",
		"status": "true",
	}
	rsp, err := s.AutoDepositStatus(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &PositionModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	}
}

func TestApiService_DepositMargin(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol": "XBTUSDTM",
		"margin": "0.1111",
		"bizNo":  IntToString(time.Now().UnixNano()),
	}
	rsp, err := s.AutoDepositStatus(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &PositionModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	}
}

func TestApiService_GetMarginMode(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.GetMarginMode("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	var data MarginModeModel
	if err := rsp.ReadData(&data); err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestApiService_ChangeMarginMode(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.ChangeMarginMode("XBTUSDTM", "CROSS")
	if err != nil {
		t.Fatal(err)
	}
	var data MarginModeModel
	if err := rsp.ReadData(&data); err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestApiService_GetCrossUserLeverage(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.GetCrossUserLeverage("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	var data MarginUserLeverage
	if err := rsp.ReadData(&data); err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestApiService_ChangeCrossUserLeverage(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.ChangeCrossUserLeverage("XBTUSDTM", "6")
	if err != nil {
		t.Fatal(err)
	}
	var data bool
	if err := rsp.ReadData(&data); err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestApiService_MaxWithdrawMarginV1(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.MaxWithdrawMarginV1("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	var data string
	if err := rsp.ReadData(&data); err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestApiService_WithdrawMarginV1(t *testing.T) {

	s := NewApiServiceFromEnv()

	r := &WithdrawMarginV1Req{
		Symbol:         "XBTUSDTM",
		WithdrawAmount: "0.1",
	}
	rsp, err := s.WithdrawMarginV1(r)
	if err != nil {
		t.Fatal(err)
	}
	var data string
	if err := rsp.ReadData(&data); err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
func TestApiService_GetPositionsHistoryV1(t *testing.T) {
	s := NewApiServiceFromEnv()

	r := &GetPositionsHistoryV1Req{
		Symbol: "PEPEUSDTM",
	}
	rsp, err := s.GetPositionsHistoryV1(r)
	if err != nil {
		t.Fatal(err)
	}
	data := GetPositionsHistoryV1Resp{}
	if err := rsp.ReadData(&data); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(data))
}

func TestApiService_GetMaxOpenSize(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := GetMaxOpenSizeReq{}
	req.Symbol = "PEPEUSDTM"
	req.Price = "0.0000000001"
	req.Leverage = "10"

	resp, err := s.GetMaxOpenSize(&req)
	if err != nil {
		t.Fatal(err)
	}
	m := &GetMaxOpenSizeResp{}
	if err := resp.ReadData(m); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(m))
}
