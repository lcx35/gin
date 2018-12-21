package middleware

import (
	"log"
	"net/http"
	"time"
	"web/libs"

	"github.com/gin-gonic/gin"
)

var issuer = libs.Conf.Read("site", "issuer")

func Tokenauth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求头中的Token
		token := c.Request.Header["Token"]
		if token == nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "请先登录"})
		}

		//解析token获取用户信息
		u, err := libs.ParseToken(token[0])
		if err != nil {
			log.Println("err:", err)
			c.Abort()
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "请重新登录"})
		}
		//剩余过期时间
		t := u.StandardClaims.ExpiresAt - time.Now().Unix()
		u_issuer := u.StandardClaims.Issuer
		if 1200 <= t && t <= 3000 {
			if u_issuer == issuer {
				//通过
				//c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "身份验证成功"})
				c.Next() //可以不写
			}
		} else if 0 < t && t < 1200 {
			if u_issuer == issuer {
				//刷新token
				token, err := u.RefreshToken()
				if err != nil {
					//失败
					c.Abort()
					c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "请重新登录"})
				}
				c.Header("Token", token)
				//c.JSON(http.StatusOK, gin.H{"status": 2, "token": token, "msg": "身份验证成功"})
				c.Next()
			}
		} else {
			//失败
			c.Abort()
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "请重新登录"})
		}
	}
}
