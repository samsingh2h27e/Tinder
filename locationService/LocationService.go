package locationservice

import (
	location "Tinder/Location"
	locationstrategies "Tinder/locationStrategies"
	"Tinder/user"
)

type LocationService struct {
	Strat locationstrategies.LocationStrategy
}

func (l *LocationService) SetStrategy(strat locationstrategies.LocationStrategy) {
	l.Strat = strat
}

func (l *LocationService) FindNearByUsers(loc location.Location, maxD float64, allusers []user.User) []user.User{
	return l.Strat.FindNearByUsers(loc,maxD,allusers)
}