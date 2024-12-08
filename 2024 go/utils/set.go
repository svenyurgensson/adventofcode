package utils

type Set[T comparable] map[T]bool

func SetOf[T comparable](elems ...T) Set[T] {
	s := make(Set[T])

	for _, elem := range elems {
		s.Add(elem)
	}

	return s
}


func (s Set[T]) Add(el T) {
	s[el] = true
}

func (s Set[T]) Remove(el T) {
	delete(s, el)
}

func (s Set[T]) Contains(el T) bool {
	_, ok := s[el]
	return ok
}

func (left Set[T]) Intersection(right Set[T]) Set[T] {
	new := make(Set[T])

	for elem := range left {
		if right.Contains(elem) {
			new.Add(elem)
		}
	}

	return new
}

func (left Set[T]) Union(right Set[T]) Set[T] {
	new := make(Set[T])

	for elem := range left {
		new.Add(elem)
	}

	for elem := range right {
		new.Add(elem)
	}

	return new
}