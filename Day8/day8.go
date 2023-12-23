package main

import (
	"fmt"
	"os"
	"strings"
)

type Location struct {
   mName string
   mLeftLocation string
   mRightLocation string
}

type Instructions struct {
   mDirections string
   mCurrentDirectionIdx int
   mLocations map[string]Location
   mCurrentLocation string
   mDestinationLocation string
}

type Path struct {
   mLocations []string
}

func main() {
   Part1()
   Part2()
}

func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   instructions := ParseInstructions(inputLines)
   numSteps := GetPart1NumSteps(instructions)

   fmt.Println("Num Steps:", numSteps)
}

func Part2() {
   fmt.Println("Part 2")

   inputLines := ReadInput()

   instructions := ParseInstructions(inputLines)
   numSteps := GetPart2NumSteps(instructions)

   fmt.Println("Num Steps:", numSteps)
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

func ParseInstructions(pLines []string) Instructions {
   instructions := Instructions{}

   instructions.mDirections = pLines[0]

   locations := ParseLocations(pLines[2:])
   instructions.mCurrentLocation = "AAA"

   instructions.mLocations = map[string]Location{}

   for _, location := range locations {
      instructions.mLocations[location.mName] = location
   }

   instructions.mDestinationLocation = "ZZZ"

   return instructions
}

func ParseLocations(pLines []string) []Location {
   locations := []Location{}

   for _, line := range pLines {
      location := ParseLocation(line)
      locations = append(locations, location)
   }

   return locations
}

func ParseLocation(pInput string) Location {
   fields := strings.Fields(pInput)

   location := Location{}
   location.mName = fields[0]
   location.mLeftLocation = fields[2][1:len(fields[2])-1]
   location.mRightLocation = fields[3][0:len(fields[3])-1]

   return location
}

func GetPart1NumSteps(pInstructions Instructions) int {
   numSteps := 0

   for pInstructions.mCurrentLocation != pInstructions.mDestinationLocation {
      pInstructions.Advance()
      numSteps++
   }

   return numSteps
}

func GetPart2NumSteps(pInstructions Instructions) int {

   part2StartingLocations := GetPart2StartingLocations(pInstructions)
   part2Instructions := GetPart2Instructions(pInstructions, part2StartingLocations)

   part2IndividualNumSteps := GetPart2IndividualNumSteps(part2Instructions)
   fmt.Println("Individual Num Steps: ", part2IndividualNumSteps)

   part2Paths := GetPart2Paths(part2Instructions, part2IndividualNumSteps)

   numSteps := part2IndividualNumSteps[0]
   stepSize := part2IndividualNumSteps[0]

   for i := 1; i < len(part2Paths); i++ {
      fmt.Println("Analyzing Instruction Range:", 1, "to", i+1)

      numSteps = GetPart2PathNumSteps(part2Paths[:i+1], numSteps, stepSize)
      stepSize = numSteps

      fmt.Println("Instruction Range Num Steps:", numSteps)
   }

   return numSteps
}

func GetPart2StartingLocations(pInstructions Instructions) []string {
   startingLocations := []string{}

   for _, location := range pInstructions.mLocations {
      if location.mName[len(location.mName)-1] == 'A' {
         startingLocations = append(startingLocations, location.mName)
      }
   }

   return startingLocations
}

func GetPart2Instructions(pInstructions Instructions, pStartingLocations []string) []Instructions {
   instructions := []Instructions{}

   for _, startingLocation := range pStartingLocations {
      newInstructions := pInstructions
      newInstructions.mCurrentLocation = startingLocation
      newInstructions.mDestinationLocation = ""

      instructions = append(instructions, newInstructions)
   }

   return instructions
}

func GetPart2IndividualNumSteps(pInstructions []Instructions) []int {
   individualNumSteps := []int{}

   instructions := make([]Instructions, len(pInstructions))
   copy(instructions, pInstructions)

   for i := range instructions {
      numSteps := 0

      for {
         currentLocation := instructions[i].mCurrentLocation

         if currentLocation[len(currentLocation)-1] == 'Z' {
            break
         }

         instructions[i].Advance()
         numSteps++
      }

      individualNumSteps = append(individualNumSteps, numSteps)
   }

   return individualNumSteps
}

func GetPart2Paths(pInstructions []Instructions, pNumSteps []int) []Path {

   paths := []Path{}

   for i := range pNumSteps {
      part2Path := GetPart2Path(pInstructions[i], pNumSteps[i])
      paths = append(paths, part2Path)
   }

   return paths
}

func GetPart2Path(pInstructions Instructions, pNumSteps int) Path {
   path := Path{}

   for i := 0; i < pNumSteps; i++ {
      pInstructions.Advance()
      path.mLocations = append(path.mLocations, pInstructions.mCurrentLocation)
   }

   return path
}

func GetPart2PathNumSteps(pPaths []Path, pInitialNumSteps int, pStepSize int) int {
   numSteps := pInitialNumSteps

   for {
      done := true

      for _, path := range pPaths {
         locationIdx := (numSteps-1) % len(path.mLocations)
         location := path.mLocations[locationIdx]

         if location[len(location)-1] != 'Z' {
            done = false
         }
      }

      if done {
         return numSteps
      }

      numSteps += pStepSize
   }
}

func (instructions *Instructions) GoLeft() bool {
   return instructions.mDirections[instructions.mCurrentDirectionIdx] == 'L'
}

func (instructions *Instructions) Advance() {
   if instructions.GoLeft() {
      instructions.mCurrentLocation = instructions.mLocations[instructions.mCurrentLocation].mLeftLocation
   } else {
      instructions.mCurrentLocation = instructions.mLocations[instructions.mCurrentLocation].mRightLocation
   }

   instructions.mCurrentDirectionIdx = (instructions.mCurrentDirectionIdx+1) % len(instructions.mDirections)
}
