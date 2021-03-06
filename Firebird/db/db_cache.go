package db

import (
	"Firebird/config"
	"fmt"
)

type CacheDataMap struct {
	userIdMap   map[int64]UserInfo
	userNameMap map[string]UserInfo

	symbolIdMap   map[int64]SymbolInfo
	symbolNameMap map[string]SymbolInfo

	scheduleIdMap    map[int64]UserSchedule
	scheduleGroupMap map[string][]UserSchedule

	configMap map[string]string
}

var cacheDataMap = CacheDataMap{}

func LoadAllToCache() {
	cache := CacheDataMap{
		userIdMap:        make(map[int64]UserInfo),
		userNameMap:      make(map[string]UserInfo),
		symbolIdMap:      make(map[int64]SymbolInfo),
		symbolNameMap:    make(map[string]SymbolInfo),
		scheduleIdMap:    make(map[int64]UserSchedule),
		scheduleGroupMap: make(map[string][]UserSchedule),
		configMap:        make(map[string]string),
	}

	// user
	userIdMap := cache.userIdMap
	userNameMap := cache.userNameMap

	userList := loadAllUserInfo()
	if len(userList) > 0 {
		for _, user := range userList {
			userIdMap[user.Id] = user
			userNameMap[user.UserName] = user
		}
	}

	// symbol
	symbolIdMap := cache.symbolIdMap
	symbolNameMap := cache.symbolNameMap

	symbolList := loadAllSymbolInfo()
	if len(symbolList) > 0 {
		for _, symbol := range symbolList {
			symbolIdMap[symbol.Id] = symbol
			symbolNameMap[symbol.SymbolName] = symbol
		}
	}

	// schedule
	scheduleIdMap := cache.scheduleIdMap
	scheduleGroupMap := cache.scheduleGroupMap

	scheduleList := loadAllSchedule()
	if len(scheduleList) > 0 {
		var key string
		var exists bool
		var groupList []UserSchedule
		for _, schedule := range scheduleList {
			scheduleIdMap[schedule.Id] = schedule
			key = getGroupKey(schedule.SymbolId, schedule.UserId)
			groupList, exists = scheduleGroupMap[key]
			if !exists {
				groupList = make([]UserSchedule, 0)
			}
			groupList = append(groupList, schedule)
			scheduleGroupMap[key] = groupList
		}
	}

	// config
	configMap := cache.configMap
	configList := loadAllConfigInfo()
	if len(configList) > 0 {
		for _, config := range configList {
			configMap[config.Ckey] = config.Cvalue
		}
	}

	cacheDataMap = cache

	// update host
	config.SetHostName(GetConfigFromCache("host_name"))
	log.Info("LoadAllToCache", len(cacheDataMap.symbolIdMap))
}

func getGroupKey(symbolId int64, userId int64) string {
	return fmt.Sprintf("%d_%d", symbolId, userId)
}

func GetUserFromCacheById(id int64) UserInfo {
	return cacheDataMap.userIdMap[id]
}

func GetUserFromCacheByName(name string) UserInfo {
	return cacheDataMap.userNameMap[name]
}

func GetAllSymbolFromCache() map[string]SymbolInfo {
	return cacheDataMap.symbolNameMap
}
func GetSymbolFromCacheById(id int64) SymbolInfo {
	return cacheDataMap.symbolIdMap[id]
}

func GetSymbolFromCacheByName(name string) SymbolInfo {
	return cacheDataMap.symbolNameMap[name]
}

func GetScheduleFromCacheById(id int64) UserSchedule {
	return cacheDataMap.scheduleIdMap[id]
}

func GetScheduleFromCacheByGroup(symbolId int64, userId int64) []UserSchedule {
	key := getGroupKey(symbolId, userId)
	return cacheDataMap.scheduleGroupMap[key]
}

func GetConfigFromCache(key string) string {
	return cacheDataMap.configMap[key]
}
