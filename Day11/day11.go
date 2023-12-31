package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Galaxy struct {
   mId int
   mRow int
   mCol int
}

type CosmicImage struct {
   mImage []string
   mGalaxies []Galaxy
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func main() {
   Part1()
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   cosmicImage := CosmicImage{}
   cosmicImage.mImage = inputLines
   cosmicImage.Expand()
   cosmicImage.FindGalaxies()
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
func (cosmicImage *CosmicImage) Expand() {
   cosmicImage.ExpandRows()
   cosmicImage.ExpandCols()
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) ExpandRows() {
   expandedImage := []string{}

   for _, row := range cosmicImage.mImage {
      expandedImage = append(expandedImage, row)
      if !strings.ContainsRune(row, '#') {
         expandedImage = append(expandedImage, row)
      }
   }

   cosmicImage.mImage = expandedImage
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) ExpandCols() {

   expandedImage := []string{}
   transposedImage := Transpose(cosmicImage.mImage)

   for _, row := range transposedImage {
      expandedImage = append(expandedImage, row)
      if !strings.ContainsRune(row, '#') {
         expandedImage = append(expandedImage, row)
      }
   }

   cosmicImage.mImage = Transpose(expandedImage)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) FindGalaxies() {
   numGalaxies := 0

   for rowIdx, row := range cosmicImage.mImage {
      for colIdx, rowRune := range row {
         if rowRune == '#' {

            galaxy := Galaxy{}
            galaxy.mId = numGalaxies
            galaxy.mRow = rowIdx
            galaxy.mCol = colIdx

            cosmicImage.mGalaxies = append(cosmicImage.mGalaxies, galaxy)
            numGalaxies++
         }
      }
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) ComputeGalaxyPaths() map[string]int {
   galaxyPaths := map[string]int{}

   for i := 0; i < len(cosmicImage.mGalaxies); i++ {
      for j := 1; j < len(cosmicImage.mGalaxies); j++ {
         if i != j {
            galaxyA := cosmicImage.mGalaxies[i]
            galaxyB := cosmicImage.mGalaxies[j]
            minId := int(math.Min(float64(galaxyA.mId), float64(galaxyB.mId)))
            maxId := int(math.Max(float64(galaxyA.mId), float64(galaxyB.mId)))
            key := fmt.Sprintf("%d,%d", minId, maxId)
            if _, ok := galaxyPaths[key] ; !ok {
               galaxyPaths[key] = ComputeGalaxyPath(galaxyA, galaxyB)
            }
         }
      }
   }

   return galaxyPaths
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Transpose(pInput []string) []string {
   numRows := len(pInput)
   numCols := len(pInput[0])

   input := [][]rune{}

   for _, row := range pInput {
      runeRow := []rune(row)
      input = append(input, runeRow)
   }

   output := make([][]rune, numCols)

   for row := 0; row < len(output); row++ {
      output[row] = make([]rune, numRows)
   }

   for row := 0; row < numRows; row++ {
      for col := 0; col < numCols; col++ {
         output[col][row] = input[row][col]
      }
   }

   outputStrs := []string{}

   for row := 0; row < len(output); row++ {
      outputStrs = append(outputStrs, string(output[row]))
   }

   return outputStrs
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func ComputeGalaxyPath(pGalaxyA Galaxy, pGalaxyB Galaxy) int {
   rowDiff := int(math.Abs(float64(pGalaxyA.mRow-pGalaxyB.mRow)))
   colDiff := int(math.Abs(float64(pGalaxyA.mCol-pGalaxyB.mCol)))

   return rowDiff + colDiff
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
