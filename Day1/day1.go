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

func main() {
   part1()
}
