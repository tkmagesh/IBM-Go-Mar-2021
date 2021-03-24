package main

import "fmt"

/*
Array - fixed size
Slice
Range
Map (key,value)
*/
func main() {
	//Array
	var nos [10]int
	for i := 0; i < 10; i++ {
		nos[i] = i * 2
	}
	for i := 0; i < 10; i++ {
		fmt.Println(nos[i])
	}

	/* strs := [2]string{"Hello", "World"} */
	strs := [...]string{"Hello", "World"}
	//fmt.Println(strs)
	fmt.Printf("%q\n", strs)

	//multi dimensional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row - %d coloum - %d", i, j)
		}
	}
	fmt.Printf("%q\n", matrix)

	//slice
	oddNos := []int{1, 3, 5, 7, 9}
	fmt.Println(oddNos)
	/*
		[lo : hi]
		[lo : lo] => empty
		[lo : lo+1] => one element
	*/

	fmt.Println(oddNos[0:4])
	fmt.Println(oddNos[1:1])
	fmt.Println(oddNos[2 : 2+1])
	fmt.Println(oddNos[:4])
	fmt.Println(oddNos[3:])

	/*
		cities := make([]string, 3)
		cities[0] = "Bengaluru"
		cities[1] = "Pune"
		cities[2] = "Mumbai"
		fmt.Println(cities)
	*/

	cities := make([]string, 0)

	cities = append(cities, "Bengaluru")
	cities = append(cities, "Pune")
	cities = append(cities, "Mumbai")
	cities = append(cities, "Delhi", "Amritsar")
	fmt.Println(cities)
	fmt.Println(len(cities))

	//cities = cities[1:]
	/* for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	} */

	//range
	for _, city := range cities {
		if len(city) == 4 {
			break
		}
		fmt.Println(city)
	}

	//map
	stateRank := map[string]int{
		"Karnataka":   3,
		"Goa":         1,
		"Maharashtra": 4,
	}

	//adding a new k,v pair to the map
	stateRank["Kerala"] = 2

	//retrieving a value using the key
	goaRank := stateRank["Goa"]
	fmt.Printf("Goa rank = %d\n", goaRank)

	//fmt.Printf("%#v\n", stateRank)
	for k, v := range stateRank {
		fmt.Println(k, v)
	}

	//removing a key value pair from map
	delete(stateRank, "Goa")
	fmt.Println("After deleting goa")
	for k, v := range stateRank {
		fmt.Println(k, v)
	}

	_, exists := stateRank["Maharashtra"]
	fmt.Println(exists)

}
