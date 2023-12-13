package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
   mWinningNumbers []int
   mPlayerNumbers []int
   mNumMatchingNumbers int
   mNumCopies int
}

func main() {
   Part1()
   Part2()
}

func Part1() {
   fmt.Println("Part 1")

   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := strings.TrimSpace(string(input))
   inputLines := strings.Split(inputStr, "\n")

   sumOfCards := SumOfCards(inputLines)

   fmt.Println("Sum of Cards:", sumOfCards)
}

func Part2() {
   fmt.Println("Part 2")

   input, err := os.ReadFile("./input.txt")
   check(err)

   inputStr := strings.TrimSpace(string(input))
   inputLines := strings.Split(inputStr, "\n")

   cards := ParseOriginalCards(inputLines)
   copies := DistributeCopies(cards)

   numCards := 0

   for _, card := range copies {
      numCards += card.mNumCopies
   }

   fmt.Println("Number of Cards:", numCards)
}

func check(pE error) {
   if pE != nil {
      panic(pE)
   }
}

func SumOfCards(pLines []string) int {

   sum := 0

   for _, line := range pLines {
      _, winningNumbers, playerNumbers, err := ParseLine(line)

      if err != nil {
         continue
      }

      numMatchingNumbers := GetNumMatchingNumbers(winningNumbers, playerNumbers)
      points := int(math.Pow(2.0, float64(numMatchingNumbers-1)))
      sum += points
   }

   return sum
}

func ParseOriginalCards(pLines []string) map[int]Card {

   cards := map[int]Card{}

   for _, line := range pLines {
      cardNumber, winningNumbers, playerNumbers, err := ParseLine(line)

      if err != nil {
         continue
      }

      card := Card{}
      card.mWinningNumbers = winningNumbers
      card.mPlayerNumbers = playerNumbers
      card.mNumMatchingNumbers = GetNumMatchingNumbers(winningNumbers, playerNumbers)
      card.mNumCopies = 1

      cards[cardNumber] = card
   }

   return cards
}

func DistributeCopies(pCards map[int]Card) []Card {

   copies := GetSortedCards(pCards)

   for i, card := range copies {
      for j := 0; j < card.mNumMatchingNumbers; j++ {
         for k := 0; k < card.mNumCopies; k++ {
            copyCardIdx := i+j+1

            if copyCardIdx < len(copies) {
               copies[copyCardIdx].mNumCopies++
            }
         }
      }
   }

   return copies
}

func ParseLine(pLine string) (int, []int, []int, error) {
   cardSplit := strings.SplitN(pLine, ":", 2)

   if len(cardSplit) != 2 {
      return -1, nil, nil, errors.New("ParseLine: Failed to split the card number")
   }

   cardNumSplit := strings.Split(cardSplit[0], " ")
   cardNumber, err := strconv.ParseInt(cardNumSplit[len(cardNumSplit)-1], 10, 0)

   if err != nil {
      return -1, nil, nil, err
   }

   numbersSplit := strings.SplitN(cardSplit[1], "|", 2)

   if len(numbersSplit) != 2 {
      return -1, nil, nil, errors.New("ParseLine: Failed to split the numbers")
   }

   winningNumberStrs := numbersSplit[0]
   winningNumbers := ParseNumbers(winningNumberStrs)

   playerNumberStrs := numbersSplit[1]
   playerNumbers := ParseNumbers(playerNumberStrs)

   return int(cardNumber), winningNumbers, playerNumbers, nil
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

func GetNumMatchingNumbers(pWinningNumbers []int, pPlayerNumbers []int) int {
   numMatchingNumbers := 0

   for _, playerNumber := range pPlayerNumbers {
      if HasNumber(pWinningNumbers, playerNumber) {
         numMatchingNumbers++
      }
   }

   return numMatchingNumbers
}

func HasNumber(pSlice []int, pTest int) bool {
   for _, element := range pSlice {
      if pTest == element {
         return true
      }
   }

   return false
}

func GetSortedCards(pCards map[int]Card) []Card {
   var cardNumbers []int

   for cardNumber, _ := range pCards {
      cardNumbers = append(cardNumbers, cardNumber)
   }

   sort.Ints(cardNumbers)

   var sortedCards []Card

   for _, cardNumber := range cardNumbers {
      sortedCards = append(sortedCards, pCards[cardNumber])
   }

   return sortedCards
}
