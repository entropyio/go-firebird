package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetParamString(c *gin.Context, key string) string {
	v := c.GetString(key)
	if "" == v {
		v = c.PostForm(key)
	}
	if "" == v {
		v = c.Query(key)
	}
	return v
}

func GetParamInt64(c *gin.Context, key string) int64 {
	v := GetParamString(c, key)
	if "" == v {
		return 0
	}
	num64, err := strconv.ParseInt(v, 10, 64)
	if nil != err {
		return 0
	}
	return num64
}

func GetParamInt(c *gin.Context, key string) int {
	v := GetParamString(c, key)
	if "" == v {
		return 0
	}
	num32, err := strconv.Atoi(v)
	if nil != err {
		return 0
	}
	return num32
}

func GetParamFloat64(c *gin.Context, key string) float64 {
	v := GetParamString(c, key)
	if "" == v {
		return 0.0
	}
	f64, err := strconv.ParseFloat(v, 64)
	if nil != err {
		return 0.0
	}
	return f64
}

func GetParamTime(c *gin.Context, key string) time.Time {
	v := GetParamString(c, key)
	if "" == v {
		return time.Time{}
	}
	st, err := time.Parse("2006-01-02 03:04:05", v)
	if nil != err {
		return time.Time{}
	}
	return st
}
