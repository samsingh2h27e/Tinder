package enums

type SwipeAction int

const (
	LEFT SwipeAction = iota
	RIGHT
)

func (s SwipeAction) String() string {
	switch s {
		case LEFT:
			return "left"
		default:
			return "right"
	}
}