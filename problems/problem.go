package problems

type Problem interface {
	PrintDescription()
	Solve() string
	GetAnswer() string
}

