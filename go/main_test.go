package main

import "testing"

func TestEgcd(t *testing.T) {
	t.Run("Return the m, a, b", func(t *testing.T) {
		r_a, r_b, r_c := egcd(33, 27)
		x_a, x_b, x_c := 3, -4, 5

		if r_a != x_a {
			t.Errorf("Expected %d, got %d", x_a, r_a)
		}
		if r_b != x_b {
			t.Errorf("Expected %d, got %d", x_b, r_b)
		}
		if r_c != x_c {
			t.Errorf("Expected %d, got %d", x_c, r_c)
		}
	})
}

func TestModInv(t *testing.T) {
	tt := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Test1",
			a:        3,
			b:        7,
			expected: 5,
		},
		{
			name:     "Test2",
			a:        7,
			b:        3,
			expected: 1,
		},
		{
			name:     "Test2",
			a:        2,
			b:        6,
			expected: 0,
		},
		{
			name:     "Test2",
			a:        -7,
			b:        19,
			expected: 8,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := modInv(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}
