package locationstrategies

import (
	location "Tinder/Location"
	"Tinder/user"
)

type LocationStrategy interface {
	FindNearByUsers(loc location.Location,maxD float64,allUser []user.User) []user.User
}