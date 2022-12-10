package main

import (
	"fmt"
	"os"
	"bufio"
    "strings"
	"strconv"
)

func containsEachOther(x1, x2, y1, y2 int)bool{
	return (x1 >= x2 && y1 <= y2) || (x1 <= x2 && y1 >= y2) 
}

func overlap(x1, x2, y1, y2 int)bool{
	return (x1 <= x2 && x2 <= y1) || (x1 <= y2 && y2 <= y1) || 
		   (x2 <= x1 && x1 <= y2) || (x2 <= y1 && y1 <= y2)
}

func parseLine(line string)(x1, x2, y1, y2 int){
	pairs := strings.Split(line, ",")
	x1, _ = strconv.Atoi(strings.Split(pairs[0], "-")[0])
	y1, _ = strconv.Atoi(strings.Split(pairs[0], "-")[1])
	x2, _ = strconv.Atoi(strings.Split(pairs[1], "-")[0])
	y2, _ = strconv.Atoi(strings.Split(pairs[1], "-")[1])
	return
}

func part1()int{
	var cnt int
	
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line := fileScanner.Text()
		
		x1, x2, y1, y2 := parseLine(line)
		if containsEachOther(x1, x2, y1, y2){
			cnt += 1
		}
    }

    readFile.Close()
	return cnt
}


func part2()int{
	var cnt int
	
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line := fileScanner.Text()
		
		x1, x2, y1, y2 := parseLine(line)
		if overlap(x1, x2, y1, y2){
			cnt += 1
		}
    }

    readFile.Close()
	return cnt
}

func main() {
	// fmt.Println(part1())
	fmt.Println(part2())
}
