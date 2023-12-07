package main

import (
	"testing"
)

func TestGetFirstDigit(t *testing.T) {
   firstDigit, err := GetFirstDigit("rhqrpdxsqhgxzknr2foursnrcfthree")
   if firstDigit != 2 || err != nil {
      t.Fatal()
   }
}

func TestGetFirstDigit2(t *testing.T) {
   firstDigit, err := GetFirstDigit("four95qvkvveight5")
   if firstDigit != 9 || err != nil {
      t.Fatal()
   }
}

func TestGetLastDigit(t *testing.T) {
   lastDigit, err := GetLastDigit("rhqrpdxsqhgxzknr2foursnrcfthree")
   if lastDigit != 2 || err != nil {
      t.Fatal()
   }
}

func TestGetLastDigit2(t *testing.T) {
   lastDigit, err := GetLastDigit("four95qvkvveight5")
   if lastDigit != 5 || err != nil {
      t.Fatal()
   }
}

func TestGetTwoDigitNumber(t *testing.T) {
   twoDigitNumber, err := GetTwoDigitNumber("rhqrpdxsqhgxzknr2foursnrcfthree")
   if twoDigitNumber != 22 || err != nil {
      t.Fatal()
   }
}

func TestGetTwoDigitNumber2(t *testing.T) {
   twoDigitNumber, err := GetTwoDigitNumber("four95qvkvveight5")
   if twoDigitNumber != 95 || err != nil {
      t.Fatal()
   }
}
