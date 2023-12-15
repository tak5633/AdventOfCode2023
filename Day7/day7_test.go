package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test_GetCardStrength(t *testing.T) {
   CheckCardStrength(t, 'A', 14)
   CheckCardStrength(t, 'K', 13)
   CheckCardStrength(t, 'Q', 12)
   CheckCardStrength(t, 'J', 11)
   CheckCardStrength(t, 'T', 10)
   CheckCardStrength(t, '9', 9)
   CheckCardStrength(t, '8', 8)
   CheckCardStrength(t, '7', 7)
   CheckCardStrength(t, '6', 6)
   CheckCardStrength(t, '5', 5)
   CheckCardStrength(t, '4', 4)
   CheckCardStrength(t, '3', 3)
   CheckCardStrength(t, '2', 2)
}

func CheckCardStrength(t *testing.T, pCard rune, pExpectedStrength int) {
   strength, err := GetCardStrength(pCard)

   if err != nil || strength != pExpectedStrength {
      t.Fatal()
   }
}

func Test_Hand_GetTypeStrength(t *testing.T) {
   hand := Hand{[]rune{'3', '2', 'T', '3', 'K'}, 765}

   if hand.GetTypeStrength() != 1 {
      t.Fatal()
   }

   hand = Hand{[]rune{'T', '5', '5', 'J', '5'}, 684}

   if hand.GetTypeStrength() != 3 {
      t.Fatal()
   }

   hand = Hand{[]rune{'K', 'K', '6', '7', '7'}, 28}

   if hand.GetTypeStrength() != 2 {
      t.Fatal()
   }

   hand = Hand{[]rune{'K', 'T', 'J', 'J', 'T'}, 220}

   if hand.GetTypeStrength() != 2 {
      t.Fatal()
   }

   hand = Hand{[]rune{'Q', 'Q', 'Q', 'J', 'A'}, 483}

   if hand.GetTypeStrength() != 3 {
      t.Fatal()
   }
}

func Test_Hands_Order(t *testing.T) {
   hands := Hands{
      Hand{[]rune{'3', '2', 'T', '3', 'K'}, 765},
      Hand{[]rune{'T', '5', '5', 'J', '5'}, 684},
      Hand{[]rune{'K', 'K', '6', '7', '7'}, 28},
      Hand{[]rune{'K', 'T', 'J', 'J', 'T'}, 220},
      Hand{[]rune{'Q', 'Q', 'Q', 'J', 'A'}, 483},
   }

   sort.Sort(Hands(hands))

   if hands[0].mBid != 765 {
      t.Fatal()
   }

   if hands[1].mBid != 220 {
      t.Fatal()
   }

   if hands[2].mBid != 28 {
      t.Fatal()
   }

   if hands[3].mBid != 684 {
      t.Fatal()
   }

   if hands[4].mBid != 483 {
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

   hands := ParseHands(lines)
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

   hands := ParseHands(lines)

   if hands[0].GetTypeStrength() != 1 {
      t.Fatal()
   }

   if hands[1].GetTypeStrength() != 6 {
      fmt.Println(hands[1].GetTypeStrength())
      t.Fatal()
   }

   sort.Sort(hands)

   if hands[0].mBid != 853 {
      t.Fatal()
   }

   if hands[1].mBid != 396 {
      t.Fatal()
   }

   totalWinnings := GetTotalWinnings(hands)

   if totalWinnings != 1645 {
      t.Fatal()
   }
}
