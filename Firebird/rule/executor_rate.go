package rule

import "Firebird/db"

type rateExecutor string

func (re rateExecutor) execute(ruleItem *db.RuleItem, paramsMap *map[string]interface{}) bool {
	return true
}
