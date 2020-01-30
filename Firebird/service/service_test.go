package service

import (
	"Firebird/db"
	"testing"
)

func TestUpdateUserBenefit(t *testing.T) {
	db.LoadAllToCache()
	syncAllSymbolKline(100)

	updateAllUserBenefit(90)
}
