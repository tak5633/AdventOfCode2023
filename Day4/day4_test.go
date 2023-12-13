package main

import (
	"fmt"
	"testing"
)

func Test_ParseNumbers(t *testing.T) {
   input := "13 92 16  5 87 78 15 94 21"
   numbers := ParseNumbers(input)

   if len(numbers) != 9 ||
      numbers[0] != 13 || numbers[1] != 92 || numbers[2] != 16 ||
      numbers[3] !=  5 || numbers[4] != 87 || numbers[5] != 78 ||
      numbers[6] != 15 || numbers[7] != 94 || numbers[8] != 21 {
      t.Fatal()
   }
}

func Test_ParseLine(t *testing.T) {
   line := "Card   1: 17 15  5 | 13 92 16  5 87 78 15 94 21"
   winningNumbers, playerNumbers := ParseLine(line)

   if len(winningNumbers) != 3 &&
      winningNumbers[0] != 17 && winningNumbers[1] != 15 && winningNumbers[2] != 5 {
      t.Fatal()
   }

   if len(playerNumbers) != 9 &&
      playerNumbers[0] != 13 && playerNumbers[1] != 92 && playerNumbers[2] != 16 &&
      playerNumbers[3] !=  5 || playerNumbers[4] != 87 || playerNumbers[5] != 78 ||
      playerNumbers[6] != 15 || playerNumbers[7] != 94 || playerNumbers[8] != 21 {
      t.Fatal()
   }
}

func Test_HasNumber(t *testing.T) {
   winningNumbers := [3]int{17, 15, 5}

   if HasNumber(winningNumbers[:], 15) == false {
      t.Fatal()
   }

   if HasNumber(winningNumbers[:], 21) == true {
      t.Fatal()
   }
}

func Test_SumOfCards_Example(t *testing.T) {
   lines := []string{
      "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
      "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
      "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
      "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
      "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
      "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
   }

   sumOfCards := SumOfCards(lines)

   if sumOfCards != 13 {
      fmt.Println(sumOfCards)
      t.Fatal()
   }
}
