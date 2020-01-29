package web

import (
	"github.com/gin-gonic/gin"
	"Firebird/logger"
	"Firebird/utils"
	"strings"
	"Firebird/service"
)

var log = logger.NewLogger("[web]")

func StartHttpServer() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use(decryptParams())
	router.POST("/user/login", userLogin)

	//以下的接口，都使用Authorize()中间件身份验证
	router.Use(authorize())

	router.POST("/user/logout", userLogout)

	router.POST("/user/account/list", listUserAccount)
	router.POST("/user/account/get", getUserAccountByUid)

	router.POST("/user/trade/list", listUserTrade)
	router.POST("/user/trade/save", addUserTrade)

	router.POST("/current/price", getCurrentPrice)
	router.POST("/current/account/get", getCurrentAccount)
	router.POST("/current/account/list", listCurrentAccounts)

	router.POST("/user/schedule/list", listUserSchedule)
	router.POST("/user/schedule/save", saveUserSchedule)
	router.POST("/user/schedule/delete", deleteUserSchedule)
	router.POST("/user/schedule/detail", detailUserSchedule)
	router.POST("/user/rule/list", listRuleItem)
	router.POST("/user/rule/save", saveRuleItem)
	router.POST("/user/rule/delete", deleteRuleItem)

	router.POST("/user/data/list", listUserData)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Info("http server started.")
}

/**
	decrtyp params handler
 */
func decryptParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.PostForm("q")
		q = utils.AesDecrypt(q)
		if q != "" {
			params := strings.Split(q, "&")
			for i := 0; i < len(params); i++ {
				kvs := strings.Split(params[i], "=")
				c.Set(kvs[0], kvs[1])
			}
		}

		c.Next()
	}
}

/**
	authentication check handler
 */
func authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := utils.GetParamString(c, "userId")
		ts := utils.GetParamString(c, "ts")       // 时间戳
		token := utils.GetParamString(c, "token") // 访问令牌
		if strings.ToLower(utils.MD5(userId+ts)) == strings.ToLower(token) {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(200, JSONResult{
				"retCode": CODE_NEED_LOGIN,
				"message": "Authentication needed",
			})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}

func getCurrentPrice(c *gin.Context) {
	symbol := c.PostForm("symbol")
	c.JSON(200, JSONResult{
		"retCode": CODE_SUCCESS,
		"message": "SUCCESS",
		"data": JSONResult{
			"symbol": symbol,
			"price":  service.GetSymbolPrice(symbol),
		},
	})
}
