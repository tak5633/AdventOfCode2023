package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handfull struct {
   mNumRedCubes int
   mNumGreenCubes int
   mNumBlueCubes int
}

type Game struct {
   mId int
   mHandfulls []Handfull
}

func (game *Game) AddHandfull(pHandfull Handfull) {
   game.mHandfulls = append(game.mHandfulls, pHandfull)
}

func (game *Game) GetMaxRedCubes() int {
   var max int = 0

   for i := 0 ; i < len(game.mHandfulls) ; i++ {
      if game.mHandfulls[i].mNumRedCubes > max {
         max = game.mHandfulls[i].mNumRedCubes
      }
   }

   return max
}

func (game *Game) GetMaxGreenCubes() int {
   var max int = 0

   for i := 0 ; i < len(game.mHandfulls) ; i++ {
      if game.mHandfulls[i].mNumGreenCubes > max {
         max = game.mHandfulls[i].mNumGreenCubes
      }
   }

   return max
}

func (game *Game) GetMaxBlueCubes() int {
   var max int = 0

   for i := 0 ; i < len(game.mHandfulls) ; i++ {
      if game.mHandfulls[i].mNumBlueCubes > max {
         max = game.mHandfulls[i].mNumBlueCubes
      }
   }

   return max
}

func ParseGames(pInputFilepath string) []Game {

   input, err := os.ReadFile(pInputFilepath)
   check(err)

   inputStr := string(input)
   inputLines := strings.Split(inputStr, "\n")

   var games []Game

   for i := 0 ; i < len(inputLines) ; i++ {
      inputLine := inputLines[i]

      if len(inputLine) > 0 {
         game, err := ParseGame(inputLine)
         if err == nil {
            games = append(games, game)
         } else {
            fmt.Println(err)
         }
      }
   }

   return games
}

func check(pE error) {
    if pE != nil {
        panic(pE)
    }
}

func ParseGame(pInput string) (Game, error) {

   var game Game

   input := strings.Trim(pInput, " ")
   gameFields := strings.SplitN(input, ":", 2)

   if len(gameFields) != 2 {
      return Game{}, errors.New("ParseGame: Not enough game fields")
   }

   id, err := ParseGameId(gameFields[0])

   if err != nil {
      return Game{}, err
   }

   game.mId = int(id)

   handfullInputs := gameFields[1]
   handfulls := ParseHandfulls(handfullInputs)

   for i := 0 ; i < len(handfulls) ; i++ {
      game.AddHandfull(handfulls[i])
   }

   return game, nil
}

func ParseGameId(pInput string) (int, error) {

   input := strings.Trim(pInput, " ")
   gameIdFields := strings.SplitN(input, " ", 2)

   if len(gameIdFields) != 2 {
      return -1, errors.New("ParseGameId: Not enough game Id fields")
   }

   id, err := strconv.ParseInt(gameIdFields[1], 10, 0)

   if err != nil {
      return -1, errors.New("ParseHandfull: Failed to parse the game Id")
   }

   return int(id), nil
}

func ParseHandfulls(pInput string) []Handfull {

   var handfulls []Handfull

   input := strings.Trim(pInput, " ")
   handfullInputs := strings.Split(input, ";")

   for i := 0 ; i < len(handfullInputs) ; i++ {
      handfull, err := ParseHandfull(handfullInputs[i])
      if err == nil {
         handfulls = append(handfulls, handfull)
      } else {
         fmt.Println(err)
      }
   }

   return handfulls
}

func ParseHandfull(pInput string) (Handfull, error) {

   var handfull Handfull

   cubeEntries := strings.Split(pInput, ",")

   for i := 0 ; i < len(cubeEntries) ; i++ {
      cubeEntry := strings.Trim(cubeEntries[i], " ")
      cubeFields := strings.Split(cubeEntry, " ")

      var trimmedCubeFields []string

      for cubeFieldIdx := 0 ; cubeFieldIdx < len(cubeFields) ; cubeFieldIdx++ {
         if len(cubeFields[cubeFieldIdx]) != 0 {
            trimmedCubeFields = append(trimmedCubeFields, cubeFields[cubeFieldIdx])
         }
      }

      if len(trimmedCubeFields) < 2 {
         return Handfull{}, errors.New("ParseHandfull: Not enough cube fields")
      }

      numCubes, err := strconv.ParseInt(trimmedCubeFields[0], 10, 0)
      if err != nil {
         return Handfull{}, errors.New("ParseHandfull: Failed to parse the number of cubes")
      }

      cubeColor := strings.ToLower(trimmedCubeFields[1])
      if cubeColor == "red" {
         handfull.mNumRedCubes += int(numCubes)
      } else if cubeColor == "green" {
         handfull.mNumGreenCubes += int(numCubes)
      } else if cubeColor == "blue" {
         handfull.mNumBlueCubes += int(numCubes)
      } else {
         return Handfull{}, errors.New("ParseHandfull: Failed to parse the cube color (" + cubeColor + ")")
      }
   }

   return handfull, nil
}

func main() {
   fmt.Println("Part 1")

   games := ParseGames("./part1Input.txt")

   possibleGameIdsSum := 0

   for i := 0 ; i < len(games) ; i++ {
      game := games[i]

      if game.GetMaxRedCubes() <= 12 && game.GetMaxGreenCubes() <= 13 && game.GetMaxBlueCubes() <= 14 {
         possibleGameIdsSum += game.mId
      }
   }

   fmt.Println("Possible Game Ids Sum:", possibleGameIdsSum)
}
