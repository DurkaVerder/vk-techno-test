package solution_test

import (
	"testing"
	"vk-techno-test/solution"
)

func comparePaths(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}

func TestFindShortPath_technoTest(t *testing.T) {
	rows := 3
	cols := 3
	matrix := [][]int{
		{1, 2, 0},
		{2, 0, 1},
		{9, 1, 0},
	}
	start := solution.Point{0, 0}
	finish := solution.Point{2, 1}
	expected := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
	}

	result := solution.FindShortPath(rows, cols, matrix, start, finish)
	if !comparePaths(result, expected) {
		t.Errorf("findShortPath() = %v, expected %v", result, expected)
	}
}

func TestFindShortPath_SimplePath(t *testing.T) {
	rows := 3
	cols := 3
	matrix := [][]int{
		{1, 3, 1},
		{1, 0, 1},
		{4, 2, 1},
	}
	start := solution.Point{0, 0}
	finish := solution.Point{2, 2}
	expected := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}

	result := solution.FindShortPath(rows, cols, matrix, start, finish)
	if !comparePaths(result, expected) {
		t.Errorf("findShortPath() = %v, expected %v", result, expected)
	}
}

func TestFindShortPath_NoPathDueToWalls(t *testing.T) {
	rows := 3
	cols := 3
	matrix := [][]int{
		{1, 0, 1},
		{1, 0, 1},
		{1, 0, 1},
	}
	start := solution.Point{0, 0}
	finish := solution.Point{2, 2}
	expected := [][]int{}

	result := solution.FindShortPath(rows, cols, matrix, start, finish)
	if !comparePaths(result, expected) {
		t.Errorf("findShortPath() = %v, expected %v", result, expected)
	}
}

func TestFindShortPath_StartEqualsFinish(t *testing.T) {
	rows := 3
	cols := 3
	matrix := [][]int{
		{1, 3, 1},
		{1, 0, 1},
		{4, 2, 1},
	}
	start := solution.Point{0, 0}
	finish := solution.Point{0, 0}
	expected := [][]int{
		{0, 0},
	}

	result := solution.FindShortPath(rows, cols, matrix, start, finish)
	if !comparePaths(result, expected) {
		t.Errorf("findShortPath() = %v, expected %v", result, expected)
	}
}

func TestFindShortPath_ComplexMaze(t *testing.T) {
	rows := 6
	cols := 6
	matrix := [][]int{
		{1, 0, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 1},
		{1, 1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 1},
	}
	start := solution.Point{0, 0}
	finish := solution.Point{5, 5}
	expected := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
		{4, 3},
		{4, 4},
		{4, 5},
		{5, 5},
	}

	result := solution.FindShortPath(rows, cols, matrix, start, finish)
	if !comparePaths(result, expected) {
		t.Errorf("findShortPath() = %v, expected %v", result, expected)
	}
}
