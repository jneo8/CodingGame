package main

import "fmt"
import "os"

/**
 * Save humans, destroy zombies!
 **/

type Hurnter struct {
	x float64
	y float64
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

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)

		hurnter := &Hurnter{x: float64(x), y: float64(y)}
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

		fmt.Fprintf(os.Stderr, "humans: %#v\n", humans)
		centerX, centerY := humans.centerPoint()
		fmt.Fprintf(os.Stderr, "humans centerpoint: %v %v\n", centerX, centerY)

		var zombieCount int
		fmt.Scan(&zombieCount)

		zombies := []Zombie{}

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
				"zombie: %#v",
				zombie,
			)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("0 0") // Your destination coordinates
	}
}
