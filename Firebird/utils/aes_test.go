package utils

import "testing"

//encode params = 6C498135C351A17E8089B298E04FB19EBBDA1083851CE794E5A25D1C47E0478EE1DAD62D2D6648EACC9F7E0D77CF0E7194DF868B81206B42A5FCE9B69A87D837

//decode params = userId=1&symbolId=0&status=2&pageNumber=1&type=1
func TestAes(t *testing.T) {
	/*
	*src 要加密的字符串
	*key 用来加密的密钥 密钥长度可以是128bit、192bit、256bit中的任意一个
	*16位key对应128bit
	 */
	msg := "userId=1中文测试&symbolId=0&status=2&pageNumber=1&type=1"

	crypted := AesEncrypt(msg)
	println(crypted)

	msg = AesDecrypt(crypted)
	println(msg)
}
