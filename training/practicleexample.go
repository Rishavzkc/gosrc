package main

import (
	"fmt"
	"math/rand"
	"time"
)

//build a hotel application.  the receptionist will welcome customers.
//Clients will ask the receptionist for an available room. He needs to know :

// 1. How many rooms are available now . 2. For how long each room is available

func main() {

	const hotelName = "Windsor International"
	const totalRooms = 134
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := (float64(roomsOccupied) / float64(totalRooms)) * 100

	var occupancyLevel string

	if occupancyRate > 70 {
		occupancyLevel = "High"
	} else if occupancyRate > 20 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("                  Occupancy Level:", occupancyLevel)
	fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
		}
	} else {
		fmt.Println(" No rooms available for today")
	}

}
