package main

import (
	"testing"

	"github.com/forestgiant/sliceutil"
)

func TestRemoveCityFromConnections(t *testing.T) {
	cSlice := []city{"Foo", "Bar", "Som", "Wer", "Io"}

	cityToRemove := city("Som")
	newcSlice := removeCityFromConnections(cityToRemove, cSlice)

	if len(cSlice)-1 != len(newcSlice) || sliceutil.Contains(newcSlice, cityToRemove) {
		t.Errorf("Failed to remove city %v from connections", cityToRemove)
	}
}

func TestDestroyCity(t *testing.T) {
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

	iSize := len(wMap)
	cityToDestroy := city("Foo")
	destroyCity(wMap, cityToDestroy)
	iSize1 := len(wMap)

	if iSize-1 != iSize1 {
		t.Errorf("Expected map without city %v, but got another", cityToDestroy)
	}

	for city, conn := range wMap {
		if city == cityToDestroy || sliceutil.Contains(conn, cityToDestroy) {
			t.Errorf("Failed to remove city %v from connections", cityToDestroy)
		}
	}
}
