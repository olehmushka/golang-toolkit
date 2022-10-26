package maptools

import "testing"

func TestPickOneByProb(t *testing.T) {
	m1 := map[string]float64{
		"one":   0.7,
		"two":   0.2,
		"three": 0.9,
		"four":  0.8,
		"five":  0.1,
		"six":   1,
	}
	_, err := PickOneByProb(m1)
	if err != nil {
		t.Fatalf("unexpected error = %+v", err)
		return
	}
}
