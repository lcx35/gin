package apis

import (
	"net/http"
	"web/libs"
	"web/models"

	"github.com/gin-gonic/gin"
)

var issuer = libs.Conf.Read("site", "issuer")

func Login(c *gin.Context) {
	req_u := models.User{}
	err := c.Bind(&req_u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	d_u := models.User{Username: req_u.Username}
	user, err := d_u.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if req_u.Password == user.Password {
		userinfo := libs.UserInfo{}
		userinfo.Id = user.Id
		userinfo.Username = user.Username
		userinfo.Password = user.Password
		userinfo.StandardClaims.Issuer = issuer
		token, err := userinfo.CreateToken()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0, "msg": "用户名或密码错误",
			})
		}
		c.Header("Token", token)
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "msg": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 0, "msg": "用户名或密码错误",
		})
	}

}
