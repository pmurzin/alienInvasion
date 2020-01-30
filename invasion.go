package main

import (
	"fmt"
	"os"
	"strconv"
)

// MOVES is the minimum moves amount for alien that is alive when program is finished
const MOVES int = 10000

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify the number of aliens\n./alien <NumOfAliens>")
		os.Exit(1)
	}

	NumOfAliens, err := strconv.Atoi(os.Args[1])
	if err != nil || NumOfAliens < 1 {
		fmt.Println("Please specify POSITIVE number of aliens :)\n./alien <NumOfAliens>")
		os.Exit(1)
	}

	worldMap := newMapFromFile("worldmap.txt")

	// aliens slice contains only ALIVE aliens
	aliens := genAliens(NumOfAliens, worldMap)

	fmt.Println("------Map of the world before alien invasion------------")
	worldMap.print()

	fmt.Println("---------------Aliens start to fight--------------------")
	for !aliensMadeAllMoves(aliens) {

		// one alien could not destroy city himself
		if len(aliens) == 1 {
			break
		}

		// first figure out whether we already have collisions at the initial stage and
		// only after that make aliens wander
		curCityPopulation := make(map[city][]int)
		for _, alien := range aliens {
			curCityPopulation[alien.location] = append(curCityPopulation[alien.location], alien.alienID)
		}

		for city, aliensIDToKill := range curCityPopulation {
			if len(aliensIDToKill) > 1 {

				// update aliens slice in order not to iterate over dead aliens
				for _, alIDtoKill := range aliensIDToKill {
					aliens = removeDeadAlien(alIDtoKill, aliens)
				}

				fmt.Println(city, "has been destroyed by aliens", intSliceToString(aliensIDToKill, ","))
				destroyCity(worldMap, city)
			}
		}

		fmt.Println("-------------Alive aliens after round-------------------")
		fmt.Println(aliens)
		fmt.Println("--------------------------------------------------------")

		// alive aliens are wandering around
		for i := range aliens {
			aliens[i].wander(worldMap)
		}
	}

	fmt.Println("-------------post apocalyptic alien world---------------")
	worldMap.print()
}
