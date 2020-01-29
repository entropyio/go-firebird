package data

// depth response data
type DepthData struct {
	Ch   string        `json:"ch"`
	Ts   int64         `json:"ts"`
	Tick DepthTickData `json:"tick"`
}

type DepthTickData struct {
	Bids    [][]float64 `json:"bids"`
	Asks    [][]float64 `json:"asks"`
	Version int64       `json:"version"`
	Ts      int64       `json:"ts"`
}

// kline response data
type KlineData struct {
	Ch   string        `json:"ch"`
	Ts   int64         `json:"ts"`
	Tick KlineTickData `json:"tick"`
}

type KlineTickData struct {
	Id     int64   `json:"id"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Amount float64 `json:"amount"`
	Vol    float64 `json:"vol"`
	Count  int64   `json:"count"`
}

// history kline response data
type HistoryKlineData struct {
	Ch     string          `json:"ch"`
	Ts     int64           `json:"ts"`
	Data   []KlineTickData `json:"data"`
	Status string          `json:"status"`
}

type LoginUser struct {
	UserId int64  `json:"userId"`
	Ts     int64  `json:"ts"`
	Token  string `json:"token"`
}
