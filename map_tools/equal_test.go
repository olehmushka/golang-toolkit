package maptools

import "testing"

func TestEqual(t *testing.T) {
	tCases := []struct {
		name string
		in   []map[string]float64
		out  bool
	}{
		{
			name: "should return true",
			in: []map[string]float64{
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
			},
			out: true,
		},
		{
			name: "should return false",
			in: []map[string]float64{
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
				},
			},
			out: false,
		},
		{
			name: "should return false",
			in: []map[string]float64{
				{
					"a": 0.22,
					"b": 0.32,
					"c": 0.12,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
				},
			},
			out: false,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := Equal(tc.in...); out != tc.out {
				tt.Errorf("expected = %t, actual = %t (in = %+v)", tc.out, out, tc.in)
			}
		})
	}
}
