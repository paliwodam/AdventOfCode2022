package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int)int{
	if x >= 0{
		return x
	}
	return -x
}

func isTouching(hx, hy, tx, ty int)bool{
	return abs(hx - tx) < 1.0 && abs(hy - ty) < 1.0
}

func moveT(hx, hy, tx, ty int)(int, int){
	if hx == tx{
		return tx, ty + (hy - ty) / 2 
	}
	if hy == ty{
		return tx + (hx - tx) / 2, ty
	}
	if abs(hx-tx) == 2{
		return tx + (hx-tx) / 2, hy
	}
	if abs(hy-ty) == 2{
		return hx, ty + (hy-ty) / 2
	}
	return tx, ty
}

func part1()int{
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	hx, hy := 500, 500
	tx, ty := 500, 500

	mapp := make([][]bool, 1000)
	for i := range mapp{
		mapp[i] = make([]bool, 1000)
	}
	mapp[500][500] = true

    for fileScanner.Scan() {
        line := fileScanner.Text()

		move := strings.Split(line, " ")
		direction := move[0]
		times, _ := strconv.Atoi(move[1])

		for i := 0; i < times; i++{
			switch direction{
			case "R":
				hx += 1
			case "L":
				hx -= 1
			case "U":
				hy += 1
			case "D":
				hy -= 1
			}
			if !isTouching(hx, hy, tx, ty){
				tx, ty = moveT(hx, hy, tx, ty)
				mapp[ty][tx] = true
			}
		}
    }

	var cnt int
	for i := range mapp{
		for j := range mapp[i]{
			if mapp[i][j]{
				cnt += 1
			}
		}
	}

    readFile.Close()
	return cnt
}

func part2()int{
	return 0
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
