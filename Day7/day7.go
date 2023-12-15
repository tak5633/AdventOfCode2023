package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
   mCards []rune
   mBid int
}

type Hands []Hand

func main() {
   Part1()
}

func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   hands := ParseHands(inputLines)
   totalWinnings := GetTotalWinnings(hands)

   fmt.Println("Total Winnings:", totalWinnings)
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

func ParseHands(pLines []string) Hands {
   hands := Hands{}

   for _, line := range pLines {
      hand, err := ParseHand(line)

      if err == nil {
         hands = append(hands, hand)
      }
   }

   return hands
}

func ParseHand(pInput string) (Hand, error) {
   handInfo := strings.SplitN(pInput, " ", 2)

   if len(handInfo) != 2 {
      return Hand{}, errors.New("ParseHand: Failed to parse hand info")
   }

   cards := []rune{}

   for _, card := range handInfo[0] {
      cards = append(cards, card)
   }

   bid, err := strconv.ParseInt(handInfo[1], 10, 0)

   if err != nil {
      return Hand{}, errors.New("ParseHand: Failed to parse bid")
   }

   return Hand{cards, int(bid)}, nil
}

func GetTotalWinnings(pHands Hands) int {
   sort.Sort(pHands)

   totalWinnings := 0

   for rank, hand := range pHands {
      totalWinnings += (rank+1)*hand.mBid
   }

   return totalWinnings
}

func (hand *Hand) GetTypeStrength() int {
   counts := hand.getCardCounts()

   // Make sure there are enough counts. Zero pad as needed.
   minNumCounts := 2

   for i := len(counts)-minNumCounts; i < 0; i++ {
      counts = append(counts, 0)
   }

   sort.Ints(counts)

   maxCount := counts[len(counts)-1]
   nextMaxCount := counts[len(counts)-2]

   if maxCount == 5 {
      return 6
   } else if maxCount == 4 {
      return 5
   } else if maxCount == 3 && nextMaxCount == 2 {
      return 4
   } else if maxCount == 3 {
      return 3
   } else if maxCount == 2 && nextMaxCount == 2 {
      return 2
   } else if maxCount == 2 && nextMaxCount == 1 {
      return 1
   }

   return 0
}

func (hand *Hand) getCardCounts() []int {
   cardMap := hand.getCardMap()

   cardCounts := []int{}

   for _, count := range cardMap {
      if count > 0 {
         cardCounts = append(cardCounts, count)
      }
   }

   return cardCounts
}

func (hand *Hand) getCardMap() map[rune]int {
   cardMap := map[rune]int{}

   for _, card := range hand.mCards {
      cardMap[card] = cardMap[card]+1
   }

   return cardMap
}

func GetCardStrength(pCard rune) (int, error) {
   if pCard == 'A' {
      return 14, nil
   } else if pCard == 'K' {
      return 13, nil
   } else if pCard == 'Q' {
      return 12, nil
   } else if pCard == 'J' {
      return 11, nil
   } else if pCard == 'T' {
      return 10, nil
   }

   strength, err := strconv.ParseInt(string(pCard), 10, 0)

   if err != nil {
      return -1, err
   }

   return int(strength), nil
}

func (hands Hands) Len() int {
   return len(hands)
}

func (hands Hands) Less(i, j int) bool {
   iTypeStrength := hands[i].GetTypeStrength()
   jTypeStrength := hands[j].GetTypeStrength()

   if iTypeStrength != jTypeStrength {
      return iTypeStrength < jTypeStrength
   }

   for c := 0; c < len(hands[i].mCards); c++ {
      iCardStrength, iErr := GetCardStrength(hands[i].mCards[c])
      jCardStrength, jErr := GetCardStrength(hands[j].mCards[c])

      if iErr == nil && jErr == nil && iCardStrength != jCardStrength {
         return iCardStrength < jCardStrength
      }
   }

   return true
}

func (hands Hands) Swap(i, j int) {
   hands[i], hands[j] = hands[j], hands[i]
}
