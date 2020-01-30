package web

import (
	"Firebird/data"
	"Firebird/db"
	"Firebird/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var loginMap = make(map[int64]*data.LoginUser)

func userLogin(c *gin.Context) {
	username := utils.GetParamString(c, "username")
	//pwd := utils.GetParamString(c, "pwd")
	user := db.GetUserInfoByName(username)
	if user.Id > 0 {
		loginUser := data.LoginUser{
			UserId: user.Id,
			Ts:     time.Now().UnixNano() / 1e6,
		}
		loginUser.Token = utils.MD5(strconv.FormatInt(loginUser.UserId, 10) + strconv.FormatInt(loginUser.Ts, 10))

		loginMap[loginUser.UserId] = &loginUser

		c.JSON(200, JSONResult{
			"retCode": CODE_SUCCESS,
			"message": "SUCCESS",
			"data":    loginUser,
		})
	} else {
		c.JSON(200, JSONResult{
			"retCode": CODE_FAILED,
			"message": "Login failed",
		})
	}
}

func userLogout(c *gin.Context) {
	userId := utils.GetParamInt64(c, "userId")
	token := utils.GetParamString(c, "token")
	if userId > 0 {
		if loginUser, ok := loginMap[userId]; ok {
			if token == loginUser.Token {
				delete(loginMap, userId)
				c.JSON(200, JSONResult{
					"retCode": CODE_SUCCESS,
					"message": "SUCCESS",
					"data":    userId,
				})
				return
			}
		}
	}

	c.JSON(200, JSONResult{
		"retCode": CODE_FAILED,
		"message": "logout failed",
	})
}
