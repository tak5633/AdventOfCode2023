package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RowCol struct {
   mRow int
   mCol int
}

type PartNumber struct {
   mNumber int
   mRange NumberRange
   mRow int
}

type NumberRange struct {
   mStartIdx int
   mEndIdx int
}

func Part1() {
   fmt.Println("Part 1")

   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := string(input)
   inputLines := strings.Split(inputStr, "\n")

   sumOfPartNumbers := SumOfPartNumbers(inputLines)

   fmt.Println("Sum of Part Numbers:", sumOfPartNumbers)
}

func Part2() {
   fmt.Println("Part 2")

   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := string(input)
   inputLines := strings.Split(inputStr, "\n")

   sumOfGearRatios := SumOfGearRatios(inputLines)

   fmt.Println("Sum of Gear Ratios:", sumOfGearRatios)
}

func check(pE error) {
   if pE != nil {
      panic(pE)
   }
}

func SumOfPartNumbers(pLines []string) int {

   sum := 0
   partNumbers := FindPartNumbers(pLines)

   for i := 0; i < len(partNumbers); i++ {
      sum += partNumbers[i].mNumber
   }

   return sum
}

func FindPartNumbers(pLines []string) []PartNumber {

   var partNumbers []PartNumber

   potentialPartNumbers := FindAllPotentialPartNumbers(pLines)

   for i := 0; i < len(potentialPartNumbers); i++ {
      numberRange := potentialPartNumbers[i].mRange
      numberRow := potentialPartNumbers[i].mRow
      numberCol := numberRange.mStartIdx
      numberLength := numberRange.mEndIdx - numberRange.mStartIdx + 1
      borderIndexes := GetBorderIndexes(numberRow, numberCol, numberLength)

      for j := 0; j < len(borderIndexes); j++ {
         borderRune, err := GetRune(pLines, borderIndexes[j].mRow, borderIndexes[j].mCol)

         if err == nil && IsSymbol(borderRune) {
            partNumbers = append(partNumbers, potentialPartNumbers[i])
         }
      }
   }

   return partNumbers
}

func FindAllPotentialPartNumbers(pLines []string) []PartNumber {

   var allPartNumbers []PartNumber

   for i := 0; i < len(pLines); i++ {
      partNumbers := FindPotentialPartNumbers(pLines[i])

      for j := 0; j < len(partNumbers); j++ {
         partNumbers[j].mRow = i
      }

      allPartNumbers = append(allPartNumbers, partNumbers...)
   }

   return allPartNumbers
}

func FindPotentialPartNumbers(pLine string) []PartNumber {

   var partNumbers []PartNumber

   numberRanges := FindNumberRanges(pLine)

   for i := 0 ; i < len(numberRanges) ; i++ {

      startIdx := numberRanges[i].mStartIdx
      endIdx := numberRanges[i].mEndIdx+1

      number, err := strconv.ParseInt(pLine[startIdx:endIdx], 10, 0)

      if err == nil {
         var partNumber PartNumber
         partNumber.mNumber = int(number)
         partNumber.mRange = numberRanges[i]

         partNumbers = append(partNumbers, partNumber)
      }
   }

   return partNumbers
}

func FindNumberRanges(pLine string) []NumberRange {

   var numberRanges []NumberRange

   for i := 0 ; i < len(pLine) ; {

      startIdx, endIdx, err := FindNumberRange(pLine[i:])

      if err != nil {
         break
      }

      numberRange := NumberRange{i+startIdx, i+endIdx}
      numberRanges = append(numberRanges, numberRange)

      i = max(i+1, i+endIdx+1) // Including i+1 guarantees the loop terminates
   }

   return numberRanges
}

func FindNumberRange(pLine string) (int, int, error) {

   for i := 0 ; i < len(pLine) ; i++ {

      startIdx := i
      endIdx := i

      potentialNumber := string(pLine[i])
      _, err := strconv.ParseInt(potentialNumber, 10, 0)

      if err != nil {
         continue
      }

      for j := i+1 ; j < len(pLine) ; j++ {
         potentialNumber = string(pLine[j])
         _, err = strconv.ParseInt(potentialNumber, 10, 0)

         if err != nil {
            break
         }

         endIdx = j
      }

      return startIdx, endIdx, nil
   }

   return 0, 0, errors.New("FindNumberRange: Failed to find number")
}

func GetBorderIndexes(pStartRow int, pStartCol int, pStringLen int) []RowCol {
   var borderIndexes []RowCol

   // Left column
   borderIndexes = append(borderIndexes, RowCol{pStartRow-1, pStartCol-1})
   borderIndexes = append(borderIndexes, RowCol{pStartRow-0, pStartCol-1})
   borderIndexes = append(borderIndexes, RowCol{pStartRow+1, pStartCol-1})

   // Middle columns
   for i := 0; i < pStringLen; i++ {
      borderIndexes = append(borderIndexes, RowCol{pStartRow-1, pStartCol+i})
      borderIndexes = append(borderIndexes, RowCol{pStartRow+1, pStartCol+i})
   }

   // Right column
   borderIndexes = append(borderIndexes, RowCol{pStartRow-1, pStartCol+pStringLen})
   borderIndexes = append(borderIndexes, RowCol{pStartRow-0, pStartCol+pStringLen})
   borderIndexes = append(borderIndexes, RowCol{pStartRow+1, pStartCol+pStringLen})

   return borderIndexes
}

func GetRune(pLines []string, pRow int, pCol int) (rune, error) {
   if (0 <= pRow && pRow < len(pLines)) && (0 <= pCol && pCol < len(pLines[pRow])) {
      return rune(pLines[pRow][pCol]), nil
   }

   return ' ', errors.New("GetRune: Failed")
}

func IsSymbol(pRune rune) bool {

   nonSymbols := [11]rune{'.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

   for i := 0; i < len(nonSymbols); i++ {
      if pRune == nonSymbols[i] {
         return false
      }
   }

   return true
}

func SumOfGearRatios(pLines []string) int {

   sum := 0

   potentialGears := FindAllPotentialGears(pLines)

   for _, gearPartNumbers := range potentialGears {
      if len(gearPartNumbers) == 2 {
         sum += gearPartNumbers[0] * gearPartNumbers[1]
      }
   }

   return sum
}

func FindAllPotentialGears(pLines []string) map[string][]int {

   potentialGears := map[string][]int{}

   partNumbers := FindPartNumbers(pLines)

   for _, partNumber := range partNumbers {
      numberRange := partNumber.mRange
      numberRow := partNumber.mRow
      numberCol := numberRange.mStartIdx
      numberLength := numberRange.mEndIdx - numberRange.mStartIdx + 1

      borderIndexes := GetBorderIndexes(numberRow, numberCol, numberLength)

      for _, borderIndex := range borderIndexes {
         borderRune, err := GetRune(pLines, borderIndex.mRow, borderIndex.mCol)

         if err == nil && borderRune == '*' {
            key := fmt.Sprint(borderIndex.mRow) + "," + fmt.Sprint(borderIndex.mCol)
            potentialGears[key] = append(potentialGears[key], partNumber.mNumber)
         }
      }
   }

   return potentialGears
}

func main() {
   Part1()
   Part2()
}
