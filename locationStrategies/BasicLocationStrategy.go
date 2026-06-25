package locationstrategies

import (
	location "Tinder/Location"
	"Tinder/user"
)

type BasicLocationStrategy struct {
}

func (b *BasicLocationStrategy) FindNearByUsers(loc location.Location, maxD float64, allUser []user.User) []user.User {
	var ans []user.User
	for _, val := range allUser {
		dis := val.Up.Loc.DistanceInKM(loc)
		if dis <= maxD {
			ans = append(ans, val)
		}
	}
	return ans
}