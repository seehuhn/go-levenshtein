package levenshtein

// Distance computes the Levenshtein distance between two strings.
// This is the minimum number of single-rune edits (insertions,
// deletions or substitutions) required to change s into t.
func Distance(s, t string) int {
	ss := []rune(s)
	tt := []rune(t)
	if len(ss) < len(tt) {
		ss, tt = tt, ss
	}

	n := len(tt)
	v0 := make([]int, n+1)
	v1 := make([]int, n+1)
	for j := range v0 {
		v0[j] = j
	}

	for i, c := range ss {
		v1[0] = i + 1
		for j, d := range tt {
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1
			var substitutionCost int
			if c == d {
				substitutionCost = v0[j]
			} else {
				substitutionCost = v0[j] + 1
			}
			v1[j+1] = min3(deletionCost, insertionCost, substitutionCost)
		}
		v0, v1 = v1, v0
	}
	return v0[n]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else if b < c {
		return b
	}
	return c
}
