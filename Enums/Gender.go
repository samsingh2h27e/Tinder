package enums

type Gender int

const (
	MALE Gender = iota
	FEMALE
	OTHERS
)

func (G Gender) String() string {
	switch G {
	case MALE:
		return "male"
	case FEMALE:
		return "female"
	default:
		return "Others"
	}
}