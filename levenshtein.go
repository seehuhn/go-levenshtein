// seehuhn.de/go/levenshtein - compute Levenshtein distances
// Copyright (C) 2018  Jochen Voss <voss@seehuhn.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
