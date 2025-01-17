package solution

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Point structure for matrix
type Point struct {
	Row, Col int
}

// Item structure for priority queue
type Item struct {
	point Point
	price int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].price < pq[j].price
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// readLine returns line from console
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// readInt returns int slice from console
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

// readMatrix returns int matrix from console
func readMatrix(rows int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = readIntArray()
	}
	return matrix
}

// readSize returns rows and cols
func readSize() (int, int) {
	nums := readIntArray()
	rows, cols := nums[0], nums[1]
	return rows, cols
}

// readStartAndFinish returns start and finish Points
func readStartAndFinish() (Point, Point) {
	nums := readIntArray()
	if len(nums) < 4 {
		log.Fatalln("Error in write start and finish points")
	}
	start := Point{nums[0], nums[1]}
	finish := Point{nums[2], nums[3]}
	return start, finish
}

// findShortPath returns short path from start to finish using Dijkstra's algorithm
func FindShortPath(rows, cols int, matrix [][]int, start, finish Point) [][]int {
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	parent := make([][]Point, rows)
	for i := range parent {
		parent[i] = make([]Point, cols)
	}

	dist[start.Row][start.Col] = matrix[start.Row][start.Col]
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{point: start, price: dist[start.Row][start.Col]})

	for pq.Len() > 0 {
		currentItem := heap.Pop(&pq).(*Item)
		currentPoint := currentItem.point

		if currentPoint == finish {
			return createPath(parent, start, finish)
		}

		for _, dir := range directions {
			newRow := currentPoint.Row + dir.Row
			newCol := currentPoint.Col + dir.Col

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && matrix[newRow][newCol] != 0 {
				newDist := dist[currentPoint.Row][currentPoint.Col] + matrix[newRow][newCol]

				if newDist < dist[newRow][newCol] {
					dist[newRow][newCol] = newDist
					parent[newRow][newCol] = currentPoint
					heap.Push(&pq, &Item{point: Point{newRow, newCol}, price: newDist})
				}
			}
		}
	}

	return [][]int{}
}

// CreatePath returns the created path from parent
func createPath(parent [][]Point, start, finish Point) [][]int {
	path := [][]int{}
	current := finish

	for current != start {
		path = append(path, []int{current.Row, current.Col})
		current = parent[current.Row][current.Col]
	}

	path = append(path, []int{start.Row, start.Col})

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// printPath print path in console
func printPath(path [][]int) {
	for _, point := range path {
		fmt.Println(point)
	}
	fmt.Println(".")
}

func StartSolution() {
	rows, cols := readSize()
	if rows <= 0 || cols <= 0 {
		log.Fatalf("Invalid matrix size: rows and cols must be positive")
	}
	matrix := readMatrix(rows)

	start, finish := readStartAndFinish()
	if start.Row < 0 || start.Row >= rows || start.Col < 0 || start.Col >= cols {
		log.Fatalf("Start point is out of bounds")
	}
	if finish.Row < 0 || finish.Row >= rows || finish.Col < 0 || finish.Col >= cols {
		log.Fatalf("Finish point is out of bounds")
	}

	shortPath := FindShortPath(rows, cols, matrix, start, finish)
	if len(shortPath) == 0 {
		log.Fatalf("Path not found")
	}
	printPath(shortPath)
}
