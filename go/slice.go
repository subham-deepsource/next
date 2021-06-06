package demo

type Int []int

const NotFound = -1

// Present checks if n is present slice.
func (s Int) Present(n int) bool {
	return s.Find(n) != NotFound
}

// Find tries to find n in slice. If present, it
// returns the position. Otherwise, it returns NotFound.
func (s Int) Find(n int) int {
	for i := range s {
		if s[i] == n {
			return i
		}
	}
	return NotFound
}

// Map applies op function for every element in slice and
// replaces "s[i]" with "op(s[i])" where "i" is the index.
func (s Int) Map(op func(n int) int) {
	for i := range s {
		s[i] = op(s[i])
	}
}

// Capacity returns the capacity of the slice.
func (s Int) Capacity() int {
	return cap(s)
}

// Length returns the length of the slice.
func (s Int) Length() int {
	return len(s)
}

// Empty checks if the slice is empty.
func (s Int) Empty() bool {
	return s.Length() == 0
}

// Front returns the first element of the slice. If the
// slice is empty, then it panics.
func (s Int) Front() int {
	return s[0]
}

// Back returns the last element of the slice. If the
// slice is empty, then it panics.
func (s Int) Back() int {
	return s[s.Length()-1]
}
