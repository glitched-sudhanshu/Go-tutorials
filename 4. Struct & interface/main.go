package main

import (
	"fmt"
)

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	company
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func milesLeft(e gasEngine) uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	milesLeft() uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type owner struct {
	name string
}

type company struct {
	title string
}

func canMakeIt(e engine, milesToCover uint8) bool {
	if e.milesLeft() >= milesToCover {
		return true
	} else {
		return false
	}
}

func main() {
	var myEngine gasEngine
	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var newEngine gasEngine = gasEngine{3, 4, owner{"Ash Ketchum"}, company{"Audi"}}
	fmt.Println(newEngine)
	newEngine.mpg = 31
	fmt.Println(newEngine)
	fmt.Printf("The miles left are %v\n", newEngine.milesLeft())
	fmt.Printf("Can we make it? %v\n", canMakeIt(newEngine, 130))

	shortEngine := gasEngine{4, 5, owner{"Ash Ketchum"}, company{"Audi"}}
	fmt.Println(shortEngine, shortEngine.ownerInfo.name, shortEngine.title, shortEngine.company.title)

	// anonymous struct
	var engine = struct {
		mpg  uint16
		name string
	}{23, "Keller"}
	fmt.Println(engine)

	electricEngine := electricEngine{4, 5}
	fmt.Printf("Can we make it? %v\n", canMakeIt(electricEngine, 130))
}
