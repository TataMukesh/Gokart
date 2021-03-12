package main

import (
	"fmt"
	"os"
	"strings"
)

/*type gokart interface{
	AddUsers()
	GetProducts()
	PlaceOrder()
	ChangeOrder()
	GetOrders()
}*/

//user datatype
type User struct {
	phone            string
	place            string
	orders           []Product
	out_for_delivery bool
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

//AddUsers function
func AddUsers(x User, a []User) bool {
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

func GetProducts(Type string, prod []Product) {
	res := []Product{}
	for i := 0; i < len(prod); i++ {
		chk := prod[i].Type
		if strings.Contains(chk, Type) {
			res = append(res, prod[i])
		}
	}

	fmt.Println(res)
}

//pladeorder
func PlaceOrder(Type string, model string, qty int, phone string, cur []Product, users []User) {
	ordered := Product{Type: Type, model: model, qty: qty}
	for i := 0; i < len(users); i++ {
		if phone == users[i].phone {
			users[i].orders = append(users[i].orders, ordered)
		}

	}

	for i := 0; i < len(cur); i++ {
		if cur[i].Type == Type && cur[i].model == model && cur[i].qty >= qty {
			cur[i].qty = cur[i].qty - qty
			fmt.Println("Order Placed Succesfully")
			return
		}
	}
	fmt.Println("No sufficient stocks")
	return

}

func CancelOrder(Type string, model string, qty int, phone string, cur []Product, users []User) {
	inc := 0
	for i := 0; i < len(users); i++ {
		if phone == users[i].phone {
			inc = users[i].orders[0].qty
			users[i].orders = nil
		}

	}

	for i := 0; i < len(cur); i++ {
		if cur[i].Type == Type && cur[i].model == model {
			cur[i].qty = cur[i].qty + inc
			fmt.Println("Order cancelled Succesfully")
			return
		}
	}

}

func GetOrders(phone string, users []User) {
	for i := 0; i < len(users); i++ {
		if phone == users[i].phone && len(users[i].orders) > 0 {
			fmt.Printf("Hello %s....here is your order\n", users[i].phone)
			for j := 0; j < len(users[i].orders); j++ {
				fmt.Printf("%d pieces of %s %s \n", users[i].orders[j].qty, users[i].orders[j].model, users[i].orders[j].Type)
			}
			return
		}

	}

	fmt.Println("Cart is empty")

}

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

func Init() ([]User, []Product, []Rep) {
	users := []User{}
	m := User{phone: "1234567890", place: "home"}
	o := User{phone: "9876543210", place: "house"}
	q := User{phone: "1111111111", place: "hotel"}
	p := User{phone: "2222222222", place: "hostel"}
	users = append(users, m, o, p, q)

	products := []Product{}
	p1 := Product{Type: "mobile", model: "samsung", qty: 30}
	p2 := Product{Type: "mobile", model: "oppo", qty: 10}
	p3 := Product{Type: "tv", model: "mi", qty: 300}
	p4 := Product{Type: "tv", model: "samsung", qty: 70}
	p5 := Product{Type: "laptop", model: "dell", qty: 3}
	products = append(products, p1, p2, p3, p4, p5)

	reps := []Rep{}
	r1 := Rep{phone: "4554455445", place: "home", free: true}
	r2 := Rep{phone: "4554455445", place: "hostel", free: false}
	r3 := Rep{phone: "4554455445", place: "house", free: false}
	reps = append(reps, r1, r2, r3)

	return users, products, reps
}

func main() {
	users, products, reps := Init()

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
			flag := AddUsers(n, users)
			if flag == true {
				users = append(users, n)
			}

		case 2:
			fmt.Println(products)
			fmt.Println("Enter type of product")
			fmt.Scanln(&Type)
			GetProducts(Type, products)

		case 3:
			fmt.Println("Enter model to purchase")
			fmt.Scanln(&model)
			fmt.Println("Enter quantity")
			fmt.Scanln(&qty)
			PlaceOrder(Type, model, qty, phone, products, users)

		case 4:
			fmt.Println("Cancelling order")
			CancelOrder(Type, model, qty, phone, products, users)

		case 5:
			GetOrders(phone, users)

		case 6:
			fmt.Println("Enter phone number")
			fmt.Scanln(&phone)
			fmt.Println("Enter place")
			fmt.Scanln(&place)
			n := Rep{phone: phone, place: place}
			flag := AddReps(n, reps)
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
					users[match].out_for_delivery = true
					break
				} else {
					fmt.Println("Delivery Rep Not Found!")
				}
			}

		case 9:
			os.Exit(3)
		}

	}
}
