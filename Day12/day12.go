package main

import (
	"errors"
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
   // part2()
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

// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func part2() {
//    log.Println("Part 2")
//
//    inputLines := ReadInput()
//
//    part2InputLines := Part2InputFormatter(inputLines)
//    allRecordData := ParseAllRecordData(part2InputLines)
//
//    sumOfAllArrangements := 0
//
//    for recordDataIdx, recordData := range allRecordData {
//       log.Println("RecordDataIdx:", recordDataIdx)
//       sumOfAllArrangements += GetRecordDataNumArrangementsPart2(recordData)
//    }
//
//    log.Println("Sum of All Arrangements:", sumOfAllArrangements)
// }

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

   // log.Println(pCurrentRecord, pRemainingRecord, pNumDamagedSpringsArray)

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
      return GetNumArrangements(currentRecord+".", remainingRecord, numDamagedSpringsArray) +
             GetNumArrangements(currentRecord+"#", remainingRecord, numDamagedSpringsArray)
   }

   currentRecord += string(pRemainingRecord[0])

   return GetNumArrangements(currentRecord, remainingRecord, numDamagedSpringsArray)
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

// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func Part2InputFormatter(pInputs []string) []string {
//    formattedInputs := []string{}
//
//    for _, input := range pInputs {
//       fields := strings.Fields(input)
//
//       if len(fields) == 2 {
//          record := fields[0] + "?"
//          newRecord := strings.Repeat(record, 5)
//          newRecord = newRecord[:len(newRecord)-1]
//
//          numContiguousDamagedSpringsListing := fields[1] + ","
//          newNumContiguousDamagedSpringsListing := strings.Repeat(numContiguousDamagedSpringsListing, 5)
//          newNumContiguousDamagedSpringsListing = newNumContiguousDamagedSpringsListing[:len(newNumContiguousDamagedSpringsListing)-1]
//
//          formattedInput := newRecord + " " + newNumContiguousDamagedSpringsListing
//          formattedInputs = append(formattedInputs, formattedInput)
//       }
//    }
//
//    return formattedInputs
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func GetRecordDataNumArrangements(pRecordData RecordData) int {
//
//    numContiguousOperationalSprings := []int{}
//    numContiguousOperationalSprings = append(numContiguousOperationalSprings, 0)
//
//    for i := 0; i < len(pRecordData.mNumContiguousDamagedSprings)-1; i++ {
//       numContiguousOperationalSprings = append(numContiguousOperationalSprings, 1)
//    }
//
//    numContiguousOperationalSprings = append(numContiguousOperationalSprings, 0)
//
//    sumOfContiguousDamagedSprings := 0
//
//    for _, numContiguousDamagedSprings := range pRecordData.mNumContiguousDamagedSprings {
//       sumOfContiguousDamagedSprings += numContiguousDamagedSprings
//    }
//
//    sumOfContiguousOperationalSprings := len(pRecordData.mNumContiguousDamagedSprings)-1
//
//    numOperationalSpringsToDistribute := len(pRecordData.mRecord) - (sumOfContiguousDamagedSprings + sumOfContiguousOperationalSprings)
//
//    matchingRecords := map[string]bool{}
//
//    return GetNumArrangements(pRecordData.mRecord, pRecordData.mNumContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func GetNumArrangements(pOriginalRecord string, pNumContiguousDamagedSprings []int, pNumOperationalSpringsToDistribute int, pNumContiguousOperationalSprings []int, pMatchingRecords map[string]bool) int {
//
//    if pNumOperationalSpringsToDistribute <= 0 {
//
//       potentialRecord := ""
//
//       for numContiguousOperationalSpringsIdx, numContiguousOperationalSprings := range pNumContiguousOperationalSprings {
//
//          operationalSprings := strings.Repeat(".", numContiguousOperationalSprings)
//          potentialRecord += operationalSprings
//
//          if numContiguousOperationalSpringsIdx < len(pNumContiguousDamagedSprings) {
//
//             numContiguousDamagedSprings := pNumContiguousDamagedSprings[numContiguousOperationalSpringsIdx]
//             damagedSprings := strings.Repeat("#", numContiguousDamagedSprings)
//
//             potentialRecord += damagedSprings
//          }
//       }
//
//       match := len(potentialRecord) == len(pOriginalRecord)
//
//       numRunesToCompare := int(math.Min(float64(len(potentialRecord)), float64(len(pOriginalRecord))))
//
//       for i := 0; i < numRunesToCompare && match; i++ {
//          matchingRune := (potentialRecord[i] == pOriginalRecord[i]) ||
//                          (potentialRecord[i] == '.' && pOriginalRecord[i] == '?') ||
//                          (potentialRecord[i] == '#' && pOriginalRecord[i] == '?')
//
//          if !matchingRune {
//             match = false
//          }
//       }
//
//       _, existingMatch := pMatchingRecords[potentialRecord]
//
//       if match && !existingMatch {
//          pMatchingRecords[potentialRecord] = true
//          return 1
//       }
//
//       return 0
//    }
//
//    numArrangements := 0
//
//    for i := 0; i < len(pNumContiguousOperationalSprings); i++ {
//       numContiguousOperationalSprings := make([]int, len(pNumContiguousOperationalSprings))
//       copy(numContiguousOperationalSprings, pNumContiguousOperationalSprings)
//
//       numContiguousOperationalSprings[i] = numContiguousOperationalSprings[i]+1
//       numArrangements += GetNumArrangements(pOriginalRecord, pNumContiguousDamagedSprings, pNumOperationalSpringsToDistribute-1, numContiguousOperationalSprings, pMatchingRecords)
//    }
//
//    return numArrangements
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func GetRecordDataNumArrangementsPart2(pRecordData RecordData) int {
//    return GetNumArrangementsPart2(pRecordData.mRecord, pRecordData.mNumContiguousDamagedSprings)
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func GetNumArrangementsPart2(pTestRecord string, pNumContiguousDamagedSprings []int) int {
//
//    testRecord := pTestRecord
//    unknownSpringIdx := strings.Index(testRecord, "?")
//
//    numTestContiguousDamagedSprings := 0
//    testContiguousDamagedSpringsIdx := 0
//
//    if unknownSpringIdx == -1 {
//
//       for testRecordRuneIdx := 0; testRecordRuneIdx < len(pTestRecord); testRecordRuneIdx++ {
//          testRecordRune := pTestRecord[testRecordRuneIdx]
//
//          if testRecordRune == '#' {
//             numTestContiguousDamagedSprings++
//
//          } else if testRecordRune == '.' {
//
//             if numTestContiguousDamagedSprings > 0 {
//                if testContiguousDamagedSpringsIdx >= len(pNumContiguousDamagedSprings) ||
//                   numTestContiguousDamagedSprings != pNumContiguousDamagedSprings[testContiguousDamagedSpringsIdx] {
//                   return 0
//                }
//                testContiguousDamagedSpringsIdx++
//             }
//             numTestContiguousDamagedSprings = 0
//          }
//
//          if testRecordRuneIdx == (len(pTestRecord)-1) {
//             if numTestContiguousDamagedSprings > 0 {
//                if testContiguousDamagedSpringsIdx < len(pNumContiguousDamagedSprings) &&
//                   numTestContiguousDamagedSprings != pNumContiguousDamagedSprings[testContiguousDamagedSpringsIdx] {
//                   return 0
//                }
//                testContiguousDamagedSpringsIdx++
//             }
//             numTestContiguousDamagedSprings = 0
//
//             if testContiguousDamagedSpringsIdx != len(pNumContiguousDamagedSprings) {
//                return 0
//             }
//          }
//       }
//
//       return 1
//    }
//
//    numArrangements := GetNumTargetArrangements(testRecord, pNumContiguousDamagedSprings, '.')
//    numArrangements += GetNumTargetArrangements(testRecord, pNumContiguousDamagedSprings, '#')
//
//    return numArrangements
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func GetNumTargetArrangements(pTestRecord string, pNumContiguousDamagedSprings []int, pTargetRune rune) int {
//
//    numArrangements := 0
//
//    targetRecord := strings.Replace(pTestRecord, "?", string(pTargetRune), 1)
//    numTestContiguousDamagedSprings := 0
//    testContiguousDamagedSpringsIdx := 0
//
//    for targetRecordRuneIdx := 0; targetRecordRuneIdx < len(targetRecord); targetRecordRuneIdx++ {
//       targetRecordRune := targetRecord[targetRecordRuneIdx]
//
//       if targetRecordRune == '#' {
//          numTestContiguousDamagedSprings++
//
//       } else if targetRecordRune == '.' {
//
//          if numTestContiguousDamagedSprings > 0 {
//             if testContiguousDamagedSpringsIdx >= len(pNumContiguousDamagedSprings) ||
//                numTestContiguousDamagedSprings != pNumContiguousDamagedSprings[testContiguousDamagedSpringsIdx] {
//                break
//             }
//
//             testContiguousDamagedSpringsIdx++
//          }
//          numTestContiguousDamagedSprings = 0
//
//       } else if targetRecordRune == '?' {
//
//          if numTestContiguousDamagedSprings > 0 {
//             if testContiguousDamagedSpringsIdx >= len(pNumContiguousDamagedSprings) ||
//                numTestContiguousDamagedSprings > pNumContiguousDamagedSprings[testContiguousDamagedSpringsIdx] {
//                break
//             }
//          }
//
//          numArrangements += GetNumArrangementsPart2(targetRecord, pNumContiguousDamagedSprings)
//          break
//       }
//
//       if targetRecordRuneIdx == (len(targetRecord)-1) {
//          numArrangements += GetNumArrangementsPart2(targetRecord, pNumContiguousDamagedSprings)
//       }
//    }
//
//    return numArrangements
// }
//
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func GetNumArrangementsPart2New(pNumContiguousDamagedSprings []int, pNumOperationalSprings int) int {
//
//    if len(pNumContiguousDamagedSprings) == 1 {
//       return pNumOperationalSprings+1
//
//    } else if len(pNumContiguousDamagedSprings) == 2 {
//       return pNumOperationalSprings
//
//    } else if len(pNumContiguousDamagedSprings) == 3 {
//       return pNumOperationalSprings-1
//    }
//
//    numArrangements := 0
//    maxNumContiguousOperationalSprings := pNumOperationalSprings-(len(pNumContiguousDamagedSprings)-2);
//
//    for numContiguousOperationalSprings := 1; numContiguousOperationalSprings <= maxNumContiguousOperationalSprings; numContiguousOperationalSprings++ {
//
//       newNumContiguousDamagedSprings := pNumContiguousDamagedSprings[1:]
//       newNumOperationalSprings := pNumOperationalSprings-numContiguousOperationalSprings
//
//       numArrangements += GetNumArrangementsPart2New(newNumContiguousDamagedSprings, newNumOperationalSprings)
//    }
//
//    return numArrangements
// }
