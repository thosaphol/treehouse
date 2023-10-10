package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	if args[1] == "" {
		fmt.Println("first argument is empty")
		return
	}

	matrix, foundFlags := createMatrix(args[1])
	if len(matrix) < 3 {
		fmt.Println("row count more than 3")
		return
	}

	if len(matrix[0]) < 3 {
		fmt.Println("column count more than 3")
		return
	}

	treeCount := 0

	rowLen := len(matrix)
	colLen := len(matrix[0])

	for i := 1; i < rowLen-1; i++ {
		columnLen := len(matrix[i])
		maxStart := 0
		//find max from left(start) column
		for j := 1; j < columnLen-1; j++ {
			if foundFlags[i][j] {
				continue
			}

			currentTree := matrix[i][j]

			treePrevius := matrix[i][j-1]
			if treePrevius > maxStart {
				maxStart = treePrevius
			}

			if maxStart < currentTree {
				treeCount++
				foundFlags[i][j] = true
			}
		}

		maxEnd := 0

		for j := columnLen - 2; j > 0; j-- {
			if foundFlags[i][j] {
				continue
			}
			currentTree := matrix[i][j]
			treeNext := matrix[i][j+1]
			if treeNext > maxEnd {
				maxEnd = treeNext
			}

			if maxEnd < currentTree {
				treeCount++
				foundFlags[i][j] = true
			}
		}
	}

	for i := 1; i < colLen-1; i++ {
		rowLen = len(matrix)

		//find max from start(left column or top row)
		maxStart := 0
		for j := 1; j < rowLen-1; j++ {

			currentTree := matrix[j][i]

			treePrevius := matrix[j-1][i]
			if treePrevius > maxStart {
				maxStart = treePrevius
			}

			if !foundFlags[j][i] && maxStart < currentTree {
				treeCount++
				foundFlags[j][i] = true
			}
		}

		//find max from end(right column or bottom row)
		maxEnd := 0
		for j := rowLen - 2; j > 0; j-- {

			currentTree := matrix[j][i]
			treeNext := matrix[j+1][i]
			if treeNext > maxEnd {
				maxEnd = treeNext
			}

			if !foundFlags[j][i] && maxEnd < currentTree {
				treeCount++
				foundFlags[j][i] = true
			}
		}
	}
	fmt.Println(treeCount + rowLen*2 + (colLen-2)*2)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createMatrix(path string) ([][]int, [][]bool) {
	mFile, err := os.Open(path)
	check(err)
	defer mFile.Close()

	var forests [][]int
	var foundFlags [][]bool

	scanner := bufio.NewScanner(mFile) //scan the contents of a file and print line by line
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "")
		rowLen := len(lineSplit)
		fRow := make([]bool, rowLen)
		var rows []int
		for i := 0; i < rowLen; i++ {
			tree, err := strconv.Atoi(lineSplit[i])
			check(err)
			rows = append(rows, tree)
		}

		forests = append(forests, rows)
		foundFlags = append(foundFlags, fRow)
	}

	err = scanner.Err()
	check(err)

	return forests, foundFlags
}
