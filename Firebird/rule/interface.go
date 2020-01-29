package rule

import "Firebird/db"

type Executor interface {
	execute(ruleItem *db.RuleItem, paramsMap *map[string]interface{}) (bool)
}
