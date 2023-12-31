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
   allPositions := CheckStartingPipes(lines, startingRow, startingCol, startingPipes)

   allNumPipes := []int{}
   allNumSteps := []float32{}

   for _, positions := range allPositions {
      numPipes := len(positions)
      allNumPipes = append(allNumPipes, numPipes)
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
   allPositions := CheckStartingPipes(lines, startingRow, startingCol, startingPipes)

   allNumPipes := []int{}
   allNumSteps := []float32{}

   for _, positions := range allPositions {
      numPipes := len(positions)
      allNumPipes = append(allNumPipes, numPipes)
      allNumSteps = append(allNumSteps, float32(numPipes)/2.0)
   }

   if len(allNumSteps) != 1 || allNumSteps[0] != 8 {
      t.Fatal()
   }
}

func Test_Part2ExampleA(t *testing.T) {
   lines := []string{
      "...........",
      ".S-------7.",
      ".|F-----7|.",
      ".||.....||.",
      ".||.....||.",
      ".|L-7.F-J|.",
      ".|..|.|..|.",
      ".L--J.L--J.",
      "...........",
   }

   startingRow, startingCol, _ := FindStartingPosition(lines)
   startingPipes := FindPossibleStartingPipes(lines, startingRow, startingCol)
   allPositions := CheckStartingPipes(lines, startingRow, startingCol, startingPipes)

   if len(allPositions) != 1 {
      t.Fatal()
   }

   numRows := len(lines)
   numCols := len(lines[0])

   numEnclosedTiles := GetNumEnclosedTiles(numRows, numCols, allPositions[0])

   if numEnclosedTiles != 4 {
      t.Fatal()
   }
}

func Test_Part2ExampleB(t *testing.T) {
   lines := []string{
      ".F----7F7F7F7F-7....",
      ".|F--7||||||||FJ....",
      ".||.FJ||||||||L7....",
      "FJL7L7LJLJ||LJ.L-7..",
      "L--J.L7...LJS7F-7L7.",
      "....F-J..F7FJ|L7L7L7",
      "....L7.F7||L7|.L7L7|",
      ".....|FJLJ|FJ|F7|.LJ",
      "....FJL-7.||.||||...",
      "....L---J.LJ.LJLJ...",
   }

   startingRow, startingCol, _ := FindStartingPosition(lines)
   startingPipes := FindPossibleStartingPipes(lines, startingRow, startingCol)
   allPositions := CheckStartingPipes(lines, startingRow, startingCol, startingPipes)

   if len(allPositions) != 1 {
      t.Fatal()
   }

   numRows := len(lines)
   numCols := len(lines[0])

   numEnclosedTiles := GetNumEnclosedTiles(numRows, numCols, allPositions[0])

   if numEnclosedTiles != 8 {
      t.Fatal()
   }
}

func Test_Part2ExampleC(t *testing.T) {
   lines := []string{
      "FF7FSF7F7F7F7F7F---7",
      "L|LJ||||||||||||F--J",
      "FL-7LJLJ||||||LJL-77",
      "F--JF--7||LJLJ7F7FJ-",
      "L---JF-JLJ.||-FJLJJ7",
      "|F|F-JF---7F7-L7L|7|",
      "|FFJF7L7F-JF7|JL---7",
      "7-L-JL7||F7|L7F-7F7|",
      "L.L7LFJ|||||FJL7||LJ",
      "L7JLJL-JLJLJL--JLJ.L",
   }

   startingRow, startingCol, _ := FindStartingPosition(lines)
   startingPipes := FindPossibleStartingPipes(lines, startingRow, startingCol)
   allPositions := CheckStartingPipes(lines, startingRow, startingCol, startingPipes)

   if len(allPositions) != 1 {
      t.Fatal()
   }

   numRows := len(lines)
   numCols := len(lines[0])

   numEnclosedTiles := GetNumEnclosedTiles(numRows, numCols, allPositions[0])

   if numEnclosedTiles != 10 {
      t.Fatal()
   }
}
