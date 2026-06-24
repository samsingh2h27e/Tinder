package location

import (
	"os/user"
)

type LocationService struct {
	Strat LocationStrategy
}

func (l *LocationService) SetStrategy(strat LocationStrategy) {
	l.Strat = strat
}

func (l *LocationService) FindNearByUsers(loc Location, maxD float64, allusers []user.User) []user.User{

}