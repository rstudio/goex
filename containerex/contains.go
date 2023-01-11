package containerex

func Contains[T comparable](sl []T, inst T) bool {
	for _, entry := range sl {
		if entry == inst {
			return true
		}
	}

	return false
}

type Eqer[T any] interface {
	Eq(T) bool
}

func ContainsEq[T Eqer[T]](sl []T, inst T) bool {
	for _, entry := range sl {
		if entry.Eq(inst) {
			return true
		}
	}

	return false
}
