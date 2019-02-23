package api

import (
	"github.com/colinaaa/hackweek/database"
	"github.com/colinaaa/hackweek/models"
	"github.com/kataras/iris"
	"log"
)

type response struct {
	Msg        string `json:"msg"`
	StatueCode int    `json:"statue_code"`
}

// Meals is the handler to get user's meals info
func Meals(ctx iris.Context) {
	num := ctx.Params().GetString("num")
	m := database.SelectByPhone(num).UserDiet
	ctx.JSON(m)
}

// SetMeal is the handler to set a meal
func SetMeal(ctx iris.Context) {
	phoneNum := ctx.Values().GetString("phoneNum")
	var m models.Meal
	err := ctx.ReadJSON(&m)
	if err != nil {
		log.Printf("setMeal read wrong: %s", err)
		ctx.JSON(
			response{
				"server wrong",
				500,
			},
		)
		return
	}
	u, err := database.AddMeal(phoneNum, m)
	if err != nil {
		ctx.JSON(
			response{
				"add wrong",
				500,
			},
		)
		return
	}
	ctx.JSON(u)
}

// Diet is the handler for getting user diet in array
func Diet(ctx iris.Context) {
	num := ctx.Params().Get("num")
	u := database.SelectByPhone(num)
	res := []string{}
	for _, meal := range u.UserDiet.Meals {
		res = append(res, meal.Food...)
	}
	ctx.JSON(res)
}
