package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
   Up Direction = iota
   Down
   Left
   Right
)

type Position struct {
   mIsStart bool
   mRow int
   mCol int
   mInputDirection Direction
   mOutputDirection Direction
   mPrevPosition *Position
   mNextPosition *Position
}

type Tile int

const (
   Inner Tile = iota
   Outer
   Loop
   NoType
)

func main() {
   Part1()
   Part2()
}

func Part1() {
   fmt.Println("Part 1")

   inputLines := ReadInput()

   startingRow, startingCol, err := FindStartingPosition(inputLines)
   check(err)

   fmt.Println("Starting Position:", startingRow, ",", startingCol)

   startingPipes := FindPossibleStartingPipes(inputLines, startingRow, startingCol)
   fmt.Println("Starting Pipes:", string(startingPipes))

   allPositions := CheckStartingPipes(inputLines, startingRow, startingCol, startingPipes)

   allNumPipes := []int{}
   allNumSteps := []float32{}

   for _, positions := range allPositions {
      numPipes := len(positions)
      allNumPipes = append(allNumPipes, numPipes)
      allNumSteps = append(allNumSteps, float32(numPipes)/2.0)
   }

   fmt.Println("All Num Pipes:", allNumPipes)
   fmt.Println("All Num Steps:", allNumSteps)
}

func Part2() {
   fmt.Println("Part 2")

   inputLines := ReadInput()

   startingRow, startingCol, err := FindStartingPosition(inputLines)
   check(err)

   fmt.Println("Starting Position:", startingRow, ",", startingCol)

   startingPipes := FindPossibleStartingPipes(inputLines, startingRow, startingCol)
   fmt.Println("Starting Pipes:", string(startingPipes))

   allPositions := CheckStartingPipes(inputLines, startingRow, startingCol, startingPipes)

   for i, positions := range allPositions {
      fmt.Println("Positions Idx:", i)

      numCumulativeClockwiseTurns := GetNumCumulativeClockwiseTurns(positions)
      fmt.Println("Num Cumulative Clockwise Turns:", numCumulativeClockwiseTurns)

      numRows := len(inputLines)
      numCols := len(inputLines[0])

      numEnclosedTiles := GetNumEnclosedTiles(numRows, numCols, positions)
      fmt.Println("Num Enclosed Tiles:", numEnclosedTiles)
   }
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

func CheckStartingPipes(pLines []string, pStartingRow int, pStartingCol int, pStartingPipes []rune) [][]*Position {
   allPositions := [][]*Position{}

   for _, startingPipe := range pStartingPipes {
      positions, err := CheckStartingPipe(pLines, pStartingRow, pStartingCol, startingPipe)

      if err == nil {
         allPositions = append(allPositions, positions)
      }
   }

   return allPositions
}

func CheckStartingPipe(pLines []string, pStartingRow int, pStartingCol int, pStartingPipe rune) ([]*Position, error) {

   row := pStartingRow
   col := pStartingCol
   direction, err := GetStartingDirection(pStartingPipe)

   if err != nil {
      return []*Position{}, err
   }

   positionMap := map[string]bool{}

   positionKey := fmt.Sprintf("%d,%d", pStartingRow, pStartingCol)
   positionMap[positionKey] = true

   positions := []*Position{}

   startingPosition := Position{}
   startingPosition.mIsStart = true
   startingPosition.mRow = row
   startingPosition.mCol = col
   startingPosition.mOutputDirection = direction
   positions = append(positions, &startingPosition)

   for {
      row, col, err = GetNextPosition(pLines, row, col, direction)

      if err != nil {
         return []*Position{}, err
      }

      if row == pStartingRow && col == pStartingCol {
         positions[0].mInputDirection = positions[len(positions)-1].mOutputDirection
         positions[0].mPrevPosition = positions[len(positions)-1]

         return positions, nil
      }

      positionKey := fmt.Sprintf("%d,%d", row, col)

      if _, ok := positionMap[positionKey] ; ok {
         return []*Position{}, errors.New("CheckStartingPipe: Position has already been visited")
      }

      if pLines[row][col] == '.' {
         return []*Position{}, errors.New("CheckStartingPipe: The end of the pipe has been reached")
      }

      position := Position{}
      position.mIsStart = false
      position.mRow = row
      position.mCol = col
      position.mInputDirection = direction

      position.mPrevPosition = positions[len(positions)-1]
      positions[len(positions)-1].mNextPosition = &position

      direction, err = GetNextDirection(pLines, row, col, direction)

      if err != nil {
         return []*Position{}, err
      }

      position.mOutputDirection = direction
      positions = append(positions, &position)
   }
}

func GetStartingDirection(pStartingPipe rune) (Direction, error) {

   switch pStartingPipe {
   case '|':
      return Up, nil
   case '-':
      return Right, nil
   case 'L':
      return Up, nil
   case 'J':
      return Up, nil
   case '7':
      return Down, nil
   case 'F':
      return Down, nil
   }

   return Up, errors.New("GetStartingDirection: Failed to determine the starting direction")
}

func GetNextPosition(pLines []string, pRow int, pCol int, pDirection Direction) (int, int, error) {

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

func GetNextDirection(pLines []string, pRow int, pCol int, pDirection Direction) (Direction, error) {

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

func GetNumEnclosedTiles(pNumRows int, pNumCols int, pPositions []*Position) int {
   tileMap := map[string]Tile{}

   // Initialize the map
   for row := 0; row < pNumRows; row++ {
      for col := 0; col < pNumCols; col++ {
         tileKey := fmt.Sprintf("%d,%d", row, col)
         tileMap[tileKey] = NoType
      }
   }

   // Identify the loop tiles
   for _, position := range pPositions {
      tileKey := fmt.Sprintf("%d,%d", position.mRow, position.mCol)
      tileMap[tileKey] = Loop
   }

   numCumulativeClockwiseTurns := GetNumCumulativeClockwiseTurns(pPositions)
   isClockwiseLoop := numCumulativeClockwiseTurns > 0

   for _, position := range pPositions {

      // Mark inner tiles by adjusting the row
      row := position.mRow
      col := position.mCol
      var err error = nil

      for {
         row, col, err = GetNextInnerRowPosition(row, col, position.mInputDirection, position.mOutputDirection, isClockwiseLoop)

         tileKey := fmt.Sprintf("%d,%d", row, col)

         if row < 0 || pNumRows <= row || col < 0 || pNumCols <= col || err != nil || tileMap[tileKey] == Loop {
            break
         }

         tileMap[tileKey] = Inner
      }

      // Mark inner tiles by adjusting the column
      row = position.mRow
      col = position.mCol
      err = nil

      for {
         row, col, err = GetNextInnerColPosition(row, col, position.mInputDirection, position.mOutputDirection, isClockwiseLoop)

         tileKey := fmt.Sprintf("%d,%d", row, col)

         if row < 0 || pNumRows <= row || col < 0 || pNumCols <= col || err != nil || tileMap[tileKey] == Loop {
            break
         }

         tileMap[tileKey] = Inner
      }
   }

   numEnclosedTiles := 0

   for _, tile := range tileMap {
      if tile == Inner {
         numEnclosedTiles++
      }
   }

   return numEnclosedTiles
}

func GetNumCumulativeClockwiseTurns(pPositions []*Position) int {

   numCumulativeClockwiseTurns := 0

   for i := 1; i < len(pPositions); i++ {
      switch pPositions[i].mInputDirection {
      case Up:
         if pPositions[i].mOutputDirection == Right {
            numCumulativeClockwiseTurns++
         } else if pPositions[i].mOutputDirection == Left {
            numCumulativeClockwiseTurns--
         }
      case Down:
         if pPositions[i].mOutputDirection == Right {
            numCumulativeClockwiseTurns--
         } else if pPositions[i].mOutputDirection == Left {
            numCumulativeClockwiseTurns++
         }
      case Left:
         if pPositions[i].mOutputDirection == Up {
            numCumulativeClockwiseTurns++
         } else if pPositions[i].mOutputDirection == Down {
            numCumulativeClockwiseTurns--
         }
      case Right:
         if pPositions[i].mOutputDirection == Up {
            numCumulativeClockwiseTurns--
         } else if pPositions[i].mOutputDirection == Down {
            numCumulativeClockwiseTurns++
         }
      }
   }

   return numCumulativeClockwiseTurns
}

func GetNextInnerRowPosition(pRow int, pCol int, pInputDirection Direction, pOutputDirection Direction, pIsClockwiseLoop bool) (int, int, error) {

   // Straight pipe
   if pInputDirection == pOutputDirection {
      if pIsClockwiseLoop {
         switch pInputDirection {
         case Left:
            return pRow-1, pCol, nil
         case Right:
            return pRow+1, pCol, nil
         }
      } else {
         switch pInputDirection {
         case Left:
            return pRow+1, pCol, nil
         case Right:
            return pRow-1, pCol, nil
         }
      }
   }

   // Bent pipe
   if pIsClockwiseLoop {
      if pInputDirection == Up && pOutputDirection == Left {
         return pRow-1, pCol, nil
      } else if pInputDirection == Right && pOutputDirection == Up {
         return pRow+1, pCol, nil
      } else if pInputDirection == Down && pOutputDirection == Right {
         return pRow+1, pCol, nil
      } else if pInputDirection == Left && pOutputDirection == Down {
         return pRow-1, pCol, nil
      }
   } else {
      if pInputDirection == Up && pOutputDirection == Right {
         return pRow-1, pCol, nil
      } else if pInputDirection == Left && pOutputDirection == Up {
         return pRow+1, pCol, nil
      } else if pInputDirection == Down && pOutputDirection == Left {
         return pRow+1, pCol, nil
      } else if pInputDirection == Right && pOutputDirection == Down {
         return pRow-1, pCol, nil
      }
   }

   return pRow, pCol, errors.New("GetNextInnerRowPosition: Cannot retrieve next inner position")
}

func GetNextInnerColPosition(pRow int, pCol int, pInputDirection Direction, pOutputDirection Direction, pIsClockwiseLoop bool) (int, int, error) {

   // Straight pipe
   if pInputDirection == pOutputDirection {
      if pIsClockwiseLoop {
         switch pInputDirection {
         case Up:
            return pRow, pCol+1, nil
         case Down:
            return pRow, pCol-1, nil
         }
      } else {
         switch pInputDirection {
         case Up:
            return pRow, pCol-1, nil
         case Down:
            return pRow, pCol+1, nil
         }
      }
   }

   // Bent pipe
   if pIsClockwiseLoop {
      if pInputDirection == Up && pOutputDirection == Left {
         return pRow, pCol+1, nil
      } else if pInputDirection == Right && pOutputDirection == Up {
         return pRow, pCol+1, nil
      } else if pInputDirection == Down && pOutputDirection == Right {
         return pRow, pCol-1, nil
      } else if pInputDirection == Left && pOutputDirection == Down {
         return pRow, pCol-1, nil
      }
   } else {
      if pInputDirection == Up && pOutputDirection == Right {
         return pRow, pCol-1, nil
      } else if pInputDirection == Left && pOutputDirection == Up {
         return pRow, pCol-1, nil
      } else if pInputDirection == Down && pOutputDirection == Left {
         return pRow, pCol+1, nil
      } else if pInputDirection == Right && pOutputDirection == Down {
         return pRow, pCol+1, nil
      }
   }

   return pRow, pCol, errors.New("GetNextInnerColPosition: Cannot retrieve next inner position")
}
