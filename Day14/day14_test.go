package main

import (
	"log"
	"testing"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_CalculateTotalColumnLoad(t *testing.T) {
	{
		col := "OOOO....##"
		load := CalculateTotalColumnLoad(col)

		expected := 34

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "OO.O.O..##"
		load := CalculateTotalColumnLoad(col)

		expected := 34

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "...OO....O"
		load := CalculateTotalColumnLoad(col)

		expected := 27

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".O...#O..O"
		load := CalculateTotalColumnLoad(col)

		expected := 17

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".O.#......"
		load := CalculateTotalColumnLoad(col)

		expected := 10

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".#.O......"
		load := CalculateTotalColumnLoad(col)

		expected := 8

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "#.#..O#.##"
		load := CalculateTotalColumnLoad(col)

		expected := 7

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "..#...O.#."
		load := CalculateTotalColumnLoad(col)

		expected := 7

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "....O#.O#."
		load := CalculateTotalColumnLoad(col)

		expected := 14

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := "....#....."
		load := CalculateTotalColumnLoad(col)

		expected := 0

		if load != expected {
			log.Println(load)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		col := ".#.O.#O..."
		load := CalculateTotalColumnLoad(col)

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
func Test_CalculateTotalLoad(t *testing.T) {

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

	load := CalculateTotalLoad(platform)

	expected := 136

	if load != expected {
		log.Println(load)
		log.Println(expected)
		t.Fatal()
	}
}
