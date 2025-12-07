package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, elem := range l {

		s[elem] = struct{}{}
	}
	return s
}

func (s Set) String() string {
	result := "{"
	switch len(s) {
	case 0:
		return result + "}"
	default:
		counter := 0
		for elem := range s {
			counter++
			if counter == len(s) {
				result += "\"" + elem + "\""
			} else {
				result += "\"" + elem + "\", "
			}
		}
	}
	result += "}"
	return result
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	for key := range s {
		if key == elem {
			return true
		}
	}
	return false
}

func (s Set) Add(elem string) {
	if !s.Has(elem) {
		s[elem] = struct{}{}
	}
}

func Subset(s1, s2 Set) bool {
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for elem := range s1 {
		if s2.Has(elem) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	for elem := range s2 {
		if !s1.Has(elem) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		if s2.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		if !s2.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		result.Add(elem)
	}
	for elem := range s2 {
		if !result.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}
