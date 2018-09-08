package main

import "fmt"
import "os"

/**
 * Save humans, destroy zombies!
 **/

type Coordinate struct {
	x int
	y int
}

type Hurnter struct {
	id         int
	coordinate Coordinate
}

type Human struct {
	id         int
	coordinate Coordinate
}

type Zombie struct {
	id          int
	zombieXNext int
	zombieYNext int
	coordinate  Coordinate
}

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)

		hurnter := &Hurnter{id: -1, coordinate: Coordinate{x: x, y: y}}
		fmt.Fprintf(os.Stderr, "Hurnter: %#v\n", hurnter)

		var humanCount int
		fmt.Scan(&humanCount)

		humans := []Human{}
		for i := 0; i < humanCount; i++ {
			var humanId, humanX, humanY int
			fmt.Scan(&humanId, &humanX, &humanY)
			human := Human{
				id:         humanId,
				coordinate: Coordinate{x: humanX, y: humanY},
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
				id:          zombieId,
				zombieXNext: zombieXNext,
				zombieYNext: zombieYNext,
				coordinate:  Coordinate{x: zombieX, y: zombieY},
			}
			zombies = append(zombies, zombie)
		}
		for _, zombie := range zombies {
			fmt.Fprintf(
				os.Stderr,
				"zombie: %v\ncoordinate: %v\nnext: %v %v",
				zombie.id,
				zombie.coordinate,
				zombie.zombieXNext,
				zombie.zombieYNext,
			)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("0 0") // Your destination coordinates
	}
}
