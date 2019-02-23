package api

import (
	"github.com/kataras/iris"
)

var port = "12380"

// Photo is the handler for upload photo
func Photo(ctx iris.Context) {
	ctx.Redirect("http://localhost:"+port, 304)
}

// Recommend is the handler to get recommend food from AI
// func Recommend(ctx iris.Context) {
// ctx.Redirect("http://localhost:"+port, 302)
// }
