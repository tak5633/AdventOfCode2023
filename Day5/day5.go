package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CategoryMap struct {
   mSource string
   mDest string
   mEntries []MapEntry
   mDestMap *CategoryMap
}

type MapEntry struct {
   mDestRangeStart int
   mSourceRangeStart int
   mRangeLength int
}

func main() {
   Part1()
}

func Part1() {
   fmt.Println("Part 1")

   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := strings.TrimSpace(string(input))
   inputLines := strings.Split(inputStr, "\n")

   seedNumbers := ParseSeedNumbers(inputLines[0])
   categoryMaps := ParseCategoryMaps(inputLines)

   minLocationNumber := math.MaxInt

   for _, seedNumber := range seedNumbers {
      seedMap, found := GetCategoryMap(categoryMaps, "seed") ; if found == true {
         seedLocationNumber := seedMap.FindNumber(seedNumber, "location")
         minLocationNumber = int(math.Min(float64(minLocationNumber), float64(seedLocationNumber)))
      }
   }

   fmt.Println("Minimum Location Number:", minLocationNumber)
}

func check(pE error) {
   if pE != nil {
      panic(pE)
   }
}

func ParseSeedNumbers(pLine string) []int {
   var seedNumbers []int

   seedInfoStrs := strings.SplitN(pLine, ":", 2)
   seedNumberStrs := strings.Split(seedInfoStrs[1], " ")

   for _, seedNumberStr := range seedNumberStrs {
      seedNumber, err := strconv.ParseInt(seedNumberStr, 10, 0) ; if err == nil {
         seedNumbers = append(seedNumbers, int(seedNumber))
      }
   }

   return seedNumbers
}

func ParseCategoryMaps(pLines []string) []CategoryMap {
   var categoryMaps []CategoryMap

   allMapLines := pLines[2:]
   allCategoryMapLines := ParseAllCategoryMapLines(allMapLines)

   for _, categoryMapLines := range allCategoryMapLines {
      categoryMap, err := ParseCategoryMap(categoryMapLines) ; if err == nil {
         categoryMaps = append(categoryMaps, categoryMap)
      }
   }

   LinkCategoryMaps(&categoryMaps)

   return categoryMaps
}

func ParseAllCategoryMapLines(pLines []string) [][]string {
   var allCategoryMapLines [][]string

   categoryMapLines := []string{}

   for i, line := range pLines {

      endOfParagraph := len(line) == 0
      endOfInput := i == len(pLines)-1

      if endOfParagraph || endOfInput {
         if len(categoryMapLines) > 0 {
            allCategoryMapLines = append(allCategoryMapLines, categoryMapLines)
         }

         categoryMapLines = []string{}
      } else {
         categoryMapLines = append(categoryMapLines, line)
      }
   }

   return allCategoryMapLines
}

func ParseCategoryMap(pLines []string) (CategoryMap, error) {
   var categoryMap CategoryMap

   headerFields := strings.Split(pLines[0], " ")
   sourceDestFields := strings.Split(headerFields[0], "-")

   if len(sourceDestFields) != 3 {
      return CategoryMap{}, errors.New("ParseCategoryMap: Failed to parse the source and destination")
   }

   categoryMap.mSource = sourceDestFields[0]
   categoryMap.mDest = sourceDestFields[2]

   categoryMap.mEntries = ParseEntries(pLines[1:])

   return categoryMap, nil
}

func ParseEntries(pInputs []string) []MapEntry {
   var mapEntries []MapEntry

   for _, input := range pInputs {
      mapEntry, err := ParseEntry(input) ; if err == nil {
         mapEntries = append(mapEntries, mapEntry)
      }
   }

   return mapEntries
}

func ParseEntry(pInput string) (MapEntry, error) {
   fields := strings.SplitN(pInput, " ", 3)

   if len(fields) != 3 {
      return MapEntry{}, errors.New("ParseEntry: Failed to parse the fields")
   }

   var mapEntry MapEntry
   destRangeStart, err := strconv.ParseInt(fields[0], 10, 0)

   if err != nil {
      return MapEntry{}, errors.New("ParseEntry: Failed to parse the destination range start")
   }

   mapEntry.mDestRangeStart = int(destRangeStart)

   sourceRangeStart, err := strconv.ParseInt(fields[1], 10, 0)

   if err != nil {
      return MapEntry{}, errors.New("ParseEntry: Failed to parse the source range start")
   }

   mapEntry.mSourceRangeStart = int(sourceRangeStart)

   rangeLength, err := strconv.ParseInt(fields[2], 10, 0)

   if err != nil {
      return MapEntry{}, errors.New("ParseEntry: Failed to parse the rangth length")
   }

   mapEntry.mRangeLength = int(rangeLength)

   return mapEntry, nil
}

func LinkCategoryMaps(pCategoryMaps *[]CategoryMap) {
   for i := 0; i < len(*pCategoryMaps); i++ {
      for j := i+1; j < len(*pCategoryMaps); j++ {
         if (*pCategoryMaps)[i].mDest == (*pCategoryMaps)[j].mSource {
            (*pCategoryMaps)[i].mDestMap = &(*pCategoryMaps)[j]
         }
      }
   }
}

func GetCategoryMap(pCategoryMaps []CategoryMap, pSource string) (CategoryMap, bool) {
   for _, categoryMap := range pCategoryMaps {
      if categoryMap.mSource == pSource {
         return categoryMap, true
      }
   }

   return CategoryMap{}, false
}

func (categoryMap *CategoryMap) FindNumber(pSourceNumber int, pTargetDest string) int {
   curCategoryMap := categoryMap
   sourceNumber := pSourceNumber

   for {
      destNumber := curCategoryMap.GetDestNumber(sourceNumber)

      if curCategoryMap.mDest == pTargetDest {
         return destNumber
      }

      curCategoryMap = curCategoryMap.mDestMap
      sourceNumber = destNumber
   }
}

func (categoryMap *CategoryMap) GetDestNumber(pSourceNumber int) int {
   for _, entry := range categoryMap.mEntries {
      sourceStart := entry.mSourceRangeStart
      sourceEnd := sourceStart+entry.mRangeLength-1

      if sourceStart <= pSourceNumber && pSourceNumber <= sourceEnd {
         sourceOffset := pSourceNumber-sourceStart
         return entry.mDestRangeStart+sourceOffset
      }
   }

   return pSourceNumber
}
