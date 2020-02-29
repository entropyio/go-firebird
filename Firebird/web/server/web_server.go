package server

import (
	"Firebird/logger"
	"Firebird/service"
	"Firebird/utils"
	"Firebird/web"
	"Firebird/web/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var log = logger.NewLogger("[web]")

func StartHttpServer(webRoot string) {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// admin webpage
	router.StaticFS("/firebird-admin", http.Dir(webRoot))

	router.Use(decryptParams())
	router.POST("/user/login", api.UserLogin)

	//以下的接口，都使用Authorize()中间件身份验证
	router.Use(authorize())

	router.POST("/user/logout", api.UserLogout)

	router.POST("/user/account/list", api.ListUserAccount)
	router.POST("/user/account/get", api.GetUserAccountByUid)
	router.POST("/user/account/save", api.SaveUserAccount)
	router.POST("/user/account/delete", api.DeleteUserAccount)

	router.POST("/user/trade/list", api.ListUserTrade)
	router.POST("/user/trade/save", api.AddUserTrade)

	router.POST("/current/price", getCurrentPrice)
	router.POST("/current/account/get", api.GetCurrentAccount)
	router.POST("/current/account/list", api.ListCurrentAccounts)

	router.POST("/user/schedule/list", api.ListUserSchedule)
	router.POST("/user/schedule/save", api.SaveUserSchedule)
	router.POST("/user/schedule/delete", api.DeleteUserSchedule)
	router.POST("/user/schedule/detail", api.DetailUserSchedule)
	router.POST("/user/rule/list", api.ListRuleItem)
	router.POST("/user/rule/save", api.SaveRuleItem)
	router.POST("/user/rule/delete", api.DeleteRuleItem)

	router.POST("/user/data/list", api.ListUserData)

	// symbol info
	router.POST("/user/symbol/list", api.ListSymbolInfo)
	router.POST("/user/symbol/save", api.SaveSymbolInfo)
	router.POST("/user/symbol/delete", api.DeleteSymbolInfo)

	// user info
	router.POST("/user/info/list", api.ListUserInfo)
	router.POST("/user/info/save", api.SaveUserInfo)
	router.POST("/user/info/delete", api.DeleteUserInfo)

	// config info
	router.POST("/config/list", api.ListConfigInfo)
	router.POST("/config/save", api.SaveConfigInfo)
	router.POST("/config/delete", api.DeleteConfigInfo)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Info("http server started.")
}

/**
decrypt params handler
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
			c.JSON(200, web.JSONResult{
				"retCode": web.CODE_NEED_LOGIN,
				"message": "Authentication needed",
			})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}

func getCurrentPrice(c *gin.Context) {
	symbol := c.PostForm("symbol")
	c.JSON(200, web.JSONResult{
		"retCode": web.CODE_SUCCESS,
		"message": "SUCCESS",
		"data": web.JSONResult{
			"symbol": symbol,
			"price":  service.GetSymbolPrice(symbol),
		},
	})
}
