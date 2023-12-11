package main

import (
	"testing"
)

func Test_FindNumberRange_NoNumber(t *testing.T) {
   startIdx, endIdx, err := FindNumberRange(".......")
   if startIdx != 0 || endIdx != 0 || err == nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("._*+$#/")
   if startIdx != 0 || endIdx != 0 || err == nil {
      t.Fatal()
   }
}

func Test_FindNumberRange_SingleDigit(t *testing.T) {
   startIdx, endIdx, err := FindNumberRange("6......")
   if startIdx != 0 || endIdx != 0 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("...6...")
   if startIdx != 3 || endIdx != 3 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("......6")
   if startIdx != 6 || endIdx != 6 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("6.6....")
   if startIdx != 0 || endIdx != 0 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("..-2..")
   if startIdx != 3 || endIdx != 3 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("..+2..")
   if startIdx != 3 || endIdx != 3 || err != nil {
      t.Fatal()
   }
}

func Test_FindNumberRange_MultipleDigits(t *testing.T) {
   startIdx, endIdx, err := FindNumberRange("61.....")
   if startIdx != 0 || endIdx != 1 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("...61..")
   if startIdx != 3 || endIdx != 4 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange(".....61")
   if startIdx != 5 || endIdx != 6 || err != nil {
      t.Fatal()
   }

   startIdx, endIdx, err = FindNumberRange("61.6...")
   if startIdx != 0 || endIdx != 1 || err != nil {
      t.Fatal()
   }
}

func Test_FindNumberRanges_NoNumbers(t *testing.T) {
   numberRanges := FindNumberRanges(".......")
   if len(numberRanges) != 0 {
      t.Fatal()
   }
}

func Test_FindNumberRanges(t *testing.T) {
   numberRanges := FindNumberRanges("61.6...")
   if len(numberRanges) != 2 {
      t.Fatal()
   }

   if numberRanges[0].mStartIdx != 0 || numberRanges[0].mEndIdx != 1 {
      t.Fatal()
   }

   if numberRanges[1].mStartIdx != 3 || numberRanges[1].mEndIdx != 3 {
      t.Fatal()
   }
}

func Test_FindPotentialPartNumbers(t *testing.T) {
   potentialPartNumbers := FindPotentialPartNumbers("61.6...")
   if len(potentialPartNumbers) != 2 {
      t.Fatal()
   }

   if potentialPartNumbers[0].mNumber != 61 ||
      potentialPartNumbers[0].mRange.mStartIdx != 0 ||
      potentialPartNumbers[0].mRange.mEndIdx != 1 {
      t.Fatal()
   }

   if potentialPartNumbers[1].mNumber != 6 ||
      potentialPartNumbers[1].mRange.mStartIdx != 3 ||
      potentialPartNumbers[1].mRange.mEndIdx != 3 {
      t.Fatal()
   }
}

func Test_FindAllPotentialPartNumbers(t *testing.T) {
   var lines [2]string
   lines[0] = "61.6..."
   lines[1] = "*+2-/48"

   potentialPartNumbers := FindAllPotentialPartNumbers(lines[:])
   if len(potentialPartNumbers) != 4 {
      t.Fatal()
   }

   if potentialPartNumbers[0].mNumber != 61 ||
      potentialPartNumbers[1].mNumber != 6 ||
      potentialPartNumbers[2].mNumber != 2 ||
      potentialPartNumbers[3].mNumber != 48 {
      t.Fatal()
   }
}

func Test_IsSymbol(t *testing.T) {
   if IsSymbol('.') == true {
      t.Fatal()
   }

   if IsSymbol('1') == true {
      t.Fatal()
   }

   if IsSymbol('9') == true {
      t.Fatal()
   }

   if IsSymbol('+') == false {
      t.Fatal()
   }

   if IsSymbol('-') == false {
      t.Fatal()
   }
}

func Test_GetRune(t *testing.T) {

   lines := []string{"1234", "5678", "9012"}

   borderRune, err := GetRune(lines, 0, 0)
   if borderRune != '1' || err != nil {
      t.Fatal()
   }

   borderRune, err = GetRune(lines, -1, -1)
   if borderRune != ' ' || err == nil {
      t.Fatal()
   }

   borderRune, err = GetRune(lines, 3, 3)
   if borderRune != ' ' || err == nil {
      t.Fatal()
   }
}

func Test_GetBorderIndexes(t *testing.T) {
   borderIndexes := GetBorderIndexes(1, 1, 1)
   if len(borderIndexes) != 8 {
      t.Fatal()
   }

   if borderIndexes[0].mRow != 0 || borderIndexes[0].mCol != 0 ||
      borderIndexes[1].mRow != 1 || borderIndexes[1].mCol != 0 ||
      borderIndexes[2].mRow != 2 || borderIndexes[2].mCol != 0 ||
      borderIndexes[3].mRow != 0 || borderIndexes[3].mCol != 1 ||
      borderIndexes[4].mRow != 2 || borderIndexes[4].mCol != 1 ||
      borderIndexes[5].mRow != 0 || borderIndexes[5].mCol != 2 ||
      borderIndexes[6].mRow != 1 || borderIndexes[6].mCol != 2 ||
      borderIndexes[7].mRow != 2 || borderIndexes[7].mCol != 2 {
      t.Fatal()
   }

   borderIndexes = GetBorderIndexes(1, 1, 3)
   if len(borderIndexes) != 12 {
      t.Fatal()
   }

   if borderIndexes[0].mRow != 0 || borderIndexes[0].mCol != 0 ||
      borderIndexes[1].mRow != 1 || borderIndexes[1].mCol != 0 ||
      borderIndexes[2].mRow != 2 || borderIndexes[2].mCol != 0 ||
      borderIndexes[3].mRow != 0 || borderIndexes[3].mCol != 1 ||
      borderIndexes[4].mRow != 2 || borderIndexes[4].mCol != 1 ||
      borderIndexes[5].mRow != 0 || borderIndexes[5].mCol != 2 ||
      borderIndexes[6].mRow != 2 || borderIndexes[6].mCol != 2 ||
      borderIndexes[7].mRow != 0 || borderIndexes[7].mCol != 3 ||
      borderIndexes[8].mRow != 2 || borderIndexes[8].mCol != 3 ||
      borderIndexes[9].mRow != 0 || borderIndexes[9].mCol != 4 ||
      borderIndexes[10].mRow != 1 || borderIndexes[10].mCol != 4 ||
      borderIndexes[11].mRow != 2 || borderIndexes[11].mCol != 4 {
      t.Fatal()
   }

   borderIndexes = GetBorderIndexes(0, 0, 1)
   if len(borderIndexes) != 8 {
      t.Fatal()
   }

   if borderIndexes[0].mRow != -1 || borderIndexes[0].mCol != -1 ||
      borderIndexes[1].mRow !=  0 || borderIndexes[1].mCol != -1 ||
      borderIndexes[2].mRow !=  1 || borderIndexes[2].mCol != -1 ||
      borderIndexes[3].mRow != -1 || borderIndexes[3].mCol !=  0 ||
      borderIndexes[4].mRow !=  1 || borderIndexes[4].mCol !=  0 ||
      borderIndexes[5].mRow != -1 || borderIndexes[5].mCol !=  1 ||
      borderIndexes[6].mRow !=  0 || borderIndexes[6].mCol !=  1 ||
      borderIndexes[7].mRow !=  1 || borderIndexes[7].mCol !=  1 {
      t.Fatal()
   }
}

func Test_Example_Simple(t *testing.T) {
   lines := []string{
      "467..114..",
      "...*......",
      "..35..633.",
      "......#...",
   }

   potentialPartNumbers := FindAllPotentialPartNumbers(lines)

   if len(potentialPartNumbers) != 4 {
      t.Fatal()
   }

   partNumbers := FindPartNumbers(lines)

   if len(partNumbers) != 3 {
      t.Fatal()
   }
}

func Test_Example(t *testing.T) {

   lines := []string{
      "467..114..",
      "...*......",
      "..35..633.",
      "......#...",
      "617*......",
      ".....+.58.",
      "..592.....",
      "......755.",
      "...$.*....",
      ".664.598..",
   }

   partNumbers := FindPartNumbers(lines)

   if len(partNumbers) != 8 {
      t.Fatal()
   }

   sumOfPartNumbers := SumOfPartNumbers(lines)

   if sumOfPartNumbers != 4361 {
      t.Fatal()
   }
}
