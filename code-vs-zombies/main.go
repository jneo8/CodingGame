package main

import (
	"fmt"
	"math"
	"os"
)

/**
 * Save humans, destroy zombies!
 **/

type Hurnter struct {
	x float64
	y float64
}

type Human struct {
	id int
	x  float64
	y  float64
}

type Zombie struct {
	id    int
	xNext float64
	yNext float64
	x     float64
	y     float64
}

type Humans []Human

type Zombies []Zombie

func distance(p1X float64, p1Y float64, p2X float64, p2Y float64) float64 {
	return math.Pow(math.Pow(p1X-p2X, 2)+math.Pow(p1Y-p2Y, 2), 0.5)
}

func findTargetZombie(ht Hurnter, humans Humans, zombies Zombies) Zombie {

	targetZombie := zombies[0]
	min := 25000.0
	// Find target Zombie
	for _, z := range zombies {
		for _, h := range humans {
			d1 := distance(z.x, z.y, h.x, h.y)
			d2 := distance(z.xNext, z.yNext, ht.x, ht.y)
			d3 := distance(z.xNext, z.yNext, h.x, h.y)
			if d1+d2 < min && d3 > 400 {
				min = d1 + d2
				targetZombie = z
			}
		}
	}
	return targetZombie
}

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)

		hurnter := Hurnter{x: float64(x), y: float64(y)}
		fmt.Fprintf(os.Stderr, "Hurnter: %#v\n", hurnter)

		var humanCount int
		fmt.Scan(&humanCount)

		humans := Humans{}
		for i := 0; i < humanCount; i++ {
			var humanId, humanX, humanY int
			fmt.Scan(&humanId, &humanX, &humanY)
			human := Human{
				id: humanId,
				x:  float64(humanX),
				y:  float64(humanY),
			}
			humans = append(humans, human)
		}

		var zombieCount int
		fmt.Scan(&zombieCount)

		zombies := Zombies{}

		for i := 0; i < zombieCount; i++ {
			var zombieId, zombieX, zombieY, zombieXNext, zombieYNext int
			fmt.Scan(&zombieId, &zombieX, &zombieY, &zombieXNext, &zombieYNext)
			zombie := Zombie{
				id:    zombieId,
				xNext: float64(zombieXNext),
				yNext: float64(zombieYNext),
				x:     float64(zombieX),
				y:     float64(zombieY),
			}
			zombies = append(zombies, zombie)
		}

		targetZombie := findTargetZombie(hurnter, humans, zombies)
		fmt.Fprintf(os.Stderr, "targetZombie: %#v\n", targetZombie)

		move := fmt.Sprintf("%v %v", int(targetZombie.xNext), int(targetZombie.yNext))
		fmt.Fprintf(os.Stderr, "move: %v\n", move)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println(move) // Your destination coordinates
	}
}
