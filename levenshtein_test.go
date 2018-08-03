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

import "testing"

func TestMin3(t *testing.T) {
	cases := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, q := range cases {
		if min3(q.a, q.b, q.c) != 1 {
			t.Errorf("min3(%d, %d, %d) failed", q.a, q.b, q.c)
		}
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		a, b string
		dist int
	}{
		{"abc", "abc", 0},
		{"abc", "abcd", 1},
		{"abc", "ac", 1},
		{"abc", "aec", 1},
		{"", "", 0},
		{"", "0123456789", 10},
		{"kitten", "sitting", 3},
		{"flaw", "lawn", 2},
		{"Bär", "Bar", 1},
		{"Bär", "Br", 1},
		{"Bär", "Bäär", 1},
		{"bufniță", "Eule", 6},
	}
	for i, c := range cases {
		d := Distance(c.a, c.b)
		if d != c.dist {
			t.Errorf("wrong dist(%q, %q) in test %d, expected %d, got %d",
				c.a, c.b, i, c.dist, d)
		}
		d = Distance(c.b, c.a)
		if d != c.dist {
			t.Errorf("wrong dist(%q, %q) in test %d, expected %d, got %d",
				c.b, c.a, i, c.dist, d)
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	s1 := "Call me Ishmael."
	s2 := "Some years ago — never mind how long precisely — having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Distance(s1, s2)
	}
}
