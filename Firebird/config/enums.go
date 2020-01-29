package config

const (
	STATUS_INVALID = 0
	STATUS_ENABLE  = 1
	STATUS_DISABLE = 2
	STATUS_DELETE  = 3
)

const (
	TRADE_BUY  = 1
	TRADE_SOLD = 2
)

const (
	RULE_TYPE_RATE  = 1
	RULE_TYPE_PRICE = 2
)

const (
	JOIN_TYPE_AND = 1
	JOIN_TYPE_OR  = 2
)

const (
	OP_TYPE_EQUAL       = 1
	OP_TYPE_NOT_EQUAL   = 2
	OP_TYPE_BIG         = 3
	OP_TYPE_BIG_EQUAL   = 4
	OP_TYPE_LESS        = 5
	OP_TYPE_LESS_EQUAL  = 6
	OP_TYPE_CONTAIN     = 7
	OP_TYPE_NOT_CONTAIN = 8
)
