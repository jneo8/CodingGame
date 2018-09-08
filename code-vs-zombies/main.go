package main

import "fmt"

//import "os"

/**
 * Save humans, destroy zombies!
 **/

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)

		var humanCount int
		fmt.Scan(&humanCount)

		for i := 0; i < humanCount; i++ {
			var humanId, humanX, humanY int
			fmt.Scan(&humanId, &humanX, &humanY)
		}
		var zombieCount int
		fmt.Scan(&zombieCount)

		for i := 0; i < zombieCount; i++ {
			var zombieId, zombieX, zombieY, zombieXNext, zombieYNext int
			fmt.Scan(&zombieId, &zombieX, &zombieY, &zombieXNext, &zombieYNext)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("0 0") // Your destination coordinates
	}
}
