package main

import "testing"

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_CosmicImage_Expand(t *testing.T) {
   lines := []string{
      "...#......",
      ".......#..",
      "#.........",
      "..........",
      "......#...",
      ".#........",
      ".........#",
      "..........",
      ".......#..",
      "#...#.....",
   }

   cosmicImage := CosmicImage{}
   cosmicImage.mImage = lines
   cosmicImage.Expand()

   expandedImage := []string{
      "....#........",
      ".........#...",
      "#............",
      ".............",
      ".............",
      "........#....",
      ".#...........",
      "............#",
      ".............",
      ".............",
      ".........#...",
      "#....#.......",
   }

   if len(cosmicImage.mImage) != len(expandedImage) {
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_CosmicImage_FindGalaxies(t *testing.T) {
   lines := []string{
      "...#......",
      ".......#..",
      "#.........",
      "..........",
      "......#...",
      ".#........",
      ".........#",
      "..........",
      ".......#..",
      "#...#.....",
   }

   cosmicImage := CosmicImage{}
   cosmicImage.mImage = lines
   cosmicImage.Expand()
   cosmicImage.FindGalaxies()

   if len(cosmicImage.mGalaxies) != 9 {
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_ComputeGalaxyPaths(t *testing.T) {
   lines := []string{
      "...#......",
      ".......#..",
      "#.........",
      "..........",
      "......#...",
      ".#........",
      ".........#",
      "..........",
      ".......#..",
      "#...#.....",
   }

   cosmicImage := CosmicImage{}
   cosmicImage.mImage = lines
   cosmicImage.Expand()
   cosmicImage.FindGalaxies()
   galaxyPaths := cosmicImage.ComputeGalaxyPaths()

   if len(galaxyPaths) != 36 {
      t.Fatal()
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_SumGalaxyPaths(t *testing.T) {
   lines := []string{
      "...#......",
      ".......#..",
      "#.........",
      "..........",
      "......#...",
      ".#........",
      ".........#",
      "..........",
      ".......#..",
      "#...#.....",
   }

   cosmicImage := CosmicImage{}
   cosmicImage.mImage = lines
   cosmicImage.Expand()
   cosmicImage.FindGalaxies()
   galaxyPaths := cosmicImage.ComputeGalaxyPaths()
   sumOfGalaxyPaths := SumGalaxyPaths(galaxyPaths)

   if sumOfGalaxyPaths != 374 {
      t.Fatal()
   }
}
