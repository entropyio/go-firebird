package service

import (
	"testing"
	"Firebird/db"
)

func TestUpdateUserBenefit(t *testing.T) {
	db.LoadAllToCache()
	syncAllSymbolKline(100)

	updateAllUserBenefit(90)
}
