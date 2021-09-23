package buda

import (
	_"fmt"
	"net/http"
	_"time"
	_"crypto/sha512"
	_"strings"
	_"crypto/hmac"
	_"encoding/base64"
	_"encoding/hex"
	_"encoding/json"
	_"strconv"
	_"io/ioutil"

	. "github.com/nntaoli-project/goex"

)

const (
	BaseURL = "https://www.buda.com/api/v2"
	MarketsEndpoint = "/markets"
	MarketEndpoint = "/markets/%d"
	MarketVolumeEndpoint = "/markets/%s/volume"
	MarketTickerEndpoint = "/markets/%s/ticker"
	MarketOrderBookEndpoint = "/markets/%s/order_book"
	MarketTradesEndpoint = "/markets/%s/trades"
	BalancesEndpoint = "/balances"
	OrdersEndpoint = "/markets/%s/orders"
	OrderEndpoint = "/orders/%d"
	WithdrawalsEndpoint = "/currencies/%s/withdrawals"
	DepositsEndpoint = "/currencies/%s/deposits"
	DepositFeeEndpoint = "/currencies/%s/fees/deposit"
	WithdrawalFeeEndpoint = "/currencies/%s/fees/withdrawal"
	ReceiveAddressEndpoint = "/currencies/%s/receive_addresses/%s"
	ElementsPerPage = "300"
)

type Buda struct {
	Key string
	Secret string
	Client *http.Client
}

type Market struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	BaseCurrency       string   `json:"base_currency"`
	QuoteCurrency      string   `json:"quote_currency"`
	MinimumOrderAmount []string `json:"minimum_order_amount"`
}

type Markets struct {
	Markets []Market `json:"markets"`
}

type MarketSingle struct {
	Market Market `json:"market"`
}

type Fee struct {
	Name string `json:"name"`
	Percent float64 `json:"percent"`
	Base []string `json:"base"`
}

type FeeSingle struct {
	Fee Fee `json:"fee"`
}

type Volume struct {
	AskVolume24H []string `json:"ask_volume_24h"`
	AskVolume7D  []string `json:"ask_volume_7d"`
	BidVolume24H []string `json:"bid_volume_24h"`
	BidVolume7D  []string `json:"bid_volume_7d"`
	MarketID     string   `json:"market_id"`
}

type VolumeSingle struct {
 	Volume Volume `json:"volume"`
}

type Ticker struct {
	LastPrice         []string `json:"last_price"`
	MaxBid            []string `json:"max_bid"`
	MinAsk            []string `json:"min_ask"`
	PriceVariation24H string `json:"price_variation_24h"`
	PriceVariation7D  string `json:"price_variation_7d"`
	Volume            []string `json:"volume"`
}

type TickerSingle struct {
	Ticker Ticker `json:"ticker"`
}

type OrderBook struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

type OrderBookSingle struct {
	OrderBook OrderBook `json:"order_book"`
}

type Trade struct {
	MarketId 	  string	 `json:"market_id"`
	Timestamp     string	 `json:"timestamp"`
	LastTimestamp string     `json:"last_timestamp"`
	Entries       [][]interface{} `json:"entries"`
}

type Trades struct {
	Trade Trade `json:"trades"`
}

type Balance struct {
	ID                    string   `json:"id"`
	Amount                []string `json:"amount"`
	AvailableAmount       []string `json:"available_amount"`
	FrozenAmount          []string `json:"frozen_amount"`
	PendingWithdrawAmount []string `json:"pending_withdraw_amount"`
	AccountID             int      `json:"account_id"`
}

type Balances struct {
	Balances []Balance `json:"balances"`
}

type BalanceSingle struct {
	Balance Balance `json:"balance"`
}

type Metadata struct {
	CurrentPage int `json:"current_page"`
	TotalCount  int `json:"total_count"`
	TotalPages  int `json:"total_pages"`
}
/*
type Order struct {
	ID             int       `json:"id"`
	Type           string    `json:"type"`
	State          string    `json:"state"`
	CreatedAt      time.Time `json:"created_at"`
	MarketID       string    `json:"market_id"`
	AccountID      int       `json:"account_id"`
	FeeCurrency    string    `json:"fee_currency"`
	PriceType      string    `json:"price_type"`
	Limit          []string  `json:"limit"`
	Amount         []string  `json:"amount"`
	OriginalAmount []string  `json:"original_amount"`
	TradedAmount   []string  `json:"traded_amount"`
	TotalExchanged []string  `json:"total_exchanged"`
	PaidFee        []string  `json:"paid_fee"`
}

type OrderSingle struct {
	Order Order `json:"order"`
}

type Orders struct {
	Orders []Order `json:"orders"`
	Meta Metadata `json:"meta"`
}
*/
type DepositData struct {
	Type    string `json:"type"`
	Address string `json:"address"`
	TxHash  string `json:"tx_hash"`
}

type Deposit struct {
	ID          int      `json:"id"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Amount      []string `json:"amount"`
	Currency    string   `json:"currency"`
	State       string   `json:"state"`
	DepositData DepositData `json:"deposit_data"`
}

type Deposits struct {
	Deposits []Deposit `json:"deposits"`
	Meta Metadata `json:"meta"`
}

type WithdrawalData struct {
	Type          string `json:"type"`
	TargetAddress string `json:"target_address"`
	TxHash        string `json:"tx_hash"`
}

type Withdrawal struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	State          string   `json:"state"`
	Amount         []string `json:"amount"`
	Fee            []string `json:"fee"`
	Currency       string   `json:"currency"`
	WithdrawalData WithdrawalData `json:"withdrawal_data"`
}

type WithdrawalSingle struct {
	Withdrawal Withdrawal `json:"withdrawal"`
}

type Withdrawals struct {
	Withdrawals []Withdrawal `json:"withdrawals"`
	Meta Metadata `json:"meta"`
}

type ReceiveAddress struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Address   string `json:"address"`
	Used      bool   `json:"used"`
}

type ReceiveAddressSingle struct {
	ReceiveAddress ReceiveAddress `json:"receive_address"`
}

////////////////////////////////////////////////////////////////////////////////
//     METODOS   NO   IMPLEMENTADOS  de la interface API de goex
////////////////////////////////////////////////////////////////////////////////

func (client *Buda) LimitBuy(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error) {
	panic("not implement")
}
func (client *Buda) LimitSell(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error) {
	panic("not implement")
}
func (client *Buda) MarketBuy(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("not implement")
}
func (client *Buda) MarketSell(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("not implement")
}
func (client *Buda) CancelOrder(orderId string, currency CurrencyPair) (bool, error) {
	panic("not implement")
}
func (client *Buda) GetOneOrder(orderId string, currency CurrencyPair) (*Order, error) {
	panic("not implement")
}
func (client *Buda) GetUnfinishOrders(currency CurrencyPair) ([]Order, error) {
	panic("not implement")
}
func (client *Buda) GetOrderHistorys(currency CurrencyPair, optional ...OptionalParameter) ([]Order, error) {
	panic("not implement")
}
func (client *Buda) GetAccount() (*Account, error) {
	panic("not implement")
}
func (client *Buda) GetKlineRecords(currency CurrencyPair, period KlinePeriod, size int, opt ...OptionalParameter) ([]Kline, error) {
	panic("not implement")
}

func (client *Buda) GetTrades(currencyPair CurrencyPair, since int64) ([]Trade, error) {
	panic("not implement")
}
func (client *Buda) GetTicker(currency CurrencyPair) (*Ticker, error) {
	panic("not implement")
}

func (client *Buda) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	panic("not implement")
}


////////////////////////////////////////////////////////////////////////////////
//         ADAPTERS: la idea es usar las funciones y modificar las respuestas para la interface
////////////////////////////////////////////////////////////////////////////////

func NewAPI(client *http.Client, accessKey string, secretKey string) *Buda {
	return &Buda{accessKey, secretKey, client}
}

func (client *Buda) GetExchangeName() string {
	return BUDA
}
/*
func (client *Buda) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	client.GetOrderBookByMarket()
//func (client *Buda) GetOrderBookByMarket(marketId string) (*OrderBook, error) {

	return &depth, nil
}
*/
////////////////////////////////////////////////////////////////////////////////
//         Lo que sigue son las FUNCIONES A RE-IMPLEMENTAR
////////////////////////////////////////////////////////////////////////////////
/*




func (client *Buda) SignRequest(params...string) (string) {
	h := hmac.New(sha512.New384, []byte(client.Secret))
	h.Write([]byte(strings.Join(params, " ")))
	return hex.EncodeToString(h.Sum(nil))
}

func (client *Buda) AuthenticatedRequest(request *http.Request) (*http.Request, error) {
	var signature string
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano()*1E6, 10)

	switch request.Method {
		case "POST": {
			var body []byte
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				return nil, err
			}
			signature = client.SignRequest(request.Method, request.URL.RequestURI(), base64.StdEncoding.EncodeToString(body), timestamp)
		}
		case "GET": {
			signature = client.SignRequest(request.Method, request.URL.RequestURI(), timestamp)
		}
	}

	request.Header.Set("X-SBTC-APIKEY", client.Key)
	request.Header.Set("X-SBTC-NONCE", timestamp)
	request.Header.Set("X-SBTC-SIGNATURE", signature)

	return request, nil
}

//deprecada
func NewAPIClient(apiKey string, apiSecret string) (*Buda, error){
 	return &Buda{Client: &http.Client{}, Key: apiKey, Secret: apiSecret}, nil
}

func (client *Buda) FormatResource(resource string) (string) {
	return fmt.Sprintf("%s%s", BaseURL, resource)
}

func (client *Buda) Get(resource string, private bool) ([]byte, error) {
	req, err := http.NewRequest("GET", client.FormatResource(resource), nil)
	if err != nil {
		return nil, err
	}

	if private {
		req, err = client.AuthenticatedRequest(req)
		if err != nil {
			return nil, err
		}
	}

	response, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (client *Buda) GetMarkets() ([]Market, error) {
	var markets Markets

	response, err := client.Get(MarketsEndpoint, false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &markets)
	if err != nil {
		return nil, err
	}

	return markets.Markets, nil
}

func (client *Buda) GetMarket(id int) (*Market, error) {
	var market MarketSingle

	data, err := client.Get(fmt.Sprintf(MarketEndpoint, id), false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &market)
	if err != nil {
		return nil, err
	}

	return &market.Market, nil
}

func (client *Buda) GetVolumeByMarket(marketId string) (*Volume, error) {
	var volume VolumeSingle

	data, err := client.Get(fmt.Sprintf(MarketVolumeEndpoint, marketId),false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &volume)
	if err != nil {
		return nil, err
	}

	return &volume.Volume, nil
}

func (client *Buda) GetTickerByMarket(marketId string) (*Ticker, error) {
	var ticker TickerSingle

	data, err := client.Get(fmt.Sprintf(MarketTickerEndpoint, marketId),false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &ticker)
	if err != nil {
		return nil, err
	}

	return &ticker.Ticker, nil
}

func (client *Buda) GetOrderBookByMarket(marketId string) (*OrderBook, error) {
	var orderBook OrderBookSingle

	data, err := client.Get(fmt.Sprintf(MarketOrderBookEndpoint, marketId),false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &orderBook)
	if err != nil {
		return nil, err
	}

	return &orderBook.OrderBook, nil
}

func (client *Buda) GetTradesByMarket(marketId string, timestamp string) (*Trade, error) {
	var trades Trades
	var url string

	if timestamp != "" {
		url = fmt.Sprintf(MarketTradesEndpoint, marketId) + "?timestamp=" + timestamp
	} else {
		url = fmt.Sprintf(MarketTradesEndpoint, marketId)
	}

	data, err := client.Get(url,false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &trades)
	if err != nil {
		return nil, err
	}

	return &trades.Trade, nil
}

func (client *Buda) GetBalances() ([]Balance, error) {
	var balances Balances

	data, err := client.Get(BalancesEndpoint,true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &balances)
	if err != nil {
		return nil, err
	}

	return balances.Balances, nil
}

func (client *Buda) GetBalanceByCurrency(currency string) (*Balance, error) {
	var balance BalanceSingle

	data, err := client.Get(BalancesEndpoint + "/" + currency, true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &balance)
	if err != nil {
		return nil, err
	}

	return &balance.Balance, nil
}

func (client *Buda) GetOrderById(id int) (*Order, error) {
	var order OrderSingle

	data, err := client.Get(fmt.Sprintf(OrderEndpoint, id), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &order)
	if err != nil {
		return nil, err
	}

	return &order.Order, nil
}

func (client *Buda) GetOrdersByMarket(marketId string) ([]Order, error) {
	var orders Orders
	var ret []Order

	data, err := client.Get(fmt.Sprintf(OrdersEndpoint + "?page=1&per=" + ElementsPerPage, marketId), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &orders)
	if err != nil {
		return nil, err
	}

	resc, errc := make(chan []Order), make(chan error)

	ret = append(ret, orders.Orders...)

	if orders.Meta.TotalPages > 1 {
		for i := orders.Meta.CurrentPage + 1; i <= orders.Meta.TotalPages; i++ {
			go func(i int) {
				data, err := client.Get(fmt.Sprintf(OrdersEndpoint + fmt.Sprintf("?page=%d", i) + "&per=" + ElementsPerPage, marketId), true)
				if err != nil {
					errc <- err
					return
				}
				err = json.Unmarshal(data, &orders)
				if err != nil {
					errc <- err
					return
				}
				resc <- orders.Orders
			}(i)
		}

		for i := orders.Meta.CurrentPage + 1; i <= orders.Meta.TotalPages; i++ {
			select {
			case res := <-resc:
				{
					ret = append(ret, res...)
				}
			case err := <-errc:
				{
					return nil, err
				}
			}
		}
	}

	return ret, nil
}

func (client *Buda) GetOrdersByMarketAndState(marketId string, state string) ([]Order, error) {
	var orders Orders
	var ret []Order

	data, err := client.Get(fmt.Sprintf(OrdersEndpoint + "?page=1&per=" + ElementsPerPage + "&state=" + state, marketId), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &orders)
	if err != nil {
		return nil, err
	}

	resc, errc := make(chan []Order), make(chan error)

	ret = append(ret, orders.Orders...)

	if orders.Meta.TotalPages > 1 {
		for i := orders.Meta.CurrentPage + 1; i <= orders.Meta.TotalPages; i++ {
			go func(i int) {
				data, err := client.Get(fmt.Sprintf(OrdersEndpoint + fmt.Sprintf("?page=%d", i) + "&per=" + ElementsPerPage + "&state=" + state, marketId), true)
				if err != nil {
					errc <- err
					return
				}
				err = json.Unmarshal(data, &orders)
				if err != nil {
					errc <- err
					return
				}
				resc <- orders.Orders
			}(i)
		}

		for i := orders.Meta.CurrentPage + 1; i <= orders.Meta.TotalPages; i++ {
			select {
			case res := <-resc:
				{
					ret = append(ret, res...)
				}
			case err := <-errc:
				{
					return nil, err
				}
			}
		}
	}

	return ret, nil
}


func (client *Buda) GetWithdrawalsByCurrency(currency string) ([]Withdrawal, error) {
	var withdrawals Withdrawals
	var ret []Withdrawal

	data, err := client.Get(fmt.Sprintf(WithdrawalsEndpoint + "?page=1&per=" + ElementsPerPage, currency), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &withdrawals)
	if err != nil {
		return nil, err
	}

	resc, errc := make(chan []Withdrawal), make(chan error)
	ret = append(ret, withdrawals.Withdrawals...)

	if withdrawals.Meta.TotalPages > 1 {

		for i := withdrawals.Meta.CurrentPage + 1; i <= withdrawals.Meta.TotalPages; i++ {
			go func(i int) {
				data, err := client.Get(fmt.Sprintf(WithdrawalsEndpoint + fmt.Sprintf("?page=%d", i) + "&per=" + ElementsPerPage, currency), true)
				if err != nil {
					errc <- err
					return
				}
				err = json.Unmarshal(data, &withdrawals)
				if err != nil {
					errc <- err
					return
				}
				resc <- withdrawals.Withdrawals
			}(i)
		}

		for i := withdrawals.Meta.CurrentPage + 1; i <= withdrawals.Meta.TotalPages; i++ {
			select {
			case res := <-resc:
				{
					ret = append(ret, res...)
				}
			case err := <-errc:
				{
					return nil, err
				}
			}
		}
	}
	return ret, nil
}

func (client *Buda) GetDepositsByCurrency(currency string) ([]Deposit, error) {

	var deposits Deposits
	var ret []Deposit

	data, err := client.Get(fmt.Sprintf(DepositsEndpoint + "?page=1&per=" + ElementsPerPage, currency), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &deposits)
	if err != nil {
		return nil, err
	}

	resc, errc := make(chan []Deposit), make(chan error)
	ret = append(ret, deposits.Deposits...)

	if deposits.Meta.TotalPages > 1 {

		for i := deposits.Meta.CurrentPage + 1; i <= deposits.Meta.TotalPages; i++ {
			go func(i int) {
				data, err := client.Get(fmt.Sprintf(DepositsEndpoint + fmt.Sprintf("?page=%d", i) + "&per=" + ElementsPerPage, currency), true)
				if err != nil {
					errc <- err
					return
				}
				err = json.Unmarshal(data, &deposits)
				if err != nil {
					errc <- err
					return
				}
				resc <- deposits.Deposits
			}(i)
		}

		for i := deposits.Meta.CurrentPage + 1; i <= deposits.Meta.TotalPages; i++ {
			select {
			case res := <-resc:
				{
					ret = append(ret, res...)
				}
			case err := <-errc:
				{
					return nil, err
				}
			}
		}
	}
	return ret, nil
}

func (client *Buda) GetDepositsByCurrencyAndState(currency string, state string) ([]Deposit, error) {

	var deposits Deposits
	var ret []Deposit

	data, err := client.Get(fmt.Sprintf(DepositsEndpoint + "?page=1&per=" + ElementsPerPage + "&state=" + state, currency), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &deposits)
	if err != nil {
		return nil, err
	}

	resc, errc := make(chan []Deposit), make(chan error)
	ret = append(ret, deposits.Deposits...)

	if deposits.Meta.TotalPages > 1 {

		for i := deposits.Meta.CurrentPage + 1; i <= deposits.Meta.TotalPages; i++ {
			go func(i int) {
				data, err := client.Get(fmt.Sprintf(DepositsEndpoint + fmt.Sprintf("?page=%d", i) + "&per=" + ElementsPerPage + "&state=" + state, currency), true)
				if err != nil {
					errc <- err
					return
				}
				err = json.Unmarshal(data, &deposits)
				if err != nil {
					errc <- err
					return
				}
				resc <- deposits.Deposits
			}(i)
		}

		for i := deposits.Meta.CurrentPage + 1; i <= deposits.Meta.TotalPages; i++ {
			select {
			case res := <-resc:
				{
					ret = append(ret, res...)
				}
			case err := <-errc:
				{
					return nil, err
				}
			}
		}
	}
	return ret, nil
}

func (client *Buda) GetWithdrawalsByCurrencyAndState(currency string, state string) ([]Withdrawal, error) {

	var withdrawals Withdrawals
	var ret []Withdrawal

	data, err := client.Get(fmt.Sprintf(WithdrawalsEndpoint + "?page=1&per=" + ElementsPerPage + "&state=" + state, currency), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &withdrawals)
	if err != nil {
		return nil, err
	}

	resc, errc := make(chan []Withdrawal), make(chan error)
	ret = append(ret, withdrawals.Withdrawals...)

	if withdrawals.Meta.TotalPages > 1 {

		for i := withdrawals.Meta.CurrentPage + 1; i <= withdrawals.Meta.TotalPages; i++ {
			go func(i int) {
				data, err := client.Get(fmt.Sprintf(WithdrawalsEndpoint + fmt.Sprintf("?page=%d", i) + "&per=" + ElementsPerPage + "&state=" + state, currency), true)
				if err != nil {
					errc <- err
					return
				}
				err = json.Unmarshal(data, &withdrawals)
				if err != nil {
					errc <- err
					return
				}
				resc <- withdrawals.Withdrawals
			}(i)
		}

		for i := withdrawals.Meta.CurrentPage + 1; i <= withdrawals.Meta.TotalPages; i++ {
			select {
			case res := <-resc:
				{
					ret = append(ret, res...)
				}
			case err := <-errc:
				{
					return nil, err
				}
			}
		}
	}
	return ret, nil
}

func (client *Buda) GetWithdrawalFeeByCurrency(currency string) (*Fee, error) {
	var fee FeeSingle

	data, err := client.Get(fmt.Sprintf(WithdrawalFeeEndpoint, currency), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &fee)
	if err != nil {
		return nil, err
	}

	return &fee.Fee, nil
}

func (client *Buda) GetDepositFeeByCurrency(currency string) (*Fee, error) {
	var fee FeeSingle

	data, err := client.Get(fmt.Sprintf(DepositFeeEndpoint, currency), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &fee)
	if err != nil {
		return nil, err
	}

	return &fee.Fee, nil
}

func (client *Buda) GetReceiveAddresses(id int, currency string) (*ReceiveAddress, error) {
	var receiveAddress ReceiveAddressSingle

	data, err := client.Get(fmt.Sprintf(ReceiveAddressEndpoint, currency, string(id)), true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &receiveAddress)
	if err != nil {
		return nil, err
	}

	return &receiveAddress.ReceiveAddress, nil

}
*/
