package CabBookingSystem

import (
	"fmt"
	"math"
)

type Location struct {
	x, y int
}

func (location *Location) toDistance(other Location) float64 {
	return math.Sqrt(math.Pow(float64(location.x-other.x), 2) + math.Pow(float64(location.y-other.y), 2))
}

type Cab struct {
	id, name    string
	location    Location
	isAvailable bool
}

type Rider struct {
	id, name string
}

type Trip struct {
	to, from  Location
	cab       *Cab
	rider     *Rider
	completed bool
}

type CabBookingSystem struct {
	cabs       map[string]*Cab
	riders     map[string]*Rider
	riderTrips map[string][]*Trip
}

func (cbs *CabBookingSystem) registerCab(id, driverName string, location Location) {
	cbs.cabs[id] = &Cab{id: id, name: driverName, location: Location{x: location.x, y: location.y},
		isAvailable: true}
}

func (cbs *CabBookingSystem) registerRider(id, riderName string) {
	cbs.riders[id] = &Rider{id: id, name: riderName}
}

func (cbs *CabBookingSystem) bookCab(riderId string, fromX, fromY, toX, toY int) {
	if _, exists := cbs.riders[riderId]; !exists {
		fmt.Println("Rider not found")
		return
	}

	var nearestCab *Cab
	minDistance := 5.0
	fromLocation := Location{fromX, fromY}
	for _, cab := range cbs.cabs {
		dist := cab.location.toDistance(fromLocation)
		if dist <= minDistance {
			minDistance = dist
			nearestCab = cab
		}
	}

	if nearestCab == nil {
		fmt.Println("No cabs found nearby")
		return
	}

	nearestCab.isAvailable = false
	cbs.riderTrips[riderId] = append(cbs.riderTrips[riderId], &Trip{from: fromLocation, to: Location{toX, toY}, cab: nearestCab, rider: cbs.riders[riderId], completed: false})
	fmt.Printf("Cab %s assigned to rider %s\n", nearestCab.id, riderId)
}

func (cbs *CabBookingSystem) endTrip(riderId string) {
	if trips, exists := cbs.riderTrips[riderId]; !exists || len(trips) == 0 {
		fmt.Println("No trips found for rider")
		return
	}

	trips := cbs.riderTrips[riderId]
	latestTrip := trips[len(trips)-1]
	if latestTrip.completed {
		fmt.Println("Latest trip already completed.")
		return
	}

	latestTrip.completed = true
	latestTrip.cab.isAvailable = true
	latestTrip.cab.location = latestTrip.to

	fmt.Printf("Trip completed for rider %s\n", riderId)
}

func (cbs *CabBookingSystem) showTrips(riderId string) {
	if len(cbs.riderTrips[riderId]) == 0 {
		fmt.Println("No trips found for rider.")
		return
	}

	for _, trip := range cbs.riderTrips[riderId] {
		fmt.Printf("Trip: Cab %s From (%d, %d) To (%d, %d) Completed %s\n", trip.cab.id, trip.from.x, trip.from.y, trip.to.x, trip.to.y, map[bool]string{true: "Yes", false: "No"}[trip.completed])
	}
}

func Example() {
	var system CabBookingSystem = CabBookingSystem{
		cabs:       make(map[string]*Cab),
		riders:     make(map[string]*Rider),
		riderTrips: make(map[string][]*Trip),
	}
	system.registerCab("cab1", "Alice", Location{0, 0})
	system.registerCab("cab2", "Bob", Location{10, 10})

	system.registerRider("r1", "John")
	system.bookCab("r1", 1, 1, 5, 5)
	system.showTrips("r1")
	system.endTrip("r1")
	system.showTrips("r1")

}
