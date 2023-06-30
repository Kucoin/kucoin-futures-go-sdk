package kumex

import (
	"net/http"
)

// A WithdrawalQuotasModel represents the quotas for a currency.
type WithdrawalQuotasModel struct {
	Currency            string  `json:"currency"`
	LimitAmount         float32 `json:"limitAmount"`
	UsedAmount          float32 `json:"usedAmount"`
	RemainAmount        float32 `json:"remainAmount"`
	AvailableAmount     float32 `json:"availableAmount"`
	WithdrawMinSize     float32 `json:"withdrawMinSize"`
	InnerWithdrawMinFee float32 `json:"innerWithdrawMinFee"`
	WithdrawMinFee      float32 `json:"withdrawMinFee"`
	IsWithdrawEnabled   bool    `json:"isWithdrawEnabled"`
	Precision           uint8   `json:"precision"`
}

// WithdrawalQuotas returns the quotas of withdrawal.
// Deprecated
func (as *ApiService) WithdrawalQuotas(currency string) (*ApiResponse, error) {
	params := map[string]string{"currency": currency}
	req := NewRequest(http.MethodGet, "/api/v1/withdrawals/quotas", params)
	return as.Call(req)
}

// ApplyWithdrawalResultModel represents the result of ApplyWithdrawal().
type ApplyWithdrawalResultModel struct {
	WithdrawalId string `json:"withdrawalId"`
}

// ApplyWithdrawal applies a withdrawal.
// Deprecated
func (as *ApiService) ApplyWithdrawal(currency, address, amount string, options map[string]string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"address":  address,
		"amount":   amount,
	}
	for k, v := range options {
		p[k] = v
	}
	req := NewRequest(http.MethodPost, "/api/v1/withdrawals", p)
	return as.Call(req)
}

// A WithdrawalModel represents a withdrawal.
type WithdrawalModel struct {
	WithdrawalId string `json:"withdrawalId"`
	Currency     string `json:"currency"`
	Status       string `json:"status"`
	Address      string `json:"address"`
	Memo         string `json:"memo"`
	IsInner      bool   `json:"isInner"`
	Amount       string `json:"amount"`
	Fee          string `json:"fee"`
	WalletTxId   string `json:"walletTxId"`
	CreatedAt    int64  `json:"createdAt"`
	Remark       string `json:"remark"`
	Reason       string `json:"reason"`
}

// A WithdrawalsModel is the set of *WithdrawalModel.
type WithdrawalsModel []*WithdrawalModel

// Withdrawals returns a list of withdrawals.
// Deprecated
func (as *ApiService) Withdrawals(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/withdrawal-list", params)
	return as.Call(req)
}

// CancelWithdrawalResultModel represents the result of CancelWithdrawal().
type CancelWithdrawalResultModel struct {
	CancelledWithdrawIds []string `json:"cancelledWithdrawIds"`
}

// CancelWithdrawal cancels a withdrawal by withdrawalId.
// Deprecated
func (as *ApiService) CancelWithdrawal(withdrawalId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/withdrawals/"+withdrawalId, nil)
	return as.Call(req)
}
