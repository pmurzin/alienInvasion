# alienInvasion

This repo contains Golang solution to the problem stated in **assignment.txt**.
It is assumed that aliens are obliged to move every round (1 hop) until either
they got stuck or dead.

**Build:**
```
go build invasion.go map.go alien.go city.go
```

**Run:**
```
./invasion <NumOfAliens>
```

**Test:**
```
go test
```
