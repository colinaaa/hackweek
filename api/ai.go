package api

import (
	"github.com/kataras/iris"
)

const port = "12380"

// Photo is the handler for upload photo
func Photo(ctx iris.Context) {
	ctx.Redirect("http://ai:"+port, 304)
}

// Recommend is the handler to get recommend food from AI
func Recommend(ctx iris.Context) {
	ctx.Redirect("http://ai:"+port, 304)
}
