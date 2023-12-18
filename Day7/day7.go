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

type CardInt interface {
   Type() rune
   Strength() int
}

type Part1Card struct {
   mType rune
}

func (card *Part1Card) Type() rune {
   return card.mType
}

func (card *Part1Card) Strength() int {

   if card.mType == 'A' {
      return 14
   } else if card.mType == 'K' {
      return 13
   } else if card.mType == 'Q' {
      return 12
   } else if card.mType == 'J' {
      return 11
   } else if card.mType == 'T' {
      return 10
   }

   strength, _ := strconv.ParseInt(string(card.mType), 10, 0)

   return int(strength)
}

type Part2Card struct {
   mType rune
}

func (card *Part2Card) Type() rune {
   return card.mType
}

func (card *Part2Card) Strength() int {

   if card.mType == 'A' {
      return 14
   } else if card.mType == 'K' {
      return 13
   } else if card.mType == 'Q' {
      return 12
   } else if card.mType == 'J' {
      return 1
   } else if card.mType == 'T' {
      return 10
   }

   strength, _ := strconv.ParseInt(string(card.mType), 10, 0)

   return int(strength)
}

type HandInt interface {
   AddCard(pRune rune)
   NumCards() int
   Card(pIndex int) CardInt
   Strength() int
   SetBid(pBid int)
   Bid() int
}

type Part1Hand struct {
   mCards []CardInt
   mBid int
}

func (hand *Part1Hand) AddCard(pRune rune) {
   card := Part1Card{pRune}
   hand.mCards = append(hand.mCards, &card)
}

func (hand *Part1Hand) NumCards() int {
   return len(hand.mCards)
}

func (hand *Part1Hand) Card(pIndex int) CardInt {
   return hand.mCards[pIndex]
}

func (hand *Part1Hand) Strength() int {

   cardMap := GetCardMap(hand)
   cardCounts := GetCardCounts(cardMap)

   return GetHandStrength(cardCounts)
}

func (hand *Part1Hand) SetBid(pBid int) {
   hand.mBid = pBid
}

func (hand *Part1Hand) Bid() int {
   return hand.mBid
}

type Part2Hand struct {
   mCards []CardInt
   mBid int
}

func (hand *Part2Hand) AddCard(pRune rune) {
   card := Part2Card{pRune}
   hand.mCards = append(hand.mCards, &card)
}

func (hand *Part2Hand) NumCards() int {
   return len(hand.mCards)
}

func (hand *Part2Hand) Card(pIndex int) CardInt {
   return hand.mCards[pIndex]
}

func (hand *Part2Hand) Strength() int {

   cardMap := GetCardMap(hand)

   // The logic with Jokers:
   // If I have 5, then no change
   // If I have 4 + 1J, then bump it to 5 (five of a kind)
   // If I have 3 + 2J, then bump it to 5 (five of a kind)
   // If I have 3 + 1J, then bump it to 4 (four of a kind)
   // If I have 2 + 3J, then bump it to 5 (five of a kind)
   // If I have 2 + 2J, then bump it to 4 (four of a kind)
   // If I have 2 + 1J, then bump it to 3 (full house, or three of a kind)
   //    - Make sure if there are 2 two pairs I bump the larger of the two pairs
   // If I have 1 + 4J, then bump it to 5 (five of a kind)
   // If I have 1 + 3J, then bump it to 4 (four of a kind)
   //   - Make sure to bump the larger of the two others
   // If I have 1 + 2J, then bump it to 3 (three of a kind)
   //   - Make sure to bump the larger of the three others
   // If I have 1 + 1J, then bump it to 2 (one pair)
   //   - Make sure to bump the larger of the four others

   // Find the number of Jokers/Jacks and clear that count to avoid biasing the max card count
   numJokers := cardMap['J']
   cardMap['J'] = 0

   // Find the max card count
   maxCardCount := 0

   for _, cardCount := range cardMap {
      maxCardCount = int(math.Max(float64(maxCardCount), float64(cardCount)))
   }

   // Increment the highest type matching the max card count by the number of Jokers/Jacks
   if cardMap['A'] == maxCardCount {
      cardMap['A'] = cardMap['A']+numJokers
   } else if cardMap['K'] == maxCardCount {
      cardMap['K'] = cardMap['K']+numJokers
   } else if cardMap['Q'] == maxCardCount {
      cardMap['Q'] = cardMap['Q']+numJokers
   } else if cardMap['T'] == maxCardCount {
      cardMap['T'] = cardMap['T']+numJokers
   } else if cardMap['9'] == maxCardCount {
      cardMap['9'] = cardMap['9']+numJokers
   } else if cardMap['8'] == maxCardCount {
      cardMap['8'] = cardMap['8']+numJokers
   } else if cardMap['7'] == maxCardCount {
      cardMap['7'] = cardMap['7']+numJokers
   } else if cardMap['6'] == maxCardCount {
      cardMap['6'] = cardMap['6']+numJokers
   } else if cardMap['5'] == maxCardCount {
      cardMap['5'] = cardMap['5']+numJokers
   } else if cardMap['4'] == maxCardCount {
      cardMap['4'] = cardMap['4']+numJokers
   } else if cardMap['3'] == maxCardCount {
      cardMap['3'] = cardMap['3']+numJokers
   } else if cardMap['2'] == maxCardCount {
      cardMap['2'] = cardMap['2']+numJokers
   }

   cardCounts := GetCardCounts(cardMap)

   return GetHandStrength(cardCounts)
}

func (hand *Part2Hand) SetBid(pBid int) {
   hand.mBid = pBid
}

func (hand *Part2Hand) Bid() int {
   return hand.mBid
}

func GetCardMap(pHand HandInt) map[rune]int {
   cardMap := map[rune]int{}

   for i := 0; i < pHand.NumCards(); i++ {
      card := pHand.Card(i)
      cardMap[card.Type()] = cardMap[card.Type()]+1
   }

   return cardMap
}

func GetCardCounts(pCardMap map[rune]int) []int {
   cardCounts := []int{}

   for _, count := range pCardMap {
      if count > 0 {
         cardCounts = append(cardCounts, count)
      }
   }

   return cardCounts
}

func GetHandStrength(pCardCounts []int) int {
   // Make sure there are enough counts. Zero pad as needed.
   minNumCounts := 2

   for i := len(pCardCounts)-minNumCounts; i < 0; i++ {
      pCardCounts = append(pCardCounts, 0)
   }

   sort.Ints(pCardCounts)

   maxCount := pCardCounts[len(pCardCounts)-1]
   nextMaxCount := pCardCounts[len(pCardCounts)-2]

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

type Hands []HandInt

func (hands Hands) Len() int {
   return len(hands)
}

func (hands Hands) Less(i, j int) bool {
   iTypeStrength := hands[i].Strength()
   jTypeStrength := hands[j].Strength()

   if iTypeStrength != jTypeStrength {
      return iTypeStrength < jTypeStrength
   }

   for c := 0; c < hands[i].NumCards(); c++ {
      iCardStrength := hands[i].Card(c).Strength()
      jCardStrength := hands[j].Card(c).Strength()

      if iCardStrength != jCardStrength {
         return iCardStrength < jCardStrength
      }
   }

   return true
}

type HandCreator func() HandInt

func Part1HandCreator() HandInt {
   return &Part1Hand{}
}

func Part2HandCreator() HandInt {
   return &Part2Hand{}
}

func (hands Hands) Swap(i, j int) {
   hands[i], hands[j] = hands[j], hands[i]
}

func main() {
   Part1()
   Part2()
}

func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   hands := ParseHands(inputLines, Part1HandCreator)
   totalWinnings := GetTotalWinnings(hands)

   fmt.Println("Total Winnings:", totalWinnings)
}

func Part2() {
   fmt.Println("Part 2")

   inputLines := ReadInput()

   hands := ParseHands(inputLines, Part2HandCreator)
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

func ParseHands(pLines []string, pCreateHand HandCreator) Hands {
   hands := Hands{}

   for _, line := range pLines {
      hand, err := ParseHand(line, pCreateHand)

      if err == nil {
         hands = append(hands, hand)
      }
   }

   return hands
}

func ParseHand(pInput string, pCreateHand HandCreator) (HandInt, error) {
   hand := pCreateHand()
   handInfo := strings.Fields(pInput)

   if len(handInfo) != 2 {
      return hand, errors.New("ParseHand: Failed to parse hand info")
   }

   validCards := "AKQJT98765432"

   for _, potentialCard := range handInfo[0] {
      if strings.ContainsRune(validCards, potentialCard) {
         hand.AddCard(potentialCard)
      } else {
         return hand, errors.New("ParseHand: Failed to parse cards")
      }
   }

   bid, err := strconv.ParseInt(handInfo[1], 10, 0)

   if err != nil {
      return hand, errors.New("ParseHand: Failed to parse bid")
   }

   hand.SetBid(int(bid))

   return hand, nil
}

func GetTotalWinnings(pHands Hands) int {
   sort.Sort(pHands)

   totalWinnings := 0

   for rank, hand := range pHands {
      totalWinnings += (rank+1)*hand.Bid()
   }

   return totalWinnings
}
