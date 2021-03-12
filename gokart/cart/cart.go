package cart

import (
	"fmt"
	"strings"
)

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
