package main

import (
	ci "day11/cosmicImage"
   "testing"
)

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

   cosmicImage := ci.CosmicImage{}
   cosmicImage.SetImage(lines)

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

   if len(cosmicImage.Image()) != len(expandedImage) {
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

   cosmicImage := ci.CosmicImage{}
   cosmicImage.SetImage(lines)

   if len(cosmicImage.Galaxies()) != 9 {
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

   cosmicImage := ci.CosmicImage{}
   cosmicImage.SetImage(lines)
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

   cosmicImage := ci.CosmicImage{}
   cosmicImage.SetImage(lines)
   galaxyPaths := cosmicImage.ComputeGalaxyPaths()
   sumOfGalaxyPaths := SumGalaxyPaths(galaxyPaths)

   if sumOfGalaxyPaths != 374 {
      t.Fatal()
   }
}