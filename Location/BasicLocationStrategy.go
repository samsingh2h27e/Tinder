package location

import "os/user"

type BasicLocationStrategy struct {
}

func (b *BasicLocationStrategy) FindNearByUsers(loc Location, maxD float64, allUser []user.User) []user.User{

}