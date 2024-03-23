package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func TestMain(t *testing.T) {
	t.Run("Main", func(t *testing.T) {
		// requester := req

		main()
	})
}*/

func TestFactorial(t *testing.T) {
	tests := []struct {
		Name     string
		Num      int
		Expected int
	}{
		{
			Name:     "0!",
			Num:      0,
			Expected: 1,
		},
		{
			Name:     "1!",
			Num:      1,
			Expected: 1,
		},
		{
			Name:     "5!",
			Num:      5,
			Expected: 120,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			expected := Fact(test.Num)
			assert.Equal(t, test.Expected, expected)
		})
	}
}

func BenchmarkFact(b *testing.B) {
	tests := []struct {
		Name     string
		Num      int
		Expected int
	}{
		{
			Name:     "0!",
			Num:      0,
			Expected: 1,
		},
		{
			Name:     "1!",
			Num:      1,
			Expected: 1,
		},
		{
			Name:     "5!",
			Num:      5,
			Expected: 120,
		},
	}

	for _, test := range tests {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fact(test.Num)
			}
		})
	}
}
