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
		{[]int{0, -203, 1, 2, 3}, true, []int{0, 1}},
		{[]int{1, 1, 2, 3}, true, []int{0, 1}},
		{[]int{1, 2, 3, 3, 4}, true, []int{2, 3}},
		{[]int{1, 2, 3, 3}, true, []int{2, 3}},
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
		report    []int
		want      bool
		wantIndex int
	}{
		{[]int{1, 2, 7, 8, 9}, false, -1},
		{[]int{1, 3, 2, 4, 5}, true, 1},
		{[]int{1, 3, 6, 7, 9}, true, -1},
		{[]int{17, 19, 20, 21, 23, 24, 29, 27}, true, 6},
		{[]int{25, 27, 32, 35, 35}, false, -1},
		{[]int{29, 31, 32, 36, 43}, false, -1},
		{[]int{37, 40, 44, 46, 49, 50}, false, -1},
		{[]int{47, 48, 49, 50, 54, 57, 57}, false, -1},
		{[]int{61, 63, 67, 69, 72, 70}, false, -1},
		{[]int{7, 6, 4, 2, 1}, true, -1},
		{[]int{70, 73, 76, 77, 77, 78, 85}, false, -1},
		{[]int{72, 75, 79, 80, 82, 86}, false, -1},
		{[]int{77, 80, 81, 84, 85, 86, 86, 90}, false, -1},
		{[]int{8, 6, 4, 4, 1}, true, 2},
		{[]int{83, 84, 91, 92, 94, 97}, false, -1},
		{[]int{86, 87, 89, 89, 92, 92}, false, -1},
		{[]int{9, 7, 6, 2, 1}, false, -1},
	}

	for _, c := range cases {
		isSafe, _ := IsReportSafeWithDampener(c.report)
		if isSafe != c.want {
			t.Fatalf("expected %v for %v", c.want, c.report)
		}
	}
}
