package matcher

import (
	enums "Tinder/Enums"
)

type MatcherFactory struct {
}

func (m *MatcherFactory) CreateMatcher(t enums.MatcherType) Matcher {
	switch t.String() {
	case "basic" :
		return &BasicMatcher{}
	case "interest" :
		return &InterestBasedMatcher{}
	default :
		return &LocationBasedMatcher{}
	}
}