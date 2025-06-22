package main

import (
   "log"
   "os"
   "strconv"
   "strings"
   "slices"
)

type StepFields struct {
   mLabel string
   mOperation rune
   mHasFocalLength bool
   mFocalLength int
}

type LenseInfo struct {
   mLabel string
   mFocalLength int
}

type Box struct {
   mIndex int
   mLenseInfoList []LenseInfo
}

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

   input := ReadInput()

   hash := GetHash(input)
   log.Println("Hash:", hash)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func part2() {

   log.Println("Part 2")

   input := ReadInput()

   boxes := make([]Box, 256)

   for i := range boxes {
      boxes[i].mIndex = i
   }

   ProcessSteps(input, boxes)

   focusingPower := 0

   for _, box := range boxes {
      focusingPower += box.CalculateFocusingPower()
   }

   log.Println("Focusing Power:", focusingPower)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func ReadInput() string {

   input, err := os.ReadFile("./input.txt")
   check(err)

   return string(input)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func ParseStep(pInput string) StepFields {
   var stepFields StepFields

   input := strings.TrimSpace(string(pInput))
   fieldsSplit := strings.Split(input, "=")

   if len(fieldsSplit) > 1 {
      stepFields.mLabel = fieldsSplit[0]
      stepFields.mOperation = '='
      focalLength, err := strconv.ParseInt(fieldsSplit[1], 10, 0)
      check(err)
      stepFields.mHasFocalLength = true
      stepFields.mFocalLength = int(focalLength)

   } else {
      fieldsSplit := strings.Split(pInput, "-")

      stepFields.mLabel = fieldsSplit[0]
      stepFields.mOperation = '-'
      stepFields.mHasFocalLength = false
   }

   return stepFields
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
func GetHash(pInput string) int {

   hash := 0
   inputs := strings.Split(pInput, ",")

   for _, input := range inputs {
      hash += GetStepHash(input)
   }

   return hash
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetStepHash(pInput string) int {
   hash := 0

   for i := range pInput {
      inputRune := rune(pInput[i])

      if inputRune != '\n' {
         hash = GetNewHash(hash, inputRune)
      }
   }

   return hash
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func GetNewHash(pHash int, pRune rune) int {
   newHash := pHash + int(pRune)
   newHash *= 17
   newHash %= 256

   return newHash
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func ProcessSteps(pInput string, pBoxes []Box) {

   inputs := strings.Split(pInput, ",")

   for _, input := range inputs {
      stepFields := ParseStep(input)
      ProcessStep(stepFields, pBoxes)
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func ProcessStep(pStepFields StepFields, pBoxes []Box) {

   boxIndex := GetStepHash(pStepFields.mLabel)

   if pStepFields.mOperation == '=' {
      pBoxes[boxIndex].AddLense(pStepFields.mLabel, pStepFields.mFocalLength)

   } else if pStepFields.mOperation == '-' {
      pBoxes[boxIndex].RemoveLense(pStepFields.mLabel)
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (box *Box) AddLense(pLabel string, pFocalLength int) {
   for i := range box.mLenseInfoList {
      if box.mLenseInfoList[i].mLabel == pLabel {
         box.mLenseInfoList[i].mFocalLength = pFocalLength
         return
      }
   }

   lenseInfo := LenseInfo{}
   lenseInfo.mLabel = pLabel
   lenseInfo.mFocalLength = pFocalLength
   box.mLenseInfoList = append(box.mLenseInfoList, lenseInfo)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (box *Box) RemoveLense(pLabel string) {
   indexToRemove := -1

   for i := range box.mLenseInfoList {
      if box.mLenseInfoList[i].mLabel == pLabel {
         indexToRemove = i
      }
   }

   if indexToRemove == -1 {
      return
   }

   box.mLenseInfoList = slices.Delete(box.mLenseInfoList, indexToRemove, indexToRemove+1)
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func (box *Box) CalculateFocusingPower() int {
   focusingPower := 0

   for lenseIndex, lenseInfo := range box.mLenseInfoList {
      focusingPower += CalculateLenseFocusingPower(box.mIndex, lenseIndex, lenseInfo.mFocalLength)
   }

   return focusingPower
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func CalculateLenseFocusingPower(pBoxIndex int, pLenseIndex int, pFocalLength int) int {
   return (pBoxIndex+1) * (pLenseIndex+1) * pFocalLength
}
