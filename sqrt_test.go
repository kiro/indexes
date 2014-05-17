package indexes

import("testing")

func TestSum(t *testing.T) {
	n := 123
	index := NewSqrtIndex(n, Sum)

	for i := 0; i < n; i++ {
		index.AddSqrt(i, 1)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			value := index.Get(i, j)
			expected := j - i
			if value != expected {
				t.Errorf("Value from %d to %d expected to be %d but was %d", i, j, expected, value)
			}
		}
	}
}

func TestMax(t *testing.T) {
	n := 100
	mod := 30
	index := NewSqrtIndex(n, Max)

	for i := 0; i < n; i++ {
		index.AddSqrt(i, i % mod)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			value := index.Get(i, j)

			expected := 0
			for k := i; k < j; k++ {
				expected = Max(expected, k % mod)
			}

			if value != expected {
				t.Errorf("Value from %d to %d expected to be %d but was %d", i, j, expected, value)
			}			
		}
	}
}