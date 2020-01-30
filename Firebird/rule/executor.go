package rule

import (
	"Firebird/config"
	"Firebird/db"
)

var (
	rateExec  rateExecutor
	priceExec priceExecutor
)

func init() {
	rateExec = rateExecutor("Rate")
	priceExec = priceExecutor("Price")
}

func Execute(ruleList []db.RuleItem, paramsMap map[string]interface{}) bool {
	if len(ruleList) == 0 || len(paramsMap) == 0 {
		return false
	}

	// both join result is false when start
	var (
		andInit    = false
		andResult  = false
		orInit     = false
		orResult   = false
		itemResult = false
	)

	for _, ruleItem := range ruleList {
		itemResult = executeRuleItem(&ruleItem, &paramsMap)
		if ruleItem.JoinType == config.JOIN_TYPE_AND {
			if !andInit {
				andResult = itemResult
				andInit = true
			} else {
				andResult = andResult && itemResult
			}
		} else if ruleItem.JoinType == config.JOIN_TYPE_OR {
			if !orInit {
				orResult = itemResult
				orInit = true
			} else {
				orResult = orResult || itemResult
			}
		}
		// or condition success, no need to check other conditions
		// for and, need to verify all conditions
		if orInit && orResult {
			return true
		}
	}
	if !andInit {
		andResult = false
	}
	if !orInit {
		orResult = false
	}

	return andResult || orResult
}

func executeRuleItem(ruleItem *db.RuleItem, paramsMap *map[string]interface{}) bool {
	ruleType := ruleItem.RuleType
	result := false
	switch ruleType {
	case config.RULE_TYPE_RATE:
		result = rateExec.execute(ruleItem, paramsMap)
		break
	case config.RULE_TYPE_PRICE:
		result = priceExec.execute(ruleItem, paramsMap)
		break
	default:
		break
	}
	return result
}
