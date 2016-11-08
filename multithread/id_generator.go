package multithread

func NewIdGenerator(start, step int) *idGenerator {
	return &idGenerator{start, step}
}

type idGenerator struct {
	current int
	step    int
}

func (ig *idGenerator) Next() int {
	current := ig.current
	ig.current += ig.step

	return current
}
