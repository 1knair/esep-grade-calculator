package grade_calculator

type number = float64

type Category int

const (
	Assignment Category = iota
	Exam
	Essay
)

type GradeCalculator struct {
	assignments []int
	exams       []int
	essays      []int
}

// New blank calc
func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: []int{},
		exams:       []int{},
		essays:      []int{},
	}
}

// signature: (name string, score number, category Category)
func (g *GradeCalculator) AddGrade(name string, score number, category Category) {
	iv := int(score)
	switch category {
	case Assignment:
		g.assignments = append(g.assignments, iv)
	case Exam:
		g.exams = append(g.exams, iv)
	case Essay:
		g.essays = append(g.essays, iv)
	}
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
	a := avg(g.assignments)
	e := avg(g.exams)
	s := avg(g.essays)
	return a*0.50 + e*0.35 + s*0.15
}

// A/B/C/D/F
func (g *GradeCalculator) GetFinalGrade() string {
	score := g.FinalNumeric()
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

// wrappers (used in one test)
func (g *GradeCalculator) Add(cat Category, score int) {
	switch cat {
	case Assignment:
		g.assignments = append(g.assignments, score)
	case Exam:
		g.exams = append(g.exams, score)
	case Essay:
		g.essays = append(g.essays, score)
	}
}

func (g *GradeCalculator) GetGrade() string {
	return g.GetFinalGrade()
}
