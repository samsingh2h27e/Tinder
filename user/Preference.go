package user

import (
	enums "Tinder/Enums"
	interest "Tinder/Interest"
)

type Preference struct {
	MinAge       int
	MaxAge       int
	MaxDis       float64
	Interests    []interest.Interest
	InterestedIn []enums.Gender
}