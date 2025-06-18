package main

import (
	"log"
	"os"
	"strings"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func main() {
   part1()
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func part1() {
   log.Println("Part 1")

   inputLines := ReadInput()
	totalLoad := CalculateTotalLoad(inputLines)

	log.Println("Total Load:", totalLoad)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func ReadInput() []string {

   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := strings.TrimSpace(string(input))
   inputLines := strings.Split(inputStr, "\n")

   return inputLines
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func check(pE error) {

   if pE != nil {
      panic(pE)
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func CalculateTotalLoad(pInput []string) int {

	numRows := len(pInput)
	numCols := len(pInput[0])

	load := 0

	for col := range numCols {
		var colRunes []rune

		for row := range numRows {
			colRunes = append(colRunes, rune(pInput[row][col]))
		}

		colString := string(colRunes)
		load += CalculateTotalColumnLoad(colString)
	}

   return load
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func CalculateTotalColumnLoad(pCol string) int {

	numRows := len(pCol)
	minTiltIndex := 0

	load := 0

	if pCol[0] == 'O' {
		load += (numRows - minTiltIndex)
	}

	if pCol[0] == 'O' {
		minTiltIndex++
	} else if pCol[0] == '#' {
		minTiltIndex = 1
	}

	for row := 1; row < numRows; row++ {
		if pCol[row] == 'O' {
			load += (numRows - minTiltIndex)
		}

		if pCol[row] == 'O' {
			minTiltIndex++
		} else if pCol[row] == '#' {
			minTiltIndex = row + 1
		}
	}

	return load
}
