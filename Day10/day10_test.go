package main

import (
	"testing"
)

func Test_FindStartingPosition(t *testing.T) {
   lines := []string{
      ".....",
      ".S-7.",
      ".|.|.",
      ".L-J.",
      ".....",
   }

   if row, col, err := FindStartingPosition(lines) ; err != nil || row != 1 || col != 1 {
      t.Fatal()
   }

   lines = []string{
      "7-F7-",
      ".FJ|7",
      "SJLL7",
      "|F--J",
      "LJ.LJ",
   }

   if row, col, err := FindStartingPosition(lines) ; err != nil || row != 2 || col != 0 {
      t.Fatal()
   }
}

func Test_FindPossibleStartingPipes(t *testing.T) {
   lines := []string{
      ".|.",
      "-S.",
      "...",
   }

   startingRow, startingCol, err := FindStartingPosition(lines)

   if err != nil {
      t.Fatal()
   }

   startingPipes := FindPossibleStartingPipes(lines, startingRow, startingCol)

   if len(startingPipes) != 1 || startingPipes[0] != 'J' {
      t.Fatal()
   }

   lines = []string{
      "|..",
      "S..",
      "|..",
   }

   startingRow, startingCol, err = FindStartingPosition(lines)

   if err != nil {
      t.Fatal()
   }

   startingPipes = FindPossibleStartingPipes(lines, startingRow, startingCol)

   if len(startingPipes) != 1 || startingPipes[0] != '|' {
      t.Fatal()
   }
}

func Test_ExampleA(t *testing.T) {
   lines := []string{
      ".....",
      ".S-7.",
      ".|.|.",
      ".L-J.",
      ".....",
   }

   startingRow, startingCol, _ := FindStartingPosition(lines)
   startingPipes := FindPossibleStartingPipes(lines, startingRow, startingCol)
   allNumPipes := CheckStartingPipes(lines, startingRow, startingCol, startingPipes)
   allNumSteps := []float32{}

   for _, numPipes := range allNumPipes {
      allNumSteps = append(allNumSteps, float32(numPipes)/2.0)
   }

   if len(allNumSteps) != 1 || allNumSteps[0] != 4 {
      t.Fatal()
   }
}

func Test_ExampleB(t *testing.T) {
   lines := []string{
      "..F7.",
      ".FJ|.",
      "SJ.L7",
      "|F--J",
      "LJ...",
   }

   startingRow, startingCol, _ := FindStartingPosition(lines)
   startingPipes := FindPossibleStartingPipes(lines, startingRow, startingCol)
   allNumPipes := CheckStartingPipes(lines, startingRow, startingCol, startingPipes)
   allNumSteps := []float32{}

   for _, numPipes := range allNumPipes {
      allNumSteps = append(allNumSteps, float32(numPipes)/2.0)
   }

   if len(allNumSteps) != 1 || allNumSteps[0] != 8 {
      t.Fatal()
   }
}
