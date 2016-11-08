package multithread

func NewIdGenerator(start, step int) *idGenerator {
	return &idGenerator{start, step}
}

type idGenerator struct {
	current int
	step    int
}

func (ig *idGenerator) Current() int {
	return ig.current
}

func (ig *idGenerator) GenerateNext() {
	ig.current += ig.step
}
