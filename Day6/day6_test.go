package main

import (
	"testing"
)

func Test_ParseInputs(t *testing.T) {
   timeInputs := ParseTimeInputs("Time:      7  15   30")

   if len(timeInputs) != 3 {
      t.Fatal()
   }

   recordDistanceInputs := ParseDistanceRecordInputs("Distance:  9  40  200")

   if len(recordDistanceInputs) != 3 {
      t.Fatal()
   }
}

func Test_ComputeNumNewRecords(t *testing.T) {
   race := Race{7, 9}
   numNewRecords := race.ComputeNumNewRecords()

   if numNewRecords != 4 {
      t.Fatal()
   }

   race = Race{15, 40}
   numNewRecords = race.ComputeNumNewRecords()

   if numNewRecords != 8 {
      t.Fatal()
   }

   race = Race{30, 200}
   numNewRecords = race.ComputeNumNewRecords()

   if numNewRecords != 9 {
      t.Fatal()
   }
}

func Test_ComputeNumNewRecordsMultiplied(t *testing.T) {
   lines := []string{
      "Time:      7  15   30",
      "Distance:  9  40  200",
   }

   races := ParseInput(lines)
   numNewRecordsMultiplied := ComputeNumNewRecordsMultiplied(races)

   if len(races) != 3 {
      t.Fatal()
   }

   if numNewRecordsMultiplied != 288 {
      t.Fatal()
   }
}
