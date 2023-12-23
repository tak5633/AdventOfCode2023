package main

import (
	"testing"
)

func Test_ParseLocation(t *testing.T) {
   location := ParseLocation("RJK = (DPP, JQR)")

   if location.mName != "RJK" || location.mLeftLocation != "DPP" || location.mRightLocation != "JQR" {
      t.Fatal()
   }
}

func Test_ParseLocations(t *testing.T) {
   lines := []string{
      "RJK = (DPP, JQR)",
      "QLH = (CXC, MXS)",
   }

   locations := ParseLocations(lines)

   if len(lines) != 2 ||
      locations[0].mName != "RJK" || locations[0].mLeftLocation != "DPP" || locations[0].mRightLocation != "JQR" ||
      locations[1].mName != "QLH" || locations[1].mLeftLocation != "CXC" || locations[1].mRightLocation != "MXS" {
      t.Fatal()
   }
}

func Test_ParseInstructions(t *testing.T) {
   lines := []string {
      "LRLRR",
      "",
      "RJK = (DPP, JQR)",
      "QLH = (CXC, MXS)",
   }

   instructions := ParseInstructions(lines)

   if instructions.mDirections != "LRLRR" ||
      instructions.mCurrentDirectionIdx != 0 ||
      instructions.mLocations["RJK"].mName != "RJK" ||
      instructions.mLocations["RJK"].mLeftLocation != "DPP" ||
      instructions.mLocations["RJK"].mRightLocation != "JQR" ||
      instructions.mLocations["QLH"].mName != "QLH" ||
      instructions.mLocations["QLH"].mLeftLocation != "CXC" ||
      instructions.mLocations["QLH"].mRightLocation != "MXS" ||
      instructions.mCurrentLocation != "AAA" ||
      instructions.mDestinationLocation != "ZZZ" {
      t.Fatal()
   }
}

func Test_Example(t *testing.T) {
   lines := []string {
      "LLR",
      "",
      "AAA = (BBB, BBB)",
      "BBB = (AAA, ZZZ)",
      "ZZZ = (ZZZ, ZZZ)",
   }

   instructions := ParseInstructions(lines)
   numSteps := GetPart1NumSteps(instructions)

   if numSteps != 6 {
      t.Fatal()
   }
}

func Test_GetPart2StartingLocations(t *testing.T) {
   lines := []string {
      "LLR",
      "",
      "AAA = (BBB, BBB)",
      "BBB = (AAA, ZZZ)",
      "TTA = (BQP, LTM)",
      "ZZZ = (ZZZ, ZZZ)",
      "KJA = (XSN, LKF)",
   }

   instructions := ParseInstructions(lines)
   startingLocations := GetPart2StartingLocations(instructions)

   if len(startingLocations) != 3 {
      t.Fatal()
   }
}

func Test_Part2_Example(t *testing.T) {
   lines := []string {
      "LR",
      "",
      "11A = (11B, XXX)",
      "11B = (XXX, 11Z)",
      "11Z = (11B, XXX)",
      "22A = (22B, XXX)",
      "22B = (22C, 22C)",
      "22C = (22Z, 22Z)",
      "22Z = (22B, 22B)",
      "XXX = (XXX, XXX)",
   }

   instructions := ParseInstructions(lines)
   numSteps := GetPart2NumSteps(instructions)

   if numSteps != 6 {
      t.Fatal()
   }
}
