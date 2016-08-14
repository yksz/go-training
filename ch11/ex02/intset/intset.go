package intset

type Set interface {
	// Has reports whether the set contains the non-negative value x.
	Has(x int) bool
	// Add adds the non-negative value x to the set.
	Add(x int)
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

type MapIntSet struct {
	set map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{make(map[int]bool)}
}

func (s *MapIntSet) Has(x int) bool {
	_, ok := s.set[x]
	return ok
}

func (s *MapIntSet) Add(x int) {
	s.set[x] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for k := range t.set {
		s.set[k] = true
	}
}
