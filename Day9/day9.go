package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvVar struct {
   mReadings []int
}

func main() {
   Part1()
}

func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   envVars := ParseEnvVars(inputLines)
   extrapolatedReadingsSum := SumExtrapolatedReadings(envVars)

   fmt.Println("Sum of Extrapolated Readings:", extrapolatedReadingsSum)
}

func ReadInput() []string {
   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := strings.TrimSpace(string(input))
   inputLines := strings.Split(inputStr, "\n")

   return inputLines
}

func check(pE error) {
   if pE != nil {
      panic(pE)
   }
}

func ParseEnvVars(pLines []string) []EnvVar {
   envVars := []EnvVar{}

   for _, line := range pLines {
      envVar := ParseEnvVar(line)
      envVars = append(envVars, envVar)
   }

   return envVars
}

func ParseEnvVar(pInput string) EnvVar {
   envVar := EnvVar{}

   fields := strings.Fields(pInput)

   for _, field := range fields {
      reading, err := strconv.ParseInt(field, 10, 0) ; if err == nil {
         envVar.mReadings = append(envVar.mReadings, int(reading))
      }
   }

   return envVar
}

func SumExtrapolatedReadings(pEnvVars []EnvVar) int {
   extrapolatedReadingsSum := 0

   for _, envVar := range pEnvVars {
      extrapolatedReadingsSum += envVar.Extrapolate()
   }

   return extrapolatedReadingsSum
}

func (envVar *EnvVar) Extrapolate() int {

   extrapolatedReading := 0

   readingDiffs := envVar.ReadingDiffs()

   for i := len(readingDiffs)-2; i >= 0; i-- {
      extrapolatedReading += readingDiffs[i][len(readingDiffs[i])-1]
   }

   return extrapolatedReading
}

func (envVar *EnvVar) ReadingDiffs() [][]int {
   readingDiffs := [][]int{}
   readingDiffs = append(readingDiffs, envVar.mReadings)

   isAllZeros := false

   for ; !isAllZeros ; {
      curReadings := readingDiffs[len(readingDiffs)-1]
      curReadingDiffs := Difference(curReadings)

      readingDiffs = append(readingDiffs, curReadingDiffs)
      isAllZeros = IsAllZeros(curReadingDiffs)
   }

   return readingDiffs
}

func Difference(pInputs []int) []int {
   diffs := []int{}

   for i := 1; i < len(pInputs); i++ {
      diff := pInputs[i] - pInputs[i-1]
      diffs = append(diffs, diff)
   }

   return diffs
}

func IsAllZeros(pInputs []int) bool {

   for _, input := range pInputs {
      if input != 0 {
         return false
      }
   }

   return true
}
