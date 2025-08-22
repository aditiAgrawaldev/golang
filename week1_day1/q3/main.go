package main

import "fmt"

type CalculateSalary interface {
	getSalary() float64
}

type FullTime struct {
	dailyRate  float64
	daysWorked int
}

type Contractor struct {
	dailyRate  float64
	daysWorked int
}

type FreeLancer struct {
	hourlyRate  float64
	hoursWorked int
}

func (f FullTime) getSalary() float64 {
	return f.dailyRate * float64(f.daysWorked)
}

func (c Contractor) getSalary() float64 {
	return c.dailyRate * float64(c.daysWorked)
}

func (fl FreeLancer) getSalary() float64 {
	return fl.hourlyRate * float64(fl.hoursWorked)
}

func main() {
	full := FullTime{dailyRate: 100, daysWorked: 30}
	contract := Contractor{dailyRate: 100, daysWorked: 20}
	lancer := FreeLancer{hourlyRate: 50, hoursWorked: 40}

	employees := []CalculateSalary{full, contract, lancer}

	for _, emp := range employees {
		fmt.Println(emp.getSalary())
	}
}
