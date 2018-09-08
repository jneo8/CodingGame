package main

import "fmt"
import "os"

/**
 * Save humans, destroy zombies!
 **/

type Hurnter struct {
	x int
	y int
}

type Human struct {
	id int
	x  int
	y  int
}

type Zombie struct {
	id    int
	xNext int
	yNext int
	x     int
	y     int
}

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)

		hurnter := &Hurnter{x: x, y: y}
		fmt.Fprintf(os.Stderr, "Hurnter: %#v\n", hurnter)

		var humanCount int
		fmt.Scan(&humanCount)

		humans := []Human{}
		for i := 0; i < humanCount; i++ {
			var humanId, humanX, humanY int
			fmt.Scan(&humanId, &humanX, &humanY)
			human := Human{
				id: humanId,
				x:  humanX,
				y:  humanY,
			}
			humans = append(humans, human)
		}

		fmt.Fprintf(os.Stderr, "humans: %#v\n", humans)

		var zombieCount int
		fmt.Scan(&zombieCount)

		zombies := []Zombie{}

		for i := 0; i < zombieCount; i++ {
			var zombieId, zombieX, zombieY, zombieXNext, zombieYNext int
			fmt.Scan(&zombieId, &zombieX, &zombieY, &zombieXNext, &zombieYNext)
			zombie := Zombie{
				id:    zombieId,
				xNext: zombieXNext,
				yNext: zombieYNext,
				x:     zombieX,
				y:     zombieY,
			}
			zombies = append(zombies, zombie)
		}
		for _, zombie := range zombies {
			fmt.Fprintf(
				os.Stderr,
				"zombie: %#v",
				zombie,
			)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("0 0") // Your destination coordinates
	}
}
