package filters

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/astaxie/beego/context"
	"beego/extensions"
)

var (
	authExcept = map[string]bool{"/v1/login": true}
)

type LoginClaims struct {
	UserID int64
	jwt.StandardClaims
}


func Auth(ctx *context.Context) {
	tokenString := ctx.Input.Header("token")

	//tokenString, err := extensions.NewJWTTokenStringWithClaims(LoginClaims{
	//	1234567,
	//	jwt.StandardClaims {
	//		ExpiresAt: 1537956658,
	//		Issuer: "test",
	//	},
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println("token:" + tokenString)

	_, isExcept := authExcept[ctx.Request.RequestURI]
	if isExcept {
		return
	}

	//if len(tokenString) == 0 {
	//	ctx.Redirect(302, "/login")
	//}

	// parse token with claims
	token, err := extensions.ParseJWTTokenWithClaims(tokenString, &LoginClaims{})

	if err != nil {
		fmt.Println(err)
		ctx.Redirect(302, "/login")
		return
	}

	// validate & extract token
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
		ctx.Redirect(302, "/login")
	}
}