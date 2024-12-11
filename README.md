# Notice: SDK Deprecation

Thank you for your support and usage of this SDK. We want to inform you that **this project is no longer actively maintained or updated**. 

To ensure you have access to the latest features, improvements, and support, we recommend transitioning to our new SDK: [**KuCoin Universal SDK**](https://github.com/Kucoin/kucoin-universal-sdk).

The KuCoin Universal SDK offers:
- A unified architecture across multiple programming languages.
- Enhanced performance and stability.
- Continued support and updates.

👉 **New SDK Repository**: [https://github.com/Kucoin/kucoin-universal-sdk](https://github.com/Kucoin/kucoin-universal-sdk)

We appreciate your understanding and encourage you to migrate to the new SDK for a better development experience. Should you have any questions or require assistance, feel free to reach out to us.

# Go SDK for KuMex API
> The detailed document [https://docs.kucoin.com/futures](https://docs.kucoin.com/futures), in order to receive the latest API change notifications, please `Watch` this repository.

[![Latest Version](https://img.shields.io/github/release/Kucoin/kucoin-futures-go-sdk.svg)](https://github.com/Kucoin/kucoin-futures-go-sdk/releases)
[![GoDoc](https://godoc.org/github.com/Kucoin/kucoin-futures-go-sdk?status.svg)](https://godoc.org/github.com/Kucoin/kucoin-futures-go-sdk)
[![Build Status](https://travis-ci.org/Kucoin/kucoin-futures-go-sdk.svg?branch=master)](https://travis-ci.org/Kucoin/kucoin-futures-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kucoin/kucoin-futures-go-sdk)](https://goreportcard.com/report/github.com/Kucoin/kucoin-futures-go-sdk)
[![Sourcegraph](https://sourcegraph.com/github.com/Kucoin/kucoin-futures-go-sdk/-/badge.svg)](https://sourcegraph.com/github.com/Kucoin/kucoin-futures-go-sdk?badge)
<!-- [![Total Lines](https://tokei.rs/b1/github/Kucoin/kucoin-futures-go-sdk)](https://github.com/Kucoin/kucoin-futures-go-sdk) -->


## Install

```bash
go get github.com/Kucoin/kucoin-futures-go-sdk
```

## Usage

### Choose environment

| Environment | BaseUri |
| -------- | -------- |
| *Production* | `https://api-futures.kucoin.com(DEFAULT)` `https://api-futures.kucoin.cc` |
| *Sandbox* | `https://api-sandbox-futures.kucoin.com` |

### Create ApiService

###### **Note** 
To reinforce the security of the API, KuCoin upgraded the API key to version 2.0, the validation logic has also been changed. It is recommended to create(https://www.kucoin.com/account/api) and update your API key to version 2.0. 
The API key of version 1.0 will be still valid until May 1, 2021.

```go
// API key version 2.0
s :=  kucoin.NewApiService( 
	// kucoin.ApiBaseURIOption("https://api.kucoin.com"), 
	kucoin.ApiKeyOption("key"),
	kucoin.ApiSecretOption("secret"),
	kucoin.ApiPassPhraseOption("passphrase"),
	kucoin.ApiKeyVersionOption(ApiKeyVersionV2)
)

// API key version 1.0
s := kucoin.NewApiService( 
	// kucoin.ApiBaseURIOption("https://api.kucoin.com"), 
	kucoin.ApiKeyOption("key"),
	kucoin.ApiSecretOption("secret"),
	kucoin.ApiPassPhraseOption("passphrase"), 
)

// Or add these options into the environmental variable
// Bash: 
// export API_BASE_URI=https://api-futures.kucoin.com
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

ch1 := kumex.NewSubscribeMessage("/contractMarket/ticker:XBTUSDTM", false)
ch2 := kumex.NewSubscribeMessage("/contractMarket/ticker:XBTUSDTM", false)
uch := kumex.NewUnsubscribeMessage("/contractMarket/ticker:XBTUSDTM", false)

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
            log.Println("Unsubscribe XBTUSDTM")
            if err = c.Unsubscribe(uch); err != nil {
                log.Printf("Error: %s", err.Error())
                // Handle error
                return
            }
        }
        if i == 10 {
            log.Println("Subscribe XBTUSDTM")
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
| ApiService.AccountOverview() | YES | https://docs.kucoin.com/futures/#get-account-overview |
| ApiService.TransactionHistory() | YES | https://docs.kucoin.com/futures/#get-transaction-history |

</details>

<details>
<summary>Deposit</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.DepositAddresses() | YES | https://docs.kucoin.com/futures/#get-deposit-address |
| ApiService.Deposits() | YES | https://docs.kucoin.com/futures/#get-deposit-list |

</details>

<details>
<summary>Withdrawal</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.WithdrawalQuotas() | YES | https://docs.kucoin.com/futures/#get-withdrawal-quotas |
| ApiService.ApplyWithdrawal() | YES | https://docs.kucoin.com/futures/#apply-withdraw |
| ApiService.Withdrawals() | YES | https://docs.kucoin.com/futures/#get-withdrawals-list |
| ApiService.CancelWithdrawal() | YES | https://docs.kucoin.com/futures/#cancel-withdrawal |

</details>

<details>
<summary>Transfer</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.TransferOut() | YES | https://docs.kucoin.com/futures/#transfer-out |
| ApiService.TransferOutV2() | YES | https://docs.kucoin.com/futures/#transfer-funds-to-kucoin-main-account |
| ApiService.TransferList() | YES | https://docs.kucoin.com/futures/#get-transfer-list |
| ApiService.CancelTransfer() | YES | https://docs.kucoin.com/futures/#cancel-transfer |

</details>

<details>
<summary>Fill</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.Fills() | YES | https://docs.kucoin.com/futures/#list-fills |
| ApiService.RecentFills() | YES | https://docs.kucoin.com/futures/#recent-fills |
| ApiService.openOrderStatistics() | YES | https://docs.kucoin.com/futures/#open-order-statistics |

</details>

<details>
<summary>Order</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.CreateOrder() | YES | https://docs.kucoin.com/futures/#place-a-new-order |
| ApiService.CancelOrder() | YES | https://docs.kucoin.com/futures/#cancel-an-order |
| ApiService.CancelOrders() | YES | https://docs.kucoin.com/futures/#cancel-all-orders |
| ApiService.StopOrders() | YES | https://docs.kucoin.com/futures/#get-untriggered-stop-order-list |
| ApiService.Orders() | YES | https://docs.kucoin.com/futures/#list-orders |
| ApiService.Order() | YES | https://docs.kucoin.com/futures/#get-an-order |
| ApiService.RecentOrders() | YES | https://docs.kucoin.com/futures/#recent-orders |

</details>

<details>
<summary>Market</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.Ticker() | NO | https://docs.kucoin.com/futures/#get-real-time-ticker |
| ApiService.Level2Snapshot() | NO | https://docs.kucoin.com/futures/#get-full-order-book-level-2 |
| ApiService.Level2MessageQuery()() | NO | https://docs.kucoin.com/futures/#level-2-pulling-messages |
| ApiService.Level3Snapshot() | NO | https://docs.kucoin.com/futures/#get-full-order-book-level-3 |
| ApiService.Level3MessageQuery() | NO | https://docs.kucoin.com/futures/#level-3-pulling-messages|
| ApiService.TradeHistory() | NO | https://docs.kucoin.com/futures/#transaction-history |
| ApiService.InterestQuery() | NO | https://docs.kucoin.com/futures/#get-interest-rate-list |
| ApiService.IndexQuery() | NO | https://docs.kucoin.com/futures/#get-index-list |
| ApiService.MarkPrice() | NO | https://docs.kucoin.com/futures/#get-current-mark-price |
| ApiService.PremiumQuery() | NO | https://docs.kucoin.com/futures/#get-premium-index |
| ApiService.FundingRate() | NO | https://docs.kucoin.com/futures/#get-current-funding-rate |

</details>

<details>
<summary>Symbol</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.ActiveContracts() | NO | https://docs.kucoin.com/futures/#get-open-contract-list |
| ApiService.Contracts() | NO | https://docs.kucoin.com/futures/#get-order-info-of-the-contract |

</details>

<details>
<summary>WebSocket Feed</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.WebSocketPublicToken() | NO | https://docs.kucoin.com/futures/#apply-connect-token |
| ApiService.WebSocketPrivateToken() | YES | https://docs.kucoin.com/futures/#apply-connect-token |
| ApiService.NewWebSocketClient() | - | https://docs.kucoin.com/futures/#websocket-feed |

</details>

<details>
<summary>Time</summary>

| API | Authentication | Description |
| -------- | -------- | -------- |
| ApiService.ServerTime() | NO | https://docs.kucoin.com/futures/#server-time |

</details>

## Run tests

```shell
# Add your API configuration items into the environmental variable first
export API_BASE_URI=https://api-futures.kucoin.com
export API_KEY=key
export API_SECRET=secret
export API_PASSPHRASE=passphrase
export API_KEY_VERSION=2

# Run tests
go test -v
```

## License

[MIT](LICENSE)

