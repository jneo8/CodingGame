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

func (h Hurnter) move(x float64, y float64) string {
	needX := x - h.x
	needY := y - h.y

	// need := needX + needY
	// if math.Abs(needX) + math.Abs(needY) > 1000.0 {
	//      need = 1000.0
	// }

	moveX := 1000.0 * (needX / (needX + needY))
	moveY := 1000.0 * (needY / (needX + needY))

	finalX := int(h.x + moveX)
	finalY := int(h.y + moveY)
	result := fmt.Sprintf("%v %v", finalX, finalY)
	return result
}

type Human struct {
	id float64
	x  float64
	y  float64
}

type Zombie struct {
	id    float64
	xNext float64
	yNext float64
	x     float64
	y     float64
}

type Humans []Human

func (humans Humans) centerPoint() (float64, float64) {
	totalX := 0.0
	totalY := 0.0
	for _, h := range humans {
		totalX += h.x
		totalY += h.y
	}

	x := totalX / float64(len(humans))
	y := totalY / float64(len(humans))
	return x, y
}

type Zombies []Zombie

func (zombies Zombies) centerPoint() (float64, float64) {
	totalX := 0.0
	totalY := 0.0
	for _, z := range zombies {
		totalX += z.x
		totalY += z.y
	}

	x := totalX / float64(len(zombies))
	y := totalY / float64(len(zombies))
	return x, y
}

func distance(p1X float64, p1Y float64, p2X float64, p2Y float64) float64 {
	return math.Pow(math.Pow(p1X-p2X, 2)+math.Pow(p1Y-p2Y, 2), 0.5)
}

func findTargetZombie(h Hurnter, humans Humans, zombies Zombies) Zombie {
	targetZombie := zombies[0]

	min := 25000.0
	for _, z := range zombies {
		for _, h := range humans {
			d := distance(z.x, z.y, h.x, h.y)
			if d < min {
				min = d
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
				id: float64(humanId),
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
				id:    float64(zombieId),
				xNext: float64(zombieXNext),
				yNext: float64(zombieYNext),
				x:     float64(zombieX),
				y:     float64(zombieY),
			}
			zombies = append(zombies, zombie)
		}
		for _, zombie := range zombies {
			fmt.Fprintf(
				os.Stderr,
				"zombie: %#v\n",
				zombie,
			)
		}

		targetZombie := findTargetZombie(hurnter, humans, zombies)
		fmt.Fprintf(os.Stderr, "targetZombie: %#v\n", targetZombie)

		move := fmt.Sprintf("%v %v", int(targetZombie.xNext), int(targetZombie.yNext))
		fmt.Fprintf(os.Stderr, "move: %v\n", move)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println(move) // Your destination coordinates
	}
}
