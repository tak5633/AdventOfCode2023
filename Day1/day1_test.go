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

func TestConvertTextToDigitsNotADigit(t *testing.T) {
   newText := ConvertTextToDigits("notAdigit")
   if newText != "notAdigit" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThree(t *testing.T) {
   newText := ConvertTextToDigits("three")
   if newText != "3hree" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeUppercase(t *testing.T) {
   newText := ConvertTextToDigits("THREE")
   if newText != "3HREE" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeMixedCase(t *testing.T) {
   newText := ConvertTextToDigits("tHrEe")
   if newText != "3HrEe" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeLeadingChars(t *testing.T) {
   newText := ConvertTextToDigits("asdfTHREE")
   if newText != "asdf3HREE" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeTrailingChars(t *testing.T) {
   newText := ConvertTextToDigits("THREEasdf")
   if newText != "3HREEasdf" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeLeadingTrailingChars(t *testing.T) {
   newText := ConvertTextToDigits("asdfTHREEfdsa")
   if newText != "asdf3HREEfdsa" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeLeadingDigit(t *testing.T) {
   newText := ConvertTextToDigits("1THREE")
   if newText != "13HREE" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeTrailingDigit(t *testing.T) {
   newText := ConvertTextToDigits("THREE1")
   if newText != "3HREE1" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsThreeLeadingTrailingDigits(t *testing.T) {
   newText := ConvertTextToDigits("1THREE2")
   if newText != "13HREE2" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsMultiple(t *testing.T) {
   newText := ConvertTextToDigits("asdfnine1THREE2eightfdsa")
   if newText != "asdf9ine13HREE28ightfdsa" {
      t.Fatal()
   }
}

func TestConvertTextToDigitsAocExample1(t *testing.T) {
   newText := ConvertTextToDigits("two1nine")
   twoDigitNumber, err := GetTwoDigitNumber(newText)
   if twoDigitNumber != 29 || err != nil {
      t.Fatal()
   }
}

func TestConvertTextToDigitsAocExample2(t *testing.T) {
   newText := ConvertTextToDigits("eightwothree")
   twoDigitNumber, err := GetTwoDigitNumber(newText)
   if twoDigitNumber != 83 || err != nil {
      t.Fatal()
   }
}

func TestConvertTextToDigitsAocExample3(t *testing.T) {
   newText := ConvertTextToDigits("abcone2threexyz")
   twoDigitNumber, err := GetTwoDigitNumber(newText)
   if twoDigitNumber != 13 || err != nil {
      t.Fatal()
   }
}

func TestConvertTextToDigitsAocExample4(t *testing.T) {
   newText := ConvertTextToDigits("xtwone3four")
   twoDigitNumber, err := GetTwoDigitNumber(newText)
   if twoDigitNumber != 24 || err != nil {
      t.Fatal()
   }
}

func TestConvertTextToDigitsAocExample5(t *testing.T) {
   newText := ConvertTextToDigits("4nineeightseven2")
   twoDigitNumber, err := GetTwoDigitNumber(newText)
   if twoDigitNumber != 42 || err != nil {
      t.Fatal()
   }
}

func TestConvertTextToDigitsAocExample6(t *testing.T) {
   newText := ConvertTextToDigits("zoneight234")
   twoDigitNumber, err := GetTwoDigitNumber(newText)
   if twoDigitNumber != 14 || err != nil {
      t.Fatal()
   }
}

func TestConvertTextToDigitsAocExample7(t *testing.T) {
   newText := ConvertTextToDigits("7pqrstsixteen")
   twoDigitNumber, err := GetTwoDigitNumber(newText)
   if twoDigitNumber != 76 || err != nil {
      t.Fatal()
   }
}
