package main

import (
	"log"

	"github.com/Kucoin/kumex-go-sdk"
)

func main() {
	//s := kumex.NewApiServiceFromEnv()
	s := kumex.NewApiService(
		kumex.ApiKeyOption(""),
		kumex.ApiSecretOption(""),
		kumex.ApiPassPhraseOption(""),
	)
	serverTime(s)
	accounts(s)
	orders(s)
	publicWebsocket(s)
	privateWebsocket(s)

}

func serverTime(s *kumex.ApiService) {
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
}

func accounts(s *kumex.ApiService) {
	p := map[string]string{}
	p["currency"] = "XBT"
	rsp, err := s.AccountOverview(p)
	if err != nil {
		// Handle error
		return
	}

	as := kumex.AccountModel{}
	if err := rsp.ReadData(&as); err != nil {
		log.Println(err)
		// Handle error
		return
	}
	log.Printf("Available balance: %f %f => %f", as.AccountEquity, as.OrderMargin, as.AvailableBalance)
}

func orders(s *kumex.ApiService) {
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
}
func publicWebsocket(s *kumex.ApiService) {
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

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	ch1 := kumex.NewSubscribeMessage("/contract/instrument:XBTUSDM", false)
	ch2 := kumex.NewSubscribeMessage("/contractMarket/level3:XBTUSDM", false)
	//uch := kumex.NewUnsubscribeMessage("/contractMarket/ticker:XBTUSDM", false)

	if err := c.Subscribe(ch1, ch2); err != nil {
		// Handle error
		return
	}

	for {
		select {
		case err := <-ec:
			c.Stop() // Stop subscribing the WebSocket feed
			log.Printf("Error: %s", err.Error())
			// Handle error
			return
		case msg := <-mc:
			log.Printf("Received: %s", kumex.ToJsonString(msg))
		}
	}
}

func privateWebsocket(s *kumex.ApiService) {
	rsp, err := s.WebSocketPrivateToken()
	if err != nil {
		// Handle error
		return
	}

	tk := &kumex.WebSocketTokenModel{}
	tk.AcceptUserMessage = true
	if err := rsp.ReadData(tk); err != nil {
		// Handle error
		return
	}
	c := s.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	ch1 := kumex.NewSubscribeMessage("/contract/position:XBTUSDM", false)
	ch2 := kumex.NewSubscribeMessage("/contractAccount/wallet", false)

	log.Println(kumex.ToJsonString(ch1))
	log.Println(kumex.ToJsonString(ch2))

	if err := c.Subscribe(ch1, ch2); err != nil {
		// Handle error
		return
	}

	for {
		select {
		case err := <-ec:
			c.Stop() // Stop subscribing the WebSocket feed
			log.Printf("Error: %s", err.Error())
			// Handle error
			return
		case msg := <-mc:
			log.Printf("Received: %s", kumex.ToJsonString(msg))
		}
	}
}
