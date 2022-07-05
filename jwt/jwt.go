/***************************
@File        : jwt.go
@Time        : 2022/02/08 17:59:49
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : jwt 认证
****************************/

package jwt

import "github.com/dgrijalva/jwt-go"

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
