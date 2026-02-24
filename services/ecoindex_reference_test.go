package services

import "testing"

func TestGradeColorMapContainsExpectedGrades(t *testing.T) {
	expected := map[string]string{
		"A": "#349A47",
		"B": "#51B84B",
		"C": "#CADB2A",
		"D": "#F6EB15",
		"E": "#FECD06",
		"F": "#F99839",
		"G": "#ED2124",
	}

	for grade, expectedColor := range expected {
		color, ok := gradeColorMap[grade]
		if !ok {
			t.Fatalf("grade %q not found in gradeColorMap", grade)
		}
		if color != expectedColor {
			t.Fatalf("unexpected color for grade %q: got %q, want %q", grade, color, expectedColor)
		}
	}
}

func TestGetColorKnownAndUnknownGrades(t *testing.T) {
	if got := GetColor("A"); got != "#349A47" {
		t.Fatalf("GetColor(\"A\") = %q, want %q", got, "#349A47")
	}

	if got := GetColor("Z"); got != "lightgrey" {
		t.Fatalf("GetColor(\"Z\") = %q, want %q", got, "lightgrey")
	}
}

