package location

import (
	"os/user"
)

type LocationStrategy interface {
	FindNearByUsers(loc Location,maxD float64,allUser []user.User) []user.User
}