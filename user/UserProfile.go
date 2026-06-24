package user

import (
	enums "Tinder/Enums"
	location "Tinder/Location"
)

type UserProfile struct {
	Name      string
	Age       int
	Gen       enums.Gender
	Bio       string
	Photos    []string
	Interests []Interest
	Loc		  location.Location
}

func (U *UserProfile) Display() {

}