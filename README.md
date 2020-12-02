# Go SDK for KuMex API
> The detailed document [https://docs.kumex.com](https://docs.kumex.com), in order to receive the latest API change notifications, please `Watch` this repository.

[![Latest Version](https://img.shields.io/github/release/Kucoin/kumex-go-sdk.svg)](https://github.com/Kucoin/kumex-go-sdk/releases)
[![GoDoc](https://godoc.org/github.com/Kucoin/kumex-go-sdk?status.svg)](https://godoc.org/github.com/Kucoin/kumex-go-sdk)
[![Build Status](https://travis-ci.org/Kucoin/kumex-go-sdk.svg?branch=master)](https://travis-ci.org/Kucoin/kumex-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kucoin/kumex-go-sdk)](https://goreportcard.com/report/github.com/Kucoin/kumex-go-sdk)
[![Sourcegraph](https://sourcegraph.com/github.com/Kucoin/kumex-go-sdk/-/badge.svg)](https://sourcegraph.com/github.com/Kucoin/kumex-go-sdk?badge)
<!-- [![Total Lines](https://tokei.rs/b1/github/Kucoin/kumex-go-sdk)](https://github.com/Kucoin/kumex-go-sdk) -->


## Install

```bash
go get github.com/Kucoin/kumex-go-sdk
```

## Usage

### Choose environment

| Environment | BaseUri |
| -------- | -------- |
| *Production* | `https://api.kumex.com(DEFAULT)` `https://api.kumex.top` |
| *Sandbox* | `https://sandbox-api.kumex.com` |

### Create ApiService

```go
s := kumex.NewApiService( 
	// kumex.ApiBaseURIOption("https://api.kumex.com"), 
	kumex.ApiKeyOption("key"),
	kumex.ApiSecretOption("secret"),
	kumex.ApiPassPhraseOption("passphrase"),
)

// Or add these options into the environmental variable
// Bash: 
// export API_BASE_URI=https://api.kumex.com
// export API_KEY=key
// export API_SECRET=secret
// export API_PASSPHRASE=passphrase
// s := NewApiServiceFromEnv()
```

### Debug mode & logging

```go
// Require package github.com/sirupsen/logrus
// Debug mode will record the logs of API and WebSocket to files.
// Default values: LogLevel=logrus.DebugLevel, LogDirectory="/tmp"
kumex.DebugMode = true
// Or export API_DEBUG_MODE=1

// Logging in your code
// kumex.SetLoggerDirectory("/tmp")
// logrus.SetLevel(logrus.DebugLevel)
logrus.Debugln("I'm a debug message")
```

### Examples
> See the test case for more examples.

#### Example of API `without` authentication

```go
rsp, err := s.ServerTime()
if err != nil {
    log.Printf("Error: %s", err.Error())
    // Handle error
    return
}

var ts int64
if err := rsp.ReadData(&ts); err != nil {
    // Handle error
    return
}
log.Printf("The server time: %d", ts)
```

#### Example of API `with` authentication

```go
// Without pagination
rsp, err := s.AccountOverview()
if err != nil {
    // Handle error
    return
}

as := kumex.AccountsModel{}
if err := rsp.ReadData(&as); err != nil {
    // Handle error
    return
}

for _, a := range as {
    log.Printf("Available balance: %s %s => %s", a.Type, a.Currency, a.Available)
}
```

```go
// Handle pagination
rsp, err := s.Orders(map[string]string{}, &kumex.PaginationParam{CurrentPage: 1, PageSize: 10})
if err != nil {
    // Handle error
    return
}

os := kumex.OrdersModel{}
pa, err := rsp.ReadPaginationData(&os)
if err != nil {
    // Handle error
    return
}
log.Printf("Total num: %d, total page: %d", pa.TotalNum, pa.TotalPage)
for _, o := range os {
    log.Printf("Order: %s, %s, %s", o.Id, o.Type, o.Price)
}
```

#### Example of WebSocket feed
> Require package [gorilla/websocket](https://github.com/gorilla/websocket)

```bash
go get github.com/gorilla/websocket github.com/pkg/errors
```

```go
rsp, err := s.WebSocketPublicToken()
if err != nil {
    // Handle error
    return
}

tk := &kumex.WebSocketTokenModel{}
if err := rsp.ReadData(tk); err != nil {
    // Handle error
    return
}

c := s.NewWebSocketClient(tk)
// c.AcceptUserMessage = true 


mc, ec, err := c.Connect()
if err != nil {
    // Handle error
    return
}

ch1 := kumex.NewSubscribeMessage("/contractMarket/ticker:XBTUSDM", false)
ch2 := kumex.NewSubscribeMessage("/contractMarket/ticker:XBTUSDM", false)
uch := kumex.NewUnsubscribeMessage("/contractMarket/ticker:XBTUSDM", false)

if err := c.Subscribe(ch1, ch2); err != nil {
    // Handle error
    return
}

var i = 0
for {
    select {
    case err := <-ec:
        c.Stop() // Stop subscribing the WebSocket feed
        log.Printf("Error: %s", err.Error())
        // Handle error
        return
    case msg := <-mc:
        // log.Printf("Received: %s", kumex.ToJsonString(m))
        t := &kumex.TickerLevel1Model{}
        if err := msg.ReadData(t); err != nil {
            log.Printf("Failure to read: %s", err.Error())
            return
        }
        log.Printf("Ticker: %s, %s, %s, %s", msg.Topic, t.Sequence, t.Price, t.Size)
        i++
        if i == 5 {
            log.Println("Unsubscribe XBTUSDM")
            if err = c.Unsubscribe(uch); err != nil {
                log.Printf("Error: %s", err.Error())
                // Handle error
                return
            }
        }
        if i == 10 {
            log.Println("Subscribe XBTUSDM")
            if err = c.Subscribe(ch2); err != nil {
                log.Printf("Error: %s", err.Error())
                // Handle error
                return
            }
        }
        if i == 15 {
            log.Println("Exit subscription")
            c.Stop() // Stop subscribing the WebSocket feed
            return
        }
    }
}
```

### API list

<details>
<summary>Account</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.AccountOverview() | YES | https://docs.kumex.com/#get-account-overview |
| ApiService.TransactionHistory() | YES | https://docs.kumex.com/#get-transaction-history |

</details>

<details>
<summary>Deposit</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.DepositAddresses() | YES | https://docs.kumex.com/#get-deposit-address |
| ApiService.Deposits() | YES | https://docs.kumex.com/#get-deposit-list |

</details>

<details>
<summary>Withdrawal</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.WithdrawalQuotas() | YES | https://docs.kumex.com/#get-withdrawal-quotas |
| ApiService.ApplyWithdrawal() | YES | https://docs.kumex.com/#apply-withdraw |
| ApiService.Withdrawals() | YES | https://docs.kumex.com/#get-withdrawals-list |
| ApiService.CancelWithdrawal() | YES | https://docs.kumex.com/#cancel-withdrawal |

</details>

<details>
<summary>Transfer</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.TransferOut() | YES | https://docs.kumex.com/#transfer-out |
| ApiService.TransferOutV2() | YES | https://docs.kumex.com/#transfer-funds-to-kucoin-main-account |
| ApiService.TransferList() | YES | https://docs.kumex.com/#get-transfer-list |
| ApiService.CancelTransfer() | YES | https://docs.kumex.com/#cancel-transfer |

</details>

<details>
<summary>Fill</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.Fills() | YES | https://docs.kumex.com/#list-fills |
| ApiService.RecentFills() | YES | https://docs.kumex.com/#recent-fills |
| ApiService.openOrderStatistics() | YES | https://docs.kumex.com/#open-order-statistics |

</details>

<details>
<summary>Order</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.CreateOrder() | YES | https://docs.kumex.com/#place-a-new-order |
| ApiService.CancelOrder() | YES | https://docs.kumex.com/#cancel-an-order |
| ApiService.CancelOrders() | YES | https://docs.kumex.com/#cancel-all-orders |
| ApiService.StopOrders() | YES | https://docs.kumex.com/#get-untriggered-stop-order-list |
| ApiService.Orders() | YES | https://docs.kumex.com/#list-orders |
| ApiService.Order() | YES | https://docs.kumex.com/#get-an-order |
| ApiService.RecentOrders() | YES | https://docs.kumex.com/#recent-orders |

</details>

<details>
<summary>Market</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.Ticker() | NO | https://docs.kumex.com/#get-real-time-ticker |
| ApiService.Level2Snapshot() | NO | https://docs.kumex.com/#get-full-order-book-level-2 |
| ApiService.Level2MessageQuery()() | NO | https://docs.kumex.com/#level-2-pulling-messages |
| ApiService.Level3Snapshot() | NO | https://docs.kumex.com/#get-full-order-book-level-3 |
| ApiService.Level3MessageQuery() | NO | https://docs.kumex.com/#level-3-pulling-messages|
| ApiService.TradeHistory() | NO | https://docs.kumex.com/#transaction-history |
| ApiService.InterestQuery() | NO | https://docs.kumex.com/#get-interest-rate-list |
| ApiService.IndexQuery() | NO | https://docs.kumex.com/#get-index-list |
| ApiService.MarkPrice() | NO | https://docs.kumex.com/#get-current-mark-price |
| ApiService.PremiumQuery() | NO | https://docs.kumex.com/#get-premium-index |
| ApiService.FundingRate() | NO | https://docs.kumex.com/#get-current-funding-rate |

</details>

<details>
<summary>Symbol</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.ActiveContracts() | NO | https://docs.kumex.com/#get-open-contract-list |
| ApiService.Contracts() | NO | https://docs.kumex.com/#get-order-info-of-the-contract |

</details>

<details>
<summary>WebSocket Feed</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.WebSocketPublicToken() | NO | https://docs.kumex.com/#apply-connect-token |
| ApiService.WebSocketPrivateToken() | YES | https://docs.kumex.com/#apply-connect-token |
| ApiService.NewWebSocketClient() | - | https://docs.kumex.com/#websocket-feed |

</details>

<details>
<summary>Time</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.ServerTime() | NO | https://docs.kumex.com/#server-time |

</details>

## Run tests

```shell
# Add your API configuration items into the environmental variable first
export API_BASE_URI=https://api.kumex.com
export API_KEY=key
export API_SECRET=secret
export API_PASSPHRASE=passphrase

# Run tests
go test -v
```

## License

[MIT](LICENSE)

