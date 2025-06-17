package main

import (
	"log"
	"math"
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
	patterns := ParsePatterns(inputLines)

	log.Println("Num Patterns:", len(patterns))

	patternSummaryValue := GetSummaryValue(patterns)
	log.Println("Pattern Summary Value:", patternSummaryValue)
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
func ParsePatterns(pInputs []string) [][]string {
	var patterns [][]string

	var pattern []string

	for _, input := range pInputs {
	   if len(input) == 0 {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
	   } else {
			pattern = append(pattern, input)
		}
	}

	patterns = append(patterns, pattern)

	return patterns
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetSummaryValue(pPatterns [][]string) int {

	patternSummaryValue := 0

	for i, pattern := range pPatterns {
		log.Println("i:", i)
		rowReflectionIndex, rowReflectionWidth := FindRowReflectionIndex(pattern)
		colReflectionIndex, colReflectionWidth := FindColReflectionIndex(pattern)

		if rowReflectionIndex != -1 {
			log.Println("Row:", rowReflectionIndex, "Width:", rowReflectionWidth)
			patternSummaryValue += (100 * (rowReflectionIndex + 1))
		}

		if colReflectionIndex != -1 {
			log.Println("Col:", colReflectionIndex, "Width:", colReflectionWidth)
			patternSummaryValue += (colReflectionIndex + 1)
		}

		log.Println("Current Pattern Summary Value:", patternSummaryValue)
	}

	return patternSummaryValue
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func FindColReflectionIndex(pPattern []string) (int, int) {

	transposedPattern := Transpose(pPattern)

	return FindRowReflectionIndex(transposedPattern)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func FindRowReflectionIndex(pPattern []string) (int, int) {

	reflectionIndex := -1
	reflectionWidth := -1

	for row := range pPattern {

		upperReflectionWidth := row + 1
		lowerReflectionWidth := len(pPattern) - row - 1
		testReflectionWidth := int(math.Min(float64(upperReflectionWidth), float64(lowerReflectionWidth)))

		rowReflectionIndex, rowReflectionWidth := FindRowReflectionIndexConstrained(pPattern, row, 0, testReflectionWidth)

		if rowReflectionWidth != -1 && rowReflectionWidth >= reflectionWidth {
			reflectionIndex = rowReflectionIndex
			reflectionWidth = rowReflectionWidth
		}
	}

	return reflectionIndex, reflectionWidth
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func FindRowReflectionIndexConstrained(pPattern []string, pRow int, pCol int, pReflectionWidth int) (int, int) {

	if pReflectionWidth == 0 {
		return -1, -1
	}

	numCols := len(pPattern[0])

	if pCol >= numCols {
		return pRow, pReflectionWidth
	}

	for rowOffset := range pReflectionWidth {

		upperRow := pRow - rowOffset
		lowerRow := (pRow + 1) + rowOffset

		if pPattern[upperRow][pCol] != pPattern[lowerRow][pCol] {
			return -1, -1
		}
	}

	return FindRowReflectionIndexConstrained(pPattern, pRow, pCol+1, pReflectionWidth)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Transpose(pPattern []string) []string {

	numRows := len(pPattern)
	numCols := len(pPattern[0])

	transposedRowPlaceholder := strings.Repeat(" ", numRows)

	var transposedPattern []string

	for range numCols {
	   transposedPattern = append(transposedPattern, transposedRowPlaceholder)
	}

	for row := range numCols {
	   for col := range numRows {
			newRunes := []rune(transposedPattern[row])
			newRunes[col] = rune(pPattern[col][row])
	      transposedPattern[row] = string(newRunes)
	   }
	}

	return transposedPattern
}
