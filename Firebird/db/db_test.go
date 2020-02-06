package db

import (
	"fmt"
	"testing"
)

var id int64

/** user info test **/
func TestInsertUserInfo(t *testing.T) {
	var userInfo = UserInfo{
		Username: "test123",
		UserDesc: "desc123",
		Status:   1,
	}
	id = InsertUserInfo(&userInfo)
	fmt.Println(id)
}

func TestQueryUserInfo(t *testing.T) {
	var userInfoQuery = UserInfoQuery{
		Id: id,
	}
	count, list := QueryUserInfo(&userInfoQuery)
	fmt.Println(count, list)
}

func TestUpdateUserInfo(t *testing.T) {
	var userInfo = UserInfo{
		Id:       id,
		Username: "test1-123",
		UserDesc: "desc1-123",
	}
	count := UpdateUserInfo(&userInfo)
	fmt.Println(count)
}

func TestDeleteUserInfo(t *testing.T) {
	count := DeleteUserInfo(id)
	fmt.Println(count)
}

/** user account test **/
func TestInsertUserAccount(t *testing.T) {
	var userAccount = UserAccount{
		UserId:   100,
		SymbolId: 100,
		Price:    2.356,
		Amount:   100,
		Status:   1,
	}
	id = InsertUserAccount(&userAccount)
	fmt.Println(id)
}

func TestQueryUserAccount(t *testing.T) {
	var userAccountQuery = UserAccountQuery{
		Id: id,
	}
	count, list := QueryUserAccount(&userAccountQuery)
	fmt.Println(count, list)
}

func TestUpdateUserAccount(t *testing.T) {
	var userAccount = UserAccount{
		Id: id,
		//YestBenifit: 23.56,
		//TotalBenifit:123.45,
		//HoldBenifit: 56.78,
		HoldAmount: 200,
	}
	count := UpdateUserAccount(&userAccount)
	fmt.Println(count)
}

func TestDeleteUserAccount(t *testing.T) {
	count := DeleteUserAccount(id)
	fmt.Println(count)
}

func TestGetUserAccountById(t *testing.T) {
	record := GetUserAccountById(1)
	fmt.Println(record)
}

func TestGetUserAccountByUId(t *testing.T) {
	record := GetUserAccountByUid(1, 1)
	fmt.Println(record)
}

/** user trade test **/
func TestInsertUserTrade(t *testing.T) {
	var userTrade = UserTrade{
		UserId:   100,
		SymbolId: 100,
		Price:    2.356,
		Amount:   100,
		Status:   1,
	}
	id = InsertUserTrade(&userTrade)
	fmt.Println(id)
}

func TestQueryUserTrade(t *testing.T) {
	var userTradeQuery = UserTradeQuery{
		Id: id,
	}
	count, list := QueryUserTrade(&userTradeQuery)
	fmt.Println(count, list)
}

func TestUpdateUserTrade(t *testing.T) {
	var userTrade = UserTrade{
		Id:     id,
		Amount: 50,
		Price:  2.51,
	}
	count := UpdateUserTrade(&userTrade)
	fmt.Println(count)
}

func TestDeleteUserTrade(t *testing.T) {
	count := DeleteUserTrade(id)
	fmt.Println(count)
}

/** symbol info test **/
func TestInsertSymbolInfo(t *testing.T) {
	var symbolInfo = SymbolInfo{
		SymbolName:  "test1",
		SymbolDesc:  "desc1",
		SymbolIcon:  "icon",
		SymbolGroup: "group",
		Status:      1,
	}
	id = InsertSymbolInfo(&symbolInfo)
	fmt.Println(id)
}

func TestQuerySymbolInfo(t *testing.T) {
	var symbolInfoQuery = SymbolInfoQuery{
		Id: id,
	}
	count, list := QuerySymbolInfo(&symbolInfoQuery)
	fmt.Println(count, list)
}

func TestUpdateSymbolInfo(t *testing.T) {
	var symbolInfo = SymbolInfo{
		Id:          id,
		SymbolName:  "test1-123",
		SymbolDesc:  "desc1-123",
		SymbolIcon:  "icon-123",
		SymbolGroup: "group-123",
	}
	count := UpdateSymbolInfo(&symbolInfo)
	fmt.Println(count)
}

func TestDeleteSymbolInfo(t *testing.T) {
	count := DeleteSymbolInfo(id)
	fmt.Println(count)
}

func TestLoadAllToCache(t *testing.T) {
	LoadAllToCache()
}
