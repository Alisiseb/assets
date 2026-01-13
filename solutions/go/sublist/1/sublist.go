package sublist

// Relation type is defined in relations.go file.
func Sublist(l1, l2 []int) Relation {
	ll1, ll2 := len(l1), len(l2)
	switch {
	case ll1 == ll2:
		if equal(l1, l2) {
			return RelationEqual
		}
		return RelationUnequal
	case ll1 == 0:
		return RelationSublist
	case ll2 == 0:
		return RelationSuperlist
	}

	// main logic
	if ll1 < ll2 {
		if hasSublist(l2, l1) {
			return RelationSublist
		}
		return RelationUnequal
	} else { // ll1 > ll2
		if hasSublist(l1, l2) {
			return RelationSuperlist
		}
		return RelationUnequal
	}
}

// بالای این قطعه تابع equal رو ایمن‌تر کن:
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func hasSublist(a, b []int) bool {
	n, m := len(a), len(b)
	if m == 0 {
		return true
	}
	for i := 0; i <= n-m; i++ {
		if equal(a[i:i+m], b) {
			return true
		}
	}
	return false
}
