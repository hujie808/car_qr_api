//package main
package utils
import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)


type UserInfo struct {
    ID uint64
    Username string
}

func jwt_token_start(user *UserInfo) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"]=user.ID
	claims["username"]=user.Username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte("hujie"))
	if err != nil {
		log.Println("Error while signing the token")
	}
	return tokenString
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("hujie"), nil
	}
}



func ParseToken(tokenss string) (user *UserInfo, err error) {
	user = &UserInfo{}
	token, err := jwt.Parse(tokenss, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	user.ID = uint64(claim["id"].(float64))
	user.Username = claim["username"].(string)
	return
}

func test_main() {
	user:=UserInfo{ID:1,Username:"wo"}
	token_str :=jwt_token_start(&user)
	users,err:=ParseToken(token_str)
	fmt.Println(users,err)
}
