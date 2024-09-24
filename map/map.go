// Package map demonstrates how to work with maps in Go.
package main

import "fmt"

func Map() {
	// a hash table - go's map
	m := make(map[int]string) // before use, map need to be initiated. this is initiated wit pair 0,"" values
	m[1] = "Apple"
	m[2] = "Samsung"
	m[3] = "Nokia"
	for key, value := range m {
		fmt.Println("Key", key, "value", value)
	}
	delete(m, 2) // this remove the second element of map m.
	id, val := m[2]
	fmt.Println("id which is looked from the map at key 2:", id, "status either the lookup is successful or not:", val)
	//or
	if id1, ok := m[1]; ok { // if the lookup operation is successful,
		fmt.Println("from the if block of map", id1, ok)
	}
	//more on map:
	employee := map[string]map[string]float64{
		"Tigst": {
			"salary": 10000,
			"age":    24,
		},
		"Abel": {
			"experience": 4,
			"salary":     41000,
			"age":        36,
		},
		"Yina": {
			"woreda": 12,
			"age":    23,
		},
	}
	if emp, ok := employee["Abel"]; ok {
		fmt.Println("Experience", emp["experience"], "Salary:", emp["salary"])
	}
}
