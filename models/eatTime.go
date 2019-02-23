package models

// EatTime is a alias of int. 0 for breakfast, 1 for lunch and
// 2 for supper. Get them by the func of their name
type EatTime int

const (
	breakfast EatTime = iota //0
	lunch                    //1
	supper                   //2
)

// Breakfast return a EatTime (int) var for breakfast
func Breakfast() EatTime {
	return breakfast
}

// Lunch return a EatTime (int) var for Lunch
func Lunch() EatTime {
	return lunch
}

// Supper return a EatTime (int) var for Supper
func Supper() EatTime {
	return supper
}

func (e EatTime) String() string {
	switch e {
	case breakfast:
		return "breakfast"
	case lunch:
		return "lunch"
	case supper:
		return "supper"
	default:
		return ""
	}
}
