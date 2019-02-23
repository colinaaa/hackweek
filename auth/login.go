package auth

import (
	"github.com/kataras/iris"
)

type loginUser struct {
	PhoneNum string `json:"phoneNum"`
	Password string `json:"password"`
}

// Login is the handler for login
func Login(ctx iris.Context) {
	// {
	// 	"phoneNum":xxx,
	// 	"password":xxx,
	// }
	var u loginUser
	err := ctx.ReadJSON(&u)
	if err != nil {
		ctx.Application().Logger().Debugf("login: read json %s", err)
		ctx.JSON(response{
			"server wrong",
			500,
		})
		return
	}
	///////////////////////////
	//TODO: invalid string////
	//////////////////////////
	token, err := generateToken(u.PhoneNum)
	if err != nil {
		ctx.Application().Logger().Debugf("token: %s", err)
	}
	if vertifyPasswd(u.Password, u.PhoneNum) {
		ctx.JSON(struct {
			response
			Token string `json:"token"`
		}{
			response{
				"ok",
				200,
			},
			token,
		})
		return
	}
	ctx.JSON(response{
		"password wrong",
		401,
	})
}
