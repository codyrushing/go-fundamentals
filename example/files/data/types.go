package main

import "fmt"

type distanceKm float32
type distanceMile float32

// this creates a method on the distanceMile type
func (d distanceMile) toKm() distanceKm {
	return distanceKm(d * 1.60934)
}

func (d distanceKm) toMile() distanceMile {
	return distanceMile(d / 1.60934)
}

func test() bool {
	dMi := distanceMile(26.2)
	dKm := distanceKm(dMi.toKm())
	fmt.Printf("Distance in miles: %v", dKm.toMile())
	fmt.Println("")
	fmt.Printf("Distance in km: %v", dMi.toKm())
	fmt.Println("")
	return true
}

func main() {
	test()
}
