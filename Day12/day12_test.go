package main

import (
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

   record = "?###????????"
   numDamagedSprings = []int{3,2,1}
   numArrangements = GetNumArrangements("", record, numDamagedSprings)

   if numArrangements != 10 {
      t.Fatal()
   }
}
