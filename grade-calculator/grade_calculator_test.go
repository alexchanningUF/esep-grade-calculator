package gradecalculator

import "testing"

func TestAvg(t *testing.T) {
	if got := avg([]float64{}); got != 0 {
		t.Fatalf("empty avg = %v, want 0", got)
	}
	if got := avg([]float64{100, 50}); got != 75 {
		t.Fatalf("avg = %v, want 75", got)
	}
}

func TestFinalNumeric_WeightedByCategory(t *testing.T) {
	g := New([]float64{100, 80}, []float64{90, 70}, []float64{60})
	// avgs: A=90, E=80, S=60 -> 90*.5 + 80*.35 + 60*.15 = 45 + 28 + 9 = 82
	if got := g.FinalNumeric(); got != 82 {
		t.Fatalf("FinalNumeric = %v, want 82", got)
	}
}

func TestLetter_Cutoffs(t *testing.T) {
	g := New(nil, nil, nil)
	cases := []struct {
		in   float64
		want string
	}{
		{90, "A"},
		{89.9999, "B"},
		{80, "B"},
		{79.9999, "C"},
		{70, "C"},
		{69.9999, "D"},
		{60, "D"},
		{59.9999, "F"},
	}
	for _, tc := range cases {
		if got := g.Letter(tc.in); got != tc.want {
			t.Fatalf("Letter(%v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestPass_IsCOrHigher(t *testing.T) {
	g := New(nil, nil, nil)
	if g.Pass(70) != true {
		t.Fatalf("70 should pass (C)")
	}
	if g.Pass(69.9999) != false {
		t.Fatalf("69.9999 should fail (D)")
	}
}

func TestSpecExample_EndToEnd(t *testing.T) {
	g := New([]float64{100, 90, 80}, []float64{85, 95}, []float64{70})
	num := g.FinalNumeric()
	if num <= 0 || num > 100 {
		t.Fatalf("FinalNumeric out of range: %v", num)
	}
	if g.Letter(num) == "" {
		t.Fatalf("Letter should not be empty")
	}
}
