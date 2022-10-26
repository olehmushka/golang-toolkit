package numbertools

import "testing"

func TestRoundFloat64(t *testing.T) {
	tCases := []struct {
		name string
		in   float64
		dec  int
		out  float64
	}{
		{
			name: "10.123456 -> 10.12",
			in:   10.123456,
			dec:  2,
			out:  10.12,
		},
		{
			name: "101.23456 -> 101.235",
			in:   101.23456,
			dec:  3,
			out:  101.235,
		},
		{
			name: "1012.3456 -> 1012.3456",
			in:   1012.3456,
			dec:  4,
			out:  1012.3456,
		},
		{
			name: "10.123456 -> 10.1",
			in:   10.123456,
			dec:  1,
			out:  10.1,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := RoundFloat64(tc.in, tc.dec); out != tc.out {
				t.Errorf("unexpected out (expected = %f, actual = %f)", tc.out, out)
			}
		})
	}
}
