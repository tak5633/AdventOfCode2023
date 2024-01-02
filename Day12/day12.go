package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RecordData struct {
   mRecord string
   mNumContiguousDamagedSprings []int
}

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
   fmt.Println("Part 1")

   inputLines := ReadInput()

   allRecordData := ParseAllRecordData(inputLines)

   sumOfAllArrangements := 0

   for _, recordData := range allRecordData {
      sumOfAllArrangements += GetRecordDataNumArrangements(recordData)
   }

   fmt.Println("Sum of All Arrangements:", sumOfAllArrangements)
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
   recordData := RecordData{}

   fields := strings.Fields(pInput)

   if len(fields) != 2 {
      return RecordData{}, errors.New("ParseRecordData: Failed to parse fields")
   }

   recordData.mRecord = fields[0]

   numContiguousDamagedSpringsStrs := strings.Split(fields[1], ",")

   for _, numContiguousDamagedSpringsStr := range numContiguousDamagedSpringsStrs {
      if numContiguousDamagedSprings, err := strconv.ParseInt(numContiguousDamagedSpringsStr, 10, 0) ; err == nil {
         recordData.mNumContiguousDamagedSprings = append(recordData.mNumContiguousDamagedSprings, int(numContiguousDamagedSprings))
      }
   }

   return recordData, nil
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetRecordDataNumArrangements(pRecordData RecordData) int {

   numContiguousOperationalSprings := []int{}
   numContiguousOperationalSprings = append(numContiguousOperationalSprings, 0)

   for i := 0; i < len(pRecordData.mNumContiguousDamagedSprings)-1; i++ {
      numContiguousOperationalSprings = append(numContiguousOperationalSprings, 1)
   }

   numContiguousOperationalSprings = append(numContiguousOperationalSprings, 0)

   sumOfContiguousDamagedSprings := 0

   for _, numContiguousDamagedSprings := range pRecordData.mNumContiguousDamagedSprings {
      sumOfContiguousDamagedSprings += numContiguousDamagedSprings
   }

   sumOfContiguousOperationalSprings := len(pRecordData.mNumContiguousDamagedSprings)-1

   numOperationalSpringsToDistribute := len(pRecordData.mRecord) - (sumOfContiguousDamagedSprings + sumOfContiguousOperationalSprings)

   matchingRecords := map[string]bool{}

   return GetNumArrangements(pRecordData.mRecord, pRecordData.mNumContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetNumArrangements(pOriginalRecord string, pNumContiguousDamagedSprings []int, pNumOperationalSpringsToDistribute int, pNumContiguousOperationalSprings []int, pMatchingRecords map[string]bool) int {

   if pNumOperationalSpringsToDistribute <= 0 {

      potentialRecord := ""

      for numContiguousOperationalSpringsIdx, numContiguousOperationalSprings := range pNumContiguousOperationalSprings {

         operationalSprings := strings.Repeat(".", numContiguousOperationalSprings)
         potentialRecord += operationalSprings

         if numContiguousOperationalSpringsIdx < len(pNumContiguousDamagedSprings) {

            numContiguousDamagedSprings := pNumContiguousDamagedSprings[numContiguousOperationalSpringsIdx]
            damagedSprings := strings.Repeat("#", numContiguousDamagedSprings)

            potentialRecord += damagedSprings
         }
      }

      match := len(potentialRecord) == len(pOriginalRecord)

      numRunesToCompare := int(math.Min(float64(len(potentialRecord)), float64(len(pOriginalRecord))))

      for i := 0; i < numRunesToCompare && match; i++ {
         matchingRune := (potentialRecord[i] == pOriginalRecord[i]) ||
                         (potentialRecord[i] == '.' && pOriginalRecord[i] == '?') ||
                         (potentialRecord[i] == '#' && pOriginalRecord[i] == '?')

         if !matchingRune {
            match = false
         }
      }

      _, existingMatch := pMatchingRecords[potentialRecord]

      if match && !existingMatch {
         pMatchingRecords[potentialRecord] = true
         return 1
      }

      return 0
   }

   numArrangements := 0

   for i := 0; i < len(pNumContiguousOperationalSprings); i++ {
      numContiguousOperationalSprings := make([]int, len(pNumContiguousOperationalSprings))
      copy(numContiguousOperationalSprings, pNumContiguousOperationalSprings)

      numContiguousOperationalSprings[i] = numContiguousOperationalSprings[i]+1
      numArrangements += GetNumArrangements(pOriginalRecord, pNumContiguousDamagedSprings, pNumOperationalSpringsToDistribute-1, numContiguousOperationalSprings, pMatchingRecords)
   }

   return numArrangements
}
