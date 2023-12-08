package main

import (
	"fmt"
	"testing"
)

func Test_GetMaxCubes(t *testing.T) {
   game := Game{}

   maxRedCubes := game.GetMaxRedCubes()
   if maxRedCubes != 0 {
      t.Fatal()
   }

   maxGreenCubes := game.GetMaxGreenCubes()
   if maxGreenCubes != 0 {
      t.Fatal()
   }

   maxBlueCubes := game.GetMaxBlueCubes()
   if maxBlueCubes != 0 {
      t.Fatal()
   }

   handfull := Handfull{1, 2, 3}
   game.AddHandfull(handfull)

   maxRedCubes = game.GetMaxRedCubes()
   if maxRedCubes != 1 {
      t.Fatal()
   }

   maxGreenCubes = game.GetMaxGreenCubes()
   if maxGreenCubes != 2 {
      t.Fatal()
   }

   maxBlueCubes = game.GetMaxBlueCubes()
   if maxBlueCubes != 3 {
      t.Fatal()
   }

   biggerHandfull := Handfull{4, 8, 12}
   game.AddHandfull(biggerHandfull)

   lessBigHandfull := Handfull{2, 4, 6}
   game.AddHandfull(lessBigHandfull)

   maxRedCubes = game.GetMaxRedCubes()
   if maxRedCubes != 4 {
      t.Fatal()
   }

   maxGreenCubes = game.GetMaxGreenCubes()
   if maxGreenCubes != 8 {
      t.Fatal()
   }

   maxBlueCubes = game.GetMaxBlueCubes()
   if maxBlueCubes != 12 {
      t.Fatal()
   }
}

func Test_ParseGameId(t *testing.T) {
   gameId, err := ParseGameId("Game 12")
   if gameId != 12 || err != nil {
      fmt.Println(gameId, err)
      t.Fatal()
   }
}

func Test_ParseHandfull_OneColor(t *testing.T) {
   handfull, err := ParseHandfull("  1  red  ")
   if handfull.mNumRedCubes != 1 || handfull.mNumGreenCubes != 0 || handfull.mNumBlueCubes != 0 || err != nil {
      fmt.Println(err)
      t.Fatal()
   }

   handfull, err = ParseHandfull("  2  green  ")
   if handfull.mNumRedCubes != 0 || handfull.mNumGreenCubes != 2 || handfull.mNumBlueCubes != 0 || err != nil {
      t.Fatal()
   }

   handfull, err = ParseHandfull("  3  blue  ")
   if handfull.mNumRedCubes != 0 || handfull.mNumGreenCubes != 0 || handfull.mNumBlueCubes != 3 || err != nil {
      t.Fatal()
   }
}

func Test_ParseHandfull_MultipleColors(t *testing.T) {
   handfull, err := ParseHandfull("  1  red  ,  2  green  ,  3  blue  ")
   if handfull.mNumRedCubes != 1 || handfull.mNumGreenCubes != 2 || handfull.mNumBlueCubes != 3 || err != nil {
      fmt.Println(err)
      t.Fatal()
   }
}

func Test_ParseHandfulls_Single(t *testing.T) {
   handfulls := ParseHandfulls(" 19 blue, 12 red")
   if len(handfulls) != 1 {
      t.Fatal()
   }
}

func Test_ParseHandfulls_Double(t *testing.T) {
   handfulls := ParseHandfulls(" 19 blue, 12 red; 19 blue, 2 green, 1 red")
   if len(handfulls) != 2 {
      t.Fatal()
   }
}

func Test_ParseHandfulls_Triple(t *testing.T) {
   handfulls := ParseHandfulls(" 19 blue, 12 red; 19 blue, 2 green, 1 red; 13 red, 11 blue")
   if len(handfulls) != 3 {
      t.Fatal()
   }
}

func Test_ParseGame_BadInput(t *testing.T) {
   _, err := ParseGame("Game 1 19 blue, 12 red; 19 blue, 2 green, 1 red; 13 red, 11 blue")
   if err == nil {
      t.Fatal()
   }
}

func Test_ParseGame(t *testing.T) {
   game, err := ParseGame("Game 1: 19 blue, 12 red; 19 blue, 2 green, 1 red; 13 red, 11 blue")
   if game.GetMaxRedCubes() != 13 || game.GetMaxGreenCubes() != 2 || game.GetMaxBlueCubes() != 19 || err != nil {
      t.Fatal()
   }
}
