package kumex

import (
	"github.com/json-iterator/go"
	"math/big"
	"net/http"
)

// A TransferOutModel represents a transfer out record.
type TransferOutModel struct {
	ApplyId string `json:"applyId"`
}

// TransferOut Transfer Funds to KuCoin-Main Account.
// Deprecated
func (as *ApiService) TransferOut(bizNo, amount string) (*ApiResponse, error) {
	p := map[string]string{
		"bizNo":  bizNo,
		"amount": amount,
	}
	req := NewRequest(http.MethodPost, "/api/v1/transfer-out", p)
	return as.Call(req)
}

type TransferOutV2Model struct {
	ApplyId string `json:"applyId"`
}

// TransferOutV2 Transfer Funds to KuCoin-Main Account.
func (as *ApiService) TransferOutV2(bizNo, amount, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"bizNo":    bizNo,
		"amount":   amount,
		"currency": currency,
	}
	req := NewRequest(http.MethodPost, "/api/v2/transfer-out", p)
	return as.Call(req)
}

// A TransferModel represents a transfer record.
type TransferModel struct {
	ApplyId   string `json:"applyId"`
	Currency  string `json:"currency"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	Reason    string `json:"reason"`
	Offset    int64  `json:"offset"`
	CreatedAt int64  `json:"createdAt"`
}

// A TransfersModel  represents a transfer list.
type TransfersModel []*TransferModel

// TransferList returns a list of deposit.
func (as *ApiService) TransferList(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/transfer-list", params)
	return as.Call(req)
}

// CancelTransferModel represents the result of CancelWithdrawal().
type CancelTransferModel struct {
	ApplyId string `json:"applyId"`
}

// CancelTransfer Cancel Transfer-Out Request.
// Deprecated
func (as *ApiService) CancelTransfer(applyId string) (*ApiResponse, error) {
	p := map[string]string{
		"applyId": applyId,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/cancel/transfer-out", p)
	return as.Call(req)
}

// TransferOutV3 The amount to be transferred will be deducted from the KuCoin Futures Account.
// Please ensure that you have sufficient funds in your KuCoin Futures Account, or the transfer will fail.
// Once the transfer arrives your KuCoin-Main Account, the endpoint will respond and return the applyId.
// This ID could be used to cancel the transfer request.
func (as *ApiService) TransferOutV3(currency, recAccountType, amount string) (*ApiResponse, error) {
	p := map[string]string{
		"currency":       currency,
		"recAccountType": recAccountType,
		"amount":         amount,
	}
	req := NewRequest(http.MethodPost, "/api/v3/transfer-out", p)
	return as.Call(req)
}

type TransferOutV3Res struct {
	ApplyId        string          `json:"applyId"`
	BizNo          string          `json:"bizNo"`
	PayAccountType string          `json:"payAccountType"`
	PayTag         string          `json:"payTag"`
	Remark         string          `json:"remark"`
	RecAccountType string          `json:"recAccountType"`
	RecTag         string          `json:"recTag"`
	RecRemark      string          `json:"recRemark"`
	RecSystem      string          `json:"recSystem"`
	Status         string          `json:"status"`
	Currency       string          `json:"currency"`
	Amount         string          `json:"amount"`
	Fee            string          `json:"fee"`
	Sn             big.Int         `json:"sn"`
	Reason         string          `json:"reason"`
	CreatedAt      jsoniter.Number `json:"createdAt"`
	UpdatedAt      jsoniter.Number `json:"updatedAt"`
}

// TransferIn The amount to be transferred will be deducted from the KuCoin Futures Account.
// Please ensure that you have sufficient funds in your KuCoin Futures Account, or the transfer will fail.
// Once the transfer arrives your KuCoin-Main Account, the endpoint will respond and return the applyId.
// This ID could be used to cancel the transfer request.
func (as *ApiService) TransferIn(currency, payAccountType, amount string) (*ApiResponse, error) {
	p := map[string]string{
		"currency":       currency,
		"payAccountType": payAccountType,
		"amount":         amount,
	}
	req := NewRequest(http.MethodPost, "/api/v1/transfer-in", p)
	return as.Call(req)
}
