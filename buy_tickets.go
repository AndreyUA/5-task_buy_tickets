package main

import (
	"fmt"
	"math/rand"
)

const distance = 62_100_000
const secondsPerDay = 86_400

func main() {
	fmt.Println("this is task for buy tickets console app")

	fmt.Println("Spaceline         Days      Trip type      Price")
	fmt.Println("================================================")

	for counter := 0; counter < 10; counter++ {
		company := ""

		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"

		case 1:
			company = "SpaceX"

		case 2:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(14) + 17
		days := distance / speed / secondsPerDay
		price := 20 + speed
		tripType := ""

		if rand.Intn(2) == 0 {
			tripType = "One way"
		} else {
			tripType = "Round-tripe"
			price *= 2
		}

		fmt.Printf("%-17v %-9v %-14v %v\n", company, days, tripType, price)
	}
}
