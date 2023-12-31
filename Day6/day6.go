package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
   mTimeMs int
   mDistanceRecordMm int
}

func main() {
   Part1()
   Part2()
}

func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   races := ParseInputDay1(inputLines)
   numNewRecordsMultiplied := ComputeNumNewRecordsMultiplied(races)

   fmt.Println("Num New Records Multiplied:", numNewRecordsMultiplied)
}

func Part2() {
   fmt.Println("Part 2")

   inputLines := ReadInput()

   races := ParseInputDay2(inputLines)
   numNewRecordsMultiplied := ComputeNumNewRecordsMultiplied(races)

   fmt.Println("Num New Records Multiplied:", numNewRecordsMultiplied)
}

func ReadInput() []string {
   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := strings.TrimSpace(string(input))
   inputLines := strings.Split(inputStr, "\n")

   return inputLines
}

func ParseInputDay1(pLines []string) []Race {
   timeInputs := ParseTimeInputsDay1(pLines[0])
   distanceRecordInputs := ParseDistanceRecordInputsDay1(pLines[1])

   return ParseRaces(timeInputs, distanceRecordInputs)
}

func ParseInputDay2(pLines []string) []Race {
   timeInputs := ParseTimeInputsDay2(pLines[0])
   distanceRecordInputs := ParseDistanceRecordInputsDay2(pLines[1])

   return ParseRaces(timeInputs, distanceRecordInputs)
}

func ParseTimeInputsDay1(pInput string) []string {
   timeInfo := strings.SplitN(pInput, ":", 2)

   if len(timeInfo) != 2 {
      return []string{}
   }

   timeInputs := strings.Split(timeInfo[1], " ")

   return RemoveEmptyStrings(timeInputs)
}

func ParseTimeInputsDay2(pInput string) []string {
   input := RemoveSpaces(pInput)

   return ParseTimeInputsDay1(input)
}

func ParseDistanceRecordInputsDay1(pInput string) []string {
   distanceRecordInfo := strings.SplitN(pInput, ":", 2)

   if len(distanceRecordInfo) != 2 {
      return []string{}
   }

   distanceRecordInputs := strings.Split(distanceRecordInfo[1], " ")

   return RemoveEmptyStrings(distanceRecordInputs)
}

func ParseDistanceRecordInputsDay2(pInput string) []string {
   input := RemoveSpaces(pInput)

   return ParseDistanceRecordInputsDay1(input)
}

func RemoveEmptyStrings(pInputs []string) []string {
   var nonEmptyInputs []string

   for _, input := range pInputs {
      if input != "" {
         nonEmptyInputs = append(nonEmptyInputs, input)
      }
   }

   return nonEmptyInputs
}

func RemoveSpaces(pInput string) string {
   output := ""
   inputSplit := strings.Split(pInput, " ")

   for _, inputElement := range inputSplit {
      if inputElement != "" {
         output += inputElement
      }
   }

   return output
}

func ParseRaces(pTimeInputs []string, pDistanceRecordInputs []string) []Race {
   var races []Race

   maxElements := int(math.Min(float64(len(pTimeInputs)), float64(len(pDistanceRecordInputs))))

   for i := 0; i < maxElements; i++ {
      race, err := ParseRace(pTimeInputs[i], pDistanceRecordInputs[i]) ; if err == nil {
         races = append(races, race)
      }
   }

   return races
}

func ParseRace(pTimeInput string, pDistanceRecordInput string) (Race, error) {
   timeMs, err := strconv.ParseInt(pTimeInput, 10, 0) ; if err == nil {
      distanceRecordMm, err := strconv.ParseInt(pDistanceRecordInput, 10, 0) ; if err == nil {
         return Race{int(timeMs), int(distanceRecordMm)}, nil
      }
   }

   return Race{}, errors.New("ParseRace: Failed to parse race")
}

func check(pE error) {
   if pE != nil {
      panic(pE)
   }
}

func ComputeNumNewRecordsMultiplied(pRaces []Race) int {
   numNewRecordsMultiplied := 1

   for _, race := range pRaces {
      raceNumNewRecords := race.ComputeNumNewRecords()
      numNewRecordsMultiplied = numNewRecordsMultiplied*raceNumNewRecords
   }

   return numNewRecordsMultiplied
}

func (race *Race) ComputeNumNewRecords() int {
   numNewRecords := 0

   for chargeTimeMs := 0; chargeTimeMs <= race.mTimeMs; chargeTimeMs++ {
      distanceMm := race.ComputeDistanceMm(chargeTimeMs) ; if distanceMm > race.mDistanceRecordMm {
         numNewRecords++
      }
   }

   return numNewRecords
}

func (race *Race) ComputeDistanceMm(pChargeTimeMs int) int {
   speedMmPerMs := pChargeTimeMs
   travelTimeMs := race.mTimeMs - pChargeTimeMs

   if travelTimeMs <= 0 {
      return 0
   }

   return travelTimeMs*speedMmPerMs
}
