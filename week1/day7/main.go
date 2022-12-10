package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxDirSize int = 100000 

func getTotalSize(fileScanner *bufio.Scanner)(totalSize, dirSize int) {
	for fileScanner.Scan(){
		line := fileScanner.Text()
		if line == "$ cd .."{
			break
		} else if line[:4] == "$ cd"{
			subTotalSize, subdirSize := getTotalSize(fileScanner)
			totalSize += subTotalSize
			dirSize += subdirSize
		} else if line != "$ ls" && line[:3] != "dir"{
			fileSize, _ := strconv.Atoi(strings.Split(line, " ")[0])
			dirSize += fileSize
		}
	}
	
	if dirSize <= maxDirSize{
		totalSize += dirSize
	}
	return totalSize, dirSize
}

func getSystemSize(fileScanner *bufio.Scanner)(dirSize int) {
	for fileScanner.Scan(){
		line := fileScanner.Text()
		if line == "$ cd .."{
			break
		} else if line[:4] == "$ cd"{
			subdirSize := getSystemSize(fileScanner)
			dirSize += subdirSize
		} else if line != "$ ls" && line[:3] != "dir"{
			fileSize, _ := strconv.Atoi(strings.Split(line, " ")[0])
			dirSize += fileSize
		}
	}
	
	return dirSize
}

func getMinDelete(toDelete int, fileScanner *bufio.Scanner)(minDelete, dirSize int) {
	for fileScanner.Scan(){
		line := fileScanner.Text()
		if line == "$ cd .."{
			break
		} else if line[:4] == "$ cd"{
			subdirMinDelete, subdirSize := getMinDelete(toDelete, fileScanner)
			dirSize += subdirSize
			if subdirMinDelete > toDelete && (minDelete == 0 || subdirMinDelete < minDelete) {
				minDelete = subdirMinDelete
			} 
		} else if line != "$ ls" && line[:3] != "dir"{
			fileSize, _ := strconv.Atoi(strings.Split(line, " ")[0])
			dirSize += fileSize
		}
	}

	
	if dirSize > toDelete && (minDelete == 0 || dirSize < minDelete){
		minDelete = dirSize
	}
	return minDelete, dirSize
}

func part1()int{
	var totalSize int
	
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	totalSize, _ = getTotalSize(fileScanner)

    readFile.Close()
	return totalSize
}


func part2()int{
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	systemSize := getSystemSize(fileScanner)
    readFile.Close()

	freeSpace := 70000000 - systemSize
	toDelete := 30000000 - freeSpace

	readFile, err = os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner = bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	minDelete, _ := getMinDelete(toDelete, fileScanner)

    readFile.Close()
	return minDelete
}

func main() {
	// fmt.Println(part1())
	fmt.Println(part2())
}
