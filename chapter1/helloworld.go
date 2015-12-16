package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// Single line comment
	/*
		Multiline
		comment
	*/

	// Initialize variable
	myString := "This is a string."
	// Use just = to change value
	myString = "This is really a string."

	// Concatenate string
	myOtherString := myString + " Sure it is."

	// Print value of variable
	fmt.Println(myOtherString)

	/* Numbers
		int8, int16, int32, int64
		uint8, uint16, ...
		float32, float64

		int and uint depends on system architecture,
		i.e. on a 64bit system int is int64
	*/
	myInt := 2
	myFloat := 1.5

	// Type conversion/casting
	myOtherFloat := float64(myInt) * myFloat
	fmt.Println(myOtherFloat)

	// Boolean
	myBool := true
	if myBool {
		fmt.Println("myBool is true")
	} else {
		fmt.Println("myBool is false")
	}

	// Slices
	mySlice := []int{1, 2, 3, 4, 5}
	// Assign and not using a variable gives runtime error
	// mySlice2 := mySlice[0:3]	// [1, 2, 3]
	// mySlice3 := mySlice[1:4]	// [2, 3, 4]
	fmt.Println(mySlice[2])		// 3
	// fmt.Println(mySlice[7]) // Gives a runtime error, not compile error

	fmt.Println(len(mySlice))	// 5
	myStringSlice := []string{"Hello", "World"}
	fmt.Println(len(myStringSlice))	// 2

	// For statement
	animals := []string{"Cat", "Dog", "Emu", "Warthog"}
	for i, animal := range animals {
		fmt.Println(animal, "is at index", i)	// No need for extra spaces
	}
	// if there are no plans to use a value e.g. i, we must assign it to
	// the blank identifier "_"
	for _, animal := range animals {
		fmt.Println(animal)	// No need for extra spaces
	}

	// Maps - key/value pairs
	starWarsYears := map[string]int{
		"A New Hope":				1977,
		"The Empire Strikes Again":	1977,
		"Return of the Jedi":		1980,
		"Attack of the Clones":		1983,
		"Revenge of the Sith":		2005,
	}
	// Add new item in map
	starWarsYears["The Force Awakens"] = 2015
	fmt.Println(len(starWarsYears))	// 6

	// Looping over maps
	// Loops on maps iterate through the map in random order!!
	for title, year := range starWarsYears {
		fmt.Println(title, "was relased in", year)
	}

	colours := map[string]string {
		"red":		"#ff0000",
		"green":	"#00ff00",
		"blue":		"#0000ff",
		"fuchsia":	"#ff00ff",
	}
	// redHexCode := colours["red"]		// "#ff0000"
	// purpleHexCode := colours["purple"]	// empty value of the maps type. In this case ""

	// Delete value from map
	delete(colours, "fuchsia")			// Returns nothing. Does nothing if key does not exist

	// To check key's existence use the special two-value asignment
	code, exists := colours["burgundy"]
	if exists {
		fmt.Println("Burgundy's hex code is", code)
	} else {
		fmt.Println("Not known hex for Burgundy")
	}
}

// Functions
func noParamsNoReturn() {
	fmt.Println("Not doin' much really...")
}

func twoParamsOneReturn(myInt int, myString string) string {
	return fmt.Sprintf("myInt: %d, myString: %s", myInt, myString)
}

func oneParamTwoReturns(myInt int) (string, int) {
	return fmt.Sprintf("Int: %d", myInt), myInt + 1
}
// a, b := oneParamTworeturns(5)	// a -> Int: 5, b -> 6

func twoSameTypedParams(myString1, myString2 string) {
	fmt.Println("String 1", myString1)
	fmt.Println("String 2", myString2)
}

// Pointers (passing by reference)
// func giveMePear(fruit *string) {
// 	*fruit = "pear"
// }
// func main() {
// 	fruit := "banana"
// 	giveMePear(&fruit)
// 	fmt.Println(fruit)	// "pear"
// }

// Accessing or altering nil pointer (no value) produces runtime error