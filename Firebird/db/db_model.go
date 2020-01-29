package db

import (
	"encoding/json"
	"time"
)

type UserInfo struct {
	Id          int64     `gorose:"id"`
	GmtCreate   time.Time `gorose:"gmt_create"`
	GmtModified time.Time `gorose:"gmt_modified"`
	Username    string    `gorose:"user_name"`
	UserDesc    string    `gorose:"user_desc"`
	Status      int       `gorose:"status"`
}

type UserInfoQuery struct {
	Id         int64
	Username   string
	Status     int
	StartTime  time.Time
	EndTime    time.Time
	PageNumber int
	PageSize   int
}

type UserAccount struct {
	Id           int64     `gorose:"id"`
	GmtCreate    time.Time `gorose:"gmt_create"`
	GmtModified  time.Time `gorose:"gmt_modified"`
	UserId       int64     `gorose:"user_id"`
	SymbolId     int64     `gorose:"symbol_id"`
	HoldPrice    float64   `gorose:"hold_price"`
	HoldAmount   float64   `gorose:"hold_amount"`
	YestBenefit  float64   `gorose:"yest_benefit"`
	TotalBenefit float64   `gorose:"total_benefit"`
	Price        float64   `gorose:"price"`
	Amount       float64   `gorose:"amount"`
	Total        float64   `gorose:"total"`
	Benefit      float64   `gorose:"benefit"`
	Rate         float64   `gorose:"rate"`
	Status       int       `gorose:"status"`
}

type UserAccountVO struct {
	Id           int64     `json:"id"`
	GmtCreate    time.Time `json:"gmtCreate"`
	GmtModified  time.Time `json:"gmtModified"`
	UserId       int64     `json:"userId"`
	SymbolId     int64     `json:"symbolId"`
	HoldPrice    float64   `json:"holdPrice"`
	HoldAmount   float64   `json:"holdAmount"`
	YestBenefit  float64   `json:"yestBenefit"`
	TotalBenefit float64   `json:"totalBenefit"`
	Price        float64   `json:"price"`
	Amount       float64   `json:"amount"`
	Total        float64   `json:"total"`
	Benefit      float64   `json:"benefit"`
	Rate         float64   `json:"rate"`
	SymbolName   string    `json:"symbolName"`
	SymbolDesc   string    `json:"symbolDesc"`
	SymbolIcon   string    `json:"symbolIcon"`
	SymbolGroup  string    `json:"symbolGroup"`
	Status       int       `json:"status"`
}

func (uav UserAccountVO) MarshalJSON() ([]byte, error) {
	type UavAlias UserAccountVO
	return json.Marshal(&struct {
		UavAlias
		GmtCreate   string `json:"gmtCreate"`
		GmtModified string `json:"gmtModified"`
	}{
		UavAlias:    (UavAlias)(uav),
		GmtCreate:   uav.GmtCreate.Format("2006-01-02 15:04:05"),
		GmtModified: uav.GmtModified.Format("2006-01-02 15:04:05"),
	})
}

type UserAccountQuery struct {
	Id         int64
	Status     int
	UserId     int64
	SymbolId   int64
	StartTime  time.Time
	EndTime    time.Time
	PageNumber int
	PageSize   int
}

type UserTrade struct {
	Id          int64     `gorose:"id"`
	GmtCreate   time.Time `gorose:"gmt_create"`
	GmtModified time.Time `gorose:"gmt_modified"`
	UserId      int64     `gorose:"user_id"`
	SymbolId    int64     `gorose:"symbol_id"`
	Price       float64   `gorose:"price"`
	Amount      float64   `gorose:"amount"`
	Type        int       `gorose:"type"`
	ScheduleId  int64     `gorose:"scheduleId"`
	HoldPrice   float64   `gorose:"hold_price"`
	HoldAmount  float64   `gorose:"hold_amount"`
	Reason      string    `gorose:"reason"`
	Status      int       `gorose:"status"`
}

type UserTradeVO struct {
	Id          int64     `json:"id"`
	GmtCreate   time.Time `json:"gmtCreate"`
	GmtModified time.Time `json:"gmtModified"`
	UserId      int64     `json:"userId"`
	SymbolId    int64     `json:"symbolId"`
	Price       float64   `json:"price"`
	Amount      float64   `json:"amount"`
	Type        int       `json:"type"`
	ScheduleId  int64     `json:"scheduleId"`
	HoldPrice   float64   `json:"holdPrice"`
	HoldAmount  float64   `json:"holdAmount"`
	SymbolName  string    `json:"symbolName"`
	SymbolDesc  string    `json:"symbolDesc"`
	SymbolIcon  string    `json:"symbolIcon"`
	SymbolGroup string    `json:"symbolGroup"`
	Reason      string    `json:"reason"`
	Status      int       `json:"status"`
}

func (utv UserTradeVO) MarshalJSON() ([]byte, error) {
	type UtvAlias UserTradeVO
	return json.Marshal(&struct {
		UtvAlias
		GmtCreate   string `json:"gmtCreate"`
		GmtModified string `json:"gmtModified"`
	}{
		UtvAlias:    (UtvAlias)(utv),
		GmtCreate:   utv.GmtCreate.Format("2006-01-02 15:04:05"),
		GmtModified: utv.GmtModified.Format("2006-01-02 15:04:05"),
	})
}

type UserTradeQuery struct {
	Id         int64
	Status     int
	UserId     int64
	SymbolId   int64
	Type       int
	ScheduleId int64
	StartTime  time.Time
	EndTime    time.Time
	PageNumber int
	PageSize   int
}

type SymbolInfo struct {
	Id          int64     `gorose:"id"`
	GmtCreate   time.Time `gorose:"gmt_create"`
	GmtModified time.Time `gorose:"gmt_modified"`
	SymbolName  string    `gorose:"symbol_name"`
	SymbolDesc  string    `gorose:"symbol_desc"`
	SymbolIcon  string    `gorose:"symbol_icon"`
	SymbolGroup string    `gorose:"symbol_group"`
	Status      int       `gorose:"status"`
}

type SymbolInfoQuery struct {
	Id         int64
	Status     int
	SymbolName string
	StartTime  time.Time
	EndTime    time.Time
	PageNumber int
	PageSize   int
}

type RuleItem struct {
	Id          int64     `gorose:"id" json:"id"`
	GmtCreate   time.Time `gorose:"gmt_create" json:"gmtCreate"`
	GmtModified time.Time `gorose:"gmt_modified" json:"gmtModified"`
	UserId      int64     `gorose:"user_id" json:"userId"`
	SymbolId    int64     `gorose:"symbol_id" json:"symbolId"`
	ScheduleId  int64     `gorose:"schedule_id" json:"scheduleId"`
	RuleType    int       `gorose:"rule_type" json:"ruleType"`
	JoinType    int       `gorose:"join_type" json:"joinType"`
	OpType      int       `gorose:"op_type" json:"opType"`
	Value       string    `gorose:"value" json:"value"`
	Status      int       `gorose:"status" json:"status"`
}

func (ri RuleItem) MarshalJSON() ([]byte, error) {
	type riAlias RuleItem
	return json.Marshal(&struct {
		riAlias
		GmtCreate   string `json:"gmtCreate"`
		GmtModified string `json:"gmtModified"`
	}{
		riAlias:     (riAlias)(ri),
		GmtCreate:   ri.GmtCreate.Format("2006-01-02 15:04:05"),
		GmtModified: ri.GmtModified.Format("2006-01-02 15:04:05"),
	})
}

type RuleItemQuery struct {
	Id         int64
	UserId     int64
	SymbolId   int64
	ScheduleId int64
	PageNumber int
	PageSize   int
	Status     int
}

type UserSchedule struct {
	Id          int64     `gorose:"id"`
	GmtCreate   time.Time `gorose:"gmt_create"`
	GmtModified time.Time `gorose:"gmt_modified"`
	UserId      int64     `gorose:"user_id"`
	SymbolId    int64     `gorose:"symbol_id"`
	Name        string    `gorose:"name"`
	Type        int       `gorose:"type"`
	Amount      float64   `gorose:"amount"`
	Success     int       `gorose:"success"`
	Failed      int       `gorose:"failed"`
	Status      int       `gorose:"status"`
}

type UserScheduleVO struct {
	Id          int64      `json:"id"`
	GmtCreate   time.Time  `json:"gmtCreate"`
	GmtModified time.Time  `json:"gmtModified"`
	UserId      int64      `json:"userId"`
	SymbolId    int64      `json:"symbolId"`
	Name        string     `json:"name"`
	Type        int        `json:"type"`
	Amount      float64    `json:"amount"`
	Success     int        `json:"success"`
	Failed      int        `json:"failed"`
	SymbolName  string     `json:"symbolName"`
	SymbolDesc  string     `json:"symbolDesc"`
	SymbolIcon  string     `json:"symbolIcon"`
	SymbolGroup string     `json:"symbolGroup"`
	Status      int        `json:"status"`
	Rules       []RuleItem `json:"rules"`
}

func (usv UserScheduleVO) MarshalJSON() ([]byte, error) {
	type UsvAlias UserScheduleVO
	return json.Marshal(&struct {
		UsvAlias
		GmtCreate   string `json:"gmtCreate"`
		GmtModified string `json:"gmtModified"`
	}{
		UsvAlias:    (UsvAlias)(usv),
		GmtCreate:   usv.GmtCreate.Format("2006-01-02 15:04:05"),
		GmtModified: usv.GmtModified.Format("2006-01-02 15:04:05"),
	})
}

type UserScheduleQuery struct {
	Id         int64
	UserId     int64
	SymbolId   int64
	Type       int
	Status     int
	StartTime  time.Time
	EndTime    time.Time
	PageNumber int
	PageSize   int
}

type UserData struct {
	Id          int64     `gorose:"id" json:"id"`
	GmtCreate   time.Time `gorose:"gmt_create" json:"gmtCreate"`
	GmtModified time.Time `gorose:"gmt_modified" json:"gmtModified"`
	UserId      int64     `gorose:"user_id" json:"userId"`
	SymbolId    int64     `gorose:"symbol_id" json:"symbolId"`
	OpenPrice   float64   `gorose:"open_price" json:"openPrice"`
	ClosePrice  float64   `gorose:"close_price" json:"closePrice"`
	HighPrice   float64   `gorose:"high_price" json:"highPrice"`
	LowPrice    float64   `gorose:"low_price" json:"lowPrice"`
	HoldPrice   float64   `gorose:"hold_price" json:"holdPrice"`
	HoldAmount  float64   `gorose:"hold_amount" json:"holdAmount"`
	HoldRate    float64   `gorose:"hold_rate" json:"holdRate"`
	HoldBenefit float64   `gorose:"hold_benefit" json:"holdBenefit"`
	Status      int       `gorose:"status" json:"status"`
}

func (ud UserData) MarshalJSON() ([]byte, error) {
	type UdAlias UserData
	return json.Marshal(&struct {
		UdAlias
		GmtCreate   string `json:"gmtCreate"`
		GmtModified string `json:"gmtModified"`
	}{
		UdAlias:     (UdAlias)(ud),
		GmtCreate:   ud.GmtCreate.Format("2006-01-02 15:04:05"),
		GmtModified: ud.GmtModified.Format("2006-01-02 15:04:05"),
	})
}

type UserDataQuery struct {
	Id         int64
	UserId     int64
	SymbolId   int64
	Status     int
	StartTime  time.Time
	EndTime    time.Time
	PageNumber int
	PageSize   int
}
