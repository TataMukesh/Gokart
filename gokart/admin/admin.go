package admin

import "fmt"

type Rep struct {
	phone string
	place string
	free  bool
}

type User struct {
	phone          string
	place          string
	orders         []Product
	outfordelivery bool
}

type Product struct {
	Type  string
	model string
	qty   int
}

//AddReps function
func AddReps(x Rep, a []Rep) bool {
	if len(x.phone) != 10 {
		fmt.Println("Phone number is not 10 digits")
		return false
	}

	for i := 0; i < len(a); i++ {
		if x.phone == a[i].phone {
			fmt.Println("User already registered")
			return false
		}
	}

	fmt.Println("Registered Succesfully!")
	return true
}

//AssignReps to assign reps for deliveries
func AssignReps(users []User, phone string, reps []Rep) {
	for i := 0; i < len(users); i++ {
		if phone == users[i].phone {
			location = reps[i].place
			match = i
		}

	}

	for i := 0; i < len(reps); i++ {
		if location == reps[i].place && reps[i].free == false {
			fmt.Println("Delivery representative assigned succesfully!")
			reps[i].free = true
			users[match].outfordelivery = true
			break
		} else {
			fmt.Println("Delivery Rep Not Found!")
		}
	}

}
