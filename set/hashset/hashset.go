package hashset

import "fmt"

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func Of[T comparable](vals ...T) *Set[T] {
	s := New[T]()
	s.Add(vals...)
	return s
}

func (s *Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s.m[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(vals ...T) {
	for _, v := range vals {
		delete(s.m, v)
	}
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *Set[T]) Has(val T) bool {
	_, ok := s.m[val]
	return ok
}

func (s *Set[T]) HasAny(vals ...T) bool {
	for _, v := range vals {
		if s.Has(v) {
			return true
		}
	}
	return false
}

func (s *Set[T]) HasAll(vals ...T) bool {
	for _, v := range vals {
		if !s.Has(v) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Union(others ...*Set[T]) *Set[T] {
	c := s.Clone()
	for _, o := range others {
		c.Add(o.Members()...)
	}
	return c
}

func (s *Set[T]) Intersection(others ...*Set[T]) *Set[T] {
	c := s.Clone()
	for _, o := range others {
		for v := range c.m {
			if !o.Has(v) {
				c.Remove(v)
			}
		}
	}
	return c
}

func (s *Set[T]) Difference(others ...*Set[T]) *Set[T] {
	c := s.Clone()
	for _, o := range others {
		c.Remove(o.Members()...)
	}
	return c
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	c := s.Clone()
	for v := range other.m {
		if s.Has(v) {
			c.Remove(v)
		} else {
			c.Add(v)
		}
	}
	return c
}

func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	return s.Intersection(other).IsEmpty()
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for v := range s.m {
		if !s.Has(v) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsProperSubset(other *Set[T]) bool {
	return s.IsSubset(other) && (s.Len() < other.Len())
}

func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	for v := range other.m {
		if !s.Has(v) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsProperSuperset(other *Set[T]) bool {
	return s.IsSuperset(other) && (s.Len() > other.Len())
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set[T]) Members() []T {
	r := make([]T, 0, s.Len())
	for v := range s.m {
		r = append(r, v)
	}
	return r
}

func (s *Set[T]) String() string {
	return fmt.Sprintf("set%v", s.Members())
}

func (s *Set[T]) Clone() *Set[T] {
	return Of[T](s.Members()...)
}
