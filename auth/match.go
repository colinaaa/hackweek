package auth

import (
	"github.com/kataras/iris"
)

// Match is a middleware to inspect whether the phoneNum of URL and
// the phoneNum of token are matched
func Match(ctx iris.Context) {
	tNum := ctx.Values().GetString("phoneNum")
	uNum := ctx.Params().GetString("num")
	if tNum != uNum {
		ctx.JSON(
			response{
				"token does not fit user",
				403,
			},
		)
		return
	}
	ctx.Next()
}
