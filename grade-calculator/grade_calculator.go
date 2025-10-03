package grade_calculator

type number = float64

type Category int

const (
	Assignment Category = iota
	Exam
	Essay
)

type entry struct {
	cat   Category
	score int
}

type GradeCalculator struct {
	items    []entry
	passFail bool
}

// New blank calc
func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{items: []entry{}}
}

// signature: (name string, score number, category Category)
func (g *GradeCalculator) AddGrade(name string, score number, category Category) {
	g.items = append(g.items, entry{cat: category, score: int(score)})
}

func (g *GradeCalculator) Add(cat Category, score int) {
	g.items = append(g.items, entry{cat: cat, score: score})
}

func avg(xs []int) float64 {
	if len(xs) == 0 {
		return 0.0
	}
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return float64(sum) / float64(len(xs))
}

// assignments 50; exams 35; essays 15
func (g *GradeCalculator) FinalNumeric() float64 {
	var a, e, s []int
	for _, it := range g.items {
		switch it.cat {
		case Assignment:
			a = append(a, it.score)
		case Exam:
			e = append(e, it.score)
		case Essay:
			s = append(s, it.score)
		}
	}
	return avg(a)*0.50 + avg(e)*0.35 + avg(s)*0.15
}

// A/B/C/D/F or pPass/fail
func (g *GradeCalculator) GetFinalGrade() string {
	score := g.FinalNumeric()
	if g.passFail {
		if score >= 70 {
			return "Pass"
		}
		return "Fail"
	}
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// wrappers
func (g *GradeCalculator) GetGrade() string { return g.GetFinalGrade() }

//  pass/fail switch
func (g *GradeCalculator) SetPassFail(enabled bool) { g.passFail = enabled }
