package multithread

func NewIdGenerator(start, step, max int) *idGenerator {
	return &idGenerator{start, step, max}
}

type idGenerator struct {
	current int
	step    int
	max     int
}

func (ig *idGenerator) Current() int {
	return ig.current
}

func (ig *idGenerator) GenerateNext() bool {
	ig.current += ig.step

	return ig.current <= ig.max
}
