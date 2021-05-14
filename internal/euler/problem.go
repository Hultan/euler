package euler

type Problem interface {
	PrintDescription()
	Solve() string
	GetAnswer() string
}

