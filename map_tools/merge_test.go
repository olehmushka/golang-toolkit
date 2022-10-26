package maptools

import "testing"

func TestMerge(t *testing.T) {
	tCases := []struct {
		name string
		in   []map[string]float64
		out  map[string]float64
	}{
		{
			name: "",
			in: []map[string]float64{
				{
					"a": 0.1,
					"b": 0.2,
					"c": 0.3,
					"d": 0.4,
				},
				{
					"c": 0.9,
					"d": 0.4,
					"e": 0.5,
					"f": 0.6,
				},
			},
			out: map[string]float64{
				"a": 0.1,
				"b": 0.2,
				"c": 0.6,
				"d": 0.4,
				"e": 0.5,
				"f": 0.6,
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := Merge(tc.in...); !Equal(out, tc.out) {
				tt.Errorf("expected = %+v, actual = %+v", tc.out, out)
			}
		})
	}
}
