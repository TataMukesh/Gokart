package main

//Init is to To Initiaize some dummy data
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
