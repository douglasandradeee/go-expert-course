package main

import "fmt"

func unique(arr []string) []string {
	result := []string{}
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range arr {
		encountered[arr[v]] = true
	}

	// Place all unique keys from
	// the map into the results array.
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func main() {
	array1 := []string{"a", "a", "A", "b", "b", "c", "a", "1"}
	fmt.Println(array1)
	unique_items := unique(array1)
	fmt.Println(unique_items)

	fmt.Printf("aaaaaa\t\naaaaaaaaa\t\n")
}
