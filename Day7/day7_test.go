package main

import (
	"sort"
	"testing"
)

func Test_Part1Card_Strength(t *testing.T) {
   Check_Part1Card_Strength(t, Part1Card{'A'}, 14)
   Check_Part1Card_Strength(t, Part1Card{'K'}, 13)
   Check_Part1Card_Strength(t, Part1Card{'Q'}, 12)
   Check_Part1Card_Strength(t, Part1Card{'J'}, 11)
   Check_Part1Card_Strength(t, Part1Card{'T'}, 10)
   Check_Part1Card_Strength(t, Part1Card{'9'}, 9)
   Check_Part1Card_Strength(t, Part1Card{'8'}, 8)
   Check_Part1Card_Strength(t, Part1Card{'7'}, 7)
   Check_Part1Card_Strength(t, Part1Card{'6'}, 6)
   Check_Part1Card_Strength(t, Part1Card{'5'}, 5)
   Check_Part1Card_Strength(t, Part1Card{'4'}, 4)
   Check_Part1Card_Strength(t, Part1Card{'3'}, 3)
   Check_Part1Card_Strength(t, Part1Card{'2'}, 2)
}

func Check_Part1Card_Strength(t *testing.T, pCard Part1Card, pExpectedStrength int) {
   strength := pCard.Strength()

   if strength != pExpectedStrength {
      t.Fatal()
   }
}

func Test_Part1Hand_Strength(t *testing.T) {
   hand := Part1Hand{}
   hand.AddCard('3')
   hand.AddCard('2')
   hand.AddCard('T')
   hand.AddCard('3')
   hand.AddCard('K')
   hand.SetBid(765)

   if hand.Strength() != 1 {
      t.Fatal()
   }

   hand = Part1Hand{}
   hand.AddCard('T')
   hand.AddCard('5')
   hand.AddCard('5')
   hand.AddCard('J')
   hand.AddCard('5')
   hand.SetBid(684)

   if hand.Strength() != 3 {
      t.Fatal()
   }

   hand = Part1Hand{}
   hand.AddCard('K')
   hand.AddCard('K')
   hand.AddCard('6')
   hand.AddCard('7')
   hand.AddCard('7')
   hand.SetBid(28)

   if hand.Strength() != 2 {
      t.Fatal()
   }

   hand = Part1Hand{}
   hand.AddCard('K')
   hand.AddCard('T')
   hand.AddCard('J')
   hand.AddCard('J')
   hand.AddCard('T')
   hand.SetBid(220)

   if hand.Strength() != 2 {
      t.Fatal()
   }

   hand = Part1Hand{}
   hand.AddCard('Q')
   hand.AddCard('Q')
   hand.AddCard('Q')
   hand.AddCard('J')
   hand.AddCard('A')
   hand.SetBid(483)

   if hand.Strength() != 3 {
      t.Fatal()
   }
}

func Test_Part1Hands_Order(t *testing.T) {
   hands := Hands{}

   hand := new(Part1Hand)
   hand.AddCard('3')
   hand.AddCard('2')
   hand.AddCard('T')
   hand.AddCard('3')
   hand.AddCard('K')
   hand.SetBid(765)
   hands = append(hands, hand)

   hand = new(Part1Hand)
   hand.AddCard('T')
   hand.AddCard('5')
   hand.AddCard('5')
   hand.AddCard('J')
   hand.AddCard('5')
   hand.SetBid(684)
   hands = append(hands, hand)

   hand = new(Part1Hand)
   hand.AddCard('K')
   hand.AddCard('K')
   hand.AddCard('6')
   hand.AddCard('7')
   hand.AddCard('7')
   hand.SetBid(28)
   hands = append(hands, hand)

   hand = new(Part1Hand)
   hand.AddCard('K')
   hand.AddCard('T')
   hand.AddCard('J')
   hand.AddCard('J')
   hand.AddCard('T')
   hand.SetBid(220)
   hands = append(hands, hand)

   hand = new(Part1Hand)
   hand.AddCard('Q')
   hand.AddCard('Q')
   hand.AddCard('Q')
   hand.AddCard('J')
   hand.AddCard('A')
   hand.SetBid(483)
   hands = append(hands, hand)

   sort.Sort(Hands(hands))

   if hands[0].Bid() != 765 {
      t.Fatal()
   }

   if hands[1].Bid() != 220 {
      t.Fatal()
   }

   if hands[2].Bid() != 28 {
      t.Fatal()
   }

   if hands[3].Bid() != 684 {
      t.Fatal()
   }

   if hands[4].Bid() != 483 {
      t.Fatal()
   }
}

func Test_GetTotalWinnings_Example(t *testing.T) {
   lines := []string{
      "32T3K 765",
      "T55J5 684",
      "KK677 28",
      "KTJJT 220",
      "QQQJA 483",
   }

   hands := ParseHands(lines, Part1HandCreator)
   totalWinnings := GetTotalWinnings(hands)

   if totalWinnings != 6440 {
      t.Fatal()
   }
}

func Test_GetTotalWinnings_Simple(t *testing.T) {
   lines := []string{
      "2A85A 853",
      "JJJJJ 396",
   }

   hands := ParseHands(lines, Part1HandCreator)

   if hands[0].Strength() != 1 {
      t.Fatal()
   }

   if hands[1].Strength() != 6 {
      t.Fatal()
   }

   sort.Sort(hands)

   if hands[0].Bid() != 853 {
      t.Fatal()
   }

   if hands[1].Bid() != 396 {
      t.Fatal()
   }

   totalWinnings := GetTotalWinnings(hands)

   if totalWinnings != 1645 {
      t.Fatal()
   }
}
