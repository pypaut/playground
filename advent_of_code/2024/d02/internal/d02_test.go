package d02

import (
	"reflect"
	"testing"
)

func TestParseReports(t *testing.T) {
	reports, err := ParseReports("../input_test")
	if err != nil {
		t.Fatal(err)
	}

	expectedReports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	if !reflect.DeepEqual(reports, expectedReports) {
		t.Fatalf("expected %v, got %v", expectedReports, reports)
	}
}

func TestIsStrictlyMonotonous(t *testing.T) {
	cases := []struct {
		report []int
		want   bool
	}{
		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 3, 3}, false},
		{[]int{1, 1, 2, 3}, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9321093}, true},
		{[]int{1, 2, 3, 3, 4}, false},
		{[]int{-203, 1, 2, 3}, true},
		{[]int{1}, true},
	}

	for _, c := range cases {
		if IsStrictlyMonotonous(c.report) != c.want {
			t.Fatalf("expected %v for %v", c.want, c.report)
		}
	}
}

func TestIsStrictlyMonotonousWithDampener(t *testing.T) {
	cases := []struct {
		report  []int
		want    bool
		indices []int
	}{
		{[]int{-203, 1, 2, 3, 2}, true, []int{4}},
		{[]int{-203, 1, 2, 3}, true, nil},
		{[]int{0, -203, 1, 2, 3}, true, []int{0}},
		{[]int{1, 1, 2, 3}, true, []int{0}},
		{[]int{1, 2, 3, 3, 4}, true, []int{2}},
		{[]int{1, 2, 3, 3}, true, []int{2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9321093}, true, nil},
		{[]int{1, 2, 3}, true, nil},
		{[]int{1}, true, nil},
	}

	for _, c := range cases {
		isMonotonous, indices := IsStrictlyMonotonousWithDampener(c.report)
		if isMonotonous != c.want {
			t.Fatalf("expected %v for %v", c.want, c.report)
		}

		if !reflect.DeepEqual(indices, c.indices) {
			t.Fatalf("expected %v, got %v for %v", c.indices, indices, c.report)
		}
	}
}

func TestDiffersByMaxThree(t *testing.T) {
	cases := []struct {
		report []int
		want   bool
	}{
		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 3, 3}, true},
		{[]int{1, 1, 2, 3}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9321093}, false},
		{[]int{1, 2, 3, 3, 4}, true},
		{[]int{-203, 1, 2, 3}, false},
		{[]int{1}, true},
		{[]int{1, 4}, true},
		{[]int{1, 5}, false},
	}

	for _, c := range cases {
		if DiffersByMaxThree(c.report) != c.want {
			t.Fatalf("expected %v for %v", c.want, c.report)
		}
	}
}

func TestIsReportSafe(t *testing.T) {
	cases := []struct {
		report []int
		want   bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, c := range cases {
		if IsReportSafe(c.report) != c.want {
			t.Fatalf("expected %v for %v", c.want, c.report)
		}
	}
}

func TestNumberOfSafeReports(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	want := 2
	got := NumberOfSafeReports(reports)

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestIsReportSafeWithDampener(t *testing.T) {
	cases := []struct {
		report []int
		want   bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, c := range cases {
		isSafe, _ := IsReportSafeWithDampener(c.report)
		if isSafe != c.want {
			t.Fatalf("expected %v for %v", c.want, c.report)
		}
	}
}
