package main

import (
	"fmt"
)

// brute force implementation
func editDistance(a, b string) int {
	var subCost int
	aLen := len(a)
	bLen := len(b)
	if aLen == 0 {
		return bLen
	} else if bLen == 0 {
		return aLen
	}

	if a[aLen-1:] != b[bLen-1:] {
		subCost = 1
	} else {
		subCost = 0
	}
	subCost = subCost + editDistance(a[0:aLen-1], b[0:bLen-1])
	dropCost := 1 + editDistance(a, b[0:bLen-1])
	addCost := 1 + editDistance(a[0:aLen-1], b)
	min := subCost
	for _, v := range []int{dropCost, addCost} {
		if v < min {
			min = v
		}
	}

	return min
}

// dynamic implementation
const (
	MATCH  int = 0
	DELETE int = 1
	INSERT int = 2
)

// Make first entry in each row the cost of
// transforming a into the empty string, which is in 0,0
func initRow(editTable [][]Cell, a string) {
	for i := 0; i < len(a)+1; i++ {
		editTable[i][0] = NewCell(len(a[0:i]), INSERT)
	}
}

// Make first entry in each column the cost of
// transforming b into the empty string, which is in 0,0
func initColumn(editTable [][]Cell, b string) {
	for i := 0; i < len(b)+1; i++ {
		editTable[0][i] = NewCell(len(b[0:i]), DELETE)
	}
}

// Constructs string comparison table
// e.g. for "brine" vs "bone"
// <> | "" | b | o | n | e
// "" | 0  | 1 | 2 | 3 | 4
// b  | 1  | 0 | 1 | 2 | 3
// r  | 2  | 1 | 1 | 2 | 3
// i  | 3  | 2 | 2 | 2 | 3
// n  | 4  | 3 | 3 | 2 | 3
// e  | 5  | 4 | 4 | 3 | 2
func dynamicEditDistance(a, b string) int {
	opts := make([]Cell, 3)
	editTable := NewCellMatrix(len(a)+1, len(b)+1)
	initRow(editTable, a)
	initColumn(editTable, b)
	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < len(b)+1; j++ {
			if a[i-1] != b[j-1] {
				opts[MATCH] = NewCell(editTable[i-1][j-1].Cost+1, MATCH)
			} else {
				opts[MATCH] = NewCell(editTable[i-1][j-1].Cost, MATCH)
			}
			opts[DELETE] = NewCell(editTable[i][j-1].Cost+1, DELETE)
			opts[INSERT] = NewCell(editTable[i-1][j].Cost+1, INSERT)
			min := opts[MATCH]
			for _, cell := range opts {
				if cell.Cost < min.Cost {
					min = cell
				}
			}
			editTable[i][j] = min
		}
	}
	//PrintCellMatrix(editTable)
	return editTable[len(a)][len(b)].Cost
}

// helper data structure
// Cost: For a cell at (i,j), how much does it cost to transform a[0:i+1] into b[0:j+1]
// ParentOp: Which operation put me in this state? This is useful for figuring out the path to a given state
type Cell struct {
	ParentOp int
	Cost     int
}

func NewCell(cost, parent int) Cell {
	return Cell{
		ParentOp: parent,
		Cost:     cost,
	}
}

func NewCellMatrix(numRows, numCols int) [][]Cell {
	table := make([][]Cell, numRows)
	for i := 0; i < numRows; i++ {
		table[i] = make([]Cell, numCols)
	}
	return table
}

// for debugging
func PrintCellMatrix(mat [][]Cell) {
	for _, row := range mat {
		for _, val := range row {
			fmt.Printf("%d\t", val.Cost)
		}
		fmt.Println("")
		for k := 0; k < len(row)*len(mat); k++ {
			fmt.Printf("-")
		}
		fmt.Println("")
	}
}

func main() {
	testCases := [][]string{
		[]string{"abc", "ab"},         // 'abc', 'ab' -> 1 because of 1 insertion OR deletion
		[]string{"apple", "adam"},     // 'apple', 'adam' -> 4 because drop 1 and modify 3
		[]string{"camel", "dog"},      // 'camel', 'dog' -> 5 because drop OR add 2 and then change 3
		[]string{"a", "a"},            // 'a', 'a' -> 0
		[]string{"damn", ""},          // 'damn', '' -> 4
		[]string{"kitten", "sitting"}, // 'kitten', 'sitting' -> 3
		[]string{"brine", "bone"},     // 'brine', 'bone' -> 2
	}
	for _, x := range testCases {
		fmt.Printf("The edit distance between %s and %s is %d\n", x[0], x[1], editDistance(x[0], x[1]))
		fmt.Printf("The dynamic edit distance between %s and %s is %d\n", x[0], x[1], dynamicEditDistance(x[0], x[1]))
		fmt.Println()
	}
	xl := "antidisestablishmentarianism" // don't try these with 'editDistance' unless you have years of your life to spare
	xxl := "pneumonoultramicroscopicsilicovolcanoconiosis"
	fmt.Printf("The dynamic edit distance between %s and %s is %d\n", xl, xxl, dynamicEditDistance(xl, xxl))
}
