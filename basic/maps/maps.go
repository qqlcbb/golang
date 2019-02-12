package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	// m2 == empty map
	m2 := make(map[string]int)
	// m3 == nil
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	fmt.Println("traversing map")
	// k 在map里面是无序的
	for k, v := range(m) {
		fmt.Println(k, v)
	}

	fmt.Println("getting value")
	couseName := m["course"]
	fmt.Println(couseName)

	couseName1, ok := m["course1"]
	fmt.Println(couseName1, ok)

	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
}
