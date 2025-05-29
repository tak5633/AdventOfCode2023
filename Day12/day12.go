package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RecordData struct {
   mRecord string
   mNumDamagedSprings []int
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func main() {
   part1()
   part2()
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func part1() {
   log.Println("Part 1")

   inputLines := ReadInput()
   allRecordData := ParseAllRecordData(inputLines)

   sumOfAllArrangements := 0

   for _, recordData := range allRecordData {
      sumOfAllArrangements += GetNumArrangements("", recordData.mRecord, recordData.mNumDamagedSprings)
   }

   log.Println("Sum of All Arrangements:", sumOfAllArrangements)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func part2() {
   log.Println("Part 2")

   inputLines := ReadInput()

   part2InputLines := Part2InputFormatter(inputLines)
   allRecordData := ParseAllRecordData(part2InputLines)

   sumOfAllArrangements := 0

   for recordDataIdx, recordData := range allRecordData {
      log.Println("RecordDataIdx:", recordDataIdx)
      sumOfAllArrangements += GetNumArrangements("", recordData.mRecord, recordData.mNumDamagedSprings)
   }

   log.Println("Sum of All Arrangements:", sumOfAllArrangements)
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
func ParseAllRecordData(pInputs []string) []RecordData {

   allRecordData := []RecordData{}

   for _, input := range pInputs {
      if recordData, err := ParseRecordData(input) ; err == nil {
         allRecordData = append(allRecordData, recordData)
      }
   }

   return allRecordData
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func ParseRecordData(pInput string) (RecordData, error) {

   fields := strings.Fields(pInput)

   if len(fields) != 2 {
      return RecordData{}, errors.New("ParseRecordData: Failed to parse fields")
   }

   recordData := RecordData{}
   recordData.mRecord = fields[0]

   numDamagedSpringsStrs := strings.Split(fields[1], ",")

   for _, numDamagedSpringsStr := range numDamagedSpringsStrs {
      if numDamagedSprings, err := strconv.ParseInt(numDamagedSpringsStr, 10, 0) ; err == nil {
         recordData.mNumDamagedSprings = append(recordData.mNumDamagedSprings, int(numDamagedSprings))
      }
   }

   return recordData, nil
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetNumArrangements(pCurrentRecord string, pRemainingRecord string, pNumDamagedSpringsArray []int) int {
   memo := map[string]int{}
   return GetNumArrangementsMemo(pCurrentRecord, pRemainingRecord, pNumDamagedSpringsArray, memo)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetNumArrangementsMemo(pCurrentRecord string, pRemainingRecord string, pNumDamagedSpringsArray []int, pMemo map[string]int) int {

   // log.Println(pCurrentRecord, pRemainingRecord, pNumDamagedSpringsArray)

   memoKey := pCurrentRecord + " " + pRemainingRecord + " " + fmt.Sprint(pNumDamagedSpringsArray)
   numMemoArrangements, memoMatch := pMemo[memoKey]

   if memoMatch {
      // log.Println("Return memo")
      // log.Println(numMemoArrangements)
      return numMemoArrangements
   }

   if pRemainingRecord == "" {
      if len(pNumDamagedSpringsArray) > 1 {
         // log.Println("Return 0")
         return 0
      }

      numDamagedSprings := 0

      if len(pNumDamagedSpringsArray) > 0 {
         numDamagedSprings = pNumDamagedSpringsArray[0]
      }

      if IsMatch(pCurrentRecord, numDamagedSprings) {
         // log.Println("Return 1")
         return 1
      }

      // log.Println("Return 0")
      return 0
   }

   currentRecord := pCurrentRecord
   numDamagedSpringsArray := pNumDamagedSpringsArray

   if IsFallingEdge(currentRecord) {

      numDamagedSprings := 0

      if len(pNumDamagedSpringsArray) > 0 {
         numDamagedSprings = pNumDamagedSpringsArray[0]
      }

      numDamagedSpringsArray = []int{}

      if len(pNumDamagedSpringsArray) > 1 {
         numDamagedSpringsArray = pNumDamagedSpringsArray[1:]
      }

      if !IsMatch(currentRecord, numDamagedSprings) {
         // log.Println("Return 0")
         return 0
      }

      currentRecord = ""
   }

   newSpring := pRemainingRecord[0]
   remainingRecord := ""

   if len(pRemainingRecord) > 1 {
      remainingRecord = pRemainingRecord[1:]
   }

   if newSpring == '?' {
      numArrangements := GetNumArrangementsMemo(currentRecord+".", remainingRecord, numDamagedSpringsArray, pMemo) +
                         GetNumArrangementsMemo(currentRecord+"#", remainingRecord, numDamagedSpringsArray, pMemo)
      pMemo[memoKey] = numArrangements
      // log.Println("Return unknown num arrangements")
      // log.Println(numArrangements)
      return numArrangements
   }

   currentRecord += string(newSpring)

   numArrangements := GetNumArrangementsMemo(currentRecord, remainingRecord, numDamagedSpringsArray, pMemo)
   pMemo[memoKey] = numArrangements
   // log.Println("Return known spring num arrangements")
   // log.Println(numArrangements)
   return numArrangements
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func IsMatch(pRecord string, pNumDamagedSprings int) bool {

   trimmedRecord := strings.Trim(pRecord, ".")
   recordSplit := strings.Split(trimmedRecord, ".")

   if len(recordSplit) != 1 {
      return false
   }

   return len(recordSplit[0]) == pNumDamagedSprings
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func IsFallingEdge(pRecord string) bool {

   return len(pRecord) >= 2 && pRecord[len(pRecord)-2] == '#' && pRecord[len(pRecord)-1] == '.'
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Part2InputFormatter(pInputs []string) []string {

   formattedInputs := []string{}

   for _, input := range pInputs {
      fields := strings.Fields(input)

      if len(fields) == 2 {
         record := fields[0] + "?"
         newRecord := strings.Repeat(record, 5)
         newRecord = newRecord[:len(newRecord)-1]

         numDamagedSpringsListing := fields[1] + ","
         newNumDamagedSpringsListing := strings.Repeat(numDamagedSpringsListing, 5)
         newNumDamagedSpringsListing = newNumDamagedSpringsListing[:len(newNumDamagedSpringsListing)-1]

         formattedInput := newRecord + " " + newNumDamagedSpringsListing
         formattedInputs = append(formattedInputs, formattedInput)
      }
   }

   return formattedInputs
}
