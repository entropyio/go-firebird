package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	// 可自定义盐值
	tokenSalt = "go-fb"
)

func MD5(data string) string {
	_md5 := md5.New()
	_md5.Write([]byte(data + tokenSalt))
	return hex.EncodeToString(_md5.Sum([]byte("")))
}
