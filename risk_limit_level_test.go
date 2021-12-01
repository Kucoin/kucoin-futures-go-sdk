package kumex

import "testing"

func TestApiService_ContractsRiskLimitLevel(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.ContractsRiskLimitLevel("ADAUSDTM")
	if err != nil {
		t.Fatal(err)
	}
	o := &ContractsRiskLimitLevelModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_ChangeRiskLimitLevel(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.ChangeRiskLimitLevel(map[string]string{"symbol": "ADAUSDTM", "level": "2"})
	if err != nil {
		t.Fatal(err)
	}
	var ret bool
	if err := rsp.ReadData(&ret); err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
