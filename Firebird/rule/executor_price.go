package rule

import "Firebird/db"

type priceExecutor string

func (pe priceExecutor) execute(ruleItem *db.RuleItem, paramsMap *map[string]interface{}) bool {
	return true
}
