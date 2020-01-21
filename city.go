package main

type city string

func removeCityFromConnections(c city, cSlice []city) []city {
	for idx, v := range cSlice {
		if v == c {
			return append(cSlice[0:idx], cSlice[idx+1:]...)
		}
	}
	return cSlice
}

func destroyCity(wMap wmap, dCity city) {

	// remove connections
	connectionsToDelete := wMap[dCity]
	for _, conn := range connectionsToDelete {
		wMap[conn] = removeCityFromConnections(dCity, wMap[conn])
	}

	// remove city
	delete(wMap, dCity)
}

func citiesFromWorldMap(wMap wmap) []city {
	var cities []city
	for c := range wMap {
		cities = append(cities, c)
	}

	return cities
}
