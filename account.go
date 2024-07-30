package kumex

import (
	"github.com/json-iterator/go"
	"net/http"
)

// An AccountModel represents an account.
type AccountModel struct {
	AccountEquity    float64 `json:"accountEquity"`
	UnrealisedPNL    float64 `json:"unrealisedPNL"`
	MarginBalance    float64 `json:"marginBalance"`
	PositionMargin   float64 `json:"positionMargin"`
	OrderMargin      float64 `json:"orderMargin"`
	FrozenFunds      float64 `json:"frozenFunds"`
	AvailableBalance float64 `json:"availableBalance"`
	Currency         string  `json:"currency"`
}

// An AccountsModel is the set of *AccountModel.
type AccountsModel []*AccountModel

// AccountOverview returns a list of accounts.
// See the Deposits section for documentation on how to deposit funds to begin trading.
func (as *ApiService) AccountOverview(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/account-overview", params)
	return as.Call(req)
}

// A TransactionHistoryModel represents a sub-account user.
type TransactionHistoryModel struct {
	Time          string `json:"time"`
	Type          string `json:"type"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	AccountEquity string `json:"accountEquity"`
	Status        string `json:"status"`
	Remarks       string `json:"remark"`
	Offset        string `json:"offset"`
	Currency      string `json:"currency"`
}

// An TransactionHistoryListModel the set of *TransactionHistoryModel.
type TransactionHistoryListModel []*TransactionHistoryModel

// TransactionHistory returns a list of ledgers.
// Account activity either increases or decreases your account balance.
// Items are paginated and sorted latest first.
func (as *ApiService) TransactionHistory(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/transaction-history", params)
	return as.Call(req)
}

// SubApiKeys This endpoint can be used to obtain a list of Futures APIs pertaining to a sub-account.
func (as *ApiService) SubApiKeys(apiKey, subName string) (*ApiResponse, error) {
	p := map[string]string{
		"apiKey":  apiKey,
		"subName": subName,
	}
	req := NewRequest(http.MethodGet, "/api/v1/sub/api-key", p)
	return as.Call(req)
}

type SubApiKeysModel []*SubApiKeyModel

type SubApiKeyModel struct {
	SubName     string          `json:"subName"`
	Remark      string          `json:"remark"`
	ApiKey      string          `json:"apiKey"`
	Permission  string          `json:"permission"`
	IpWhitelist string          `json:"ipWhitelist"`
	CreatedAt   jsoniter.Number `json:"createdAt"`
}

// CreateSubApiKey This endpoint can be used to create Futures APIs for sub-accounts.
func (as *ApiService) CreateSubApiKey(p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/sub/api-key", p)
	return as.Call(req)
}

type CreateSubApiKeyRes struct {
	SubName     string          `json:"subName"`
	Remark      string          `json:"remark"`
	ApiKey      string          `json:"apiKey"`
	Permission  string          `json:"permission"`
	IpWhitelist string          `json:"ipWhitelist"`
	CreatedAt   jsoniter.Number `json:"createdAt"`
	ApiSecret   string          `json:"apiSecret"`
	Passphrase  string          `json:"passphrase"`
}

// ModifySubApiKey TThis endpoint can be used to modify sub-account Futures APIs.
func (as *ApiService) ModifySubApiKey(p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/sub/api-key/update", p)
	return as.Call(req)
}

type ModifySubApiKeyRes struct {
	SubName     string `json:"subName"`
	Permission  string `json:"permission"`
	IpWhitelist string `json:"ipWhitelist"`
	ApiKey      string `json:"apiKey"`
}

// DeleteSubApiKey This endpoint can be used to delete sub-account Futures APIs.
func (as *ApiService) DeleteSubApiKey(apiKey, passphrase, subName string) (*ApiResponse, error) {
	p := map[string]string{
		"apiKey":     apiKey,
		"passphrase": passphrase,
		"subName":    subName,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/sub/api-key", p)
	return as.Call(req)
}

type DeleteSubApiKeyRes struct {
	ApiKey  string `json:"apiKey"`
	SubName string `json:"subName"`
}

// SubAccountsBalance Get All Sub-Accounts Balance - Futures
func (as *ApiService) SubAccountsBalance(currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/account-overview-all", p)
	return as.Call(req)
}

type SubAccountBalanceModel struct {
	Summary struct {
		AccountEquityTotal    jsoniter.Number `json:"accountEquityTotal"`
		UnrealisedPNLTotal    jsoniter.Number `json:"unrealisedPNLTotal"`
		MarginBalanceTotal    jsoniter.Number `json:"marginBalanceTotal"`
		PositionMarginTotal   jsoniter.Number `json:"positionMarginTotal"`
		OrderMarginTotal      jsoniter.Number `json:"orderMarginTotal"`
		FrozenFundsTotal      jsoniter.Number `json:"frozenFundsTotal"`
		AvailableBalanceTotal jsoniter.Number `json:"availableBalanceTotal"`
		Currency              string          `json:"currency"`
	} `json:"summary"`
	Accounts []struct {
		AccountName      string          `json:"accountName"`
		AccountEquity    jsoniter.Number `json:"accountEquity"`
		UnrealisedPNL    jsoniter.Number `json:"unrealisedPNL"`
		MarginBalance    jsoniter.Number `json:"marginBalance"`
		PositionMargin   jsoniter.Number `json:"positionMargin"`
		OrderMargin      jsoniter.Number `json:"orderMargin"`
		FrozenFunds      jsoniter.Number `json:"frozenFunds"`
		AvailableBalance jsoniter.Number `json:"availableBalance"`
		Currency         string          `json:"currency"`
	} `json:"accounts"`
}
