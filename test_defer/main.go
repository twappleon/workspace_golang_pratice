package main

import "fmt"

func main() {
	// defer func() {
	// 	fmt.Println("defer main")
	// }()

	// fmt.Println("main")

	// test()
	// test2()


	map3 := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
		"e": "5",
		"f": "6",
		"g": "7",
	}

	delete(map3, "a")

	for k, v := range map3 {
		defer func() {
			fmt.Println("map3 test", k, v)
		}()
	}
}

func test() {
	defer func() {
		fmt.Println("defer test")
	}()

	fmt.Println("test")
}

func test2() {
	defer func() {
		fmt.Println("defer test2")
	}()

	fmt.Println("test2")
}
