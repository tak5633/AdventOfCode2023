package main

import (
	"testing"
)

func Test_ParseEnvVar(t *testing.T) {
   input := "12 15 18 21 24 27 30 33 36 39 42 45 48 51 54 57 60 63 66 69 72"
   envVar := ParseEnvVar(input)

   if len(envVar.mReadings) != 21 ||
      envVar.mReadings[0] != 12 || envVar.mReadings[1] != 15 || envVar.mReadings[2] != 18 ||
      envVar.mReadings[3] != 21 || envVar.mReadings[4] != 24 || envVar.mReadings[5] != 27 ||
      envVar.mReadings[6] != 30 || envVar.mReadings[7] != 33 || envVar.mReadings[8] != 36 ||
      envVar.mReadings[9] != 39 || envVar.mReadings[10] != 42 || envVar.mReadings[11] != 45 ||
      envVar.mReadings[12] != 48 || envVar.mReadings[13] != 51 || envVar.mReadings[14] != 54 ||
      envVar.mReadings[15] != 57 || envVar.mReadings[16] != 60 || envVar.mReadings[17] != 63 ||
      envVar.mReadings[18] != 66 || envVar.mReadings[19] != 69 || envVar.mReadings[20] != 72 {

      t.Fatal()
   }
}

func Test_ParseEnvVars(t *testing.T) {
   lines := []string{
      "12 15 18 21 24 27 30 33 36 39 42 45 48 51 54 57 60 63 66 69 72",
      "1 7 22 48 87 141 212 302 413 547 706 892 1107 1353 1632 1946 2297 2687 3118 3592 4111",
   }
   envVars := ParseEnvVars(lines)

   if len(envVars) != 2 ||
      envVars[0].mReadings[0] != 12 || envVars[0].mReadings[1] != 15 || envVars[0].mReadings[2] != 18 ||
      envVars[0].mReadings[3] != 21 || envVars[0].mReadings[4] != 24 || envVars[0].mReadings[5] != 27 ||
      envVars[0].mReadings[6] != 30 || envVars[0].mReadings[7] != 33 || envVars[0].mReadings[8] != 36 ||
      envVars[0].mReadings[9] != 39 || envVars[0].mReadings[10] != 42 || envVars[0].mReadings[11] != 45 ||
      envVars[0].mReadings[12] != 48 || envVars[0].mReadings[13] != 51 || envVars[0].mReadings[14] != 54 ||
      envVars[0].mReadings[15] != 57 || envVars[0].mReadings[16] != 60 || envVars[0].mReadings[17] != 63 ||
      envVars[0].mReadings[18] != 66 || envVars[0].mReadings[19] != 69 || envVars[0].mReadings[20] != 72 ||

      envVars[1].mReadings[0] != 1 || envVars[1].mReadings[1] != 7 || envVars[1].mReadings[2] != 22 ||
      envVars[1].mReadings[3] != 48 || envVars[1].mReadings[4] != 87 || envVars[1].mReadings[5] != 141 ||
      envVars[1].mReadings[6] != 212 || envVars[1].mReadings[7] != 302 || envVars[1].mReadings[8] != 413 ||
      envVars[1].mReadings[9] != 547 || envVars[1].mReadings[10] != 706 || envVars[1].mReadings[11] != 892 ||
      envVars[1].mReadings[12] != 1107 || envVars[1].mReadings[13] != 1353 || envVars[1].mReadings[14] != 1632 ||
      envVars[1].mReadings[15] != 1946 || envVars[1].mReadings[16] != 2297 || envVars[1].mReadings[17] != 2687 ||
      envVars[1].mReadings[18] != 3118 || envVars[1].mReadings[19] != 3592 || envVars[1].mReadings[20] != 4111{

      t.Fatal()
   }
}

func Test_Difference(t *testing.T) {
   inputs := []int{0, 3, 6, 9, 12, 15}

   difference := Difference(inputs)

   if len(difference) != 5 ||
      difference[0] != 3 ||
      difference[1] != 3 ||
      difference[2] != 3 ||
      difference[3] != 3 ||
      difference[4] != 3 {
      t.Fatal()
   }
}

func Test_IsAllZeros(t *testing.T) {
   input := []int{0, 0, 0, 0}

   if IsAllZeros(input) != true {
      t.Fatal()
   }

   input = []int{0, 1, 0, 0}

   if IsAllZeros(input) != false {
      t.Fatal()
   }

   input = []int{-1, 0, 0, 0}

   if IsAllZeros(input) != false {
      t.Fatal()
   }

   input = []int{0, 0, 0, 1}

   if IsAllZeros(input) != false {
      t.Fatal()
   }
}

func Test_ReadingDiffs(t *testing.T) {
   envVar := EnvVar{}
   envVar.mReadings = []int{0, 3, 6, 9, 12, 15}

   if diffs := envVar.ReadingDiffs() ; len(diffs) != 3 ||
      diffs[0][0] != 0 || diffs[0][1] != 3 || diffs[0][2] != 6 || diffs[0][3] != 9 || diffs[0][4] != 12 || diffs[0][5] != 15 ||
      diffs[1][0] != 3 || diffs[1][1] != 3 || diffs[1][2] != 3 || diffs[1][3] != 3 || diffs[1][4] != 3 ||
      diffs[2][0] != 0 || diffs[2][1] != 0 || diffs[2][2] != 0 || diffs[2][3] != 0 {
      t.Fatal()
   }
}

func Test_EnvVar_Extrapolate(t *testing.T) {
   envVar := EnvVar{}
   envVar.mReadings = []int{0, 3, 6, 9, 12, 15}

   extrapolatedReading := envVar.Extrapolate()

   if extrapolatedReading != 18 {
      t.Fatal()
   }
}

func Test_EnvVar_ExtrapolateBackwards(t *testing.T) {
   envVar := EnvVar{}
   envVar.mReadings = []int{10, 13, 16, 21, 30, 45}

   extrapolatedReading := envVar.ExtrapolateBackwards()

   if extrapolatedReading != 5 {
      t.Fatal()
   }
}
