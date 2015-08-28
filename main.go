package main

import (
	"fmt"
	"log"
	"os/user"
)

type Developer struct {
	Name    string
	Company string
}

func (d *Developer) sayGo() {
	fmt.Println(d.Name, "says go!")
}

type Gopher interface {
	sayGo()
}

func main() {
	dev, err := developer()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %v!, good to see %v at Codepot\n", dev.Name, dev.Company)

	fmt.Println("Hello for all your fellows!!")

	companies := make(map[string][]Gopher)

	for _, fellow := range fellows() {
		fmt.Printf("Hello %v!\n", fellow.Name)
		companies[fellow.Company] = append(companies[fellow.Company], &fellow)

		fellow.sayGo()
	}

	for company := range companies {
		fmt.Printf("Good to see %v at Codepot!\n", company)
	}
}

func developer() (*Developer, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &Developer{Name: currentUser.Username, Company: "Allegro"}, nil
}

func fellows() []Developer {
	return []Developer{
		Developer{"Piotrek", "Allegro"},
		Developer{"Janek", "SoftwareMill"},
	}
}
