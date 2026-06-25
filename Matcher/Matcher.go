package matcher

import "Tinder/user"

type Matcher interface {
	CalcMatchScore(u1 user.User,u2 user.User) int
}