package main

import (
	"log"
	"testing"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_IsMatch(t *testing.T) {

   record := "#"
   numDamagedSprings := 1
   isMatch := IsMatch(record, numDamagedSprings)

   if isMatch != true {
      t.Fatal()
   }

   record = ".#"
   numDamagedSprings = 1
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != true {
      t.Fatal()
   }

   record = ".#."
   numDamagedSprings = 1
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != true {
      t.Fatal()
   }

   record = "..#.."
   numDamagedSprings = 1
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != true {
      t.Fatal()
   }

   record = "##"
   numDamagedSprings = 1
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != false {
      t.Fatal()
   }

   record = "#"
   numDamagedSprings = 2
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != false {
      t.Fatal()
   }

   record = "##"
   numDamagedSprings = 2
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != true {
      t.Fatal()
   }

   record = "#.#"
   numDamagedSprings = 1
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != false {
      t.Fatal()
   }

   record = ".#.#."
   numDamagedSprings = 1
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != false {
      t.Fatal()
   }

   record = "."
   numDamagedSprings = 0
   isMatch = IsMatch(record, numDamagedSprings)

   if isMatch != true {
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_GetNumArrangements(t *testing.T) {

   record := "???.###"
   numDamagedSprings := []int{1, 1, 3}
   numArrangements := GetNumArrangements("", record, numDamagedSprings)

   if numArrangements != 1 {
      t.Fatal()
   }

   record = ".??..??...?##."
   numDamagedSprings = []int{1, 1, 3}
   numArrangements = GetNumArrangements("", record, numDamagedSprings)

   if numArrangements != 4 {
      t.Fatal()
   }

   record = "?#?#?#?#?#?#?#?"
   numDamagedSprings = []int{1,3,1,6}
   numArrangements = GetNumArrangements("", record, numDamagedSprings)

   if numArrangements != 1 {
      t.Fatal()
   }

   record = "????.#...#..."
   numDamagedSprings = []int{4,1,1}
   numArrangements = GetNumArrangements("", record, numDamagedSprings)

   if numArrangements != 1 {
      t.Fatal()
   }

   record = "????.######..#####."
   numDamagedSprings = []int{1,6,5}
   numArrangements = GetNumArrangements("", record, numDamagedSprings)

   if numArrangements != 4 {
      t.Fatal()
   }

   log.Println("Here")
   record = "?###????????"
   numDamagedSprings = []int{3,2,1}
   numArrangements = GetNumArrangements("", record, numDamagedSprings)

   if numArrangements != 10 {
      log.Println(numArrangements)
      t.Fatal()
   }

//    originalRecord = "?###????????"
//    numContiguousDamagedSprings = []int{3, 2, 1}
//    numOperationalSpringsToDistribute = 4
//    numContiguousOperationalSprings = []int{0, 1, 1, 0}
//    matchingRecords = map[string]bool{}
//    numArrangements = GetNumArrangements(originalRecord, numContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)
//
//    if numArrangements != 10 {
//       t.Fatal()
//    }
}

// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func Test_Part1_GetRecordDataNumArrangements(t *testing.T) {
//    input := "???.### 1,1,3"
//    recordData, err := ParseRecordData(input)
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements := GetRecordDataNumArrangements(recordData)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    input = ".??..??...?##. 1,1,3"
//    recordData, err = ParseRecordData(input)
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangements(recordData)
//
//    if numArrangements != 4 {
//       t.Fatal()
//    }
//
//    input = "?#?#?#?#?#?#?#? 1,3,1,6"
//    recordData, err = ParseRecordData(input)
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangements(recordData)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    input = "????.#...#... 4,1,1"
//    recordData, err = ParseRecordData(input)
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangements(recordData)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    input = "????.######..#####. 1,6,5"
//    recordData, err = ParseRecordData(input)
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangements(recordData)
//
//    if numArrangements != 4 {
//       t.Fatal()
//    }
//
//    input = "?###???????? 3,2,1"
//    recordData, err = ParseRecordData(input)
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangements(recordData)
//
//    if numArrangements != 10 {
//       t.Fatal()
//    }
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func Test_GetNumArrangementsPart2(t *testing.T) {
//    originalRecord := "???.###"
//    numContiguousDamagedSprings := []int{1, 1, 3}
//    numArrangements := GetNumArrangementsPart2(originalRecord, numContiguousDamagedSprings)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    originalRecord = ".??..??...?##."
//    numContiguousDamagedSprings = []int{1, 1, 3}
//    numArrangements = GetNumArrangementsPart2(originalRecord, numContiguousDamagedSprings)
//
//    if numArrangements != 4 {
//       t.Fatal()
//    }
//
//    originalRecord = "?#?#?#?#?#?#?#?"
//    numContiguousDamagedSprings = []int{1, 3, 1, 6}
//    numArrangements = GetNumArrangementsPart2(originalRecord, numContiguousDamagedSprings)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    originalRecord = "????.#...#..."
//    numContiguousDamagedSprings = []int{4, 1, 1}
//    numArrangements = GetNumArrangementsPart2(originalRecord, numContiguousDamagedSprings)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    originalRecord = "????.######..#####."
//    numContiguousDamagedSprings = []int{1, 6, 5}
//    numArrangements = GetNumArrangementsPart2(originalRecord, numContiguousDamagedSprings)
//
//    if numArrangements != 4 {
//       t.Fatal()
//    }
//
//    originalRecord = "?###????????"
//    numContiguousDamagedSprings = []int{3, 2, 1}
//    numArrangements = GetNumArrangementsPart2(originalRecord, numContiguousDamagedSprings)
//
//    if numArrangements != 10 {
//       t.Fatal()
//    }
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func Test_Part2_GetRecordDataNumArrangements(t *testing.T) {
//    input := "???.### 1,1,3"
//    part2Input := Part2InputFormatter([]string{input})
//    recordData, err := ParseRecordData(part2Input[0])
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements := GetRecordDataNumArrangementsPart2(recordData)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    input = ".??..??...?##. 1,1,3"
//    part2Input = Part2InputFormatter([]string{input})
//    recordData, err = ParseRecordData(part2Input[0])
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangementsPart2(recordData)
//
//    if numArrangements != 16384 {
//       t.Fatal()
//    }
//
//    input = "?#?#?#?#?#?#?#? 1,3,1,6"
//    part2Input = Part2InputFormatter([]string{input})
//    recordData, err = ParseRecordData(part2Input[0])
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangementsPart2(recordData)
//
//    if numArrangements != 1 {
//       t.Fatal()
//    }
//
//    input = "????.#...#... 4,1,1"
//    part2Input = Part2InputFormatter([]string{input})
//    recordData, err = ParseRecordData(part2Input[0])
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangementsPart2(recordData)
//
//    if numArrangements != 16 {
//       t.Fatal()
//    }
//
//    input = "????.######..#####. 1,6,5"
//    part2Input = Part2InputFormatter([]string{input})
//    recordData, err = ParseRecordData(part2Input[0])
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangementsPart2(recordData)
//
//    if numArrangements != 2500 {
//       t.Fatal()
//    }
//
//    input = "?###???????? 3,2,1"
//    part2Input = Part2InputFormatter([]string{input})
//    recordData, err = ParseRecordData(part2Input[0])
//
//    if err != nil {
//       t.Fatal()
//    }
//
//    numArrangements = GetRecordDataNumArrangementsPart2(recordData)
//
//    if numArrangements != 506250 {
//       t.Fatal()
//    }
// }
//
// //--------------------------------------------------------------------------------------------------
// //
// //--------------------------------------------------------------------------------------------------
// func Test_Part2_GetNumArrangementsPart2New(t *testing.T) {
//    numContiguousDamagedSprings := []int{1, 1, 1}
//    numOperationalSprings := 5
//
//    numArrangements := GetNumArrangementsPart2New(numContiguousDamagedSprings, numOperationalSprings)
//
//    if numArrangements != 4 {
//       t.Fatal()
//    }
//
//    numContiguousDamagedSprings = []int{1, 1, 1}
//    numOperationalSprings = 7
//
//    numArrangements = GetNumArrangementsPart2New(numContiguousDamagedSprings, numOperationalSprings)
//
//    if numArrangements != 6 {
//       t.Fatal()
//    }
//
//    numContiguousDamagedSprings = []int{1, 1, 1, 1}
//    numOperationalSprings = 6
//
//    numArrangements = GetNumArrangementsPart2New(numContiguousDamagedSprings, numOperationalSprings)
//
//    if numArrangements != 10 {
//       fmt.Println(numArrangements)
//       t.Fatal()
//    }
// }
