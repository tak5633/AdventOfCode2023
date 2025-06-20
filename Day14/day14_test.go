package main

import (
	"log"
	"testing"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_CalculateColumnStringLoad(t *testing.T) {

	{
		col := "........##"
		load := CalculateColumnStringLoad(col)

		expected := 0

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".....O...."
		load := CalculateColumnStringLoad(col)

		expected := 5

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "...O.#...O"
		load := CalculateColumnStringLoad(col)

		expected := 8

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "...#.....O"
		load := CalculateColumnStringLoad(col)

		expected := 1

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".#....O.OO"
		load := CalculateColumnStringLoad(col)

		expected := 7

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "#.#.O.#.##"
		load := CalculateColumnStringLoad(col)

		expected := 6

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "..#.OO..#."
		load := CalculateColumnStringLoad(col)

		expected := 11

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "....O#.O#."
		load := CalculateColumnStringLoad(col)

		expected := 9

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".O..#..O.."
		load := CalculateColumnStringLoad(col)

		expected := 12

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".#...#OOOO"
		load := CalculateColumnStringLoad(col)

		expected := 10

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_TiltColumnString(t *testing.T) {

	{
		col := "OO.O.O..##"
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 34

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "...OO....O"
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 27

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".O...#O..O"
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 17

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".O.#......"
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 10

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".#.O......"
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 8

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "#.#..O#.##"
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 7

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "..#...O.#."
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 7

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "....O#.O#."
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 14

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "....#....."
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 0

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".#.O.#O..."
		col = TiltColumnString(col)
		load := CalculateColumnStringLoad(col)

		expected := 12

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_TiltAndCalculateLoad(t *testing.T) {

	var platform []string

	platform = append(platform, "O....#....")
	platform = append(platform, "O.OO#....#")
	platform = append(platform, ".....##...")
	platform = append(platform, "OO.#O....O")
	platform = append(platform, ".O.....O#.")
	platform = append(platform, "O.#..O.#.#")
	platform = append(platform, "..O..#O..O")
	platform = append(platform, ".......O..")
	platform = append(platform, "#....###..")
	platform = append(platform, "#OO..#....")

	platform = Tilt(platform)
	load := CalculateLoad(platform)

	expected := 136

	if load != expected {
		log.Println(load)
		log.Println(expected)
		t.Fatal()
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_Tilt(t *testing.T) {

	var platform []string

	platform = append(platform, "O....#....")
	platform = append(platform, "O.OO#....#")
	platform = append(platform, ".....##...")
	platform = append(platform, "OO.#O....O")
	platform = append(platform, ".O.....O#.")
	platform = append(platform, "O.#..O.#.#")
	platform = append(platform, "..O..#O..O")
	platform = append(platform, ".......O..")
	platform = append(platform, "#....###..")
	platform = append(platform, "#OO..#....")

	var expectedPlatform []string

   expectedPlatform = append(expectedPlatform, "OOOO.#.O..")
   expectedPlatform = append(expectedPlatform, "OO..#....#")
   expectedPlatform = append(expectedPlatform, "OO..O##..O")
   expectedPlatform = append(expectedPlatform, "O..#.OO...")
   expectedPlatform = append(expectedPlatform, "........#.")
   expectedPlatform = append(expectedPlatform, "..#....#.#")
   expectedPlatform = append(expectedPlatform, "..O..#.O.O")
   expectedPlatform = append(expectedPlatform, "..O.......")
   expectedPlatform = append(expectedPlatform, "#....###..")
   expectedPlatform = append(expectedPlatform, "#....#....")

	tiltedPlatform := Tilt(platform)

   if len(tiltedPlatform) != len(expectedPlatform) {
		log.Println(len(tiltedPlatform))
		log.Println(len(expectedPlatform))
      t.Fatal()
   }

   if len(tiltedPlatform[0]) != len(expectedPlatform[0]) {
		log.Println(len(tiltedPlatform[0]))
		log.Println(len(expectedPlatform[0]))
      t.Fatal()
   }

   for row := range len(tiltedPlatform) {
      for col := range len(tiltedPlatform[0]) {
         if tiltedPlatform[row][col] != expectedPlatform[row][col] {
				log.Println(row, col)
				log.Println(string(tiltedPlatform[row][col]))
				log.Println(string(expectedPlatform[row][col]))
            t.Fatal()
         }
      }
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_CalculateLoadAfterSpins(t *testing.T) {

	var platform []string

	platform = append(platform, "O....#....")
	platform = append(platform, "O.OO#....#")
	platform = append(platform, ".....##...")
	platform = append(platform, "OO.#O....O")
	platform = append(platform, ".O.....O#.")
	platform = append(platform, "O.#..O.#.#")
	platform = append(platform, "..O..#O..O")
	platform = append(platform, ".......O..")
	platform = append(platform, "#....###..")
	platform = append(platform, "#OO..#....")

	{
		numSpins := 1000000000
		load := CalculateLoadAfterSpins(platform, numSpins)

		expected := 64

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_RotateClockwise90(t *testing.T) {

	var platform []string

	platform = append(platform, "O....#....")
	platform = append(platform, "O.OO#....#")
	platform = append(platform, ".....##...")
	platform = append(platform, "OO.#O....O")
	platform = append(platform, ".O.....O#.")
	platform = append(platform, "O.#..O.#.#")
	platform = append(platform, "..O..#O..O")
	platform = append(platform, ".......O..")
	platform = append(platform, "#....###..")
	platform = append(platform, "#OO..#....")

	var expected []string

   expected = append(expected, "##..O.O.OO")
   expected = append(expected, "O....OO...")
   expected = append(expected, "O..O#...O.")
   expected = append(expected, "......#.O.")
   expected = append(expected, "......O.#.")
   expected = append(expected, "##.#O..#.#")
   expected = append(expected, ".#.O...#..")
   expected = append(expected, ".#O.#O....")
   expected = append(expected, ".....#....")
   expected = append(expected, "...O#.O.#.")

	spinPlatform := RotateClockwise90(platform)

   if len(spinPlatform) != len(expected) {
      t.Fatal()
   }

   if len(spinPlatform[0]) != len(expected[0]) {
      t.Fatal()
   }

   for row := range len(spinPlatform) {
      for col := range len(spinPlatform[0]) {
         if spinPlatform[row][col] != expected[row][col] {
            t.Fatal()
         }
      }
   }
}
