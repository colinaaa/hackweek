package models

// Place is an alias for int. 0 for out, 1 for canteen, 2 for home
type Place int

const (
	out     Place = iota //0
	canteen              //1
	home                 //2
)

// Out returns a Place (int) var for Out
func Out() Place {
	return out
}

// Canteen returns a Place (int) var for Canteen
func Canteen() Place {
	return canteen
}

// Home returns a Place (int) var for Home
func Home() Place {
	return home
}

func (p Place) String() string {
	switch p {
	case out:
		return "out"
	case canteen:
		return "canteen"
	case home:
		return "home"
	default:
		return ""
	}
}
