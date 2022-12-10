package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
    "strings"
)

func allDifferent(all[] string)bool{
	for i := 0; i < len(all); i++ {
		for j := i+1; j < len(all); j++ {
			if all[i] == all[j] {
				return false
			}
		} 
	}
	return true
}


func part(n int)int{
	if n > 14{
		fmt.Println("Sorry, but unfortunatly I didn't prepere for this situation")
		os.Exit(1)
	}
	filebuffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
 
	var s[14] string;
	lastNums := s[:n]

	for i := 0; i < n-1; i++{
		data.Scan()
		lastNums[i+1] = data.Text()
	}
	
	cnt := n
	for data.Scan() {
		letter := data.Text()
		for i := 0; i < n-1; i++{
			lastNums[i] = lastNums[i+1]
		}
		lastNums[n-1] = letter
		if allDifferent(lastNums){
			return cnt
		}
		cnt += 1
	}
	return cnt
}

func main() {
	// fmt.Println(part1())
	fmt.Println(part(14))
}
