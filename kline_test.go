package kumex

import (
	"testing"
	"time"
)

func TestApiService_KLines(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.KLines("XBTUSDTM", "5", time.Now().UnixNano()/1e6-7*24*3600, time.Now().UnixNano()/1e6)
	if err != nil {
		t.Fatal(err)
	}
	l := KLinesModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		if len(*c) != 6 {
			t.Error("Invalid length of rate")
		}
	}
}
