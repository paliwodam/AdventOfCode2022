package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MaxUintConst = ^uint(0) 
	MinUintConst = 0 
	MaxIntConst = int(MaxUintConst >> 1) 
	MinIntConst = -MaxIntConst - 1
)


func MaxIntSlice(s []int) int {
	N := len(s)
	max := MinIntConst
	
	for i := 0; i < N; i++ {
		if (s[i] > max) { max = s[i] }
	}

	return max
}

func getForest()[][]int{
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	forest := make([][]int, 100)
	for i := range forest{
		forest[i] = make([]int, 100)
	}

	i := 0
    for fileScanner.Scan() {
        line := fileScanner.Text()
		forest[i] = forest[i][:len(line)]
		for j := range line{
			forest[i][j] = int(line[j]) - 48
		}
		i += 1
    }
	forest = forest[:i]

    readFile.Close()
	return forest
}

func getReversedForest(forest [][]int) [][]int{
	rForest := make([][]int, len(forest[0]), 100)
	for i := range rForest{
		rForest[i] = make([]int, len(forest), 100)
	}
	return rForest
}

func part1()int{
	forest := getForest()
	rForest := getReversedForest(forest)

	for i := range forest{
		for j := range forest{
			rForest[j][i] = forest[i][j]
		}
	}

	n := len(forest)
	m := len(rForest)
	visibleTrees := n * 2 + m * 2 - 4
	for i := 1; i < n - 1; i++{
		for j := 1; j < m - 1; j++{
			tree := forest[i][j]
			visibleL := MaxIntSlice(forest[i][:j]) < tree
			visibleR := MaxIntSlice(forest[i][j+1:m]) < tree
			visibleT := MaxIntSlice(rForest[j][:i]) < tree
			visibleB := MaxIntSlice(rForest[j][i+1:n]) < tree

			if visibleL || visibleR || visibleT || visibleB{
				visibleTrees += 1
			}
		}
	}
	return visibleTrees
}

func treesNum(height int, vison []int, direction byte)int{
	var cnt int 

	if direction == 'l' || direction == 't' {
		for i := len(vison) - 1; i >= 0; i--{
			cnt += 1
			if vison[i] >= height{
				break
			}
		}
	} else if direction == 'r' || direction == 'b' {
		for i := 0; i < len(vison); i++{
			cnt += 1
			if vison[i] >= height{
				break
			}
		}
	}

	return cnt
}

func part2()int{
	forest := getForest()
	rForest := getReversedForest(forest)

	for i := range forest{
		for j := range forest{
			rForest[j][i] = forest[i][j]
		}
	}

	n := len(forest)
	m := len(rForest)
	visibleTrees := 0
	for i := 1; i < n - 1; i++{
		for j := 1; j < m - 1; j++{
			tree := forest[i][j]
			visibleL := treesNum(tree, forest[i][:j], 'l')
			visibleR := treesNum(tree, forest[i][j+1:m], 'r')
			visibleT := treesNum(tree, rForest[j][:i], 't')
			visibleB := treesNum(tree, rForest[j][i+1:n], 'b')
			result := visibleL * visibleR * visibleT * visibleB

			if result > visibleTrees {
				visibleTrees = result
			}
		}
	}
	return visibleTrees
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}