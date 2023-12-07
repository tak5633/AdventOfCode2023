package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFirstDigit(pString string) (int64, error) {
   for i := 0 ; i < len(pString) ; i++ {
      if theInt, err := strconv.ParseInt(pString[i:i+1], 10, 0) ; err == nil {
         return theInt, nil
      }
   }

   return -1, errors.New("GetFirstDigit: Failed to find a digit")
}

func GetLastDigit(pString string) (int64, error) {
   for i := len(pString)-1 ; i >= 0 ; i-- {
      if theInt, err := strconv.ParseInt(pString[i:i+1], 10, 0) ; err == nil {
         return theInt, nil
      }
   }

   return -1, errors.New("GetFirstDigit: Failed to find a digit")
}

func GetTwoDigitNumber(pString string) (int64, error) {
   firstDigit, err := GetFirstDigit(pString)
   if err != nil {
      return -1, err
   }

   lastDigit, err := GetLastDigit(pString)
   if err != nil {
      return -1, err
   }

   return (firstDigit*10) + lastDigit, nil
}

func ConvertTextToDigits(pInput string) string {
   output := ""

   for inputIdx := 0 ; inputIdx < len(pInput) ; {
      convertedInput, _ := TryToGetDigit(pInput[inputIdx:])

      // Some numbers share a single character suffix/prefix
      // To prevent the first number from impacting the second number increment the index by one
      // Examples:
      // - twone
      // - eightwo
      // - eighthree
      // - oneight
      // - threeight
      // - fiveight
      // - nineight
      // - sevenine
      output += convertedInput
      inputIdx += 1
   }

   return output
}

func TryToGetDigit(pInput string) (string, int) {
   digitTexts := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

   for digitTextsIdx := 0 ; digitTextsIdx < len(digitTexts) ; digitTextsIdx++ {
      digitText := digitTexts[digitTextsIdx]

      if len(digitText) <= len(pInput) && digitText == strings.ToLower(pInput[:len(digitText)]) {
         return strconv.Itoa(digitTextsIdx+1), len(digitText)
      }
   }

   return string(pInput[0]), 1
}

func check(pE error) {
    if pE != nil {
        panic(pE)
    }
}

func part1() {

   fmt.Println("Part 1")

   input, err := os.ReadFile("./part1Input.txt")
   check(err)

   inputStr := string(input)
   inputLines := strings.Split(inputStr, "\n")

   var sum int = 0

   for i := 0 ; i < len(inputLines) ; i++ {
      inputLine := inputLines[i]

      if len(inputLine) > 0 {
         twoDigitNumber, err := GetTwoDigitNumber(inputLine)
         check(err)
         sum += int(twoDigitNumber)
      }
   }

   fmt.Println("Sum:", sum)
}

func part2() {

   fmt.Println("Part 2")

   input, err := os.ReadFile("./part1Input.txt")
   check(err)

   inputStr := string(input)
   inputLines := strings.Split(inputStr, "\n")

   var sum int = 0

   for i := 0 ; i < len(inputLines) ; i++ {
      inputLine := inputLines[i]
      convertedInputLine := ConvertTextToDigits(inputLine)

      if len(convertedInputLine) > 0 {
         twoDigitNumber, err := GetTwoDigitNumber(convertedInputLine)
         check(err)
         sum += int(twoDigitNumber)
      }
   }

   fmt.Println("Sum:", sum)
}

func main() {
   part1()
   part2()
}
