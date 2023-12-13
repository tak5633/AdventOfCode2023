package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
   Part1()
}

func Part1() {
   fmt.Println("Part 1")

   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := string(input)
   inputLines := strings.Split(inputStr, "\n")

   sumOfCards := SumOfCards(inputLines)

   fmt.Println("Sum of Cards:", sumOfCards)
}

func check(pE error) {
   if pE != nil {
      panic(pE)
   }
}

func SumOfCards(pLines []string) int {

   sum := 0

   for _, line := range pLines {
      if len(line) > 0 {
         numMatchingNumbers := 0
         winningNumbers, playerNumbers := ParseLine(line)

         for _, playerNumber := range playerNumbers {
            if HasNumber(winningNumbers, playerNumber) {
               numMatchingNumbers++
            }
         }

         points := int(math.Pow(2.0, float64(numMatchingNumbers-1)))
         sum += points
      }
   }

   return sum
}

func ParseLine(pLine string) ([]int, []int) {
   cardSplit := strings.SplitN(pLine, ":", 2)
   numbersSplit := strings.SplitN(cardSplit[1], "|", 2)

   winningNumberStrs := numbersSplit[0]
   winningNumbers := ParseNumbers(winningNumberStrs)

   playerNumberStrs := numbersSplit[1]
   playerNumbers := ParseNumbers(playerNumberStrs)

   return winningNumbers, playerNumbers
}

func ParseNumbers(pInput string) []int {
   var numbers []int

   numberStrs := strings.Split(pInput, " ")

   for _, numberStr := range numberStrs {
      number, err := strconv.ParseInt(numberStr, 10, 0)

      if err == nil {
         numbers = append(numbers, int(number))
      }
   }

   return numbers
}

func HasNumber(pSlice []int, pTest int) bool {
   for _, element := range pSlice {
      if pTest == element {
         return true
      }
   }

   return false
}
