package web

const (
	CODE_SUCCESS = 0
	CODE_FAILED  = -1

	CODE_NEED_LOGIN = 10
)

// JSONResult is a shortcut for map[string]interface{}
type JSONResult map[string]interface{}
