package main

import (
	"testing"

	"github.com/forestgiant/sliceutil"
)

func TestGenAliens(t *testing.T) {
	wMap := wmap{
		"Foo":   {"Baz", "Sef", "Bar", "Qu-ux"},
		"Bee":   {"Baz", "Bar"},
		"Qu-ux": {"Foo", "Cro", "Gon"},
		"Sef":   {"Foo", "Cro", "Olsa"},
		"Gon":   {"Baz", "Qu-ux"},
		"Bar":   {"Foo", "Bee", "Olsa"},
		"Baz":   {"Bee", "Foo", "Gon"},
		"Cro":   {"Sef", "Qu-ux"},
		"Olsa":  {"Sef", "Bar"},
	}

	aliens := genAliens(5, wMap)

	if len(aliens) != 5 {
		t.Errorf("Expected aliens slice length of 5, but got %v", len(aliens))
	}

	cities := citiesFromWorldMap(wMap)
	for _, al := range aliens {
		if !sliceutil.Contains(cities, al.location) {
			t.Errorf("Alien %v was placed to unknown destination", al.alienID)
		}
	}
}

func TestLocateAlien(t *testing.T) {
	cities := []city{"Foo", "Bar", "Azz", "Uhe", "Oli", "Ur"}
	c := locateAlien(cities)

	if !sliceutil.Contains(cities, c) {
		t.Errorf("Unknown destination was chosen")
	}

}

func TestRemoveDeadAlien(t *testing.T) {

	aliens := []alien{
		{
			alienID:    0,
			movesCount: 10,
			location:   "Nda",
			gotStuck:   false,
		},
		{
			alienID:    1,
			movesCount: 10,
			location:   "Nae",
			gotStuck:   false,
		},
		{
			alienID:    2,
			movesCount: 10,
			location:   "Nui",
			gotStuck:   false,
		},
	}

	al := alien{
		alienID:    1,
		movesCount: 10,
		location:   "Nae",
		gotStuck:   false,
	}

	aliens = removeDeadAlien(al, aliens)

	if sliceutil.Contains(aliens, al) {
		t.Errorf("Expected alien %v to be removed from aliens slice, but zombie is still there!", al.alienID)
	}
}

func TestWander(t *testing.T) {
	wMap := wmap{
		"Foo":   {"Baz", "Sef", "Bar", "Qu-ux"},
		"Bee":   {"Baz", "Bar"},
		"Qu-ux": {"Foo", "Cro", "Gon"},
		"Sef":   {"Foo", "Cro", "Olsa"},
		"Gon":   {"Baz", "Qu-ux"},
		"Bar":   {"Foo", "Bee", "Olsa"},
		"Baz":   {"Bee", "Foo", "Gon"},
		"Cro":   {"Sef", "Qu-ux"},
		"Olsa":  {"Sef", "Bar"},
	}

	// alien is able to move
	al := alien{
		alienID:    0,
		movesCount: 10,
		location:   "Foo",
		gotStuck:   false,
	}

	al.wander(wMap)

	if al.movesCount != 11 {
		t.Errorf("Alien movesCount has not been updated")
	}

	if al.location == "Foo" {
		t.Errorf("Alien location has not been updated")
	}

	if !sliceutil.Contains(wMap["Foo"], al.location) {
		t.Errorf("Incorrect updated alien location")
	}

	// alien has got stuck
	al2 := alien{
		alienID:    0,
		movesCount: 10,
		location:   "Foo",
		gotStuck:   true,
	}

	al2.wander(wMap)

	if al2.movesCount != 10 {
		t.Errorf("movesCount of stuck alien was updated")
	}

	if al2.location != "Foo" {
		t.Errorf("location of stuck alien was updated")
	}
}
