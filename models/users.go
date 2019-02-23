package models

// Sex is a alias for int. 0 stands for male and 1 stands for female
type Sex int

const (
	male Sex = iota
	female
)

// Male returns a Sex (int) var for Male
func Male() Sex {
	return male
}

// Female returns a Sex (int) var for Female
func Female() Sex {
	return female
}

// UserInfo contains Name, PhoneNum, Sex
type UserInfo struct {
	Name       string `json:"name" bson:"name"`
	PhoneNum   string `json:"phoneNum" bson:"phoneNum"`
	Sex        `json:"sex" bson:"sex"`
	PasswdHash string `json:"-" bson:"passwd"`
}

// UserDiet contains Meals and num
type UserDiet struct {
	Meals []Meal `json:"meals" bson:"meals"`
	Num   int    `json:"num" bson:"num"`
}

// User Info contains UserInfo and UserDiet
type User struct {
	UserInfo `json:"userInfo" bson:"userInfo"`
	UserDiet `json:"userDiet" bson:"userDiet"`
}

// TestUser is a test for user
var TestUser User
var testInfo UserInfo
var testDiet UserDiet

func init() {
	testInfo = UserInfo{
		"test",
		"123456789",
		Male(),
		"!@#$%^",
	}
	testDiet = UserDiet{
		testMeals,
		2,
	}
	TestUser = User{
		testInfo,
		testDiet,
	}
}

// SetUser accepts name, phone, hash and sex and returns User
func SetUser(name, phone, hash string, sex int) *User {
	var u = User{
		UserInfo{
			name,
			phone,
			Sex(sex),
			hash,
		},
		UserDiet{
			[]Meal{},
			0,
		},
	}
	return &u
}
