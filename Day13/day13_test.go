package main

import (
	"log"
	"testing"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_FindColReflectionIndex(t *testing.T) {

   var pattern []string

	pattern = append(pattern, "#.##..##.")
	pattern = append(pattern, "..#.##.#.")
	pattern = append(pattern, "##......#")
	pattern = append(pattern, "##......#")
	pattern = append(pattern, "..#.##.#.")
	pattern = append(pattern, "..##..##.")
	pattern = append(pattern, "#.#.##.#.")

	requiredNumSmudges := 0
   reflectionIndex, _ := FindColReflectionIndex(pattern, requiredNumSmudges)
   expected := 4

   if reflectionIndex != expected {
      log.Println(reflectionIndex)
      log.Println(expected)
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_FindRowReflectionIndex(t *testing.T) {

   var pattern []string

   pattern = append(pattern, "#...##..#")
   pattern = append(pattern, "#....#..#")
   pattern = append(pattern, "..##..###")
   pattern = append(pattern, "#####.##.")
   pattern = append(pattern, "#####.##.")
   pattern = append(pattern, "..##..###")
   pattern = append(pattern, "#....#..#")

	requiredNumSmudges := 0
   reflectionIndex, _ := FindRowReflectionIndex(pattern, requiredNumSmudges)
   expected := 3

   if reflectionIndex != expected {
      log.Println(reflectionIndex)
      log.Println(expected)
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_GetSummaryValue(t *testing.T) {

	var patterns [][]string

	{
		var pattern []string

		pattern = append(pattern, "#.##..##.")
		pattern = append(pattern, "..#.##.#.")
		pattern = append(pattern, "##......#")
		pattern = append(pattern, "##......#")
		pattern = append(pattern, "..#.##.#.")
		pattern = append(pattern, "..##..##.")
		pattern = append(pattern, "#.#.##.#.")

		patterns = append(patterns, pattern)
	}

	{
		var pattern []string

		pattern = append(pattern, "#...##..#")
		pattern = append(pattern, "#....#..#")
		pattern = append(pattern, "..##..###")
		pattern = append(pattern, "#####.##.")
		pattern = append(pattern, "#####.##.")
		pattern = append(pattern, "..##..###")
		pattern = append(pattern, "#....#..#")

		patterns = append(patterns, pattern)
	}

	requiredNumSmudges := 0
	patternSummaryValue := GetSummaryValue(patterns, requiredNumSmudges)

	if patternSummaryValue != 405 {
      log.Println(patternSummaryValue)
	   t.Fatal()
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_FindRowReflectionIndex2_1(t *testing.T) {

   var pattern []string

	pattern = append(pattern, "#.##..##.")
	pattern = append(pattern, "..#.##.#.")
	pattern = append(pattern, "##......#")
	pattern = append(pattern, "##......#")
	pattern = append(pattern, "..#.##.#.")
	pattern = append(pattern, "..##..##.")
	pattern = append(pattern, "#.#.##.#.")

	requiredNumSmudges := 1
   reflectionIndex, _ := FindRowReflectionIndex(pattern, requiredNumSmudges)
   expected := 2

   if reflectionIndex != expected {
      log.Println(reflectionIndex)
      log.Println(expected)
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_FindRowReflectionIndex2_2(t *testing.T) {

   var pattern []string

   pattern = append(pattern, "#...##..#")
   pattern = append(pattern, "#....#..#")
   pattern = append(pattern, "..##..###")
   pattern = append(pattern, "#####.##.")
   pattern = append(pattern, "#####.##.")
   pattern = append(pattern, "..##..###")
   pattern = append(pattern, "#....#..#")

	requiredNumSmudges := 1
   reflectionIndex, _ := FindRowReflectionIndex(pattern, requiredNumSmudges)
   expected := 0

   if reflectionIndex != expected {
      log.Println(reflectionIndex)
      log.Println(expected)
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_GetSummaryValue2(t *testing.T) {

	var patterns [][]string

	{
		var pattern []string

		pattern = append(pattern, "#.##..##.")
		pattern = append(pattern, "..#.##.#.")
		pattern = append(pattern, "##......#")
		pattern = append(pattern, "##......#")
		pattern = append(pattern, "..#.##.#.")
		pattern = append(pattern, "..##..##.")
		pattern = append(pattern, "#.#.##.#.")

		patterns = append(patterns, pattern)
	}

	{
		var pattern []string

		pattern = append(pattern, "#...##..#")
		pattern = append(pattern, "#....#..#")
		pattern = append(pattern, "..##..###")
		pattern = append(pattern, "#####.##.")
		pattern = append(pattern, "#####.##.")
		pattern = append(pattern, "..##..###")
		pattern = append(pattern, "#....#..#")

		patterns = append(patterns, pattern)
	}

	requiredNumSmudges := 1
	patternSummaryValue := GetSummaryValue(patterns, requiredNumSmudges)

	if patternSummaryValue != 400 {
      log.Println(patternSummaryValue)
	   t.Fatal()
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_Transpose(t *testing.T) {

   var pattern []string

	pattern = append(pattern, "#.##..##.")
	pattern = append(pattern, "..#.##.#.")
	pattern = append(pattern, "##......#")
	pattern = append(pattern, "##......#")
	pattern = append(pattern, "..#.##.#.")
	pattern = append(pattern, "..##..##.")
	pattern = append(pattern, "#.#.##.#.")

	transposedPattern := Transpose(pattern)

   var expectedTransposedPattern []string

	expectedTransposedPattern = append(expectedTransposedPattern, "#.##..#")
	expectedTransposedPattern = append(expectedTransposedPattern, "..##...")
	expectedTransposedPattern = append(expectedTransposedPattern, "##..###")
	expectedTransposedPattern = append(expectedTransposedPattern, "#....#.")
	expectedTransposedPattern = append(expectedTransposedPattern, ".#..#.#")
	expectedTransposedPattern = append(expectedTransposedPattern, ".#..#.#")
	expectedTransposedPattern = append(expectedTransposedPattern, "#....#.")
	expectedTransposedPattern = append(expectedTransposedPattern, "##..###")
	expectedTransposedPattern = append(expectedTransposedPattern, "..##...")

	if len(transposedPattern) != len(expectedTransposedPattern) {
      t.Fatal()
	}

	if len(transposedPattern[0]) != len(expectedTransposedPattern[0]) {
      t.Fatal()
	}

	for row := range len(transposedPattern) {
	   for col := range len(transposedPattern[0]) {
			if transposedPattern[row][col] != expectedTransposedPattern[row][col] {
			   t.Fatal()
			}
	   }
	}
}
