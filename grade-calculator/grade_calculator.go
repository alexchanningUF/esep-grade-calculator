package gradecalculator

import "math"

type GradeCalculator struct {
	assignments []float64
	exams       []float64
	essays      []float64
}

func New(assignments, exams, essays []float64) *GradeCalculator {
	return &GradeCalculator{assignments: assignments, exams: exams, essays: essays}
}

func avg(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	var s float64
	for _, x := range xs {
		s += x
	}
	return s / float64(len(xs))
}

// FinalNumeric computes the final numeric grade (0–100).
// Fixed weights per spec: A=0.50, Exams=0.35, Essays=0.15.
func (g *GradeCalculator) FinalNumeric() float64 {
	const wa, we, ws = 0.50, 0.35, 0.15
	a := avg(g.assignments)
	e := avg(g.exams)
	s := avg(g.essays)
	return a*wa + e*we + s*ws
}

// Letter maps a numeric grade to A/B/C/D/F using inclusive lower bounds.
func (g *GradeCalculator) Letter(num float64) string {
	// Compare on a consistent precision to avoid float flukes near cutoffs.
	n := math.Round(num*1000) / 1000
	switch {
	case n >= 90:
		return "A"
	case n >= 80:
		return "B"
	case n >= 70:
		return "C"
	case n >= 60:
		return "D"
	default:
		return "F"
	}
}

func (g *GradeCalculator) Pass(num float64) bool {
	return g.Letter(num) == "A" || g.Letter(num) == "B" || g.Letter(num) == "C"
}
