package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
   Up = iota
   Down
   Left
   Right
)

func main() {
   Part1()
}

func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   startingRow, startingCol, err := FindStartingPosition(inputLines)
   check(err)

   fmt.Println("Starting Position:", startingRow, ",", startingCol)

   startingPipes := FindPossibleStartingPipes(inputLines, startingRow, startingCol)
   fmt.Println("Starting Pipes:", string(startingPipes))

   allNumPipes := CheckStartingPipes(inputLines, startingRow, startingCol, startingPipes)
   fmt.Println("All Num Pipes:", allNumPipes)

   allNumSteps := []float32{}

   for _, numPipes := range allNumPipes {
      allNumSteps = append(allNumSteps, float32(numPipes)/2.0)
   }

   fmt.Println("All Num Steps:", allNumSteps)
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

func FindStartingPosition(pLines []string) (int, int, error) {
   for row := 0; row < len(pLines); row++ {
      for col := 0; col < len(pLines[row]); col++ {
         if pLines[row][col] == 'S' {
            return row, col, nil
         }
      }
   }

   return -1, -1, errors.New("FindStartingPosition: Failed to find the starting position")
}

func FindPossibleStartingPipes(pLines []string, pStartingRow int, pStartingCol int) []rune {

   var startingPipes []rune

   numRows := len(pLines)
   numCols := len(pLines[0])

   left := '.'

   if pStartingCol > 0 {
      left = rune(pLines[pStartingRow][pStartingCol-1])
   }

   right := '.'

   if pStartingCol < (numCols-1) {
      right = rune(pLines[pStartingRow][pStartingCol+1])
   }

   top := '.'

   if pStartingRow > 0 {
      top = rune(pLines[pStartingRow-1][pStartingCol])
   }

   bottom := '.'

   if pStartingRow < (numRows-1) {
      bottom = rune(pLines[pStartingRow+1][pStartingCol])
   }

   if (left == '-' || left == 'L' || left == 'F') && (top == '|' || top == '7' || top == 'F') {
      startingPipes = append(startingPipes, 'J')
   }

   if (left == '-' || left == 'L' || left == 'F') && (right == '-' || right == 'J' || right == '7') {
      startingPipes = append(startingPipes, '-')
   }

   if (left == '-' || left == 'L' || left == 'F') && (bottom == '|' || bottom == 'L' || bottom == 'J') {
      startingPipes = append(startingPipes, '7')
   }

   if (top == '|' || top == '7' || top == 'F') && (right == '-' || right == 'J' || right == '7') {
      startingPipes = append(startingPipes, 'L')
   }

   if (top == '|' || top == '7' || top == 'F') && (bottom == '|' || bottom == 'L' || bottom == 'J') {
      startingPipes = append(startingPipes, '|')
   }

   if (right == '-' || right == 'J' || right == '7') && (bottom == '|' || bottom == 'L' || bottom == 'J') {
      startingPipes = append(startingPipes, 'F')
   }

   return startingPipes
}

func CheckStartingPipes(pLines []string, pStartingRow int, pStartingCol int, pStartingPipes []rune) []int {
   allNumPipes := []int{}

   for _, startingPipe := range pStartingPipes {
      numPipes, err := CheckStartingPipe(pLines, pStartingRow, pStartingCol, startingPipe)

      if err == nil {
         allNumPipes = append(allNumPipes, numPipes)
      }
   }

   return allNumPipes
}

func CheckStartingPipe(pLines []string, pStartingRow int, pStartingCol int, pStartingPipe rune) (int, error) {

   row := pStartingRow
   col := pStartingCol
   direction := Up

   switch pStartingPipe {
   case '|':
      direction = Up
   case '-':
      direction = Right
   case 'L':
      direction = Up
   case 'J':
      direction = Up
   case '7':
      direction = Down
   case 'F':
      direction = Down
   }

   positionMap := map[string]bool{}

   positionKey := fmt.Sprintf("%d,%d", pStartingRow, pStartingCol)
   positionMap[positionKey] = true

   numPipes := 0
   var err error = nil

   for {
      row, col, err = GetNextPosition(pLines, row, col, direction)

      if err != nil {
         return -1, err
      }

      numPipes += 1

      if row == pStartingRow && col == pStartingCol {
         return numPipes, nil
      }

      positionKey := fmt.Sprintf("%d,%d", row, col)

      if _, ok := positionMap[positionKey] ; ok {
         return -1, errors.New("CheckStartingPipe: Position has already been visited")
      }

      if pLines[row][col] == '.' {
         return -1, errors.New("CheckStartingPipe: The end of the pipe has been reached")
      }

      direction, err = GetNextDirection(pLines, row, col, direction)

      if err != nil {
         return -1, err
      }
   }
}

func GetNextPosition(pLines []string, pRow int, pCol int, pDirection int) (int, int, error) {

   nextPositionError := errors.New("GetNextPosition: The next position exceeds the bounds")

   numRows := len(pLines)
   numCols := len(pLines[0])

   switch pDirection {
   case Up:
      if pRow == 0 {
         return -1, -1, nextPositionError
      }
      return pRow-1, pCol, nil
   case Down:
      if pRow == (numRows-1) {
         return -1, -1, nextPositionError
      }
      return pRow+1, pCol, nil
   case Left:
      if pCol == 0 {
         return -1, -1, nextPositionError
      }
      return pRow, pCol-1, nil
   case Right:
      if pCol == (numCols-1) {
         return -1, -1, nextPositionError
      }
      return pRow, pCol+1, nil
   }

   return -1, -1, nextPositionError
}

func GetNextDirection(pLines []string, pRow int, pCol int, pDirection int) (int, error) {

   pipe := pLines[pRow][pCol]

   switch pipe {
   case '|':
      if pDirection == Up {
         return Up, nil
      } else if pDirection == Down {
         return Down, nil
      }
   case '-':
      if pDirection == Left {
         return Left, nil
      } else if pDirection == Right {
         return Right, nil
      }
   case 'L':
      if pDirection == Down {
         return Right, nil
      } else if pDirection == Left {
         return Up, nil
      }
   case 'J':
      if pDirection == Down {
         return Left, nil
      } else if pDirection == Right {
         return Up, nil
      }
   case '7':
      if pDirection == Right {
         return Down, nil
      } else if pDirection == Up {
         return Left, nil
      }
   case 'F':
      if pDirection == Up {
         return Right, nil
      } else if pDirection == Left {
         return Down, nil
      }
   }

   return -1, errors.New("GetNextDirection: The next direction could not be determined")
}
