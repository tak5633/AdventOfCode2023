package ci

import (
	"fmt"
	"math"
	"strings"
)

type Galaxy struct {
   mId int
   mRow int
   mCol int
}

type CosmicImage struct {
   mImage []string
   mEmptyRows []int
   mEmptyCols []int
   mExpansionMultiplier int
   mGalaxies []Galaxy
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func NewCosmicImage(pImage []string, pExpansionMultiplier int) CosmicImage {
   cosmicImage := CosmicImage{}

   cosmicImage.mImage = pImage
   cosmicImage.mExpansionMultiplier = pExpansionMultiplier

   cosmicImage.expand()
   cosmicImage.findGalaxies()

   return cosmicImage
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) Galaxies() []Galaxy {
   return cosmicImage.mGalaxies
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) expand() {
   cosmicImage.expandRows()
   cosmicImage.expandCols()
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) expandRows() {

   for rowIdx, row := range cosmicImage.mImage {
      if !strings.ContainsRune(row, '#') {
         cosmicImage.mEmptyRows = append(cosmicImage.mEmptyRows, rowIdx)
      }
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) expandCols() {

   transposedImage := Transpose(cosmicImage.mImage)

   for rowIdx, row := range transposedImage {
      if !strings.ContainsRune(row, '#') {
         cosmicImage.mEmptyCols = append(cosmicImage.mEmptyCols, rowIdx)
      }
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) findGalaxies() {
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
               galaxyPaths[key] = cosmicImage.ComputeGalaxyPath(galaxyA, galaxyB)
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
func (cosmicImage *CosmicImage) ComputeGalaxyPath(pGalaxyA Galaxy, pGalaxyB Galaxy) int {
   rowDiff := int(math.Abs(float64(pGalaxyA.mRow-pGalaxyB.mRow)))
   emptyRows := cosmicImage.emptyRows(pGalaxyA.mRow, pGalaxyB.mRow)

   colDiff := int(math.Abs(float64(pGalaxyA.mCol-pGalaxyB.mCol)))
   emptyCols := cosmicImage.emptyCols(pGalaxyA.mCol, pGalaxyB.mCol)

   return rowDiff + colDiff + (len(emptyRows)*(cosmicImage.mExpansionMultiplier-1)) + (len(emptyCols)*(cosmicImage.mExpansionMultiplier-1))
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) emptyRows(pRowIdxA int, pRowIdxB int) []int {
   emptyRows := []int{}

   for _, rowIdx := range cosmicImage.mEmptyRows {
      if (pRowIdxA < rowIdx && rowIdx < pRowIdxB) ||
         (pRowIdxB < rowIdx && rowIdx < pRowIdxA) {
         emptyRows = append(emptyRows, rowIdx)
      }
   }

   return emptyRows
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (cosmicImage *CosmicImage) emptyCols(pColIdxA int, pColIdxB int) []int {
   emptyCols := []int{}

   for _, colIdx := range cosmicImage.mEmptyCols {
      if (pColIdxA < colIdx && colIdx < pColIdxB) ||
         (pColIdxB < colIdx && colIdx < pColIdxA) {
         emptyCols = append(emptyCols, colIdx)
      }
   }

   return emptyCols
}
