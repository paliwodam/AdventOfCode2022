package main

import(
	"fmt"
	"os"
    "bufio"
    "strings"
)

const rock = "rock"
const paper = "paper"
const scissors = "scissors"
const win = "win"
const loss = "loss"
const draw = "draw"

var elfMoves = map[string]string{
		"A": rock,
		"B": paper,
		"C": scissors,
	}
var winningMoves = map[string]string{
		rock: scissors,
		paper: rock,
		scissors: paper,
	}
var defaultValues = map[string]int{
		rock: 1,
		paper: 2,
		scissors: 3,
		loss: 0,
		draw: 3,
		win: 6,
	}


func part1()int {
	myMoves := map[string]string{
		"X": rock,
		"Y": paper, 
		"Z": scissors,
	}

	var score int

	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line := fileScanner.Text()
		moves := strings.Fields(line)
		
		elf := elfMoves[moves[0]]
		me := myMoves[moves[1]]

		score += defaultValues[me]

		switch elf{
		case me:
			score += defaultValues[draw]
		case winningMoves[me]:
			score += defaultValues[win]
		default:
			score += defaultValues[loss]
		}	
    }

    readFile.Close()
	return score
}

func part2()int{
	results := map[string]string{
		"X": loss, 
		"Y": draw,
		"Z": win,
	}
	
	lossingMoves := make(map[string]string, len(winningMoves))
	for key, val := range winningMoves {
		lossingMoves[val] = key
	}
	
	var score int

	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line := fileScanner.Text()
		moves := strings.Fields(line)
		
		elf := elfMoves[moves[0]]
		result := results[moves[1]]

		score += defaultValues[result]

		switch result{
		case draw:
			score += defaultValues[elf]
		case loss:
			score += defaultValues[winningMoves[elf]]
		default:
			score += defaultValues[lossingMoves[elf]]
		}	
    }

    readFile.Close()
	return score
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}