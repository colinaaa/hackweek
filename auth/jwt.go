package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

type jwtToken struct {
	Token string `json:"token"`
}

// VertifyToken is a middleware to vertify jwt
func VertifyToken(ctx iris.Context) {
	t := ctx.GetHeader("Authorization")
	if len(t) < 7 {
		ctx.JSON(
			response{
				"broken token",
				401,
			},
		)
		return
	}
	t = t[6:]
	token, err := jwt.Parse(
		t,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				ctx.JSON(
					response{
						"token wrong",
						401,
					},
				)
				return nil, fmt.Errorf("token method wrong")
			}
			return []byte("eat something"), nil
		})
	if err != nil {
		ctx.JSON(
			response{
				"token wrong",
				401,
			},
		)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Application().Logger().Infof("good token from %s\n", claims["phoneNum"])
		ctx.Values().Set("phoneNum", claims["phoneNum"])
		ctx.Next()
		return
	}
	ctx.Application().Logger().Info("invalid token")
	ctx.JSON(
		response{
			"token wrong",
			401,
		},
	)
}
