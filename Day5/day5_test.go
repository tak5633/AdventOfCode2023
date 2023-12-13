package main

import (
	"fmt"
	"testing"
)

func Test_ParseEntry(t *testing.T) {
   mapEntry, err := ParseEntry("2642418175 2192252668 3835256")

   if err != nil ||
mapEntry.mDestRangeStart != 2642418175 ||
      mapEntry.mSourceRangeStart != 2192252668 ||
      mapEntry.mRangeLength != 3835256 {
      t.Fatal()
   }
}

func Test_ParseEntries(t *testing.T) {
   inputs := []string{
      "2642418175 2192252668 3835256",
      "2646253431 2276158914 101631202",
      "2640809144 3719389865 1609031",
   }

   mapEntries := ParseEntries(inputs)

   if len(mapEntries) != 3 {
      t.Fatal()
   }

   if mapEntries[0].mDestRangeStart != 2642418175 ||
      mapEntries[0].mSourceRangeStart != 2192252668 ||
      mapEntries[0].mRangeLength != 3835256 {
      t.Fatal()
   }

   if mapEntries[1].mDestRangeStart != 2646253431 ||
      mapEntries[1].mSourceRangeStart != 2276158914 ||
      mapEntries[1].mRangeLength != 101631202 {
      t.Fatal()
   }

   if mapEntries[2].mDestRangeStart != 2640809144 ||
      mapEntries[2].mSourceRangeStart != 3719389865 ||
      mapEntries[2].mRangeLength != 1609031 {
      t.Fatal()
   }
}

func Test_ParseCategoryMap(t *testing.T) {
   lines := []string{
      "seed-to-soil map:",
      "2642418175 2192252668 3835256",
      "2646253431 2276158914 101631202",
      "2640809144 3719389865 1609031",
   }

   categoryMap, err := ParseCategoryMap(lines)

   if err != nil ||
      categoryMap.mSource != "seed" ||
      categoryMap.mDest != "soil" ||
      len(categoryMap.mEntries) != 3 {
      t.Fatal()
   }

   if categoryMap.mEntries[0].mDestRangeStart != 2642418175 ||
      categoryMap.mEntries[0].mSourceRangeStart != 2192252668 ||
      categoryMap.mEntries[0].mRangeLength != 3835256 {
      t.Fatal()
   }

   if categoryMap.mEntries[1].mDestRangeStart != 2646253431 ||
      categoryMap.mEntries[1].mSourceRangeStart != 2276158914 ||
      categoryMap.mEntries[1].mRangeLength != 101631202 {
      t.Fatal()
   }

   if categoryMap.mEntries[2].mDestRangeStart != 2640809144 ||
      categoryMap.mEntries[2].mSourceRangeStart != 3719389865 ||
      categoryMap.mEntries[2].mRangeLength != 1609031 {
      t.Fatal()
   }
}

func Test_ParseCategoryMaps(t *testing.T) {
   lines := []string{
      "seeds: 2906422699 6916147 3075226163 146720986 689152391 244427042 279234546 382175449 1105311711 2036236 3650753915 127044950 3994686181 93904335 1450749684 123906789 2044765513 620379445 1609835129 60050954",
      "",
      "seed-to-soil map:",
      "2642418175 2192252668 3835256",
      "2646253431 2276158914 101631202",
      "2640809144 3719389865 1609031",
      "",
      "soil-to-fertilizer map:",
      "1486714106 1238503832 507721065",
      "637816737 149749818 437782225",
      "1182620803 2675299784 39248251",
   }

   categoryMaps := ParseCategoryMaps(lines)

   if len(categoryMaps) != 2 {
      t.Fatal()
   }
}

func Test_LinkCategoryMaps(t *testing.T) {

   categoryMaps := []CategoryMap{}

   categoryMap1 := CategoryMap{}
   categoryMap1.mSource = "seed"
   categoryMap1.mDest = "soil"
   categoryMaps = append(categoryMaps, categoryMap1)

   categoryMap2 := CategoryMap{}
   categoryMap2.mSource = "soil"
   categoryMap2.mDest = "fertilizer"
   categoryMaps = append(categoryMaps, categoryMap2)

   LinkCategoryMaps(&categoryMaps)

   if categoryMaps[0].mDestMap != &categoryMaps[1] {
      t.Fatal()
   }
}

func Test_ParseSeedNumbers(t *testing.T) {
   seedNumbers := ParseSeedNumbers("seeds: 2906422699 6916147 3075226163 146720986 689152391")

   if len(seedNumbers) != 5 {
      fmt.Println(len(seedNumbers))
      t.Fatal()
   }
}

func Test_GetCategoryMap(t *testing.T) {

   categoryMaps := []CategoryMap{}

   categoryMap1 := CategoryMap{}
   categoryMap1.mSource = "seed"
   categoryMap1.mDest = "soil"
   categoryMaps = append(categoryMaps, categoryMap1)

   categoryMap2 := CategoryMap{}
   categoryMap2.mSource = "soil"
   categoryMap2.mDest = "fertilizer"
   categoryMaps = append(categoryMaps, categoryMap2)

   soilMap, found := GetCategoryMap(categoryMaps, "soil")

   if found == false || soilMap.mSource != "soil" || soilMap.mDest != "fertilizer" {
      t.Fatal()
   }
}

func Test_GetDestNumber_Example(t *testing.T) {
   lines := []string{
      "seeds: 79 14 55 13",
      "",
      "seed-to-soil map:",
      "50 98 2",
      "52 50 48",
      "",
      "soil-to-fertilizer map:",
      "0 15 37",
      "37 52 2",
      "39 0 15",
      "",
      "fertilizer-to-water map:",
      "49 53 8",
      "0 11 42",
      "42 0 7",
      "57 7 4",
      "",
      "water-to-light map:",
      "88 18 7",
      "18 25 70",
      "",
      "light-to-temperature map:",
      "45 77 23",
      "81 45 19",
      "68 64 13",
      "",
      "temperature-to-humidity map:",
      "0 69 1",
      "1 0 69",
      "",
      "humidity-to-location map:",
      "60 56 37",
      "56 93 4",
   }

   categoryMaps := ParseCategoryMaps(lines)
   seedMap, found := GetCategoryMap(categoryMaps, "seed")

   if found == false {
      t.Fatal()
   }

   destNumber := seedMap.GetDestNumber(10) ; if destNumber != 10 {
      t.Fatal()
   }

   destNumber = seedMap.GetDestNumber(98) ; if destNumber != 50 {
      t.Fatal()
   }

   destNumber = seedMap.GetDestNumber(99) ; if destNumber != 51 {
      t.Fatal()
   }

   destNumber = seedMap.GetDestNumber(51) ; if destNumber != 53 {
      t.Fatal()
   }

   destNumber = seedMap.GetDestNumber(97) ; if destNumber != 99 {
      t.Fatal()
   }
}

func Test_FindNumber_Example(t *testing.T) {
   lines := []string{
      "seeds: 79 14 55 13",
      "",
      "seed-to-soil map:",
      "50 98 2",
      "52 50 48",
      "",
      "soil-to-fertilizer map:",
      "0 15 37",
      "37 52 2",
      "39 0 15",
      "",
      "fertilizer-to-water map:",
      "49 53 8",
      "0 11 42",
      "42 0 7",
      "57 7 4",
      "",
      "water-to-light map:",
      "88 18 7",
      "18 25 70",
      "",
      "light-to-temperature map:",
      "45 77 23",
      "81 45 19",
      "68 64 13",
      "",
      "temperature-to-humidity map:",
      "0 69 1",
      "1 0 69",
      "",
      "humidity-to-location map:",
      "60 56 37",
      "56 93 4",
   }

   categoryMaps := ParseCategoryMaps(lines)
   seedMap, found := GetCategoryMap(categoryMaps, "seed")

   if found == false {
      t.Fatal()
   }

   locationNumber := seedMap.FindNumber(79, "location") ; if locationNumber != 82 {
      t.Fatal()
   }

   locationNumber = seedMap.FindNumber(14, "location") ; if locationNumber != 43 {
      t.Fatal()
   }

   locationNumber = seedMap.FindNumber(55, "location") ; if locationNumber != 86 {
      t.Fatal()
   }

   locationNumber = seedMap.FindNumber(13, "location") ; if locationNumber != 35 {
      t.Fatal()
   }
}
