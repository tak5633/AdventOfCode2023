package main

import (
	"log"
	"testing"
)

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_GetStepHash(t *testing.T) {

	{
		input := "HASH"
		hash := GetStepHash(input)
		expected := 52

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "rn=1"
		hash := GetStepHash(input)
		expected := 30

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}
	
	{
		input := "cm-"
		hash := GetStepHash(input)
		expected := 253

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "qp=3"
		hash := GetStepHash(input)
		expected := 97

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "cm=2"
		hash := GetStepHash(input)
		expected := 47

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "qp-"
		hash := GetStepHash(input)
		expected := 14

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "pc=4"
		hash := GetStepHash(input)
		expected := 180

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "ot=9"
		hash := GetStepHash(input)
		expected := 9

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "ab=5"
		hash := GetStepHash(input)
		expected := 197

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "pc-"
		hash := GetStepHash(input)
		expected := 48

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "pc=6"
		hash := GetStepHash(input)
		expected := 214

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		input := "ot=7"
		hash := GetStepHash(input)
		expected := 231

		if hash != expected {
		   log.Println(hash)
			log.Println(expected)
			t.Fatal()
		}
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_GetHash(t *testing.T) {

	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	hash := GetHash(input)
	expected := 1320

	if hash != expected {
		log.Println(hash)
		log.Println(expected)
		t.Fatal()
	}
}
