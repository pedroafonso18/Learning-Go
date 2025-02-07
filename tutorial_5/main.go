package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(engine engine, miles uint8) bool {
	if miles <= engine.milesLeft() {
		return true
	} else {
		return false
	}
}

func main() {
	var myEngine gasEngine = gasEngine{30, 10}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	canMakeIt(myEngine, 200)

}
