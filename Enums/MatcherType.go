package enums

type MatcherType int

const (
	BASIC MatcherType = iota
	INTEREST
	LOCATION
)

func (m MatcherType) String() string {
	switch m {
	case BASIC:
		return "basic"
	case INTEREST:
		return "interest"
	default:
		return "location"
	} 
}