package main

import (
	"fmt"
	"maps"
)

func main() {
	// basic map
	basicMap := make(map[string]int) // map that has string key mapping to int value
	basicMap["x"] = 10               // x maps to 10
	basicMap["y"] = 5                // y maps to 5
	fmt.Println(basicMap["x"])       // look up what x is in the map
	fmt.Println("Map is: ", basicMap)
	fmt.Printf("this map has %d key(s)\n", len(basicMap))
	delete(basicMap, "x") // remove key of "x" from the map
	fmt.Println("Map is: ", basicMap)

	// get all keys, all values?
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	song1 := "There's "

	// Method 1: Iterate through key, value pairs with basic range (should be the go-to)
	for _, element := range elements {
		song1 += element + "" + " and "
	}
	fmt.Println("Song: ", song1)

	// Method 2: If you're looking at just key or just values, you can use the maps package .Keys and .Values functions
	song2 := "There's "
	for symbol := range maps.Keys(elements) {
		fmt.Println("Symbol: ", symbol)
	}
	for elementName := range maps.Values(elements) {
		song2 += elementName + "" + " and "
	}
	fmt.Println("Song with maps package: ", song2)

	// look up an element not in map
	_, ok := elements["Un"] // ok is a bool meaning if lookup is successful
	if !ok {
		fmt.Println("Un is not in our mini periodic table")
	}

	// map to map (nested dict)

	elementDB := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	fmt.Println("element DB: ", elementDB)

	// add a new element to the DB
	elementDB["Na"] = map[string]string{
		"name":  "Sodium",
		"state": "solid",
	}
	fmt.Println("element DB after adding sodium: ", elementDB)

	// lookup an element in the DB
	el, ok := elementDB["F"]
	if ok {
		fmt.Println("name: ", el["name"])
		fmt.Println("state: ", el["state"])
	}
}
