package matcher

import "Tinder/user"

type InterestBasedMatcher struct {
}

func (i *InterestBasedMatcher) CalcMatchScore(u1 user.User, u2 user.User) int{
	if u1.HasLiked(u2.UserID) && u2.HasLiked(u1.UserID) {
		return 100
	}else if u1.HasLiked(u2.UserID) {
		return 50
	}else {
		return 0
	}
}