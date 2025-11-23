package main

import (
	"fmt"
	"sort"
)

func main() {
	parkingLot := make(map[int]interface{})

	var command string
	var valueRaw string
	var value interface{}
	var hours int

	for {
		fmt.Print("Enter command (example: create_parking_lot 5): ")
		command, valueRaw, hours = "", "", 0
		fmt.Scanln(&command, &valueRaw, &hours)

		if command == "exit" {
			fmt.Println("Exiting application")
			break
		}

		switch command {
		case "create_parking_lot":
			var capInt int
			_, err := fmt.Sscanf(valueRaw, "%d", &capInt)
			if err == nil && capInt > 0 {
				value = capInt
				for i := 1; i <= capInt; i++ {
					parkingLot[i] = nil
				}
				fmt.Printf("%s %d\n", command, capInt)
			} else {
				fmt.Println("Unknown command or invalid parking lot number")
			}
		case "park":
			if len(parkingLot) == 0 {
				fmt.Println("Parking lot not created")
				continue
			}
			if valueRaw == "" {
				fmt.Println("Invalid car plate, must be string")
			} else {
				value = valueRaw
				found := false
				count := len(parkingLot)
				for i := 1; i <= count; i++ {
					if parkingLot[i] == nil {
						parkingLot[i] = value
						fmt.Printf("Allocated slot number: %d\n", i)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Sorry, parking lot is full")
				}
			}
		case "leave":
			if len(parkingLot) == 0 {
				fmt.Println("Parking lot not created")
				continue
			}
			if valueRaw == "" || hours <= 0 {
				fmt.Println("Car number must be a string and hours must be a positive number")
			} else {
				value = valueRaw
				found := false

				for slot, parkingValue := range parkingLot {
					if value == parkingValue {
						var fee int
						if hours <= 2 {
							fee = 10
						} else {
							fee = 10 + (hours-2)*10
						}
						fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", value, slot, fee)
						parkingLot[slot] = nil
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Registration number", value, "not found")
				}
			}
		case "status":
			if len(parkingLot) == 0 {
				fmt.Println("Parking lot not created")
				continue
			} else {
				fmt.Println("Slot No. Registration No.")

				var keys []int
				for key := range parkingLot {
					keys = append(keys, key)
				}
				sort.Ints(keys)
				for _, key := range keys {
					if parkingLot[key] != nil {
						fmt.Printf("%d %v\n", key, parkingLot[key])
					}
				}
			}
		default:
		}
	}
}
