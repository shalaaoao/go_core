package main

import "fmt"

func main() {
	var m map[string]int

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d\n",
		len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。

	// 在空map中添加数据，需要先初始化
	m = map[string]int{
		"a": 1,
	}

	// 然后就能m["two"] = 2 添加数据了
}
