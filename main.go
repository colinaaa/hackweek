package main

import (
	"github.com/colinaaa/hackweek/api"

	"github.com/colinaaa/hackweek/auth"

	"github.com/colinaaa/hackweek/database"
	"github.com/kataras/iris"
)

var app = iris.New()

func main() {
	app.Run(iris.Addr(":8080"))
}
func apiMiddleWare(ctx iris.Context) {
	ctx.Next()
}
func init() {
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("%s %s", ctx.Method(), ctx.Path())
		ctx.Next()
	})
	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.JSON(
			struct {
				Msg        string `json:"msg"`
				StatusCode int    `json:"status_code"`
			}{
				"not found",
				404,
			},
		)
	})
	apiRouter := app.Party("/api", apiMiddleWare)
	authRouter := apiRouter.Party(
		"/user/{num:string regexp([0-9]+$)}",
		auth.VertifyToken,
		auth.Match,
	)
	// Get user info
	authRouter.Get("/", func(ctx iris.Context) {
		num := ctx.Params().Get("num")
		u := database.SelectByPhone(num)
		ctx.JSON(u)
	})
	// register routers
	apiRouter.Post("/user", auth.Register)
	apiRouter.Post("/login", auth.Login)
	apiRouter.Post("/photo", api.Photo)
	authRouter.Get("/meals", api.Meals)
	authRouter.Post("/meals", api.SetMeal)
	authRouter.Get("/diet", api.Diet)
	// authRouter.Get("/recommend", api.Recommend)
}
