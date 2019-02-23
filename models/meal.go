package models

import (
	"time"
)

// TimeStamp is int64 that represent unix timestamp
// It can be transform into yyyy-mm-dd by using
// time.Unix(timestamp)
type TimeStamp int64

func (t TimeStamp) String() string {
	return time.Unix(int64(t), 0).Format("2006-01-02")
}

// Meal contains Name, Date, Time, Food, PlaceCategory, PlaceExact
type Meal struct {
	Name          string   `json:"name" bson:"name"`
	Date          string   `json:"date" bson:"date"`
	Time          EatTime  `json:"time" bson:"time"`
	Food          []string `json:"food" bson:"food"`
	PlaceCategory Place    `json:"placeCategory" bson:"placeCategory"`
	PlaceExact    string   `json:"placeExact" bson:"placeExact"`
}

// testMeal1 is a test model for test
var testMeal1 Meal

// testMeal2 is a test model for test
var testMeal2 Meal
var testMeals []Meal

func init() {
	var foods []string
	foods = append(foods, "aa")
	foods = append(foods, "bbb")
	testMeal1 = Meal{
		"test",
		TimeStamp(time.Now().Unix()).String(),
		Breakfast(),
		foods,
		Out(),
		"test place",
	}
	testMeal2 = Meal{
		"test2",
		TimeStamp(time.Now().Unix()).String(),
		Supper(),
		foods,
		Out(),
		"test2 place",
	}
	testMeals = append(testMeals, testMeal1)
	testMeals = append(testMeals, testMeal2)
}
