package main

import "fmt"

func main() {
//	version_1()
	version_2()
}

/*
func version_1() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])
}
*/

func version_2() {
	// Elements
	element := map[string]map[string]string{
		"H" : map[string]string{
				"name" : "Hydrogen",
				"state" : "Gas",
		},
		"He" : map[string]string{
				"name" : "Helium",
				"state" : "Gas",
		},
		"Li" : map[string]string{
				"name" : "Lithium",
				"state" : "Solid",
		},
		"Be" : map[string]string{
				"name" : "Berylium",
				"state" : "Solid",
		},
		"B" : map[string]string{
				"name" : "Boron",
				"state" : "Solid",
		},
		"C" : map[string]string{
				"name" : "Carbon",
				"state" : "Solid",
		},
		"N" : map[string]string{
				"name" : "Nitrogen",
				"state" : "Gas",
		},
		"O" : map[string]string{
				"name" : "Oxygen",
				"state" : "Gas",
		},
		"F" : map[string]string{
				"name" : "Flouride",
				"state" : "Gas",
		},
		"Ne" : map[string]string{
				"name" : "Neon",
				"state" : "Gas",
		},
	}

	var input string

	fmt.Println("Enter a symbol: ")
	fmt.Scanf("%s", &input)

	if el, ok := element[input]; ok {
		fmt.Println("Element:", el["name"])
		fmt.Println("State:", el["state"])
	}
}
