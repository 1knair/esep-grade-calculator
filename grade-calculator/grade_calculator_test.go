package grade_calculator

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 30, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestBoundaries(t *testing.T) {
	makeAll := func(v int) (*GradeCalculator, string) {
		g := NewGradeCalculator()
		g.AddGrade("a", number(v), Assignment)
		g.AddGrade("e", number(v), Exam)
		g.AddGrade("s", number(v), Essay)
		switch {
		case v >= 90:
			return g, "A"
		case v >= 80:
			return g, "B"
		case v >= 70:
			return g, "C"
		case v >= 60:
			return g, "D"
		default:
			return g, "F"
		}
	}

	for _, v := range []int{90, 80, 70, 60, 59} {
		g, expected := makeAll(v)
		if got := g.GetFinalGrade(); got != expected {
			t.Fatalf("at %d expected %s, got %s", v, expected, got)
		}
	}
}

func TestWrappers(t *testing.T) {
	g := NewGradeCalculator()
	g.Add(Assignment, 85)
	g.Add(Exam, 85)
	g.Add(Essay, 85)
	if got := g.GetGrade(); got != "B" {
		t.Fatalf("expected B for all-85 via wrappers, got %s", got)
	}
}

func TestFloatInput(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a", 89.9, Assignment)
	g.AddGrade("e", 89.9, Exam)
	g.AddGrade("s", 89.9, Essay)
	// no rounding up
	if got := g.GetFinalGrade(); got != "B" {
		t.Fatalf("expected B for 89.9 everywhere, got %s", got)
	}
}

func TestWeighting(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a1", number(100), Assignment)
	g.AddGrade("a2", number(100), Assignment)
	out := g.FinalNumeric()
	if out < 49.9 || out > 50.1 {
		t.Fatalf("expected ~50, got %v", out)
	}
}

func TestEmpty(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("only-assignments", number(60), Assignment)
	if got := g.GetFinalGrade(); got != "F" {
		t.Fatalf("expected F when only assignments=60 and others empty, got %s", got)
	}
}
func TestPassMode_Pass(t *testing.T) {
	g := NewGradeCalculator()
	g.SetPassFail(true)
	
		g.AddGrade("a", 100, Assignment) 
	g.AddGrade("e", 100, Exam)
	if got := g.GetFinalGrade(); got != "Pass" {
		t.Fatalf("expected Pass, got %s", got)
	}
}

func TestPassMode_Fail(t *testing.T) {
	g := NewGradeCalculator()
	g.SetPassFail(true)
	g.AddGrade("a", 50, Assignment)
	g.AddGrade("e", 40, Exam)
	g.AddGrade("s", 40, Essay)
	if got := g.GetFinalGrade(); got != "Fail" {
		t.Fatalf("expected Fail, got %s", got)
	}
}

