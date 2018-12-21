package libs

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var conf = Configs{}
var Key = []byte(conf.Read("site", "sitekeys"))

type UserInfo struct {
	Id       int
	Username string
	Password string
	jwt.StandardClaims
}

func (u *UserInfo) CreateToken() (string, error) {
	u.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * 50).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return token.SignedString(Key)
}

func ParseToken(tokenString string) (*UserInfo, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &UserInfo{}, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	fmt.Println(err)
	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*UserInfo); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (u *UserInfo) RefreshToken() (string, error) {
	return u.CreateToken()
}
