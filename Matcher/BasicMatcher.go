package matcher

import "Tinder/user"

type BasicMatcher struct {
}

func (b *BasicMatcher) CalcMatchScore(u1 user.User, u2 user.User) int {
	return 100
}