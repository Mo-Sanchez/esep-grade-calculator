package esepunittests

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
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}


}

func TestGradeTypeString(t *testing.T){
		if Assignment.String() != "assignment"{
			t.Errorf("expected assignment")
		}

		if Exam.String() != "exam"{
			t.Errorf("expected exam")
		}

		if Essay.String() != "essay"{
			t.Errorf("expected essay")
		}


	}


func TestGetGradeEmptyReturnsF(t *testing.T) {
	expected := "F"
	if got := NewGradeCalculator().GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeC(t *testing.T) {
	expected := "C"
	g := NewGradeCalculator()
	g.AddGrade("assignment", 70, Assignment)
	g.AddGrade("exam", 70, Exam)
	g.AddGrade("essay", 70, Essay)
	if got := g.GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeD(t *testing.T) {
	expected := "D"
	g := NewGradeCalculator()
	g.AddGrade("assignment", 60, Assignment)
	g.AddGrade("exam", 60, Exam)
	g.AddGrade("essay", 60, Essay)
	if got := g.GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}