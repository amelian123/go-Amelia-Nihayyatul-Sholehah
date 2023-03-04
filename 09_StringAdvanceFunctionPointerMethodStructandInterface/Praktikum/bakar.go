package main

import (
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float64
}

const fuelEfficiency = 1.5 // km per milliliter

func (c *Car) CalculateEstimatedDistance() float64 {
	return c.FuelIn * fuelEfficiency
}

func main() {
	myCar := Car{
		Type:   "Car",
		FuelIn: 15000, // dalam milliliter
	}
	estimatedDistance := myCar.CalculateEstimatedDistance()
	fmt.Printf("Perkiraan jarak yang bisa ditempuh %s dengan bahan bakar %.0f ml adalah %.2f km\n", myCar.Type, myCar.FuelIn, estimatedDistance)
}
