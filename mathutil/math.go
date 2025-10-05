package mathutil

type Model interface {
	ID() int
}

func Add(a, b int) int {
	return a + b
}

type Data map[int]Model
