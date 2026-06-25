package matcher

import "Tinder/user"

type LocationBasedMatcher struct {
}

func (l *LocationBasedMatcher) CalcMatchScore(u1 user.User, u2 user.User) int{
	return 100
}