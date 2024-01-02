package main

import (
	"testing"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_GetNumArrangements(t *testing.T) {
   originalRecord := "???.###"
   numContiguousDamagedSprings := []int{1, 1, 3}
   numOperationalSpringsToDistribute := 0
   numContiguousOperationalSprings := []int{0, 1, 1, 0}
   matchingRecords := map[string]bool{}
   numArrangements := GetNumArrangements(originalRecord, numContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)

   if numArrangements != 1 {
      t.Fatal()
   }

   originalRecord = ".??..??...?##."
   numContiguousDamagedSprings = []int{1, 1, 3}
   numOperationalSpringsToDistribute = 7
   numContiguousOperationalSprings = []int{0, 1, 1, 0}
   matchingRecords = map[string]bool{}
   numArrangements = GetNumArrangements(originalRecord, numContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)

   if numArrangements != 4 {
      t.Fatal()
   }

   originalRecord = "?#?#?#?#?#?#?#?"
   numContiguousDamagedSprings = []int{1, 3, 1, 6}
   numOperationalSpringsToDistribute = 1
   numContiguousOperationalSprings = []int{0, 1, 1, 1, 0}
   matchingRecords = map[string]bool{}
   numArrangements = GetNumArrangements(originalRecord, numContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)

   if numArrangements != 1 {
      t.Fatal()
   }

   originalRecord = "????.#...#..."
   numContiguousDamagedSprings = []int{4, 1, 1}
   numOperationalSpringsToDistribute = 5
   numContiguousOperationalSprings = []int{0, 1, 1, 0}
   matchingRecords = map[string]bool{}
   numArrangements = GetNumArrangements(originalRecord, numContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)

   if numArrangements != 1 {
      t.Fatal()
   }

   originalRecord = "????.######..#####."
   numContiguousDamagedSprings = []int{1, 6, 5}
   numOperationalSpringsToDistribute = 5
   numContiguousOperationalSprings = []int{0, 1, 1, 0}
   matchingRecords = map[string]bool{}
   numArrangements = GetNumArrangements(originalRecord, numContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)

   if numArrangements != 4 {
      t.Fatal()
   }

   originalRecord = "?###????????"
   numContiguousDamagedSprings = []int{3, 2, 1}
   numOperationalSpringsToDistribute = 4
   numContiguousOperationalSprings = []int{0, 1, 1, 0}
   matchingRecords = map[string]bool{}
   numArrangements = GetNumArrangements(originalRecord, numContiguousDamagedSprings, numOperationalSpringsToDistribute, numContiguousOperationalSprings, matchingRecords)

   if numArrangements != 10 {
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_GetRecordDataNumArrangements(t *testing.T) {
   input := "???.### 1,1,3"
   recordData, err := ParseRecordData(input)

   if err != nil {
      t.Fatal()
   }

   numArrangements := GetRecordDataNumArrangements(recordData)

   if numArrangements != 1 {
      t.Fatal()
   }

   input = ".??..??...?##. 1,1,3"
   recordData, err = ParseRecordData(input)

   if err != nil {
      t.Fatal()
   }

   numArrangements = GetRecordDataNumArrangements(recordData)

   if numArrangements != 4 {
      t.Fatal()
   }

   input = "?#?#?#?#?#?#?#? 1,3,1,6"
   recordData, err = ParseRecordData(input)

   if err != nil {
      t.Fatal()
   }

   numArrangements = GetRecordDataNumArrangements(recordData)

   if numArrangements != 1 {
      t.Fatal()
   }

   input = "????.#...#... 4,1,1"
   recordData, err = ParseRecordData(input)

   if err != nil {
      t.Fatal()
   }

   numArrangements = GetRecordDataNumArrangements(recordData)

   if numArrangements != 1 {
      t.Fatal()
   }

   input = "????.######..#####. 1,6,5"
   recordData, err = ParseRecordData(input)

   if err != nil {
      t.Fatal()
   }

   numArrangements = GetRecordDataNumArrangements(recordData)

   if numArrangements != 4 {
      t.Fatal()
   }

   input = "?###???????? 3,2,1"
   recordData, err = ParseRecordData(input)

   if err != nil {
      t.Fatal()
   }

   numArrangements = GetRecordDataNumArrangements(recordData)

   if numArrangements != 10 {
      t.Fatal()
   }
}
