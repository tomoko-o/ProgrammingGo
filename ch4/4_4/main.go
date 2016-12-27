package main

import "time"

type Employee struct {
	ID            int
	Name, Address string
	// Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var delbert Employee
	delbert.Salary -= 5000
	position := &delbert.Position
	*position = "Senior" + *position
}
