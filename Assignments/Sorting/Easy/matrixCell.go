// Matrix cell in distance order
// https://leetcode.com/problems/matrix-cells-in-distance-order/description/

package main

import (
	"fmt"
	"math"
	"sort"
)

func allCellsDistOrder(rows, cols, rCenter, cCenter int) [][]int {
	cells := [][]int{}

	// generate all coordinates
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cells = append(cells, []int{i, j})
		}
	}

	// sort the manhattan distance
	sort.Slice(cells, func(a, b int) bool {
		distA := int(math.Abs(float64(cells[a][0]-rCenter)) + math.Abs(float64(cells[a][1]-cCenter)))
		distB := int(math.Abs(float64(cells[b][0]-rCenter)) + math.Abs(float64(cells[b][1]-cCenter)))
		return distA < distB
	})

	return cells
}

func main() {
	rows := 1
	cols := 2
	rCenter := 0
	cCenter := 0
	fmt.Println(allCellsDistOrder(rows, cols, rCenter, cCenter))
}
