package main
 
import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)
 

func part1() int {
    readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    var maxSum, currSum int

    for fileScanner.Scan() {

        line := fileScanner.Text()
        if line == "" {
            currSum = 0
        }

        value, _ := strconv.Atoi(line)
        currSum += value
        if currSum > maxSum {
            maxSum = currSum
        }
    }
    
    readFile.Close()
    return maxSum
}

func part2() int {
    readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    var maxSum, sndMaxSum, thrMaxSum, currSum int

    for fileScanner.Scan() {

        line := fileScanner.Text()
        if line == "" {
            currSum = 0
        }

        value, _ := strconv.Atoi(line)
        currSum += value
         
        switch {
        case currSum >= maxSum:
            thrMaxSum = sndMaxSum
            sndMaxSum = maxSum
            maxSum = currSum
        case currSum >= sndMaxSum:
            thrMaxSum = sndMaxSum
            sndMaxSum = currSum
        case currSum >= thrMaxSum:
            thrMaxSum = currSum
        }
    }
    
    readFile.Close()
    return maxSum + sndMaxSum + thrMaxSum
}

func main() {
    fmt.Println(part1())
    fmt.Println(part2())
}