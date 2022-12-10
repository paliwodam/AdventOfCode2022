package main

import (
	"fmt"
	"os"
	"bufio"
)

func getPriority(item uint8)int {
	if item >= 'A' && item <= 'Z' {
		return int(item) - 38
	}
	if item >= 'a'&& item <= 'z' {
		return int(item) - 96
	}
	return 0
}

func part1()int{
	var priorities int
	
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
		backpack := make(map[uint8]int)
        line := fileScanner.Text()
		
		for i := 0; i < (len(line) / 2); i++ {
			backpack[line[i]] = 1
		}
		for i := (len(line) / 2); i < len(line); i++ {item := line[i] 
			_, in := backpack[item]
			if in {
				priorities += getPriority(item)
				break
			}
		}
    }

    readFile.Close()
	return priorities
}


func part2()int{
	var priorities int
	var backpacks [3]map[uint8]bool
	for i := 0; i < 3; i++{
		backpacks[i] = make(map[uint8]bool)
	}
	
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	k := 0
    for fileScanner.Scan() {
        line := fileScanner.Text()

		for i := 0; i < len(line); i++ {
			backpacks[k % 3][line[i]] = true
		}
		
		k += 1
		if k % 3 == 0 {
			for i := uint8('a'); i <= uint8('z'); i++{
				_, in1 := backpacks[0][i] 
				_, in2 := backpacks[1][i] 
				_, in3 := backpacks[2][i]
				if in1 && in2 && in3 {
					priorities += getPriority(i)
				}
			}

			for i := uint8('A'); i <= uint8('Z'); i++{
				_, in1 := backpacks[0][i] 
				_, in2 := backpacks[1][i] 
				_, in3 := backpacks[2][i]
				if in1 && in2 && in3 {
					priorities += getPriority(i)
				}
			}

			for i := 0; i < 3; i++{
				backpacks[i] = make(map[uint8]bool)
			}
		}
    }

    readFile.Close()
	return priorities
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
