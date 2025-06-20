package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"strings"
)

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

   load := CalculateLoad(inputLines)
   log.Println("Load:", load)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func part2() {

   log.Println("Part 2")

   inputLines := ReadInput()

   numSpins := 1000000000
   load := CalculateLoadAfterSpins(inputLines, numSpins)
   log.Println("Load:", load)
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
func CalculateLoadAfterSpins(pPlatform []string, pNumSpins int) int {

   load := 0
   platform := pPlatform

   memo := map[string][]string{}

   var memoLoop []string

   for spin := range pNumSpins {
      // log.Println("Spin:", spin)

      // if spin % 1000 == 0 {
      //    log.Println("Spin:", spin)
      // }

      hash := CalculateHash(platform)
      // log.Println("Hash:", hash)

      if len(memoLoop) > 1 && hash == memoLoop[0] {
         remainingSpins := pNumSpins - spin - 1

         memoLoopIndex := remainingSpins % len(memoLoop)
         memoLoopHash := memoLoop[memoLoopIndex]
         memoLoopPlatform := memo[memoLoopHash]

         log.Println("Required Num Spins:", spin)

         return CalculateLoad(memoLoopPlatform)
      }

      _, memoMatch := memo[hash]

      if memoMatch {
         // log.Println("Memo match!")
         memoLoop = append(memoLoop, hash)
      }

      platform = Tilt(platform)
      platform = RotateClockwise90(platform)

      platform = Tilt(platform)
      platform = RotateClockwise90(platform)

      platform = Tilt(platform)
      platform = RotateClockwise90(platform)

      platform = Tilt(platform)
      platform = RotateClockwise90(platform)

      memo[hash] = platform

      load = CalculateLoad(platform)

      // Print(platform)
      // log.Println("Load:", load)
   }

   return load
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Tilt(pPlatform []string) []string {

   numRows := len(pPlatform)
   numCols := len(pPlatform[0])

   tiltedPlatformRunes := make([][]rune, numRows)

   for row := range numRows {
      tiltedPlatformRunes[row] = make([]rune, numCols)
   }

   for col := range numCols {

      tiltedCol := TiltColumn(pPlatform, col)

      for row := range numRows {
         tiltedPlatformRunes[row][col] = rune(tiltedCol[row])
      }
   }

   return FromRunes(tiltedPlatformRunes)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func TiltColumn(pPlatform []string, pCol int) string {

   return TiltColumnString(GetColumn(pPlatform, pCol))
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func TiltColumnString(pColString string) string {

   numRows := len(pColString)
   minTiltIndex := 0

   tiltedColRunes := make([]rune, numRows)
   tiltedColRunes[0] = rune(pColString[0])

   if pColString[0] == 'O' {
      minTiltIndex++
   } else if pColString[0] == '#' {
      minTiltIndex = 1
   }

   for row := 1; row < numRows; row++ {
      tiltedColRunes[row] = rune(pColString[row])

      if pColString[row] == 'O' {
         tiltedColRunes[row] = '.'
         tiltedColRunes[minTiltIndex] = 'O'
      }

      if pColString[row] == 'O' {
         minTiltIndex++
      } else if pColString[row] == '#' {
         minTiltIndex = row + 1
      }
   }

   return string(tiltedColRunes)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func CalculateLoad(pPlatform []string) int {

   load := 0
   numCols := len(pPlatform[0])

   for col := range numCols {
      load += CalculateColumnLoad(pPlatform, col)
   }

   return load
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func CalculateColumnLoad(pPlatform []string, pCol int) int {

   return CalculateColumnStringLoad(GetColumn(pPlatform, pCol))
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func CalculateColumnStringLoad(pColString string) int {

   numRows := len(pColString)

   load := 0

   for row := range numRows {
      if pColString[row] == 'O' {
         load += (numRows - row)
      }
   }

   return load
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func CalculateHash(pPlatform []string) string {

   hasher := md5.New()
   numRows := len(pPlatform)

   for row := range numRows {
      hasher.Write([]byte(pPlatform[row]))
   }

   return hex.EncodeToString(hasher.Sum(nil))
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func RotateClockwise(pPlatform []string, pNumRotations int) []string {

   rotatedPlatform := pPlatform

   for range pNumRotations {
      rotatedPlatform = RotateClockwise90(rotatedPlatform)
   }

   return rotatedPlatform
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func RotateClockwise90(pPlatform []string) []string{

   numRows := len(pPlatform)
   numCols := len(pPlatform[0])

   var rotatedPlatform []string

   for col := range numCols {
      var rotatedRow []rune

      for row := numRows-1; row >= 0; row-- {
         rotatedRow = append(rotatedRow, rune(pPlatform[row][col]))
      }

      rotatedPlatform = append(rotatedPlatform, string(rotatedRow))
   }

   return rotatedPlatform
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetColumn(pPlatform []string, pCol int) string {

   var colRunes []rune
   numRows := len(pPlatform)

   for row := range numRows {
      colRunes = append(colRunes, rune(pPlatform[row][pCol]))
   }

   return string(colRunes)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func FromRunes(pPlatformRunes [][]rune) []string {

   var platform []string

   numRows := len(pPlatformRunes)
   numCols := len(pPlatformRunes[0])

   for row := range numRows {
      tiltedRow := ""

      for col := range numCols {
         tiltedRow += string(pPlatformRunes[row][col])
      }

      platform = append(platform, tiltedRow)
   }

   return platform
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Print(pPlatform []string) {

   log.Println("Platform:")

   for _, platformRow := range pPlatform {
      log.Println(platformRow)
   }
}
