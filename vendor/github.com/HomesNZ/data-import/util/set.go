package util

// Set is simply a map of keys to empty structs, which allows you check if a key exists in a set. Each key is mapped to
// an empty struct to avoid unnecessary memory allocation.
type Set map[string]struct{}

// NewSet initializes a new Set.
func NewSet(keys ...string) Set {
	s := Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

// Contains returns true if the key is present in the set.
func (s Set) Contains(key string) bool {
	_, ok := s[key]
	return ok
}
