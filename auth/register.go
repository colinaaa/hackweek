package auth

import (
	"github.com/colinaaa/hackweek/models"

	"github.com/colinaaa/hackweek/database"
	"github.com/kataras/iris"
)

type registerUser struct {
	Name     string `json:"name"`
	PhoneNum string `json:"phoneNum"`
	Password string `json:"password"`
	Sex      int    `json:"sex"`
}

// Register is the handler for register
func Register(ctx iris.Context) {
	var r registerUser
	// Register
	// {
	// 	name string
	// 	phoneNum string,
	// 	password string,
	// 	sex int,
	// }
	err := ctx.ReadJSON(&r)
	if err != nil {
		ctx.Application().Logger().Infof("register: %s", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	if existedUser(r.PhoneNum) {
		ctx.JSON(response{
			"phoneNum used",
			400,
		})
		return
	}
	u := models.SetUser(
		r.Name,
		r.PhoneNum,
		string(generatePasswd(r.Password)),
		r.Sex,
	)
	//////////////////////
	// TODO: phote used //
	//////////////////////
	oid, err := database.InsertUser(u)
	if err != nil {
		ctx.Application().Logger().Infof("Login: %s", err)
		ctx.JSON(response{
			"server wrong",
			500,
		})
	}
	ctx.JSON(
		response{
			"ok",
			200,
		},
	)
	ctx.Application().Logger().Infof("register user, oid: %s", oid)
	return
}
func existedUser(phoneNum string) bool {
	u := database.SelectByPhone(phoneNum)
	if u.UserInfo == (models.UserInfo{}) {
		return false
	}
	return true
}
