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

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func AreStepFieldsEqual(pA StepFields, pB StepFields) bool {
   return pA.mLabel == pB.mLabel && pA.mOperation == pB.mOperation && pA.mHasFocalLength == pB.mHasFocalLength && pA.mFocalLength == pB.mFocalLength
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_ParseStep(t *testing.T) {

   {
      input := "rn=1"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "rn"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 1

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "cm-"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "cm"
      expected.mOperation = '-'
      expected.mHasFocalLength = false

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "qp=3"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "qp"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 3

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "cm=2"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "cm"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 2

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "qp-"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "qp"
      expected.mOperation = '-'
      expected.mHasFocalLength = false

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "pc=4"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "pc"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 4

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "ot=9"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "ot"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 9

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "ab=5"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "ab"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 5

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "pc-"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "pc"
      expected.mOperation = '-'
      expected.mHasFocalLength = false

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "pc=6"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "pc"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 6

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }

   {
      input := "ot=7"
      stepFields := ParseStep(input)

      var expected StepFields
      expected.mLabel = "ot"
      expected.mOperation = '='
      expected.mHasFocalLength = true
      expected.mFocalLength = 7

      if !AreStepFieldsEqual(stepFields, expected) {
         log.Println(stepFields)
         log.Println(expected)
         t.Fatal()
      }
   }
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_AddLense(t *testing.T) {
	box := Box{}

	{
		label := "ot"
		focalLength := 9
		box.AddLense(label, focalLength)

      if len(box.mLenseInfoList) != 1 ||
		   box.mLenseInfoList[0].mLabel != label ||
		   box.mLenseInfoList[0].mFocalLength != focalLength {

			log.Println(box.mLenseInfoList[0])
         t.Fatal()
      }
	}

	{
		label := "ab"
		focalLength := 5
		box.AddLense(label, focalLength)

      if len(box.mLenseInfoList) != 2 ||
		   box.mLenseInfoList[1].mLabel != label ||
		   box.mLenseInfoList[1].mFocalLength != focalLength {

			log.Println(box.mLenseInfoList[0])
         t.Fatal()
      }
	}

	{
		label := "ot"
		focalLength := 7
		box.AddLense(label, focalLength)

      if len(box.mLenseInfoList) != 2 ||
		   box.mLenseInfoList[0].mLabel != label ||
		   box.mLenseInfoList[0].mFocalLength != focalLength {

			log.Println(box.mLenseInfoList[0])
         t.Fatal()
      }
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_RemoveLense(t *testing.T) {

	{
		// Beginning
		box := Box{}

		box.AddLense("qp", 0)

		if len(box.mLenseInfoList) != 1 ||
		   box.mLenseInfoList[0].mLabel != "qp" {
			t.Fatal()
		}

		box.RemoveLense("nomatch")

		if len(box.mLenseInfoList) != 1 ||
		   box.mLenseInfoList[0].mLabel != "qp" {
			t.Fatal()
		}

		box.RemoveLense("qp")

		if len(box.mLenseInfoList) != 0 {
			t.Fatal()
		}
	}

	{
		// Middle
		box := Box{}

		box.AddLense("kt", 0)
		box.AddLense("qp", 0)
		box.AddLense("ck", 0)

		if len(box.mLenseInfoList) != 3 ||
		   box.mLenseInfoList[0].mLabel != "kt" ||
		   box.mLenseInfoList[1].mLabel != "qp" ||
		   box.mLenseInfoList[2].mLabel != "ck" {
			t.Fatal()
		}

		box.RemoveLense("nomatch")

		if len(box.mLenseInfoList) != 3 ||
		   box.mLenseInfoList[0].mLabel != "kt" ||
		   box.mLenseInfoList[1].mLabel != "qp" ||
		   box.mLenseInfoList[2].mLabel != "ck" {
			t.Fatal()
		}

		box.RemoveLense("qp")

		if len(box.mLenseInfoList) != 2 ||
		   box.mLenseInfoList[0].mLabel != "kt" ||
		   box.mLenseInfoList[1].mLabel != "ck" {
			t.Fatal()
		}
	}

	{
		// End
		box := Box{}

		box.AddLense("kt", 0)
		box.AddLense("qp", 0)

		if len(box.mLenseInfoList) != 2 ||
		   box.mLenseInfoList[0].mLabel != "kt" ||
		   box.mLenseInfoList[1].mLabel != "qp" {
			t.Fatal()
		}

		box.RemoveLense("nomatch")

		if len(box.mLenseInfoList) != 2 ||
		   box.mLenseInfoList[0].mLabel != "kt" ||
		   box.mLenseInfoList[1].mLabel != "qp" {
			t.Fatal()
		}

		box.RemoveLense("qp")

		if len(box.mLenseInfoList) != 1 ||
		   box.mLenseInfoList[0].mLabel != "kt" {
			t.Fatal()
		}
	}
}

//--------------------------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------------------------
func Test_CalculateFocusingPower(t *testing.T) {

	{
		box := Box{}
		box.mIndex = 0

		{
			var lenseInfo LenseInfo
			lenseInfo.mFocalLength = 1
			box.mLenseInfoList = append(box.mLenseInfoList, lenseInfo)
		}

		{
			var lenseInfo LenseInfo
			lenseInfo.mFocalLength = 2
			box.mLenseInfoList = append(box.mLenseInfoList, lenseInfo)
		}

		focusingPower := box.CalculateFocusingPower()
		expected := 5

		if focusingPower != expected {
			log.Println(focusingPower)
			log.Println(expected)
			t.Fatal()
		}
	}

	{
		box := Box{}
		box.mIndex = 3

		{
			var lenseInfo LenseInfo
			lenseInfo.mFocalLength = 7
			box.mLenseInfoList = append(box.mLenseInfoList, lenseInfo)
		}

		{
			var lenseInfo LenseInfo
			lenseInfo.mFocalLength = 5
			box.mLenseInfoList = append(box.mLenseInfoList, lenseInfo)
		}

		{
			var lenseInfo LenseInfo
			lenseInfo.mFocalLength = 6
			box.mLenseInfoList = append(box.mLenseInfoList, lenseInfo)
		}

		focusingPower := box.CalculateFocusingPower()
		expected := 140

		if focusingPower != expected {
			log.Println(focusingPower)
			log.Println(expected)
			t.Fatal()
		}
	}
}
