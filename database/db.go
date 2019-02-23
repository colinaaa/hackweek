package database

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/colinaaa/hackweek/models"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var db *mongo.Database
var user *mongo.Collection

func init() {
	client, err := mongo.NewClient("mongodb://localhost")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("eat")
	user = db.Collection("user")
}

// InsertUser recives a User struct and insert into db
// and returns inserted id
func InsertUser(u *models.User) (string, error) {
	collection, err := user.Clone()
	if err != nil {
		log.Printf("insert collection: %s", err)
		return "", err
	}
	res, err := collection.InsertOne(context.Background(), u)
	if err != nil {
		log.Printf("insert: %s", err)
		return "", err
	}
	return fmt.Sprint(res.InsertedID), nil
}

// SelectByPhone accepts a string for phoneNum and returns ptr to User
func SelectByPhone(num string) *models.User {
	collection, err := user.Clone()
	if err != nil {
		log.Printf("select collection: %s", err)
		return nil
	}
	var u = models.User{}
	err = collection.FindOne(
		context.Background(),
		bson.D{{
			"userInfo.phoneNum",
			num,
		}},
	).Decode(&u)
	if err != nil {
		log.Println(err)
	}
	return &u
}

// AddMeal receives the phoneNum of user and a Meal struct for the user
// and returns the ptr of User and error
func AddMeal(num string, m models.Meal) (*models.User, error) {
	filter := bson.D{{
		"userInfo.phoneNum",
		num,
	},
	}
	update := bson.D{{
		"$push",
		bson.D{{
			"userDiet.meals",
			m,
		}},
	}, {
		"$inc",
		bson.D{{
			"userDiet.num",
			1,
		}},
	}}
	option := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res := user.FindOneAndUpdate(context.Background(), filter, update, option)
	if res.Err() != nil {
		log.Printf("update: %s", res.Err())
		return nil, res.Err()
	}
	var u models.User
	err := res.Decode(&u)
	if err != nil {
		log.Printf("update: decode: %s", err)
		return nil, err
	}
	return &u, nil
}
