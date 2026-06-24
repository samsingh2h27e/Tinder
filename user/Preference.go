package user

import enums "Tinder/Enums"

type Preference struct {
	MinAge       int
	MaxAge       int
	MaxDis       float64
	Interests    []Interest
	InterestedIn []enums.Gender
}