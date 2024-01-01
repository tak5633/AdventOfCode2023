package main

import (
	ci "day11/cosmicImage"
	"fmt"
	"os"
	"strings"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func main() {
   Part1()
   Part2()
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   cosmicImage := ci.NewCosmicImage(inputLines, 2)
   galaxyPaths := cosmicImage.ComputeGalaxyPaths()
   sumOfGalaxyPaths := SumGalaxyPaths(galaxyPaths)

   fmt.Println("Sum of Galaxy Paths:", sumOfGalaxyPaths)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Part2() {
   fmt.Println("Part 2")

   inputLines := ReadInput()

   cosmicImage := ci.NewCosmicImage(inputLines, 1000000)
   galaxyPaths := cosmicImage.ComputeGalaxyPaths()
   sumOfGalaxyPaths := SumGalaxyPaths(galaxyPaths)

   fmt.Println("Sum of Galaxy Paths:", sumOfGalaxyPaths)
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
func SumGalaxyPaths(pInput map[string]int) int {
   sum := 0

   for _, pathLength := range pInput {
      sum += pathLength
   }

   return sum
}
