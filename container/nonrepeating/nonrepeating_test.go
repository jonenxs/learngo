package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//Normal cases
		{"abcabca", 3},
		{"asdfgfdsa", 5},

		// Edge cases
		{"", 0},
		{"a", 1},
		{"aaa", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"我是谁我是时候", 5},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "我是谁我是时候"
	ans := 5

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d",
				actual, s, ans)
		}
	}
}
