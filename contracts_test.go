package kumex

import "testing"

func TestApiService_ActiveContracts(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.ActiveContracts()
	if err != nil {
		t.Fatal(err)
	}
	c := ContractsModels{}
	if err := rsp.ReadData(&c); err != nil {
		t.Fatal(err)
	}
	for _, o := range c {
		t.Log(ToJsonString(o))
		switch {
		case o.BaseCurrency == "":
			t.Error("Empty key 'baseCurrency'")
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		//case o.FairMethod == "":
		//	t.Error("Empty key 'fairMethod'")
		case o.IndexSymbol == "":
			t.Error("Empty key 'indexSymbol'")
		case o.MaxLeverage < 0:
			t.Error("Empty key 'maxLeverage'")
		}
	}

}

func TestApiService_Contracts(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.Contracts("XBTUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	o := &ContractsModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.BaseCurrency == "":
		t.Error("Empty key 'baseCurrency'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.FairMethod == "":
		t.Error("Empty key 'fairMethod'")
	case o.IndexSymbol == "":
		t.Error("Empty key 'indexSymbol'")
	case o.MaxLeverage < 0:
		t.Error("Empty key 'maxLeverage'")
	}
}
