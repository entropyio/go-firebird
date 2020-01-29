package utils

import "time"

func GetDateTime(date string) (dtime time.Time) {
	dtime, err := time.Parse("2006-01-02 03:04:05", date)
	if nil != err {
		log.Errorf("parse date string error: %s", date)
	}

	return dtime
}

func GetDateTimeStr(dtime time.Time) (date string) {
	return dtime.Format("2006-01-02 15:04:05")
}

func GetDateStr(dtime time.Time) (date string) {
	return dtime.Format("2006-01-02")
}

func GetTimeStr(dtime time.Time) (date string) {
	return dtime.Format("15:04:05")
}
