package kumex

import "net/http"

// KLineModel represents the k lines for a symbol.
type KLineModel []interface{}

// A KLinesModel is the set of *KLineModel.
type KLinesModel []*KLineModel

// Data are returned in grouped buckets based on requested type.
// Parameter #2 granularity is the type of granularity patterns(minute).
// Parameter #3 #4 startAt, endAt is millisecond.
func (as *ApiService) KLines(symbol, granularity string, startAt, endAt int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/kline/query", map[string]string{
		"symbol":      symbol,
		"granularity": granularity,
		"startAt":     IntToString(startAt),
		"endAt":       IntToString(endAt),
	})
	return as.Call(req)
}
