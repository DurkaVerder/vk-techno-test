package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readIntArray() []int {
	parts := strings.Fields(readLine())
	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("Invalid data: %s", err.Error())
		}
		result[i] = num
	}
	return result
}

func readMatrix(rows int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = readIntArray()
	}
	return matrix
}

func readSize() (int, int) {
	nums := readIntArray()
	rows, cols := nums[0], nums[1]
	return rows, cols
}

func readStartAndFinish() ([]int, []int) {
	nums := readIntArray()
	start := []int{nums[0], nums[1]}
	finish := []int{nums[2], nums[3]}
	return start, finish
}

func findShortPath(rows, cols int, matrix [][]int, start, finish []int) [][]int {

	



	return [][]int{}
}

func main() {
	rows, cols := readSize()
	matrix := readMatrix(rows)
	start, finish := readStartAndFinish()

	shortPath := findShortPath(rows, cols, matrix, start, finish)
}
