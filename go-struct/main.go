package main

import "fmt"

/*----------------- Struct and Interface------------------*/

/*
type gasEngine struct {
	mpg uint8
	gallon uint8
	// ownerInfo owner
	owner
}*/

// type owner struct {
// 	name string
// }
type gasEngine struct {
	mpg uint8
	gallon uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallon * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.	kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if (miles <= e.milesLeft()) {
		fmt.Println("U can make it")
	} else {
		fmt.Println("Need to fuel up first")
	}
}


func main(){
	// var myEngine gasEngine = gasEngine{25, 10, owner{"Man"}}
	// myEngine.mpg = 20
	// var myEngine2 = struct {
	// 	mpg uint8
	// 	gallon uint8
	// }{25,15}

	// var myEngine gasEngine = gasEngine{25,10}
	// fmt.Println(myEngine.mpg, myEngine.gallon)

	var bikeElectric electricEngine = electricEngine{25, 10}
	var bikeGas gasEngine = gasEngine{25, 10}

	fmt.Println(bikeElectric.milesLeft())
	fmt.Println(bikeGas.milesLeft())

	

	
}