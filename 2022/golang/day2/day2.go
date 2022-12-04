package day2

import (
	"bufio"
	"os"
)

func totalScore(array []string) int {
	points := 0

	for i := 0; i < len(array); i++ {
		opponent := string(array[i][0])
		tip := string(array[i][2])

		res, _ := mapResults[opponent+tip]

		p, _ := mapPoints[res]
		points = points + p

		str := res + opponent
		p, _ = mapPoints[str]
		points = points + p
	}

	return points
}

var decode = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

// L0, D3, V3
var mapPoints = map[string]int{
	"A":  1, //rock
	"B":  2, //paper
	"C":  3, //scissors
	"AA": 3, //rock x rock = 3
	"AB": 0, //rock x paper = 0
	"AC": 6, //rock x scissors = 6
	"BA": 6, //paper  x rock = 6
	"BB": 3, //paper  x paper = 3
	"BC": 0, //paper  x scissors = 0
	"CA": 0, //scissors  x rock = 0
	"CB": 6, //scissors  x paper = 6
	"CC": 3, //scissors  x scissors = 3
}

// x - loss
// y - draw
// z - win
var mapResults = map[string]string{
	"AX": "C", //rock win scissors
	"AY": "A", //rock draw rock
	"AZ": "B", //rock loose to paper
	"BX": "A", //paper win rock
	"BY": "B", //paper draw
	"BZ": "C", //paper loose to scissors
	"CX": "B", //scissors win paper
	"CY": "C", //scissors draw scissors
	"CZ": "A", //scissors loose to rock
}

func readInput(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	array := make([]string, 0)

	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}

	err = readFile.Close()
	if err != nil {
		panic(err)
	}

	return array
}
