package main

import (
	"fmt"
	"math/rand"
	"time"
)

type alien struct {
	alienID    int
	movesCount int
	location   city
	gotStuck   bool
}

func genAliens(numOfAliens int, wMap wmap) []alien {

	cities := citiesFromWorldMap(wMap)

	aliens := make([]alien, numOfAliens)
	for i := 0; i < numOfAliens; i++ {
		aliens[i] = alien{
			alienID:    i,
			movesCount: 0,
			location:   locateAlien(cities),
			gotStuck:   false,
		}
	}

	return aliens
}

func locateAlien(cities []city) city {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return cities[r.Intn(len(cities))]
}

func (a alien) print() {
	fmt.Printf("%+v", a)
}

func (a *alien) wander(wMap wmap) {
	// if alien got stuck there is no sense in moving him again,
	// so just don't allow him to move
	if a.gotStuck {
		return
	}

	if len(wMap[a.location]) > 0 {
		a.location = locateAlien(wMap[a.location])
		a.movesCount++
	} else {
		a.gotStuck = true
		fmt.Println("Alien", a.alienID, "got stuck at", a.location)
	}
}

func aliensMadeAllMoves(aliens []alien) bool {
	aliensMoves := true
	for _, alien := range aliens {
		if alien.movesCount < MOVES && !alien.gotStuck {
			aliensMoves = false
			break
		}
	}

	return aliensMoves
}

func removeDeadAlien(al alien, aSlice []alien) []alien {
	for idx, v := range aSlice {
		if v == al {
			return append(aSlice[0:idx], aSlice[idx+1:]...)
		}
	}
	return aSlice
}
