package main

import (
	"fmt"
	c "gokart\cart"
	a "gokart\admin"
	"os"
)

//user datatype
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

type Rep struct {
	phone string
	place string
	free  bool
}

func main() {
	users, products, reps := Init()

	fmt.Println(users)
	fmt.Println(products)
	fmt.Println(reps)

	var phone string = ""
	var place string = ""
	var Type string = ""
	var model string = ""
	var qty int = 0
	var choice int

	fmt.Printf("------------MENU----------\n 1)Add User\n 2)View Products\n 3)Place Order\n 4)Cancel Order\n 5)View Ordered items\n 6)Add Delivery Reps\n 7)Viwe all delivery reps\n 8)Assign Delivery Reps\n 9)Exit")

	for {

		fmt.Println("\nEnter your choice")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter phone number")
			fmt.Scanln(&phone)
			fmt.Println("Enter place")
			fmt.Scanln(&place)
			n := User{phone: phone, place: place}
			flag := c.AddUsers(n, users)
			if flag == true {
				users = append(users, n)
			}

		case 2:
			fmt.Println(products)
			fmt.Println("Enter type of product")
			fmt.Scanln(&Type)
			c.GetProducts(Type, products)

		case 3:
			fmt.Println("Enter model to purchase")
			fmt.Scanln(&model)
			fmt.Println("Enter quantity")
			fmt.Scanln(&qty)
			c.PlaceOrder(Type, model, qty, phone, products, users)

		case 4:
			fmt.Println("Cancelling order")
			c.CancelOrder(Type, model, qty, phone, products, users)

		case 5:
			c.GetOrders(phone, users)

		case 6:
			fmt.Println("Enter phone number")
			fmt.Scanln(&phone)
			fmt.Println("Enter place")
			fmt.Scanln(&place)
			n := Rep{phone: phone, place: place}
			flag := a.AddReps(n, reps)
			if flag == true {
				reps = append(reps, n)
			}

		case 7:
			fmt.Println(reps)

		case 8:
			var location string
			var match int
			fmt.Println(users)
			fmt.Scanln(&phone)
			a.AssignReps(users, phone, reps)
			
		case 9:
			os.Exit(3)
		}

	}
}
