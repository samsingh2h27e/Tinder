package userprofile

import (
	enums "Tinder/Enums"
	interest "Tinder/Interest"
	location "Tinder/Location"
	"fmt"
	"strconv"
)

type UserProfile struct {
	Name      string
	Age       int
	Gen       enums.Gender
	Bio       string
	Photos    []string
	Interests []interest.Interest
	Loc		  location.Location
}

func (U *UserProfile) Display() {
	fmt.Println("Name: "+ U.Name)
	fmt.Println("Age: " + strconv.Itoa(U.Age))
	fmt.Println("Gender: " + U.Gen.String())
	fmt.Println("Bio: "+ U.Bio)
	for _,val := range U.Photos {
		fmt.Println(val)
	}
	for _,val := range U.Interests {
		fmt.Println(val)
	}
}