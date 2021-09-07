package main

import "fmt"

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address Address
}

type Address struct {
	City   string
	Street string
	Home   string
}

func printInfo(s Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Rate:", s.Rate)
	fmt.Println("Active:", s.Active)
	fmt.Println("Address:", s.Address.City, s.Address.Street, s.Address.Home)
}

func createSubscriber(name string) Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 11.99
	s.Active = true
	return s
}

// Здесь указатель используем тк функция будет работать с самим подпесчиком
func applyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

func addAddress(city string, street string, home string, s *Subscriber) {
	s.Address.City = city
	s.Address.Street = street
	s.Address.Home = home
}

func main() {
	subscriber1 := createSubscriber("Test")
	printInfo(subscriber1)
	applyDiscount(&subscriber1)
	addAddress("Moscow", "Pushkina", "Kolotushkina", &subscriber1)
	fmt.Println()
	printInfo(subscriber1)
}
