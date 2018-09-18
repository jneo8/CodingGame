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

func findTargetZombie(ht Hurnter, humans Humans, zombies Zombies, humanCount float64, zombieCount float64) (int, int) {

	targetZombie := zombies[0]
	targetHuman := humans[0]
	min := 25000.0
	// Find target Zombie
	moveX := 0.0
	moveY := 0.0
	for _, z := range zombies {
		for _, h := range humans {
			d1 := distance(z.x, z.y, h.x, h.y)
			d2 := distance(z.xNext, z.yNext, ht.x, ht.y)
			d3 := distance(z.xNext, z.yNext, h.x, h.y)
			d4 := distance(h.x, h.y, ht.x, ht.y)

			giveup := false
			if d3 < 400 && d2 > 3000 {
				giveup = true
			}

			if d1+d2 < min && giveup == false {
				min = d1 + d2
				targetZombie = z
				targetHuman = h

				if d1 < d4 {
					moveX = (targetZombie.xNext + targetHuman.x) / 2
					moveY = (targetZombie.yNext + targetHuman.y) / 2
				} else {
					moveX = targetZombie.xNext
					moveY = targetZombie.yNext
				}
			}
		}
	}
	return int(moveX), int(moveY)
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

		moveX, moveY := findTargetZombie(
			hurnter,
			humans,
			zombies,
			float64(humanCount),
			float64(zombieCount),
		)

		move := fmt.Sprintf("%v %v", moveX, moveY)

		fmt.Fprintf(os.Stderr, "move: %v\n", move)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println(move) // Your destination coordinates
	}
}
